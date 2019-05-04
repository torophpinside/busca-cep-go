package model

type CepRepository interface {
	GetCep(cepData string) (*Cep, bool)
	SaveCep(cepData string) (*Cep, bool)
}
