package employee

import (
	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type repository struct{}

type Repository interface {
	GetAll() ([]database.Employee, error)
	Create(id uint64, cardNumberId uint64, firstName string, lastName string, wareHouseId uint64) (database.Employee, error)
}

func (r *repository) GetAll() ([]database.Employee, error) {
	var employees []database.Employee
	return employees, nil
}

func (r *repository) Create(id uint64, cardNumberId uint64, firstName string, lastName string, wareHouseId uint64) (database.Employee, error) {
	var employee []database.Employee
	e := database.Employee{id, cardNumberId, firstName, lastName, wareHouseId}
	employee = append(employee, e)
	return e, nil
}

func NewRepository() Repository {
	return &repository{}
}
