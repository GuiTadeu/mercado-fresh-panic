package products

import (
	"errors"
	"testing"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/stretchr/testify/assert"
)

func Test_GetAll_ShouldReturnProductList(t *testing.T) {

	expectedResult := []db.Product{{}, {}, {}}

	mockRepository := mockProductRepository{
		result: expectedResult,
		err:    nil,
	}

	service := NewProductService(mockRepository)
	result, err := service.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_GetAll_ShouldReturnError(t *testing.T) {

	expectedResult := []db.Product{}
	expectedError := errors.New("Deu ruim no banco!")

	mockRepository := mockProductRepository{
		result: expectedResult,
		err:  expectedError,
	}

	service := NewProductService(mockRepository)
	result, err := service.GetAll()

	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedError, err)
}

func Test_Delete_ShouldReturnOk(t *testing.T) {

	mockRepository := mockProductRepository{
		result: db.Product{},
		err:  nil,
	}

	service := NewProductService(mockRepository)
	err := service.Delete(1)

	assert.Equal(t, nil, err)
}

func Test_Delete_ShouldReturnError(t *testing.T) {

	expectedError := errors.New("Deu ruim no banco!")

	mockRepository := mockProductRepository{
		result: db.Product{},
		err:  expectedError,
	}

	service := NewProductService(mockRepository)
	err := service.Delete(1)

	assert.Equal(t, nil, err)
}