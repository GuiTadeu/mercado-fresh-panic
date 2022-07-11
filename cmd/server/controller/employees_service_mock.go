package controller

import (
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/employees"
)

type mockEmployeeService struct {
	result any
	err    error
	employeeExists bool
}

func (m mockEmployeeService) GetAll() ([]db.Employee, error) {
	if m.err != nil {
		return []db.Employee{}, m.err
	}
	return m.result.([]db.Employee), nil
}

func (m mockEmployeeService) Get(id uint64) (db.Employee, error) {
	if m.err != nil {
		return db.Employee{}, m.err
	}
	return m.result.(db.Employee), nil
}

func (m mockEmployeeService) Delete(id uint64) error {
	if m.err != nil {
		return m.err
	}
	return nil
}

func (m mockEmployeeService) ExistsProductCode(code string) bool {
	return m.err == employees.EmployeeNotFoundError
}

func (m mockEmployeeService) Create(
	cardNumberId string, firstName string, lastName string, wareHouseId uint64,
) (db.Employee, error) {
	if m.err != nil {
		return db.Employee{}, m.err
	}
	return m.result.(db.Employee), nil
}

func (m mockEmployeeService) Update(
	id uint64, cardNumberId string, firstName string, lastName string, wareHouseId uint64,
) (db.Employee, error) {
	if m.err != nil {
		return db.Employee{}, m.err
	}
	return m.result.(db.Employee), nil
}

func (m mockEmployeeService) ReportInboundOrders(id uint64) (db.ReportInboundOrders, error) {
	if m.err != nil{
		return db.ReportInboundOrders{}, m.err
	}
	return m.result.(db.ReportInboundOrders), nil
}

func (m mockEmployeeService) ReportsInboundOrders() ([]db.ReportInboundOrders, error) {
	if m.err != nil {
		return []db.ReportInboundOrders{}, m.err
	}
	return m.result.([]db.ReportInboundOrders), nil
}

func (m mockEmployeeService) ExistsEmployee(id uint64) (bool) {
	return m.employeeExists
}

