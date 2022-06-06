package employee

import "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"

type Service interface {
	Create(id uint64, cardNumberId uint64, firstName string, lastName string, wareHouseId uint64) (database.Employee, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s service) Create(id uint64, cardNumberId uint64, firstName string, lastName string, wareHouseId uint64) (database.Employee, error) {
	employee, _ := s.repository.Create(id, cardNumberId, firstName, lastName, wareHouseId)
	return employee, nil
}
