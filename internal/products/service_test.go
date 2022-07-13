package products

import (
	"errors"
	"testing"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/stretchr/testify/assert"
)

func Test_Create_Ok(t *testing.T) {

	expectedResult := db.Product{
		Id:                      13,
		Code:                    "STN",
		Description:             "Disco da Xuxa",
		Width:                   100,
		Height:                  100,
		Length:                  200,
		NetWeight:               50.0,
		ExpirationRate:          0,
		RecommendedFreezingTemp: 1100,
		FreezingRate:            0,
		ProductTypeId:           1,
		SellerId:                666,
	}

	mockRepository := MockProductRepository{
		Result:            expectedResult,
		err:               nil,
		existsProductCode: false,
	}

	service := NewProductService(mockRepository)
	result, err := service.Create("STN", "Disco da Xuxa", 100, 100, 200, 50, 0, 1100, 0, 1, 666)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Create_ShouldReturnErrorWhenCodeAlreadyExists(t *testing.T) {

	expectedError := ExistsProductCodeError

	mockRepository := MockProductRepository{
		Result:            db.Product{},
		err:               expectedError,
		existsProductCode: true,
	}

	service := NewProductService(mockRepository)
	_, err := service.Create("STN", "Disco da Xuxa", 100, 100, 200, 50, 0, 1100, 0, 1, 666)

	assert.Equal(t, expectedError, err)
}

func Test_Get_OK(t *testing.T) {

	expectedResult := db.Product{}

	mockProductRepository := MockProductRepository{
		GetById: expectedResult,
		err:     nil,
	}

	service := NewProductService(mockProductRepository)
	result, err := service.Get(1)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Get_ShouldReturnErrorWhenIdNotExists(t *testing.T) {

	expectedError := ProductNotFoundError

	mockProductRepository := MockProductRepository{
		err: expectedError,
	}

	service := NewProductService(mockProductRepository)
	_, err := service.Get(1)

	assert.Equal(t, expectedError, err)
}

func Test_GetAll_OK(t *testing.T) {

	expectedResult := []db.Product{{}, {}, {}}

	mockRepository := MockProductRepository{
		Result: expectedResult,
		err:    nil,
	}

	service := NewProductService(mockRepository)
	result, err := service.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_GetAll_ShouldReturnErrorWhenDatabaseFails(t *testing.T) {

	expectedResult := []db.Product{}
	expectedError := errors.New("Deu ruim no banco!")

	mockRepository := MockProductRepository{
		Result: expectedResult,
		err:    expectedError,
	}

	service := NewProductService(mockRepository)
	result, err := service.GetAll()

	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedError, err)
}

func Test_Update_OK(t *testing.T) {

	getById := db.Product{
		Id:                      13,
		Code:                    "kkk",
		Description:             "kkk",
		Width:                   321,
		Height:                  321,
		Length:                  321,
		NetWeight:               321,
		ExpirationRate:          321,
		RecommendedFreezingTemp: 321,
		FreezingRate:            321,
		ProductTypeId:           666,
		SellerId:                666,
	}

	expectedResult := db.Product{
		Id:                      13,
		Code:                    "abc",
		Description:             "abc",
		Width:                   666,
		Height:                  666,
		Length:                  666,
		NetWeight:               666,
		ExpirationRate:          666,
		RecommendedFreezingTemp: 666,
		FreezingRate:            666,
		ProductTypeId:           666,
		SellerId:                666,
	}

	mockProductRepository := MockProductRepository{
		Result:            expectedResult,
		GetById:           getById,
		existsProductCode: false,
	}

	service := NewProductService(mockProductRepository)
	result, _ := service.Update(13, "abc", "abc", 666, 666, 666, 666, 666, 666, 666)

	assert.Equal(t, expectedResult, result)
}

func Test_Update_ShouldReturnErrorWhenIdNotExists(t *testing.T) {

	expectedError := ProductNotFoundError

	mockProductRepository := MockProductRepository{
		err:               expectedError,
		GetById:           db.Product{},
		existsProductCode: false,
	}

	service := NewProductService(mockProductRepository)
	_, err := service.Update(13, "abc", "abc", 666, 666, 666, 666, 666, 666, 666)

	assert.Equal(t, expectedError, err)
}

func Test_Update_ShouldReturnErrorWhenCodeAlreadyExists(t *testing.T) {

	expectedError := ExistsProductCodeError

	getById := db.Product{
		Id:                      13,
		Code:                    "abc",
		Description:             "abc",
		Width:                   666,
		Height:                  666,
		Length:                  666,
		NetWeight:               666,
		ExpirationRate:          666,
		RecommendedFreezingTemp: 666,
		FreezingRate:            666,
		ProductTypeId:           666,
		SellerId:                666,
	}

	mockProductRepository := MockProductRepository{
		err:               expectedError,
		GetById:           getById,
		existsProductCode: true,
	}

	service := NewProductService(mockProductRepository)
	_, err := service.Update(13, "abc", "abc", 666, 666, 666, 666, 666, 666, 666)

	assert.Equal(t, expectedError, err)
}

func Test_Delete_Ok(t *testing.T) {

	mockRepository := MockProductRepository{
		Result: db.Product{},
		err:    nil,
	}

	service := NewProductService(mockRepository)
	err := service.Delete(1)

	assert.Equal(t, nil, err)
}

func Test_Delete_ShouldReturnErrorWhenIdNotExists(t *testing.T) {

	expectedError := ProductNotFoundError

	mockRepository := MockProductRepository{
		Result: db.Product{},
		err:    expectedError,
	}

	service := NewProductService(mockRepository)
	err := service.Delete(1)

	assert.Equal(t, expectedError, err)
}
