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

func voteTestDB(t *testing.T) *gorm.DB {

	db := EnsureDatabaseExists(t)

	err := db.AutoMigrate(&entities.User{}, &entities.RefreshToken{},
		&entities.Poll{}, &entities.PollCategory{}, &entities.Vote{})
	if err != nil {
		t.Fatalf("failed to migrate tables: %v", err)
	}

	return db
}

// drop tables
func voteCleanupDB(t *testing.T, db *gorm.DB) {
	if err := db.Migrator().DropTable(&entities.User{}, &entities.RefreshToken{},
		&entities.Poll{}, &entities.PollCategory{}, &entities.Vote{}); err != nil {
		t.Fatalf("failed to drop tables: %v", err)
	}
}

func TestAddVoteCommand_HappyPath_ReturnsTrue(t *testing.T) {

	// Arrange
	db := voteTestDB(t)
	defer voteCleanupDB(t, db)
	go websocket.HandleBroadcast()

	uow := repository.NewUnitOfWork(db)
	validate := validator.New()
	addVoteCommand := commands.NewAddVoteCommand(uow, validate)

	user := &entities.User{
		FirstName:    "XXXX",
		LastName:     "XXXX",
		Email:        "hekow17447@myweblaw.com",
		PasswordHash: "Unifi2024",
	}
	if err := db.Create(user).Error; err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	poll := &entities.Poll{
		Title:     "Poll",
		ExpiresAt: time.Now().Add(24 * time.Hour),
		IsEnded:   false,
		CreatorID: user.ID,
	}
	if err := db.Create(poll).Error; err != nil {
		t.Fatalf("failed to create poll: %v", err)
	}

	category := &entities.PollCategory{
		Name:   "Category",
		PollID: poll.ID,
	}
	if err := db.Create(category).Error; err != nil {
		t.Fatalf("failed to create category: %v", err)
	}

	request := &requests.AddVoteRequest{
		PollID:         poll.ID,
		PollCategoryID: category.ID,
	}

	// Act
	result, errCode := addVoteCommand.AddVote(request, user)

	// Assert
	assert.Nil(t, errCode)
	assert.True(t, result)
}

func TestAddVoteCommand_EndedPoll_ReturnsError(t *testing.T) {

	// Arrange
	db := voteTestDB(t)
	defer voteCleanupDB(t, db)

	uow := repository.NewUnitOfWork(db)
	validate := validator.New()
	addVoteCommand := commands.NewAddVoteCommand(uow, validate)

	user := &entities.User{
		FirstName:    "XXXX",
		LastName:     "XXXX",
		Email:        "hekow17447@myweblaw.com",
		PasswordHash: "Unifi2024",
	}
	if err := db.Create(user).Error; err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	poll := &entities.Poll{
		Title:     "Poll",
		ExpiresAt: time.Now(),
		IsEnded:   true,
		CreatorID: user.ID,
	}
	if err := db.Create(poll).Error; err != nil {
		t.Fatalf("failed to create poll: %v", err)
	}

	category := &entities.PollCategory{
		Name:   "Category",
		PollID: poll.ID,
	}
	if err := db.Create(category).Error; err != nil {
		t.Fatalf("failed to create category: %v", err)
	}

	request := &requests.AddVoteRequest{
		PollID:         poll.ID,
		PollCategoryID: category.ID,
	}

	// Act
	result, errCode := addVoteCommand.AddVote(request, user)

	// Assert
	assert.False(t, result)
	assert.NotNil(t, errCode)
	assert.Equal(t, utility.PollExpired, errCode)
}

func TestAddVoteCommand_AlreadyVoted_ReturnsError(t *testing.T) {

	// Arrange
	db := voteTestDB(t)
	defer voteCleanupDB(t, db)

	uow := repository.NewUnitOfWork(db)
	validate := validator.New()
	addVoteCommand := commands.NewAddVoteCommand(uow, validate)

	user := &entities.User{
		FirstName:    "XXXX",
		LastName:     "XXXX",
		Email:        "hekow17447@myweblaw.com",
		PasswordHash: "Unifi2024",
	}
	if err := db.Create(user).Error; err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	poll := &entities.Poll{
		Title:     "Poll",
		ExpiresAt: time.Now(),
		IsEnded:   false,
		CreatorID: user.ID,
	}
	if err := db.Create(poll).Error; err != nil {
		t.Fatalf("failed to create poll: %v", err)
	}

	category := &entities.PollCategory{
		Name:   "Category",
		PollID: poll.ID,
	}
	if err := db.Create(category).Error; err != nil {
		t.Fatalf("failed to create category: %v", err)
	}

	vote := &entities.Vote{
		UserID:         user.ID,
		PollCategoryID: category.ID,
	}
	if err := db.Create(vote).Error; err != nil {
		t.Fatalf("failed to create vote: %v", err)
	}

	request := &requests.AddVoteRequest{
		PollID:         poll.ID,
		PollCategoryID: category.ID,
	}

	// Act
	result, errCode := addVoteCommand.AddVote(request, user)

	// Assert
	assert.False(t, result)
	assert.NotNil(t, errCode)
	assert.Equal(t, utility.AlreadyVoted, errCode)
}

func TestAddVoteCommand_InvalidCategoryID(t *testing.T) {
	// Arrange
	db := voteTestDB(t)
	defer voteCleanupDB(t, db)

	uow := repository.NewUnitOfWork(db)
	validate := validator.New()
	addVoteCommand := commands.NewAddVoteCommand(uow, validate)

	user := &entities.User{
		FirstName:    "XXXX",
		LastName:     "XXXX",
		Email:        "hekow17447@myweblaw.com",
		PasswordHash: "Unifi2024",
	}
	if err := db.Create(user).Error; err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	poll := &entities.Poll{
		Title:     "Poll",
		ExpiresAt: time.Now(),
		IsEnded:   false,
		CreatorID: user.ID,
	}
	if err := db.Create(poll).Error; err != nil {
		t.Fatalf("failed to create poll: %v", err)
	}

	request := &requests.AddVoteRequest{
		PollID:         poll.ID,
		PollCategoryID: 9999,
	}

	// Act
	result, errCode := addVoteCommand.AddVote(request, user)

	// Assert
	assert.False(t, result)
	assert.NotNil(t, errCode)
	assert.Equal(t, utility.InvalidCategoryID, errCode)
}
