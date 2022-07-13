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
	mockRepository := MockWarehouseRepository{
		Result:     expectedResult,
		Err:        nil,
		FindByCode: false,
	}

	service := NewService(mockRepository)
	result, err := service.Create("SC", "psn", "45674458", 5, 1.1, "1")

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Create_ShouldReturnError(t *testing.T) {
	expectedError := ExistsWarehouseCodeError

	mockRepository := MockWarehouseRepository{
		Result:     database.Warehouse{},
		Err:        expectedError,
		FindByCode: true,
	}
	service := NewService(mockRepository)
	_, err := service.Create("SC", "psn", "45674458", 5, 1.1, "1")

	assert.Equal(t, expectedError, err)
}

func Test_GetAll_ShouldReturnWarehouseList(t *testing.T) {

	expectedResult := []database.Warehouse{{}, {}, {}}

	mockRepository := MockWarehouseRepository{
		Result: expectedResult,
		Err:    nil,
	}

	service := NewService(mockRepository)
	result, err := service.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_GetAll_ShouldReturnError(t *testing.T) {

	expectedResult := []database.Warehouse{}
	expectedError := errors.New("The bank went bad!")

	mockRepository := MockWarehouseRepository{
		Result: expectedResult,
		Err:    expectedError,
	}

	service := NewService(mockRepository)
	result, err := service.GetAll()

	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedError, err)
}

func Test_Delete_ShouldReturnOk(t *testing.T) {

	mockRepository := MockWarehouseRepository{
		Result: database.Warehouse{},
		Err:    nil,
	}

	service := NewService(mockRepository)
	err := service.Delete(1)

	assert.Equal(t, nil, err)
}

func Test_Delete_ShouldReturnError(t *testing.T) {

	expectedError := errors.New("The bank went bad!")

	mockRepository := MockWarehouseRepository{
		Result: database.Warehouse{},
		Err:    expectedError,
	}

	service := NewService(mockRepository)
	err := service.Delete(1)

	assert.Equal(t, expectedError, err)
}

func Test_Get_ShouldReturnOK(t *testing.T) {

	expectedResult := database.Warehouse{}

	mockWarehouseRepository := MockWarehouseRepository{
		Result: expectedResult,
		Err:    nil,
	}

	service := NewService(mockWarehouseRepository)
	result, err := service.Get(1)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Get_ShouldReturnError(t *testing.T) {

	expectedError := WarehouseNotFoundError

	mockWarehouseRepository := MockWarehouseRepository{
		Err: expectedError,
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

	mockRepository := MockWarehouseRepository{
		Result:     expectedResult,
		Err:        nil,
		FindByCode: false,
		GetById:    getById,
	}

	service := NewService(mockRepository)
	result, _ := service.Update(8, "SCgf", "psnss", "45674458785", 5, 1.2)

	assert.Equal(t, expectedResult, result)
}

func Test_Update_ShouldReturnError(t *testing.T) {

	expectedError := WarehouseNotFoundError

	mockWarehouseRepository := MockWarehouseRepository{
		Err:        expectedError,
		GetById:    database.Warehouse{},
		FindByCode: false,
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

	mockWarehouseRepository := MockWarehouseRepository{
		Err:        expectedError,
		GetById:    getById,
		FindByCode: true,
	}

	service := NewService(mockWarehouseRepository)
	_, err := service.Update(8, "SCgf", "psnss", "45674458785", 5, 1.2)

	assert.Equal(t, expectedError, err)

}
