package findcep

import (
	"busca-cep-go/buscaCep/cep/domain/model"
	"busca-cep-go/buscaCep/cep/domain/repository"
	"errors"
)

type Application struct {
	CepRepository repository.CepRepository
}

func GetInstance(repository repository.CepRepository) *Application {
	return &Application{CepRepository: repository}
}

func (application *Application) FindCep(cepData string) (*model.Cep, error) {
	cep, err := application.CepRepository.GetCep(cepData)
	if cep == nil {
		cep, err = application.CepRepository.SaveCep(cepData)
		if err != false {
			return nil, errors.New("error_save_cep")
		}
		if cep == nil {
			return nil, errors.New("cep_not_found")
		}
	}

	return cep, nil
}
