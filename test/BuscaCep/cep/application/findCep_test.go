package application

import (
	"busca-cep-go/buscaCep/cep/application/findCep"
	"busca-cep-go/buscaCep/cep/domain/model"
	"errors"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
	"testing"
)

type CepRepositoryMock1 struct {
	mock.Mock
}

func (r *CepRepositoryMock1) GetCep(cep string) (*model.Cep, bool) {
	return new(model.Cep), false
}

func (r *CepRepositoryMock1) SaveCep(cep string) (*model.Cep, bool) {
	return new(model.Cep), false
}

func TestFindCepGetSuccess(t *testing.T) {
	assert := assert.New(t)

	cepRepositoryMockData := new(CepRepositoryMock1)
	cepRepositoryMockData.On("GetCep", "123456").Return(new(model.Cep), false)

	app := findCep.GetInstance(cepRepositoryMockData)
	cep, err := app.FindCep("123456")

	assert.Equal(new(model.Cep), cep, "equal")
	assert.Equal(nil, err, "no error")
}

type CepRepositoryMock2 struct {
	mock.Mock
}

func (r *CepRepositoryMock2) GetCep(cep string) (*model.Cep, bool) {
	return new(model.Cep), false
}

func (r *CepRepositoryMock2) SaveCep(cep string) (*model.Cep, bool) {
	return new(model.Cep), false
}

func TestFindCepSaveSuccess(t *testing.T) {
	assert := assert.New(t)

	cepRepositoryMockData := new(CepRepositoryMock2)
	cepRepositoryMockData.On("GetCep", "123456").Return(nil, true)
	cepRepositoryMockData.On("SaveCep", "123456").Return(new(model.Cep), false)

	app := findCep.GetInstance(cepRepositoryMockData)
	cep, err := app.FindCep("123456")

	assert.Equal(new(model.Cep), cep, "equal")
	assert.Equal(nil, err, "no error")
}

type CepRepositoryMock3 struct {
	mock.Mock
}

func (r *CepRepositoryMock3) GetCep(cep string) (*model.Cep, bool) {
	return nil, true
}

func (r *CepRepositoryMock3) SaveCep(cep string) (*model.Cep, bool) {
	return nil, true
}

func TestFindCepNotFound(t *testing.T) {
	assert := assert.New(t)

	cepRepositoryMockData := new(CepRepositoryMock3)
	cepRepositoryMockData.On("GetCep", "123456").Return(nil, true)
	cepRepositoryMockData.On("SaveCep", "123456").Return(nil, true)

	app := findCep.GetInstance(cepRepositoryMockData)
	cep, err := app.FindCep("123456")

	assert.NotEqual(new(model.Cep), cep, "cep not found")
	assert.Equal(errors.New("cep_not_found"), err, "message error")
}
