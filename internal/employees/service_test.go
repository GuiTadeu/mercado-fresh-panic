package employees

import (
	"errors"
	"testing"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/stretchr/testify/assert"
)

func Test_Create_Ok(t *testing.T) {

	expectedResult := db.Employee{
		Id:           1,
		FirstName:    "testMock",
		LastName:     "lastNameMock",
		CardNumberId: "1000",
		WarehouseId:  1,
	}

	mockRepository := MockEmployeeRepository{
		Result:             expectedResult,
		Err:                nil,
		ExistsEmployeeCode: false,
	}

	service := NewEmployeeService(mockRepository)
	result, err := service.Create("1000", "testMock", "lastNameMock", 1)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Create_ShouldReturnErrorWhenCodeAlreadyExists(t *testing.T) {

	expectedError := ExistsCardNumberIdError

	mockRepository := MockEmployeeRepository{
		Result:             db.Employee{},
		Err:                expectedError,
		ExistsEmployeeCode: true,
	}

	service := NewEmployeeService(mockRepository)
	_, err := service.Create("1000", "testMock", "lastNameMock", 1)

	assert.Equal(t, expectedError, err)
}

func Test_Get_OK(t *testing.T) {
	expectedResult := db.Employee{}

	mockRepository := MockEmployeeRepository{
		GetById: expectedResult,
		Err:     nil,
	}

	service := NewEmployeeService(mockRepository)
	result, err := service.Get(1)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Get_ShouldReturnErrorWhenIdNotExists(t *testing.T) {
	expectedError := EmployeeNotFoundError

	mockRepository := MockEmployeeRepository{
		Err: expectedError,
	}

	service := NewEmployeeService(mockRepository)
	_, err := service.Get(1)

	assert.Equal(t, expectedError, err)
}

func Test_GetAll_OK(t *testing.T) {

	expectedResult := []db.Employee{{}, {}, {}}

	mockRepository := MockEmployeeRepository{
		Result: expectedResult,
		Err:    nil,
	}

	service := NewEmployeeService(mockRepository)
	result, err := service.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_GetAll_ShouldReturnErrorWhenDatabaseFails(t *testing.T) {

	expectedResult := []db.Employee{}
	expectedError := errors.New("Falha no banco!")

	mockRepository := MockEmployeeRepository{
		Result: expectedResult,
		Err:    expectedError,
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

	mockProductRepository := MockEmployeeRepository{
		Result:            expectedResult,
		GetById:           getById,
		ExistsEmployeeCode: false,
	}

	service := NewEmployeeService(mockProductRepository)
	result, err := service.Update(10, "5", "Nelson", "Lord", 3)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Update_ShouldReturnErrorWhenIdNotExists(t *testing.T) {

	mockProductRepository := MockEmployeeRepository{
		ExistsEmployeeCode: false,
		Err: EmployeeNotFoundError,
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

	mockProductRepository := MockEmployeeRepository{
		ExistsEmployeeCode: true,
		Err: ExistsCardNumberIdError,
		GetById: getById,
	}

	service := NewEmployeeService(mockProductRepository)
	result, err := service.Update(10, "5", "Nelson", "Lord", 3)

	assert.Empty(t, result)
	assert.Equal(t, ExistsCardNumberIdError, err)
}

func Test_Delete_Ok(t *testing.T) {

	mockRepository := MockEmployeeRepository{
		Result: db.Employee{},
		Err:    nil,
	}

	service := NewEmployeeService(mockRepository)
	err := service.Delete(1)

	assert.Equal(t, nil, err)
}

func Test_Delete_ShouldReturnErrorWhenIdNotExists(t *testing.T) {

	expectedError := EmployeeNotFoundError

	mockRepository := MockEmployeeRepository{
		Result: db.Employee{},
		Err:    expectedError,
	}

	service := NewEmployeeService(mockRepository)
	err := service.Delete(1)

	assert.Equal(t, expectedError, err)
}

func Test_Count_Inbound_Orders_By_Employee_Id_Ok(t *testing.T) {
	
	expectedResult := db.ReportInboundOrders{
		Id:           1,
		FirstName:    "testMock",
		LastName:     "lastNameMock",
		CardNumberId: "1000",
		WarehouseId:  1,
		InboundOrdersCount: 2,
	}

	mockRepository := MockEmployeeRepository{
		Result: expectedResult,
		ExistsEmployeeCode: true,
	}

	service := NewEmployeeService(mockRepository)
	result, err := service.CountInboundOrdersByEmployeeId(1)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)


}

func Test_Count_Inbound_Orders_Ok(t *testing.T) {
	
	expectedResult := []db.ReportInboundOrders{
		{
			Id:           1,
			FirstName:    "testMock",
			LastName:     "lastNameMock",
			CardNumberId: "1000",
			WarehouseId:  1,
			InboundOrdersCount: 2,
		},
		{
			Id:           2,
			FirstName:    "testMock",
			LastName:     "lastNameMock",
			CardNumberId: "1000",
			WarehouseId:  1,
			InboundOrdersCount: 2,
		},
	}

	mockRepository := MockEmployeeRepository{
		Result: expectedResult,
	}

	service := NewEmployeeService(mockRepository)
	result, err := service.CountInboundOrders()

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)


}

func Test_Count_Inbound_Orders_By_Employee_Id_Not_Found(t *testing.T) {
	expectedError := EmployeeNotFoundError
	
	mockRepository := MockEmployeeRepository{
		ExistsEmployeeCode: false,
	}

	service := NewEmployeeService(mockRepository)
	result, err := service.CountInboundOrdersByEmployeeId(1)

	assert.Empty(t, result)
	assert.Equal(t, expectedError, err)


}
