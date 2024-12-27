package commands

import (
	"gin/api/requests"
	repo "gin/application/repository/contracts"
	"gin/application/usecase/poll/commands/contracts"
	"gin/application/utility"
	"gin/domain/entities"
	"time"

	"github.com/go-playground/validator/v10"
)

type UpdatePollCommand struct {
	UnitOfWork repo.IUnitOfWork
	Validator  *validator.Validate
}

func NewUpdatePollCommand(UnitOfWork repo.IUnitOfWork, Validator *validator.Validate) contracts.IUpdatePollCommand {
	return &UpdatePollCommand{UnitOfWork: UnitOfWork, Validator: Validator}
}

func (r UpdatePollCommand) UpdatePoll(userID uint, request *requests.UpdatePollRequest) (bool, *utility.ErrorCode) {

	if err := r.Validator.Struct(request); err != nil {
		return false, utility.ValidationError.WithDescription(err.Error())
	}

	uof, err := r.UnitOfWork.Begin()
	if err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}
	defer uof.Rollback()

	// check if poll exists
	poll, err := r.UnitOfWork.IPollRepository().GetPollWithCategories(request.PollID)
	if err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}
	if poll == nil {
		return false, utility.InvalidPollID
	}

	// check if ended
	if poll.IsEnded {
		return false, utility.AlreadyEnded
	}

	// check if user is the owner
	if poll.CreatorID != userID {
		return false, utility.NotPollOwner
	}

	// update
	poll.Title = request.Title
	poll.Description = request.Description
	expiresAt, err := time.Parse(time.RFC3339, request.ExpiresAt)
	if err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}
	poll.ExpiresAt = expiresAt

	// delete categories if any
	if len(request.DeleteCategories) > 0 {

		for _, categoryID := range request.DeleteCategories {
			if !contains(&poll.Categories, categoryID) {
				return false, utility.InvalidCategoryID
			}
		}

		for _, categoryID := range request.DeleteCategories {
			if err := r.UnitOfWork.IPollCategoryRepository().SoftDelete(categoryID); err != nil {
				return false, utility.InternalServerError.WithDescription(err.Error())
			}
		}
	}

	// add categories if any
	if len(request.NewCategories) > 0 {
		for _, category := range request.NewCategories {
			newCategory := entities.PollCategory{
				Name:   category,
				PollID: poll.ID,
			}

			if err := r.UnitOfWork.IPollCategoryRepository().Create(&newCategory); err != nil {
				return false, utility.InternalServerError.WithDescription(err.Error())
			}
		}
	}

	if err := r.UnitOfWork.IPollRepository().Update(poll); err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	// commit to db
	if err := uof.Commit(); err != nil {
		return false, utility.InternalServerError.WithDescription(err.Error())
	}

	return true, nil
}

func contains(s *[]entities.PollCategory, id uint) bool {
	for _, category := range *s {
		if category.ID == id {
			return true
		}
	}
	return false
}
