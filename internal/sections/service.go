package sections

import (
	"net/http"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/imdario/mergo"
)

type SectionService interface {
	GetAll() ([]db.Section, error)
	Get(id uint64) (db.Section, error)
	Delete(id uint64) error
	ExistsSectionNumber(number uint64) bool

	Create(number uint64, currentTemperature float32, minimumTemperature float32, currentCapacity uint32,
		minimumCapacity uint32, maximumCapacity uint32, warehouseId uint64, productTypeId uint64) (db.Section, error)

	Update(foundSection db.Section, number uint64, currentTemperature float32, minimumTemperature float32,
		currentCapacity uint32, minimumCapacity uint32, maximumCapacity uint32) (db.Section, error)
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

func (s *sectionService) Create(
	number uint64, currentTemperature float32, minimumTemperature float32,
	currentCapacity uint32, minimumCapacity uint32, maximumCapacity uint32,
	warehouseId uint64, productTypeId uint64,
) (db.Section, error) {
	return s.sectionRepository.Create(
		number, currentTemperature, minimumTemperature, currentCapacity,
		minimumCapacity, maximumCapacity, warehouseId, productTypeId,
	)
}

func (s *sectionService) Update(
	foundSection db.Section, newNumber uint64, newCurrentTemperature float32,
	newMinimumTemperature float32, newCurrentCapacity uint32,
	newMinimumCapacity uint32, newMaximumCapacity uint32,
) (db.Section, error) {

	id := foundSection.Id
	updatedSection := db.Section{
		Id:                 id,
		Number:             newNumber,
		CurrentTemperature: newCurrentTemperature,
		MinimumTemperature: newMinimumTemperature,
		CurrentCapacity:    newCurrentCapacity,
		MinimumCapacity:    newMinimumCapacity,
		MaximumCapacity:    newMaximumCapacity,
	}

	mergo.Merge(&foundSection, updatedSection, mergo.WithOverride)
	return s.sectionRepository.Update(id, foundSection)
}

func (s *sectionService) Delete(id uint64) error {
	_, err := s.Get(id)
	if err != nil {
		return &web.CustomError{Status: http.StatusNotFound, Err: err}
	}

	return s.sectionRepository.Delete(id)
}

func (s *sectionService) ExistsSectionNumber(number uint64) bool {
	return s.sectionRepository.ExistsSectionNumber(number)
}
