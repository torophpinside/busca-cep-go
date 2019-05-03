package model

type CepRepository interface {
	GetCep(cepData string) *Cep
	SaveCep(cepData string) (*Cep, bool)
}
