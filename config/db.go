package config

import (
	"FinalProjectGolangH8/comment"
	"FinalProjectGolangH8/photo"
	socialmedia "FinalProjectGolangH8/socialMedia"
	"FinalProjectGolangH8/user"
	"FinalProjectGolangH8/utils"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "production" {
		username := os.Getenv("DATABASE_USERNAME")
		password := os.Getenv("DATABASE_PASSWORD")
		host := os.Getenv("DATABASE_HOST")
		port := os.Getenv("DATABASE_PORT")
		database := os.Getenv("DATABASE_NAME")
		// production
		dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=require"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}

		// db.AutoMigrate(&models.User{}, &models.Order{}, &models.Category{}, &models.Product{}, &models.OrderDetail{}, &models.Confrimation{}, &models.Cart{}, &models.Review{})

		return db
	} else {
		dsn := "host=localhost user=postgres password=udang dbname=mygram port=5432 sslmode=disable TimeZone=Asia/Jakarta"

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err.Error())
		}

		db.AutoMigrate(&user.User{}, &socialmedia.SocialMedia{}, &photo.Photo{}, &comment.Comment{})

		return db
	}
}
