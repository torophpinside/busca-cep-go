package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	cepModel "busca-cep-go/buscaCep/cep/domain/model"

	"github.com/joho/godotenv"
)

var gormClient *gorm.DB

func init() {
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"),
	)
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		panic("mysql_not_connected")
	}
	gormClient = db

	db.Debug().AutoMigrate(&cepModel.Cep{})
}

func GetGorm() *gorm.DB {
	return gormClient
}
