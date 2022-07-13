package sections

import (
	"testing"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/stretchr/testify/assert"
)

func Test_Create_Ok(t *testing.T) {
	expectedResult := db.Section{
		Id:                 4,
		Number:             4,
		CurrentTemperature: 99.5,
		MinimumTemperature: 9.0,
		CurrentCapacity:    900,
		MinimumCapacity:    90,
		MaximumCapacity:    900,
		WarehouseId:        2,
		ProductTypeId:      2,
	}

	mockRepository := MockSectionRepository{
		Result:              expectedResult,
		err:                 nil,
		existsSectionNumber: false,
	}

	service := NewService(mockRepository)
	result, err := service.Create(4, 99.5, 9.0, 900, 90, 900, 2, 2)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Create_Conflict(t *testing.T) {

	expectedResult := ErrExistsSectionNumberError

	mockRepository := MockSectionRepository{
		Result:              db.Section{},
		err:                 expectedResult,
		existsSectionNumber: true,
	}

	service := NewService(mockRepository)
	_, err := service.Create(1, 99.5, 9.0, 900, 90, 900, 2, 2)

	assert.Equal(t, expectedResult, err)
}

func Test_GetAll_FindAll(t *testing.T) {

	expectedResult := []db.Section{{}, {}, {}}

	mockRepository := MockSectionRepository{
		Result: expectedResult,
		err:    nil,
	}

	service := NewService(mockRepository)
	result, err := service.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Get_FindByIdExistent(t *testing.T) {

	expectedResult := db.Section{}

	mockRepository := MockSectionRepository{
		Result: expectedResult,
		err:    nil,
	}

	service := NewService(mockRepository)
	result, err := service.Get(1)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Get_FindByIdNonExistent(t *testing.T) {

	expectedError := ErrSectionNotFoundError

	mockRepository := MockSectionRepository{
		err: expectedError,
	}

	service := NewService(mockRepository)
	_, err := service.Get(4)

	assert.Equal(t, expectedError, err)
}

func Test_Update_Ok(t *testing.T) {

	getById := db.Section{
		Id:                 4,
		Number:             4,
		CurrentTemperature: 99.5,
		MinimumTemperature: 9.0,
		CurrentCapacity:    900,
		MinimumCapacity:    90,
		MaximumCapacity:    900,
		WarehouseId:        2,
		ProductTypeId:      2,
	}
	expectedResult := db.Section{
		Id:                 4,
		Number:             4,
		CurrentTemperature: 99.5,
		MinimumTemperature: 9.0,
		CurrentCapacity:    900,
		MinimumCapacity:    90,
		MaximumCapacity:    900,
		WarehouseId:        2,
		ProductTypeId:      2,
	}

	mockRepository := MockSectionRepository{
		Result:              expectedResult,
		GetById:             getById,
		existsSectionNumber: false,
	}

	service := NewService(mockRepository)
	result, _ := service.Update(4, 4, 99.5, 9.0, 900, 90, 900)

	assert.Equal(t, expectedResult, result)
}

func Test_Update_NonExistent(t *testing.T) {

	expectedError := ErrSectionNotFoundError

	mockRepository := MockSectionRepository{
		err:                 expectedError,
		GetById:             db.Section{},
		existsSectionNumber: false,
	}

	service := NewService(mockRepository)
	_, err := service.Update(4, 4, 99.5, 9.0, 900, 90, 900)

	assert.Equal(t, expectedError, err)
}

func Test_Update_ShouldReturnErrWhenCodeAlreadyExists(t *testing.T) {

	expectedError := ErrExistsSectionNumberError

	getById := db.Section{
		Id:                 4,
		Number:             4,
		CurrentTemperature: 99.5,
		MinimumTemperature: 9.0,
		CurrentCapacity:    900,
		MinimumCapacity:    90,
		MaximumCapacity:    900,
		WarehouseId:        2,
		ProductTypeId:      2,
	}

	mockRepository := MockSectionRepository{
		err:                 expectedError,
		GetById:             getById,
		existsSectionNumber: true,
	}

	service := NewService(mockRepository)
	_, err := service.Update(4, 4, 99.5, 9.0, 900, 90, 900)

	assert.Equal(t, expectedError, err)
}

func Test_Delete_Ok(t *testing.T) {

	mockRepository := MockSectionRepository{
		Result: db.Section{},
		err:    nil,
	}

	service := NewService(mockRepository)
	err := service.Delete(1)

	assert.Equal(t, nil, err)
}

func Test_Delete_NonExistent(t *testing.T) {

	expectedError := ErrSectionNotFoundError

	mockRepository := MockSectionRepository{
		err: expectedError,
	}

	service := NewService(mockRepository)
	err := service.Delete(5)

	assert.Equal(t, expectedError, err)
}
