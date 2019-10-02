package config

import (
	"busca-cep-go/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"

	cepModel "busca-cep-go/buscaCep/cep/domain/model"

	"time"
)

var gormClient *gorm.DB

func init() {
	err := utils.LoadEnv()
	if err != nil {
		log.Fatalf("error_loading_env:%s", err.Error())
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

	db.Debug().AutoMigrate(&cepModel.Cep{})
	db.DB().SetMaxOpenConns(700)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(1 * time.Second)
	gormClient = db
}

func GetGorm() *gorm.DB {
	return gormClient
}
