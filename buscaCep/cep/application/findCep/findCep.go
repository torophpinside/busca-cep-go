package findCep

import (
	"busca-cep-go/buscaCep/cep/domain/model"
	"errors"
)

type Application struct {
	CepRepository model.CepRepository
}

func GetInstance(repository model.CepRepository) *Application {
	return &Application{CepRepository: repository}
}

func (application *Application) FindCep(cepData string) (*model.Cep, error) {
	cep, err := application.CepRepository.GetCep(cepData)
	if cep == nil {
		cep, err = application.CepRepository.SaveCep(cepData)
		if err != false {
			return nil, errors.New("cep_not_found")
		}
		if cep == nil {
			return nil, errors.New("cep_not_found")
		}
	}

	return cep, nil
}
