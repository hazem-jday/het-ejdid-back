// @/config/database.go
package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hazem-jday/het-ejdid-back/entities"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Go ORM Database
var Database *gorm.DB

//Connexion à la base
func Connect() error {
	var err error
	godotenv.Load(".env")
	DATABASE_URI := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("PORT"), os.Getenv("DATABASE"))
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 newLogger,
	})

	if err != nil {
		panic(err)
	}

	//Initialisation des tableaux
	Database.AutoMigrate(&entities.User{})
	Database.AutoMigrate(&entities.Article{})
	Database.AutoMigrate(&entities.Like{})
	Database.AutoMigrate(&entities.Dislike{})
	Database.AutoMigrate(&entities.Save{})
	Database.AutoMigrate(&entities.Meteo{})
	return nil
}
