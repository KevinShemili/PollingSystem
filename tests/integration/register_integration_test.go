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
	"gorm.io/gorm"
)

// db connection for testing
func registerTestDB(t *testing.T) *gorm.DB {

	db := EnsureDatabaseExists(t)

	err := db.AutoMigrate(&entities.User{}, &entities.RefreshToken{})
	if err != nil {
		t.Fatalf("failed to migrate tables: %v", err)
	}

	return db
}

// drop tables
func registerCleanupDB(t *testing.T, db *gorm.DB) {
	err := db.Migrator().DropTable(&entities.User{}, &entities.RefreshToken{})
	if err != nil {
		t.Fatalf("failed to drop tables: %v", err)
	}
}

func TestRegisterCommand_HappyPath_ReturnsTrue(t *testing.T) {

	// Arrange
	db := registerTestDB(t)
	defer registerCleanupDB(t, db)

	uow := repository.NewUnitOfWork(db)
	validate := validator.New()
	registerCommand := commands.NewRegisterCommand(uow, validate)

	request := &requests.RegisterRequest{
		FirstName: "XXXX",
		LastName:  "XXXX",
		Age:       37,
		Email:     "hekow17447@myweblaw.com",
		Password:  "Unifi2024",
	}

	// Act
	result, errCode := registerCommand.Register(request)

	// Assert
	assert.Nil(t, errCode)
	assert.True(t, result)
}

func TestRegisterCommand_DuplicateEmail_ReturnsError(t *testing.T) {

	// Arrange
	db := registerTestDB(t)
	defer registerCleanupDB(t, db)

	uow := repository.NewUnitOfWork(db)
	validate := validator.New()
	registerCommand := commands.NewRegisterCommand(uow, validate)

	testUser := &entities.User{
		FirstName:    "XXXX",
		LastName:     "XXXX",
		Age:          37,
		Email:        "hekow17447@myweblaw.com",
		PasswordHash: "Unifi2024",
	}

	if err := db.Create(testUser).Error; err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	request := &requests.RegisterRequest{
		FirstName: "xxxxxxx",
		LastName:  "xxxxxxxx",
		Age:       37,
		Email:     "hekow17447@myweblaw.com",
		Password:  "xxxxxxxx",
	}

	// Act
	result, errCode := registerCommand.Register(request)

	// Assert
	assert.False(t, result)
	assert.NotNil(t, errCode)
	assert.Equal(t, utility.DuplicateEmail, errCode)
}

func TestRegisterCommand_InvalidPasswordFormat_ReturnsError(t *testing.T) {

	// Arrange
	db := registerTestDB(t)
	defer registerCleanupDB(t, db)

	uow := repository.NewUnitOfWork(db)
	validate := validator.New()
	registerCommand := commands.NewRegisterCommand(uow, validate)

	request := &requests.RegisterRequest{
		FirstName: "xxxxxxxx",
		LastName:  "xxxxxxxx",
		Age:       30,
		Email:     "xxxxxx@mail.com",
		Password:  "wrong",
	}

	// Act
	result, errCode := registerCommand.Register(request)

	// Assert
	assert.False(t, result)
	assert.NotNil(t, errCode)
	assert.Equal(t, utility.PasswordFormat, errCode)
}

func TestRegisterCommand_InvalidEmailFormat_ReturnsError(t *testing.T) {

	// Arrange
	db := registerTestDB(t)
	defer registerCleanupDB(t, db)

	uow := repository.NewUnitOfWork(db)
	validate := validator.New()
	registerCommand := commands.NewRegisterCommand(uow, validate)

	request := &requests.RegisterRequest{
		FirstName: "xxxxxxxxxxxx",
		LastName:  "xxxxxxxxx",
		Age:       30,
		Email:     "wrong",
		Password:  "Unifi2024",
	}

	// Act
	result, errCode := registerCommand.Register(request)

	// Assert
	assert.False(t, result)
	assert.NotNil(t, errCode)
	assert.Equal(t, utility.EmailFormat, errCode)
}
