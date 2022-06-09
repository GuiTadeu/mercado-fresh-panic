package employees

import (
	"errors"
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/imdario/mergo"
)

var (
	ExistsCardNumberIdError = errors.New("card_number_id already exists")
	EmployeeNotFoundError   = errors.New("employee not found")
)

type EmployeeService interface {
	Create(cardNumberId string, firstName string, lastName string, wareHouseId uint64) (db.Employee, error)
	GetAll() ([]db.Employee, error)
	Get(id uint64) (db.Employee, error)
	Update(id uint64, cardNumberId string, firstName string, lastName string, wareHouseId uint64) (db.Employee, error)
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

func (s employeeService) Create(cardNumberId string, firstName string, lastName string, wareHouseId uint64) (db.Employee, error) {
	if s.existsEmployeeCardNumberId(cardNumberId) {
		return db.Employee{}, ExistsCardNumberIdError
	}

	employee, err := s.employeeRepository.Create(cardNumberId, firstName, lastName, wareHouseId)
	if err != nil {
		return db.Employee{}, err
	}
	return employee, nil
}

func (s employeeService) existsEmployeeCardNumberId(cardNumberId string) bool {
	return s.employeeRepository.ExistsEmployeeCardNumberId(cardNumberId)
}

func (s employeeService) Get(id uint64) (db.Employee, error) {
	return s.employeeRepository.Get(id)
}

func (s employeeService) Update(id uint64, cardNumberId string, firstName string, lastName string, wareHouseId uint64) (db.Employee, error) {
	employee, err := s.Get(id)
	if err != nil {
		return db.Employee{}, err
	}

	if s.existsEmployeeCardNumberId(cardNumberId) {
		return db.Employee{}, ExistsCardNumberIdError
	}

	data := db.Employee{Id: id, CardNumberId: cardNumberId, FirstName: firstName, LastName: lastName, WarehouseId: wareHouseId}

	err = mergo.Merge(&employee, data, mergo.WithOverride)
	if err != nil {
		return db.Employee{}, err
	}

	return s.employeeRepository.Update(employee.Id, employee.CardNumberId, employee.FirstName, employee.LastName, employee.WarehouseId)

}

func (s *employeeService) GetAll() ([]db.Employee, error) {
	return s.employeeRepository.GetAll()
}

func (s *employeeService) Delete(id uint64) error {
	_, err := s.Get(id)
	if err != nil {
		return EmployeeNotFoundError
	}
	return s.employeeRepository.Delete(id)
}
