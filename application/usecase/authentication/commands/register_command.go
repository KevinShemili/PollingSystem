package commands

import (
	"gin/api/requests"
	repo "gin/application/repository/contracts"
	"gin/application/usecase/authentication/commands/contracts"
	"gin/application/utility"
	"gin/domain/entities"
	"gin/infrastructure/mail"
	"log"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type RegisterCommand struct {
	UnitOfWork repo.IUnitOfWork
	Validator  *validator.Validate
}

func NewRegisterCommand(UnitOfWork repo.IUnitOfWork, Validator *validator.Validate) contracts.IRegisterCommand {
	return &RegisterCommand{UnitOfWork: UnitOfWork, Validator: Validator}
}

func (r RegisterCommand) Register(request *requests.RegisterRequest) (bool, *utility.ErrorCode) {

	// validate request
	if err := r.Validator.Struct(request); err != nil {
		return false, utility.ValidationError.WithDescription(err.Error())
	}

	// check if email already exists
	duplicate, err := r.UnitOfWork.IUserRepository().GetByEmail(request.Email)
	if err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}
	if duplicate != nil {
		return false, utility.DuplicateEmail
	}

	// validate email and password format
	if flag := utility.ValidateEmail(request.Email); !flag {
		return false, utility.EmailFormat
	}
	if flag := utility.ValidatePassword(request.Password); !flag {
		return false, utility.PasswordFormat
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	// create user
	if err := r.UnitOfWork.IUserRepository().Create(&entities.User{
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		Age:          request.Age,
		Email:        request.Email,
		PasswordHash: string(hash),
	}); err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	// send email
	go func() {
		if err := mail.SendEmail(
			request.Email,
			"Registration Successful",
			mail.GetTemplatePath("register_template.html"),
			map[string]string{
				"FirstName": request.FirstName,
			},
		); err != nil {
			log.Printf("Failed to send email. %v", err)
		}
	}()

	return true, nil
}
