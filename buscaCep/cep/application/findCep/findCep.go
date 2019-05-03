package findCep

import (
	"busca-cep-go/buscaCep/cep/domain/model"
)

type Application struct {
	CepRepository model.CepRepository
}

func GetInstance(repository model.CepRepository) *Application {
	return &Application{CepRepository: repository}
}

func (application *Application) FindCep(cepData string) *model.Cep {
	cep := application.CepRepository.GetCep(cepData)
	if cep == nil {
		cep, _ = application.CepRepository.SaveCep(cepData)
	}

	return cep
}
