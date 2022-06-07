package employee

import (
	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/imdario/mergo"
)

type EmployeeService interface {
	Create(cardNumberId string, firstName string, lastName string, wareHouseId uint64) (database.Employee, error)
	GetAll() ([]database.Employee, error)
	Get(id uint64) (database.Employee, error)
	Update(id uint64, cardNumberId string, firstName string, lastName string, wareHouseId uint64) (database.Employee, error)
	Delete(id uint64) error
}

type employeeService struct {
	employeeRepository EmployeeRepository
}

func NewService(r EmployeeRepository) *employeeService {
	return &employeeService{
		employeeRepository: r,
	}
}

func (s employeeService) Create(cardNumberId string, firstName string, lastName string, wareHouseId uint64) (database.Employee, error) {
	employee, _ := s.employeeRepository.Create(cardNumberId, firstName, lastName, wareHouseId)
	return employee, nil
}

func (s employeeService) Get(id uint64) (database.Employee, error) {
	return s.employeeRepository.Get(id)
}

func (s employeeService) Update(id uint64, cardNumberId string, firstName string, lastName string, wareHouseId uint64) (database.Employee, error) {
	data := database.Employee{Id: id, CardNumberId: cardNumberId, FirstName: firstName, LastName: lastName, WarehouseId: wareHouseId}
	employee, err := s.Get(id)
	if err != nil {
		return database.Employee{}, err
	}
	mergo.Merge(&employee, data, mergo.WithOverride)

	e, err := s.employeeRepository.Update(employee.Id, employee.CardNumberId, employee.FirstName, employee.LastName, employee.WarehouseId)
	if err != nil {
		return database.Employee{}, err
	}
	return e, nil
}

func (s *employeeService) GetAll() ([]database.Employee, error) {
	return s.employeeRepository.GetAll()
}

func (s *employeeService) Delete(id uint64) error {
	return s.employeeRepository.Delete(id)
}
