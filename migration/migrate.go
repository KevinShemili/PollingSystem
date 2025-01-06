package main

import (
	"gin/domain/entities"
	"gin/infrastructure/database"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	loadEnvironmentVariables()

	ensureDatabaseExists()

	database, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}

	dbContext := database.GetDBContext()

	err = dbContext.AutoMigrate(
		&entities.User{},
		&entities.RefreshToken{},
		&entities.Poll{},
		&entities.PollCategory{},
		&entities.Vote{},
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Migration Complete.")

	seedData(dbContext)
}

func loadEnvironmentVariables() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ensureDatabaseExists() {

	serverName := os.Getenv("SERVER_STRING")
	databaseName := os.Getenv("DATABASE_NAME")

	server, err := gorm.Open(postgres.Open(serverName), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	var result int64
	server.Raw("SELECT 1 FROM pg_database WHERE datname = ?", databaseName).Count(&result)

	if result == 0 {
		createDatabase(server, databaseName)
	}
}

func createDatabase(server *gorm.DB, databaseName string) {

	err := server.Exec("CREATE DATABASE " + databaseName).Error
	if err != nil {
		log.Fatal(err.Error())
	}
}

func seedData(db *gorm.DB) {

	if err := seedUsers(db); err != nil {
		log.Fatal(err)
	}

	if err := seedPollsAndCategoriesAndVotes(db); err != nil {
		log.Fatal(err)
	}
}

func seedUsers(db *gorm.DB) error {

	// already seeded
	var count int64
	if err := db.Model(&entities.User{}).Count(&count).Error; err != nil {
		log.Fatal(err.Error())
	}
	if count > 0 {
		return nil
	}

	users := []entities.User{
		{
			FirstName: "User 1",
			LastName:  "User 1",
			Email:     "user1@gmail.com",
			Age:       30,
			PasswordHash: func() string {
				hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
				if err != nil {
					log.Fatal(err)
				}
				return string(hash)
			}(),
		},
		{
			FirstName: "User 2",
			LastName:  "User 2",
			Email:     "user2@gmail.com",
			Age:       25,
			PasswordHash: func() string {
				hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
				if err != nil {
					log.Fatal(err)
				}
				return string(hash)
			}(),
		},
	}

	if err := db.Create(&users).Error; err != nil {
		log.Fatal(err.Error())
	}

	return nil
}

func seedPollsAndCategoriesAndVotes(db *gorm.DB) error {

	// already seeded
	var count int64
	if err := db.Model(&entities.Poll{}).Count(&count).Error; err != nil {
		log.Fatal(err.Error())
	}
	if count > 0 {
		return nil
	}

	polls := []entities.Poll{
		{
			Title:       "Favorite Programming Language",
			Description: "Vote for your favorite programming language.",
			ExpiresAt:   time.Now().AddDate(0, 0, 7),
			CreatorID:   1,
			Categories: []entities.PollCategory{
				{
					Name: "Go",
					Votes: []entities.Vote{
						{UserID: 1},
					},
				},
				{
					Name: "C#",
					Votes: []entities.Vote{
						{UserID: 2},
					},
				},
				{
					Name: "Python",
				},
			},
		},
		{
			Title:       "Best Movie of 2024",
			Description: "Vote for the best movie of 2024.",
			ExpiresAt:   time.Now().AddDate(0, 1, 0),
			CreatorID:   1,
			Categories: []entities.PollCategory{
				{
					Name: "Gladiator 2",
					Votes: []entities.Vote{
						{UserID: 1},
					},
				},
				{
					Name: "The Beekeeper",
					Votes: []entities.Vote{
						{UserID: 2},
					},
				},
			},
		},
		{
			Title:       "Favorite Food",
			Description: "Choose your favorite food.",
			ExpiresAt:   time.Now().AddDate(0, 1, 0),
			CreatorID:   2,
			Categories: []entities.PollCategory{
				{
					Name: "Pizza",
					Votes: []entities.Vote{
						{UserID: 1},
						{UserID: 2},
					},
				},
				{
					Name: "Pasta",
				},
				{
					Name: "Burger",
				},
			},
		},
	}

	if err := db.Create(&polls).Error; err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
