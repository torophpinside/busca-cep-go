package model

import (
	model "busca-cep-go/buscaCep/cep/domain/model"
	"busca-cep-go/buscaCep/cep/domain/repository"
	"busca-cep-go/config"
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis"
)

type CepRepositoryImpl struct {
	Database *redis.Client
}

func GetRedisInstance() repository.CepRepository {
	redisClient := config.GetRedis()
	return &CepRepositoryImpl{Database: redisClient}
}

func (cepRepository *CepRepositoryImpl) GetCep(cepData string) (*model.Cep, bool) {
	cep := &model.Cep{}

	cepKey := "cep_key_" + cepData
	val, err := cepRepository.Database.Get(cepKey).Result()
	if err == nil && val != "" {
		err = json.Unmarshal([]byte(val), cep)
		return nil, true
	}

	return cep, false
}

func (cepRepository *CepRepositoryImpl) SaveCep(cepData string) (*model.Cep, bool) {
	cep := &model.Cep{}

	cepKey := "cep_key_" + cepData
	resp, err := http.Get("https://viacep.com.br/ws/" + cepData + "/json/")
	if err != nil {
		return nil, true
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(cep)
	if err != nil {
		return nil, true
	}
	out, err := json.Marshal(cep)
	if err != nil {
		return nil, true
	}

	err = config.GetRedis().Set(cepKey, string(out), 0).Err()
	if err != nil {
		return nil, true
	}

	return cep, false
}
