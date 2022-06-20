package products

import (
	"errors"
	"testing"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/stretchr/testify/assert"
)

/*
find_all Se a lista tiver "n" elementos, retornará uma quantidade do total de elementos

find_by_id_non_existent Se o elemento procurado por id não existir, retorna null

find_by_id_existent Se o elemento procurado por id existir, ele retornará as informações do elemento solicitado

update_existent Quando a atualização dos dados for bem sucedida, o produto será devolvido com as informações
atualizadas

update_non_existent Se o produto a ser atualizado não existir, será retornado null.

delete_non_existent Quando o produto não existe, null será retornado.

delete_ok Se a exclusão for bem-sucedida, o item não aparecerá na lista.
*/

/*
service
func XYZ()
a
b -> repository
c
d
e -> repository
f
g

func XYZ()
a
b -> service
c
d
e -> service
f
g


result, expectedResult

mockRepository


*/

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

	mockRepository := mockProductRepository{
		result:            expectedResult,
		err:               nil,
		existsProductCode: false,
	}

	service := NewProductService(mockRepository)
	result, err := service.Create("STN", "Disco da Xuxa", 100, 100, 200, 50, 0, 1100, 0, 1, 666)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Create_Conlict(t *testing.T) {

	expectedError := ExistsProductCodeError

	mockRepository := mockProductRepository{
		result:            db.Product{},
		err:               expectedError,
		existsProductCode: true,
	}

	service := NewProductService(mockRepository)
	_, err := service.Create("STN", "Disco da Xuxa", 100, 100, 200, 50, 0, 1100, 0, 1, 666)

	assert.Equal(t, expectedError, err)
}

func Test_Get_OK(t *testing.T) {

	expectedResult := db.Product{}

	mockProductRepository := mockProductRepository{
		result: expectedResult,
		err:    nil,
	}

	service := NewProductService(mockProductRepository)
	result, err := service.Get(1)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Get_Error(t *testing.T) {

	expectedError := ProductNotFoundError

	mockProductRepository := mockProductRepository{
		err: expectedError,
	}

	service := NewProductService(mockProductRepository)
	_, err := service.Get(1)

	assert.Equal(t, expectedError, err)
}

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
		err:    expectedError,
	}

	service := NewProductService(mockRepository)
	result, err := service.GetAll()

	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedError, err)
}

func Test_Delete_ShouldReturnOk(t *testing.T) {

	mockRepository := mockProductRepository{
		result: db.Product{},
		err:    nil,
	}

	service := NewProductService(mockRepository)
	err := service.Delete(1)

	assert.Equal(t, nil, err)
}

func Test_Update_ShouldReturnOK(t *testing.T) {

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

	mockProductRepository := mockProductRepository{
		result:            expectedResult,
		getById:           getById,
		existsProductCode: false,
	}

	service := NewProductService(mockProductRepository)
	result, _ := service.Update(13, "abc", "abc", 666, 666, 666, 666, 666, 666, 666)

	assert.Equal(t, expectedResult, result)
}

func Test_Update_ShouldReturnErrWhenIdNotExists(t *testing.T) {

	expectedError := ProductNotFoundError

	mockProductRepository := mockProductRepository{
		err:               expectedError,
		getById:           db.Product{},
		existsProductCode: false,
	}

	service := NewProductService(mockProductRepository)
	_, err := service.Update(13, "abc", "abc", 666, 666, 666, 666, 666, 666, 666)

	assert.Equal(t, expectedError, err)
}

func Test_Update_ShouldReturnErrWhenCodeAlreadyExists(t *testing.T) {

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

	mockProductRepository := mockProductRepository{
		err:               expectedError,
		getById:           getById,
		existsProductCode: true,
	}

	service := NewProductService(mockProductRepository)
	_, err := service.Update(13, "abc", "abc", 666, 666, 666, 666, 666, 666, 666)

	assert.Equal(t, expectedError, err)
}
