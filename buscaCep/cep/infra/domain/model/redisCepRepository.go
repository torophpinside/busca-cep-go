package model

import (
	model "busca-cep-go/buscaCep/cep/domain/model"
	"busca-cep-go/config"
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis"
)

type CepRepositoryImpl struct {
	Database *redis.Client
}

func GetInstance() model.CepRepository {
	redisClient := config.GetRedis()
	return &CepRepositoryImpl{Database: redisClient}
}

func (cepRepository *CepRepositoryImpl) GetCep(cepData string) *model.Cep {
	cep := &model.Cep{}

	cepKey := "cep_key_" + cepData
	val, err := cepRepository.Database.Get(cepKey).Result()
	if err == nil && val != "" {
		err = json.Unmarshal([]byte(val), cep)
		return cep
	}

	return nil
}

func (cepRepository *CepRepositoryImpl) SaveCep(cepData string) (*model.Cep, bool) {
	cep := &model.Cep{}

	cepKey := "cep_key_" + cepData
	resp, err := http.Get("https://viacep.com.br/ws/" + cepData + "/json/")
	if err != nil {
		return nil, false
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(cep)
	if err != nil {
		return nil, false
	}
	out, err := json.Marshal(cep)
	if err != nil {
		return nil, false
	}

	err = config.GetRedis().Set(cepKey, string(out), 0).Err()
	if err != nil {
		return nil, false
	}

	return cep, true
}
