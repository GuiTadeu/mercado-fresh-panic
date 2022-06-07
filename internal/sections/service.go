package sections

import (
	"errors"
	"net/http"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
)

type SectionService interface {
	GetAll() ([]db.Section, error)
	Get(id uint64) (db.Section, error)
	Create(number uint64, currentTemperature float32, minimumTemperature float32, currentCapacity uint32, minimumCapacity uint32, maximumCapacity uint32, warehouseId uint64, productTypeId uint64) (db.Section, error)
	Update(id uint64, number uint64, currentTemperature float32, minimumTemperature float32, currentCapacity uint32, minimumCapacity uint32, maximumCapacity uint32, warehouseId uint64, productTypeId uint64) (db.Section, error)
	Delete(id uint64) error
}

type sectionService struct {
	sectionRepository SectionRepository
}

func NewService(r SectionRepository) SectionService {
	return &sectionService{
		sectionRepository: r,
	}
}

func (s *sectionService) GetAll() ([]db.Section, error) {
	return s.sectionRepository.GetAll()
}

func (s *sectionService) Get(id uint64) (db.Section, error) {
	return s.sectionRepository.Get(id)
}

func (s *sectionService) Create(number uint64, currentTemperature float32, minimumTemperature float32, currentCapacity uint32, minimumCapacity uint32, maximumCapacity uint32, warehouseId uint64, productTypeId uint64) (db.Section, error) {
	sections, err := s.GetAll()
	if err != nil {
		return db.Section{}, &web.CustomError{Status: http.StatusInternalServerError, Err: err}
	}

	for _, v := range sections {
		if v.Number == number {
			return db.Section{}, &web.CustomError{Status: http.StatusConflict, Err: errors.New("Section number already exists")}
		}
	}

	return s.sectionRepository.Create(number, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId)
}

func (s *sectionService) Update(id uint64, number uint64, currentTemperature float32, minimumTemperature float32, currentCapacity uint32, minimumCapacity uint32, maximumCapacity uint32, warehouseId uint64, productTypeId uint64) (db.Section, error) {

	_, err := s.Get(id)
	if err != nil {
		return db.Section{}, &web.CustomError{Status: http.StatusNotFound, Err: err}
	}

	return s.sectionRepository.Update(id, number, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId)
}

func (s *sectionService) Delete(id uint64) error {
	_, err := s.Get(id)
	if err != nil {
		return &web.CustomError{Status: http.StatusNotFound, Err: err}
	}

	return s.sectionRepository.Delete(id)
}
