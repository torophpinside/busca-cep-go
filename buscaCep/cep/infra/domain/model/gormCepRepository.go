package model

import (
	model "busca-cep-go/buscaCep/cep/domain/model"
	"busca-cep-go/config"
	"encoding/json"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type GormCepRepository struct {
	Database *gorm.DB
}

func GetGormInstance() model.CepRepository {
	gormClient := config.GetGorm()
	return &GormCepRepository{Database: gormClient}
}

func (cepRepository *GormCepRepository) GetCep(cepData string) *model.Cep {
	cep := &model.Cep{}

	err := cepRepository.Database.Where("cep LIKE ?", cepData).First(&cep).Error
	if err == nil {
		return cep
	}

	return nil
}

func (cepRepository *GormCepRepository) SaveCep(cepData string) (*model.Cep, bool) {
	cep := &model.Cep{}

	resp, err := http.Get("https://viacep.com.br/ws/" + cepData + "/json/")
	if err != nil {
		return nil, false
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(cep)
	if err != nil {
		return nil, false
	}
	cep.Cep = strings.Replace(cep.Cep, "-", "", -1)
	cep.DeletedAt = nil
	err = cepRepository.Database.Save(&cep).Error
	if err != nil {
		return nil, false
	}

	return cep, true
}
