package employee

import "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"

type EmployeeService interface {
	GetAll() ([]database.Employee, error)
	Create(id uint64, cardNumberId uint64, firstName string, lastName string, wareHouseId uint64) (database.Employee, error)
	Delete(id uint64) error
	GetNextId() uint64
}

type employeeService struct {
	employeeRepository EmployeeRepository
}

func NewService(r EmployeeRepository) *employeeService {
	return &employeeService{
		employeeRepository: r,
	}
}

func (s *employeeService) GetAll() ([]database.Employee, error) {
	return s.employeeRepository.GetAll()
}

func (s employeeService) Create(id uint64, cardNumberId uint64, firstName string, lastName string, wareHouseId uint64) (database.Employee, error) {
	employee, _ := s.employeeRepository.Create(id, cardNumberId, firstName, lastName, wareHouseId)
	return employee, nil
}

func (s *employeeService) GetNextId() uint64 {
	employees, err := s.employeeRepository.GetAll()
	if err != nil {
		return 1
	}

	if len(employees) == 0 {
		return 1
	}

	return employees[len(employees)-1].Id + 1
}

func (s *employeeService) Delete(id uint64) error {
	return s.employeeRepository.Delete(id)
}
