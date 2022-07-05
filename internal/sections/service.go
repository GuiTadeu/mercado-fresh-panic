package sections

import (
	"errors"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/imdario/mergo"
)

var (
	ErrExistsSectionNumberError = errors.New("section number already exists")
	ErrSectionNotFoundError     = errors.New("section not found")
)

type SectionService interface {
	GetAll() ([]db.Section, error)
	Get(id uint64) (db.Section, error)
	Delete(id uint64) error
	ExistsSectionNumber(number uint64) bool

	Create(number uint64, currentTemperature float32, minimumTemperature float32, currentCapacity uint32,
		minimumCapacity uint32, maximumCapacity uint32, warehouseId uint64, productTypeId uint64) (db.Section, error)

	Update(id uint64, number uint64, currentTemperature float32, minimumTemperature float32,
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

	if s.ExistsSectionNumber(number) {
		return db.Section{}, ErrExistsSectionNumberError
	}

	section, err := s.sectionRepository.Create(
		number, currentTemperature, minimumTemperature, currentCapacity,
		minimumCapacity, maximumCapacity, warehouseId, productTypeId,
	)

	if err != nil {
		return db.Section{}, err
	}

	return section, nil
}

func (s *sectionService) Update(
	id uint64, newNumber uint64, newCurrentTemperature float32,
	newMinimumTemperature float32, newCurrentCapacity uint32,
	newMinimumCapacity uint32, newMaximumCapacity uint32,
) (db.Section, error) {

	foundSection, err := s.Get(id)
	if err != nil {
		return db.Section{}, ErrSectionNotFoundError
	}

	if s.ExistsSectionNumber(newNumber) {
		return db.Section{}, ErrExistsSectionNumberError
	}

	updatedSection := db.Section{
		Id:                 id,
		Number:             newNumber,
		CurrentTemperature: newCurrentTemperature,
		MinimumTemperature: newMinimumTemperature,
		CurrentCapacity:    newCurrentCapacity,
		MinimumCapacity:    newMinimumCapacity,
		MaximumCapacity:    newMaximumCapacity,
	}

	err = mergo.Merge(&foundSection, updatedSection, mergo.WithOverride)
	if err != nil {
		return db.Section{}, err
	}

	return s.sectionRepository.Update(id, foundSection)
}

func (s *sectionService) Delete(id uint64) error {

	_, err := s.Get(id)
	if err != nil {
		return ErrSectionNotFoundError
	}

	return s.sectionRepository.Delete(id)
}

func (s *sectionService) ExistsSectionNumber(number uint64) bool {
	return s.sectionRepository.ExistsSectionNumber(number)
}
