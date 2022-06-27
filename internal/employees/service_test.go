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
	getById := db.Employee{
		Id: 10,
		CardNumberId: "1",
		FirstName: "Nelson",
		LastName: "Nerd",
		WarehouseId: 1,
	}

	expectedResult := db.Employee{
		Id: 10,
		CardNumberId: "5",
		FirstName: "Nelson",
		LastName: "Lord",
		WarehouseId: 3,
	}

	mockProductRepository := mockEmployeeRepository{
		result:            expectedResult,
		getById:           getById,
		existsEmployeeCode: false,
	}

	service := NewEmployeeService(mockProductRepository)
	result, err := service.Update(10, "5", "Nelson", "Lord", 3)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Update_ShouldReturnErrorWhenIdNotExists(t *testing.T) {

	mockProductRepository := mockEmployeeRepository{
		existsEmployeeCode: false,
		err: EmployeeNotFoundError,
	}

	service := NewEmployeeService(mockProductRepository)
	result, err := service.Update(10, "5", "Nelson", "Lord", 3)

	assert.Empty(t, result)
	assert.Equal(t, EmployeeNotFoundError, err)
}

func Test_Update_ShouldReturnErrorWhenCodeAlreadyExists(t *testing.T) {
	getById := db.Employee{
		Id: 10,
		CardNumberId: "1",
		FirstName: "Nelson",
		LastName: "Nerd",
		WarehouseId: 1,
	}

	mockProductRepository := mockEmployeeRepository{
		existsEmployeeCode: true,
		err: ExistsCardNumberIdError,
		getById: getById,
	}

	service := NewEmployeeService(mockProductRepository)
	result, err := service.Update(10, "5", "Nelson", "Lord", 3)

	assert.Empty(t, result)
	assert.Equal(t, ExistsCardNumberIdError, err)
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
