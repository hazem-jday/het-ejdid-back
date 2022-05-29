package config

import (
	"fmt"
	"het-ejdid-back/entities"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Go ORM Database
var Database *gorm.DB

//Connexion à la base
func Connect() error {
	var err error
	fmt.Printf("%s %s %s %s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	DATABASE_URI := fmt.Sprintf("%s:%s@tcp(db:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
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
