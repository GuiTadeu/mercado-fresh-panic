package employees

import (
	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type EmployeeRepository interface {
	GetAll() ([]database.Employee, error)
	Create(cardNumberId string, firstName string, lastName string, wareHouseId uint64) (database.Employee, error)
	getNextId() uint64
	Get(id uint64) (database.Employee, error)
	Update(id uint64, cardNumberId string, firstName string, lastName string, wareHouseId uint64) (database.Employee, error)
	Delete(id uint64) error
	ExistsEmployeeCardNumberId(cardNumberId string) bool
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
	return database.Employee{}, EmployeeNotFoundError
}

func (r *employeeRepository) GetAll() ([]database.Employee, error) {
	return r.employees, nil
}

func (r *employeeRepository) Update(id uint64, cardNumberId string, firstName string, lastName string, wareHouseId uint64) (database.Employee, error) {
	updatedEmployee := database.Employee{Id: id, CardNumberId: cardNumberId, FirstName: firstName, LastName: lastName, WarehouseId: wareHouseId}
	for i := range r.employees {
		if r.employees[i].Id == id {
			r.employees[i] = updatedEmployee
			return updatedEmployee, nil
		}
	}
	return database.Employee{}, EmployeeNotFoundError
}

func (r *employeeRepository) Delete(id uint64) error {
	for i := range r.employees {
		if r.employees[i].Id == id {
			r.employees = append(r.employees[:i], r.employees[i+1:]...)
			return nil
		}
	}
	return EmployeeNotFoundError
}

func (r *employeeRepository) ExistsEmployeeCardNumberId(cardNumberId string) bool {
	for _, employee := range r.employees {
		if employee.CardNumberId == cardNumberId {
			return true
		}
	}
	return false
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
