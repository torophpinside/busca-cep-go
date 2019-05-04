package model

import (
	model "busca-cep-go/buscaCep/cep/domain/model"
	service "busca-cep-go/buscaCep/core/infra/service"
	"busca-cep-go/config"
	"encoding/json"
	"strings"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type GormCepRepository struct {
	Database *gorm.DB
	Cache    *redis.Client
}

func GetGormInstance() model.CepRepository {
	gormClient := config.GetGorm()
	redisClient := config.GetRedis()
	return &GormCepRepository{Database: gormClient, Cache: redisClient}
}

func (cepRepository *GormCepRepository) GetCep(cepData string) (*model.Cep, bool) {
	cep := &model.Cep{}

	err := cepRepository.Database.Where("cep LIKE ?", cepData).First(&cep).Error
	if err != nil {
		return nil, true
	}

	return cep, false
}

func (cepRepository *GormCepRepository) SaveCep(cepData string) (*model.Cep, bool) {
	cep := &model.Cep{}

	cepKey := "cep_key_" + cepData
	val, err := cepRepository.Cache.Get(cepKey).Result()
	if err == redis.Nil && val == "" {
		cep, _ = callWs("https://viacep.com.br/ws/"+cepData+"/json/", cepKey, cepRepository)
	} else if err == nil && val == "" {
		cep, _ = callWs("https://viacep.com.br/ws/"+cepData+"/json/", cepKey, cepRepository)
	} else {
		if val == "error" {
			return nil, true
		}
	}

	return cep, false
}

func callWs(url string, key string, cepRepository *GormCepRepository) (*model.Cep, bool) {
	var cep model.Cep = &model.Cep{}

	resp := service.Call(url, cep)
	if resp == nil {
		cepRepository.Cache.Set(key, "error", 0).Err()
		return nil, true
	} else if cep == nil {
		cepRepository.Cache.Set(key, "error", 0).Err()
		return nil, true
	}

	err := json.NewDecoder(resp.Body).Decode(cep)
	if err != nil {
		return nil, true
	}
	cep.Cep = strings.Replace(cep.Cep, "-", "", -1)
	cep.DeletedAt = nil

	err = cepRepository.Database.Save(&cep).Error
	if err != nil {
		return nil, true
	}

	out, err := json.Marshal(cep)
	if err != nil {
		return nil, true
	}

	err = cepRepository.Cache.Set(key, string(out), 0).Err()
	if err != nil {
		return nil, true
	}

	return cep, false
}
