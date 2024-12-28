package integration

import (
	"gin/api/requests"
	"gin/application/repository"
	"gin/application/usecase/authentication/commands"
	"gin/application/utility"
	"gin/domain/entities"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// db connection for testing
func loginTestDB(t *testing.T) *gorm.DB {

	db := EnsureDatabaseExists(t)

	err := db.AutoMigrate(&entities.User{}, &entities.RefreshToken{})
	if err != nil {
		t.Fatalf("failed to migrate tables: %v", err)
	}

	return db
}

// drop tables
func loginCleanupDB(t *testing.T, db *gorm.DB) {
	err := db.Migrator().DropTable(&entities.User{}, &entities.RefreshToken{})
	if err != nil {
		t.Fatalf("failed to drop tables: %v", err)
	}
}

func TestLoginCommand_HappyPath_ReturnsResult(t *testing.T) {

	// Arrange
	db := loginTestDB(t)
	defer loginCleanupDB(t, db)

	uow := repository.NewUnitOfWork(db)
	validate := validator.New()
	loginCommand := commands.NewLoginCommand(uow, validate)

	password := "Unifi2024"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	testUser := &entities.User{
		Email:        "hekow17447@myweblaw.com",
		PasswordHash: string(hashedPassword),
	}

	err := db.Create(testUser).Error
	if err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	request := &requests.LoginRequest{
		Email:    "hekow17447@myweblaw.com",
		Password: password,
	}

	// Act
	result, errCode := loginCommand.Login(request)

	// Assert
	assert.Nil(t, errCode)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.JWTToken)
	assert.NotEmpty(t, result.RefreshToken)
}

func TestLoginCommand_InvalidEmail_ReturnsError(t *testing.T) {

	// Arrange
	db := loginTestDB(t)
	defer loginCleanupDB(t, db)

	uow := repository.NewUnitOfWork(db)
	validate := validator.New()
	loginCommand := commands.NewLoginCommand(uow, validate)

	// No user with this email
	request := &requests.LoginRequest{
		Email:    "wrongemail@mail.com",
		Password: "xxxxxxxxxxx",
	}

	// Act
	result, errCode := loginCommand.Login(request)

	// Assert
	assert.Nil(t, result)
	assert.NotNil(t, errCode)
	assert.Equal(t, utility.IncorrectEmail, errCode)
}
