package employee

import (
	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/imdario/mergo"
)

type EmployeeService interface {
	Create(cardNumberId string, firstName string, lastName string, wareHouseId uint64) (database.Employee, error)
	Get(id uint64) (database.Employee, error)
	Update(id uint64, cardNumberId string, firstName string, lastName string, wareHouseId uint64) (database.Employee, error)
}

type employeeService struct {
	repository Repository
}

func NewEmployeeService(r Repository) EmployeeService {
	return &employeeService{
		repository: r,
	}
}

func (s employeeService) Create(cardNumberId string, firstName string, lastName string, wareHouseId uint64) (database.Employee, error) {
	employee, _ := s.repository.Create(cardNumberId, firstName, lastName, wareHouseId)
	return employee, nil
}

func (s employeeService) Get(id uint64) (database.Employee, error) {
	return s.repository.Get(id)
}

func (s employeeService) Update(id uint64, cardNumberId string, firstName string, lastName string, wareHouseId uint64) (database.Employee, error) {
	data := database.Employee{Id: id, CardNumberId: cardNumberId, FirstName: firstName, LastName: lastName, WarehouseId: wareHouseId}
	employee, err := s.Get(id)
	if err != nil {
		return database.Employee{}, err
	}
	mergo.Merge(&employee, data, mergo.WithOverride)

	e, err := s.repository.Update(employee.Id, employee.CardNumberId, employee.FirstName, employee.LastName, employee.WarehouseId)
	if err != nil {
		return database.Employee{}, err
	}
	return e, nil
}
