package repository

import "busca-cep-go/buscaCep/cep/domain/model"

type CepRepository interface {
	GetCep(cepData string) (*model.Cep, bool)
	SaveCep(cepData string) (*model.Cep, bool)
}
