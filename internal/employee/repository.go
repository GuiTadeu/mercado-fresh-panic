package employee

import (
	"fmt"
	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type EmployeeRepository interface {
	GetAll() ([]database.Employee, error)
	Create(id uint64, cardNumberId uint64, firstName string, lastName string, wareHouseId uint64) (database.Employee, error)
	Delete(id uint64) error
}

func NewRepository(employees []database.Employee) EmployeeRepository {
	return &employeeRepository{
		employees: employees,
	}
}

type employeeRepository struct {
	employees []database.Employee
}

func (r *employeeRepository) GetAll() ([]database.Employee, error) {
	return r.employees, nil
}

func (r *employeeRepository) Create(id uint64, cardNumberId uint64, firstName string, lastName string, wareHouseId uint64) (database.Employee, error) {
	e := database.Employee{
		Id:           id,
		CardNumberId: cardNumberId,
		FirstName:    firstName,
		LastName:     lastName,
		WarehouseId:  wareHouseId,
	}
	r.employees = append(r.employees, e)
	return e, nil
}

func (r *employeeRepository) Delete(id uint64) error {
	for i := range r.employees {
		if r.employees[i].Id == id {
			r.employees = append(r.employees[:i], r.employees[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Section not found")
}
