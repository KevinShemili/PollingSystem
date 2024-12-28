package integration

import (
	"gin/api/requests"
	"gin/application/repository"
	"gin/application/usecase/poll/commands"
	"gin/application/utility"
	"gin/domain/entities"
	"gin/infrastructure/websocket"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func createPollTestDB(t *testing.T) *gorm.DB {

	db := EnsureDatabaseExists(t)

	err := db.AutoMigrate(&entities.User{}, &entities.RefreshToken{},
		&entities.Poll{}, &entities.PollCategory{})
	if err != nil {
		t.Fatalf("failed to migrate tables: %v", err)
	}

	return db
}

// drop tables
func createPollCleanupDB(t *testing.T, db *gorm.DB) {
	if err := db.Migrator().DropTable(&entities.User{}, &entities.RefreshToken{},
		&entities.Poll{}, &entities.PollCategory{}); err != nil {
		t.Fatalf("failed to drop tables: %v", err)
	}
}

func TestCreatePollCommand_HappyPath(t *testing.T) {

	// Arrange
	db := createPollTestDB(t)
	defer createPollCleanupDB(t, db)
	go websocket.HandleBroadcast()

	uow := repository.NewUnitOfWork(db)
	validate := validator.New()
	createPollCommand := commands.NewCreatePollCommand(uow, validate)

	testUser := &entities.User{
		FirstName:    "XXXX",
		LastName:     "XXXX",
		Age:          30,
		Email:        "hekow17447@myweblaw.com",
		PasswordHash: "Unifi2024",
	}
	if err := db.Create(testUser).Error; err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	expiresAt := time.Now().Add(24 * time.Hour).Format(time.RFC3339)
	request := &requests.CreatePollRequest{
		Title:      "Test",
		ExpiresAt:  expiresAt,
		Categories: []string{"Category1", "Category2"},
	}

	// Act
	result, errCode := createPollCommand.CreatePoll(request, testUser)

	// Assert
	assert.Nil(t, errCode)
	assert.NotNil(t, result)
	assert.Equal(t, "Test", result.Title)
	assert.Equal(t, request.Categories, result.Categories)
}

func TestCreatePollCommand_ExpiredTime(t *testing.T) {
	// Arrange
	db := createPollTestDB(t)
	defer createPollCleanupDB(t, db)

	uow := repository.NewUnitOfWork(db)
	validate := validator.New()
	createPollCommand := commands.NewCreatePollCommand(uow, validate)

	testUser := &entities.User{
		FirstName:    "John",
		LastName:     "Doe",
		Age:          30,
		Email:        "john.doe@example.com",
		PasswordHash: "hashedpassword",
	}
	err := db.Create(testUser).Error
	if err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	expiresAt := time.Now().Add(-24 * time.Hour).Format(time.RFC3339)
	request := &requests.CreatePollRequest{
		Title:      "Test Poll",
		ExpiresAt:  expiresAt,
		Categories: []string{"Category 1", "Category 2"},
	}

	// Act
	result, errCode := createPollCommand.CreatePoll(request, testUser)

	// Assert
	assert.Nil(t, result)
	assert.NotNil(t, errCode)
	assert.Equal(t, utility.DateShouldBeFuture, errCode)
}
