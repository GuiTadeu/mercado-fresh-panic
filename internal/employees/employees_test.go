package employees

import (
	"errors"
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

type employeeRepository2 struct {
	employees []db.Employee
}

func (r *employeeRepository2) createEmployeeMock() {
	employee := db.Employee{
		Id:           1,
		FirstName:    "testMock",
		LastName:     "lastNameMock",
		CardNumberId: "1",
		WarehouseId:  1,
	}
	r.employees = append(r.employees, employee)
}

func Test_Create_Ok(t *testing.T) {
}

func Test_Create_ShouldReturnErrorWhenCodeAlreadyExists(t *testing.T) {
}

func Test_Get_OK(t *testing.T) {
}

func Test_Get_ShouldReturnErrorWhenIdNotExists(t *testing.T) {
}

func Test_GetAll_OK(t *testing.T) {

	expectedResult := []db.Employee{{}, {}, {}}

	mockRepository := mockEmployeeRepository{
		result: expectedResult,
		err:    nil,
	}

	service := NewService(mockRepository)
	result, err := service.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_GetAll_ShouldReturnErrorWhenDatabaseFails(t *testing.T) {

	expectedResult := []db.Employee{}
	expectedError := errors.New("Deu ruim no banco!")

	mockRepository := mockEmployeeRepository{
		result: expectedResult,
		err:    expectedError,
	}

	service := NewService(mockRepository)
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

	service := NewService(mockRepository)
	err := service.Delete(1)

	assert.Equal(t, nil, err)
}

func Test_Delete_ShouldReturnErrorWhenIdNotExists(t *testing.T) {

	expectedError := errors.New("employee not found")

	mockRepository := mockEmployeeRepository{
		result: db.Employee{},
		err:    expectedError,
	}

	service := NewService(mockRepository)
	err := service.Delete(1)

	assert.Equal(t, expectedError, err)
}
