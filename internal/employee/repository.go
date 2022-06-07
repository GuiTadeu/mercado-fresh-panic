package employee

import (
	"errors"
	"fmt"
	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type Repository interface {
	GetAll() ([]database.Employee, error)
	Create(cardNumberId string, firstName string, lastName string, wareHouseId uint64) (database.Employee, error)
	getNextId() uint64
	Get(id uint64) (database.Employee, error)
	Update(id uint64, cardNumberId string, firstName string, lastName string, wareHouseId uint64) (database.Employee, error)
}

type repository struct {
	employees []database.Employee
}

func NewRepository(employees []database.Employee) Repository {
	return &repository{
		employees: employees,
	}
}

func (r *repository) GetAll() ([]database.Employee, error) {
	return r.employees, nil
}

func (r *repository) Create(cardNumberId string, firstName string, lastName string, wareHouseId uint64) (database.Employee, error) {
	e := database.Employee{Id: r.getNextId(), CardNumberId: cardNumberId, FirstName: firstName, LastName: lastName, WarehouseId: wareHouseId}
	r.employees = append(r.employees, e)
	return e, nil
}

func (r *repository) Get(id uint64) (database.Employee, error) {
	for _, employee := range r.employees {
		if employee.Id == id {
			return employee, nil
		}
	}
	return database.Employee{}, errors.New("section not found")
}

func (r *repository) getNextId() uint64 {
	employees, err := r.GetAll()
	if err != nil {
		return 1
	}

	if len(employees) == 0 {
		return 1
	}

	return employees[len(employees)-1].Id + 1
}

func (r *repository) Update(id uint64, cardNumberId string, firstName string, lastName string, wareHouseId uint64) (database.Employee, error) {
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
