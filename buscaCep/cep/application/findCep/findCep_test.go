package findCep

import (
	"busca-cep-go/buscaCep/cep/infra/domain/model"
	"testing"
)

var application = GetInstance(model.GetGormInstance())

func TestFindCepEmpty(t *testing.T) {
	cep, err := application.FindCep("")
	if err == nil {
		t.Errorf("err should not be nil")
	}

	if cep != nil {
		t.Errorf("cep should be nil, %v returned", cep)
	}
}

func TestFindCepNotExist(t *testing.T) {
	cep, err := application.FindCep("123456")
	if err == nil {
		t.Errorf("err should not be nil")
	}

	if cep != nil {
		t.Errorf("cep should be nil, %v returned", cep)
	}
}

func TestFindCepSuccess(t *testing.T) {
	cep, err := application.FindCep("88036100")
	if err != nil {
		t.Errorf("err should be nil")
	}

	if cep == nil {
		t.Errorf("cep should not be nil")
	}

	if cep.UF != "SC" {
		t.Errorf("cep UF should be %v, returned %v", "SC", cep.UF)
	}
}
