// @/config/database.go
package config

import (
	"fmt"
	"os"

	"github.com/hazem-jday/het-ejdid-back/entities"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Go ORM Database
var Database *gorm.DB

//Connexion à la base
func Connect() error {
	var err error
	godotenv.Load(".env")
	DATABASE_URI := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("PORT"), os.Getenv("DATABASE"))
	print(DATABASE_URI)

	Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	//Initialisation des tableaux
	Database.AutoMigrate(&entities.User{})
	Database.AutoMigrate(&entities.Article{})
	Database.AutoMigrate(&entities.Like{})
	Database.AutoMigrate(&entities.Dislike{})

	return nil
}
