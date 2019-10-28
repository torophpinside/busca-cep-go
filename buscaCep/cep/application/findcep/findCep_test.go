package findcep

import (
	"busca-cep-go/buscaCep/cep/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type CepRepositoryMock struct {
	mock.Mock
}

func (m *CepRepositoryMock) GetCep(cepData string) (*model.Cep, bool) {
	args := m.Called(cepData)
	if args.Get(0) != nil {
		return &model.Cep{}, args.Bool(1)
	}
	return nil, args.Bool(1)
}

func (m *CepRepositoryMock) SaveCep(cepData string) (*model.Cep, bool) {
	args := m.Called(cepData)
	if args.Get(0) != nil {
		return &model.Cep{}, args.Bool(1)
	}
	return nil, args.Bool(1)
}

func TestApplication_FindCep(t *testing.T) {
	cr := new(CepRepositoryMock)
	s := "12345678"
	cr.On("GetCep", s).Return(&model.Cep{}, false)

	app := GetInstance(cr)
	cep, err := app.FindCep(s)
	if err != nil {
		t.Error("error getting cep")
	}
	cr.AssertExpectations(t)
	assert.Equal(t, &model.Cep{}, cep)
}

func TestApplication_FindCepNotFoundSave(t *testing.T) {
	cr := new(CepRepositoryMock)
	s := "12345678"
	cr.On("GetCep", s).Return(nil, true)
	cr.On("SaveCep", s).Return(&model.Cep{}, false)

	app := GetInstance(cr)
	cep, err := app.FindCep(s)
	if err != nil {
		t.Error("error getting cep")
	}
	cr.AssertExpectations(t)
	assert.Equal(t, &model.Cep{}, cep)
}

func TestApplication_FindCepNotFoundNoSaveError(t *testing.T) {
	cr := new(CepRepositoryMock)
	s := "12345678"
	cr.On("GetCep", s).Return(nil, true)
	cr.On("SaveCep", s).Return(nil, true)

	app := GetInstance(cr)
	_, err := app.FindCep(s)

	cr.AssertExpectations(t)
	assert.Error(t, err, "error_save_cep")
}

func TestApplication_FindCepNotFoundNoSaveCepNotFounr(t *testing.T) {
	cr := new(CepRepositoryMock)
	s := "12345678"
	cr.On("GetCep", s).Return(nil, true)
	cr.On("SaveCep", s).Return(nil, false)

	app := GetInstance(cr)
	_, err := app.FindCep(s)

	cr.AssertExpectations(t)
	assert.Error(t, err, "cep_not_found")
}
