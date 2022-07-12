package inboundorders

import (
	"errors"
	"testing"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/employees"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/warehouses"
	"github.com/stretchr/testify/assert"
)

func Test_Create_Ok(t *testing.T) {
	expectedResult := db.InboundOrder{
		Id: 1,
		OrderDate: "2022-04-04",
		OrderNumber: "order#1",
		EmployeeId: 1,
		ProductBatchId: 1,
		WarehouseId: 1,
	}

	mockWarehouse := db.Warehouse{
		Id: 1,
		Code: "1",
		Address: "rua alameda",
		Telephone: "19999999999",
		MinimunCapacity: 10,
		MinimumTemperature: 1,
		LocalityID: "1",
	}

	mockEmployeeRepository := employees.MockEmployeeRepository{
		ExistsEmployeeCode: true,
		Err: nil,
	}

	mockWarehouseRepository := warehouses.MockWarehouseRepository{
		GetById: mockWarehouse,
		Err: nil,
	}

	mockInboundOrdersRepository := MockInboundOrdersRepository{
		result: expectedResult,
		err: nil,
	}
	service := NewInboundOrderService(mockEmployeeRepository, mockWarehouseRepository, mockInboundOrdersRepository)
	result, err := service.Create("2022-04-04", "order#1", 1, 1, 1)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)

}

func Test_Create_Employee_Not_Found(t *testing.T) {
	expectedError := EmployeeNotFoundError

	expectedResult := db.InboundOrder{
		Id: 1,
		OrderDate: "2022-04-04",
		OrderNumber: "order#1",
		EmployeeId: 1,
		ProductBatchId: 1,
		WarehouseId: 1,
	}

	mockEmployeeRepository := employees.MockEmployeeRepository{
		ExistsEmployeeCode: false,
	}

	mockWarehouseRepository := warehouses.MockWarehouseRepository{}

	mockInboundOrdersRepository := MockInboundOrdersRepository{
		result: expectedResult,
		err: expectedError,
	}
	service := NewInboundOrderService(mockEmployeeRepository, mockWarehouseRepository, mockInboundOrdersRepository)
	result, err := service.Create("2022-04-04", "order#1", 1, 1, 1)

	assert.Empty(t, result)
	assert.Equal(t, expectedError, err)

}

func Test_Create_Warehouse_Not_Found(t *testing.T) {
	expectedError := WarehouseNotFoundError

	expectedResult := db.InboundOrder{
		Id: 1,
		OrderDate: "2022-04-04",
		OrderNumber: "order#1",
		EmployeeId: 1,
		ProductBatchId: 1,
		WarehouseId: 1,
	}

	mockEmployeeRepository := employees.MockEmployeeRepository{
		ExistsEmployeeCode: true,
		Err: nil,
	}

	mockWarehouseRepository := warehouses.MockWarehouseRepository{
		Err: expectedError,
	}

	mockInboundOrdersRepository := MockInboundOrdersRepository{
		result: expectedResult,
		err: expectedError,
	}
	service := NewInboundOrderService(mockEmployeeRepository, mockWarehouseRepository, mockInboundOrdersRepository)
	result, err := service.Create("2022-04-04", "order#1", 1, 1, 1)

	assert.Empty(t, result)
	assert.Equal(t, expectedError, err)
}

func Test_Create_Internal_Server_Error(t *testing.T) {
	expectedError := errors.New("Internal Server Error")

	mockWarehouse := db.Warehouse{
		Id: 1,
		Code: "1",
		Address: "rua alameda",
		Telephone: "19999999999",
		MinimunCapacity: 10,
		MinimumTemperature: 1,
		LocalityID: "1",
	}

	mockEmployeeRepository := employees.MockEmployeeRepository{
		ExistsEmployeeCode: true,
		Err: nil,
	}

	mockWarehouseRepository := warehouses.MockWarehouseRepository{
		GetById: mockWarehouse,
		Err: nil,
	}

	mockInboundOrdersRepository := MockInboundOrdersRepository{
		err: expectedError,
	}
	service := NewInboundOrderService(mockEmployeeRepository, mockWarehouseRepository, mockInboundOrdersRepository)
	result, err := service.Create("2022-04-04", "order#1", 1, 1, 1)

	assert.Empty(t, result)
	assert.Equal(t, expectedError, err)
}