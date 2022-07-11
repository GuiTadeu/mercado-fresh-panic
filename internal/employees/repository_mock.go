package employees

import (
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type mockEmployeeRepository struct {
	result             any
	err                error
	existsEmployeeCode bool
	getById            db.Employee
}

func (m mockEmployeeRepository) GetAll() ([]db.Employee, error) {
	if m.err != nil {
		return []db.Employee{}, m.err
	}
	return m.result.([]db.Employee), nil
}

func (m mockEmployeeRepository) Get(id uint64) (db.Employee, error) {
	if (m.getById == db.Employee{} && m.err != nil) {
		return db.Employee{}, m.err
	}
	return m.getById, nil
}

func (m mockEmployeeRepository) Delete(id uint64) error {
	if m.err != nil {
		return m.err
	}
	return nil
}

func (m mockEmployeeRepository) ExistsEmployeeCardNumberId(cardNumberId string) (bool, error) {
	return m.existsEmployeeCode, nil
}

func (m mockEmployeeRepository) Create(
	cardNumberId string, firstName string,
	lastName string, wareHouseId uint64) (db.Employee, error) {
	if m.err != nil || m.existsEmployeeCode {
		return db.Employee{}, m.err
	}
	return m.result.(db.Employee), nil
}

func (m mockEmployeeRepository) Update(updatedEmployee db.Employee) (db.Employee, error) {
	if (m.result.(db.Employee) != db.Employee{}) {
		return updatedEmployee, nil
	}
	return db.Employee{}, m.err
}

func (m mockEmployeeRepository) ReportInboundOrders(id uint64) (db.ReportInboundOrders, error) {
	if m.err != nil{
		return db.ReportInboundOrders{}, m.err
	}
	return m.result.(db.ReportInboundOrders), nil
}

func (m mockEmployeeRepository) ReportsInboundOrders() ([]db.ReportInboundOrders, error) {
	if m.err != nil {
		return []db.ReportInboundOrders{}, m.err
	}
	return m.result.([]db.ReportInboundOrders), nil
}

func (m mockEmployeeRepository) ExistsEmployee(id uint64) (bool) {
	return m.existsEmployeeCode
}