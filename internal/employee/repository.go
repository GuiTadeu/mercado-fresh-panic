package employee

import (
	"errors"
	"fmt"
	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type EmployeeRepository interface {
	GetAll() ([]database.Employee, error)
	Create(cardNumberId string, firstName string, lastName string, wareHouseId uint64) (database.Employee, error)
	getNextId() uint64
	Get(id uint64) (database.Employee, error)
	Update(id uint64, cardNumberId string, firstName string, lastName string, wareHouseId uint64) (database.Employee, error)
	Delete(id uint64) error
}

type employeeRepository struct {
	employees []database.Employee
}

func NewRepository(employees []database.Employee) EmployeeRepository {
	return &employeeRepository{
		employees: employees,
	}
}

func (r *employeeRepository) Create(cardNumberId string, firstName string, lastName string, wareHouseId uint64) (database.Employee, error) {
	e := database.Employee{Id: r.getNextId(), CardNumberId: cardNumberId, FirstName: firstName, LastName: lastName, WarehouseId: wareHouseId}
	r.employees = append(r.employees, e)
	return e, nil
}
func (r *employeeRepository) Get(id uint64) (database.Employee, error) {
	for _, employee := range r.employees {
		if employee.Id == id {
			return employee, nil
		}
	}
	return database.Employee{}, errors.New("section not found")
}

func (r *employeeRepository) GetAll() ([]database.Employee, error) {
	return r.employees, nil
}

func (r *employeeRepository) Update(id uint64, cardNumberId string, firstName string, lastName string, wareHouseId uint64) (database.Employee, error) {
	e := database.Employee{Id: id, CardNumberId: cardNumberId, FirstName: firstName, LastName: lastName, WarehouseId: wareHouseId}
	updated := false
	for i := range r.employees {
		if r.employees[i].Id == id {
			r.employees[i] = e
			updated = true
			break
		}
	}
	if !updated {
		return database.Employee{}, fmt.Errorf("produto %d não encontrado", id)
	}
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

func (r *employeeRepository) getNextId() uint64 {
	employees, err := r.GetAll()
	if err != nil {
		return 1
	}

	if len(employees) == 0 {
		return 1
	}

	return employees[len(employees)-1].Id + 1
}
