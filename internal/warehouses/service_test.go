package warehouses

import (
	"errors"
	"testing"

	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/stretchr/testify/assert"
)

func Test_Create_ShouldReturnOK(t *testing.T) {

	expectedResult := database.Warehouse{
		Id:                 8,
		Code:               "SC",
		Address:            "psn",
		Telephone:          "45674458",
		MinimunCapacity:    5,
		MinimumTemperature: 1.1,
		LocalityID:         "1",
	}
	mockRepository := mockWarehouseRepository{
		result:   expectedResult,
		err:      nil,
		findByCode: false,
	}

	service := NewService(mockRepository)
	result, err := service.Create("SC", "psn", "45674458", 5, 1.1, "1")

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Create_ShouldReturnError(t *testing.T) {
	expectedError := ExistsWarehouseCodeError

	mockRepository := mockWarehouseRepository{
		result:   database.Warehouse{},
		err:      expectedError,
		findByCode: true,
	}
	service := NewService(mockRepository)
	_, err := service.Create("SC", "psn", "45674458", 5, 1.1, "1")

	assert.Equal(t, expectedError, err)
}

func Test_GetAll_ShouldReturnWarehouseList(t *testing.T) {

	expectedResult := []database.Warehouse{{}, {}, {}}

	mockRepository := mockWarehouseRepository{
		result: expectedResult,
		err:    nil,
	}

	service := NewService(mockRepository)
	result, err := service.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_GetAll_ShouldReturnError(t *testing.T) {

	expectedResult := []database.Warehouse{}
	expectedError := errors.New("The bank went bad!")

	mockRepository := mockWarehouseRepository{
		result: expectedResult,
		err:    expectedError,
	}

	service := NewService(mockRepository)
	result, err := service.GetAll()

	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedError, err)
}

func Test_Delete_ShouldReturnOk(t *testing.T) {

	mockRepository := mockWarehouseRepository{
		result: database.Warehouse{},
		err:    nil,
	}

	service := NewService(mockRepository)
	err := service.Delete(1)

	assert.Equal(t, nil, err)
}

func Test_Delete_ShouldReturnError(t *testing.T) {

	expectedError := errors.New("The bank went bad!")

	mockRepository := mockWarehouseRepository{
		result: database.Warehouse{},
		err:    expectedError,
	}

	service := NewService(mockRepository)
	err := service.Delete(1)

	assert.Equal(t, expectedError, err)
}

func Test_Get_ShouldReturnOK(t *testing.T) {

	expectedResult := database.Warehouse{}

	mockWarehouseRepository := mockWarehouseRepository{
		result: expectedResult,
		err:    nil,
	}

	service := NewService(mockWarehouseRepository)
	result, err := service.Get(1)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Get_ShouldReturnError(t *testing.T) {

	expectedError := WarehouseNotFoundError

	mockWarehouseRepository := mockWarehouseRepository{
		err: expectedError,
	}

	service := NewService(mockWarehouseRepository)
	_, err := service.Get(1)

	assert.Equal(t, expectedError, err)

}

func Test_Update_ShouldReturnOK(t *testing.T) {

	getById := database.Warehouse{
		Id:                 8,
		Code:               "SC",
		Address:            "psn",
		Telephone:          "45674458",
		MinimunCapacity:    5,
		MinimumTemperature: 1.1,
	}

	expectedResult := database.Warehouse{
		Id:                 8,
		Code:               "SCgf",
		Address:            "psnss",
		Telephone:          "45674458785",
		MinimunCapacity:    5,
		MinimumTemperature: 1.2,
	}

	mockRepository := mockWarehouseRepository{
		result:   expectedResult,
		err:      nil,
		findByCode: false,
		getById:  getById,
	}

	service := NewService(mockRepository)
	result, _ := service.Update(8, "SCgf", "psnss", "45674458785", 5, 1.2)

	assert.Equal(t, expectedResult, result)
}

func Test_Update_ShouldReturnError(t *testing.T) {

	expectedError := WarehouseNotFoundError

	mockWarehouseRepository := mockWarehouseRepository{
		err:      expectedError,
		getById:  database.Warehouse{},
		findByCode: false,
	}

	service := NewService(mockWarehouseRepository)
	_, err := service.Update(8, "SCgf", "psnss", "45674458785", 5, 1.2)

	assert.Equal(t, expectedError, err)

}

func Test_Update_ShouldReturnErrorExists(t *testing.T) {

	expectedError := ExistsWarehouseCodeError

	getById := database.Warehouse{
		Id:                 8,
		Code:               "SC",
		Address:            "psn",
		Telephone:          "45674458",
		MinimunCapacity:    5,
		MinimumTemperature: 1.1,
	}

	mockWarehouseRepository := mockWarehouseRepository{
		err:      expectedError,
		getById:  getById,
		findByCode: true,
	}

	service := NewService(mockWarehouseRepository)
	_, err := service.Update(8, "SCgf", "psnss", "45674458785", 5, 1.2)

	assert.Equal(t, expectedError, err)

}
