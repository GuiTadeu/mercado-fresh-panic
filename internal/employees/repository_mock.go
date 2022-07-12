package employees

import (
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type MockEmployeeRepository struct {
	Result             any
	Err                error
	ExistsEmployeeCode bool
	GetById            db.Employee
}

func (m MockEmployeeRepository) GetAll() ([]db.Employee, error) {
	if m.Err != nil {
		return []db.Employee{}, m.Err
	}
	return m.Result.([]db.Employee), nil
}

func (m MockEmployeeRepository) Get(id uint64) (db.Employee, error) {
	if (m.GetById == db.Employee{} && m.Err != nil) {
		return db.Employee{}, m.Err
	}
	return m.GetById, nil
}

func (m MockEmployeeRepository) Delete(id uint64) error {
	if m.Err != nil {
		return m.Err
	}
	return nil
}

func (m MockEmployeeRepository) ExistsEmployeeCardNumberId(cardNumberId string) (bool, error) {
	return m.ExistsEmployeeCode, nil
}

func (m MockEmployeeRepository) Create(
	cardNumberId string, firstName string,
	lastName string, wareHouseId uint64) (db.Employee, error) {
	if m.Err != nil || m.ExistsEmployeeCode {
		return db.Employee{}, m.Err
	}
	return m.Result.(db.Employee), nil
}

func (m MockEmployeeRepository) Update(updatedEmployee db.Employee) (db.Employee, error) {
	if (m.Result.(db.Employee) != db.Employee{}) {
		return updatedEmployee, nil
	}
	return db.Employee{}, m.Err
}

func (m MockEmployeeRepository) CountInboundOrdersByEmployeeId(id uint64) (db.ReportInboundOrders, error) {
	if m.Err != nil{
		return db.ReportInboundOrders{}, m.Err
	}
	return m.Result.(db.ReportInboundOrders), nil
}

func (m MockEmployeeRepository) CountInboundOrders() ([]db.ReportInboundOrders, error) {
	if m.Err != nil {
		return []db.ReportInboundOrders{}, m.Err
	}
	return m.Result.([]db.ReportInboundOrders), nil
}

func (m MockEmployeeRepository) ExistsEmployee(id uint64) (bool) {
	return m.ExistsEmployeeCode
}