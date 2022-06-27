package employees

import (
	"errors"
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Create_Ok(t *testing.T) {

	expectedResult := db.Employee{
		Id:           1,
		FirstName:    "testMock",
		LastName:     "lastNameMock",
		CardNumberId: "1000",
		WarehouseId:  1,
	}

	mockRepository := mockEmployeeRepository{
		result:             expectedResult,
		err:                nil,
		existsEmployeeCode: false,
	}

	service := NewEmployeeService(mockRepository)
	result, err := service.Create("1000", "testMock", "lastNameMock", 1)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Create_ShouldReturnErrorWhenCodeAlreadyExists(t *testing.T) {

	expectedError := ExistsCardNumberIdError

	mockRepository := mockEmployeeRepository{
		result:             db.Employee{},
		err:                expectedError,
		existsEmployeeCode: true,
	}

	service := NewEmployeeService(mockRepository)
	_, err := service.Create("1000", "testMock", "lastNameMock", 1)

	assert.Equal(t, expectedError, err)
}

func Test_Get_OK(t *testing.T) {
	expectedResult := db.Employee{}

	mockRepository := mockEmployeeRepository{
		getById: expectedResult,
		err:     nil,
	}

	service := NewEmployeeService(mockRepository)
	result, err := service.Get(1)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Get_ShouldReturnErrorWhenIdNotExists(t *testing.T) {
	expectedError := EmployeeNotFoundError

	mockRepository := mockEmployeeRepository{
		err: expectedError,
	}

	service := NewEmployeeService(mockRepository)
	_, err := service.Get(1)

	assert.Equal(t, expectedError, err)
}

func Test_GetAll_OK(t *testing.T) {

	expectedResult := []db.Employee{{}, {}, {}}

	mockRepository := mockEmployeeRepository{
		result: expectedResult,
		err:    nil,
	}

	service := NewEmployeeService(mockRepository)
	result, err := service.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_GetAll_ShouldReturnErrorWhenDatabaseFails(t *testing.T) {

	expectedResult := []db.Employee{}
	expectedError := errors.New("Falha no banco!")

	mockRepository := mockEmployeeRepository{
		result: expectedResult,
		err:    expectedError,
	}

	service := NewEmployeeService(mockRepository)
	result, err := service.GetAll()

	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedError, err)
}

func Test_Update_OK(t *testing.T) {
}

func Test_Update_ShouldReturnErrorWhenIdNotExists(t *testing.T) {
}

func Test_Update_ShouldReturnErrorWhenCodeAlreadyExists(t *testing.T) {
}

func Test_Delete_Ok(t *testing.T) {

	mockRepository := mockEmployeeRepository{
		result: db.Employee{},
		err:    nil,
	}

	service := NewEmployeeService(mockRepository)
	err := service.Delete(1)

	assert.Equal(t, nil, err)
}

func Test_Delete_ShouldReturnErrorWhenIdNotExists(t *testing.T) {

	expectedError := EmployeeNotFoundError

	mockRepository := mockEmployeeRepository{
		result: db.Employee{},
		err:    expectedError,
	}

	service := NewEmployeeService(mockRepository)
	err := service.Delete(1)

	assert.Equal(t, expectedError, err)
}
