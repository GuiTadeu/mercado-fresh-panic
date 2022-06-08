package sections

import (
	"errors"
	"fmt"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type SectionRepository interface {

	GetAll() ([]db.Section, error)
	Get(id uint64) (db.Section, error)
	Update(id uint64, updatedSection db.Section) (db.Section, error)
	Delete(id uint64) error
	ExistsSectionNumber(number uint64) bool

	Create(number uint64, currentTemperature float32, minimumTemperature float32, currentCapacity uint32,
		minimumCapacity uint32, maximumCapacity uint32, warehouseId uint64, productTypeId uint64) (db.Section, error)
}

func NewRepository(sections []db.Section) SectionRepository {
	return &sectionRepository{
		sections: sections,
	}
}

type sectionRepository struct {
	sections []db.Section
}

func (r *sectionRepository) GetAll() ([]db.Section, error) {
	return r.sections, nil
}

func (r *sectionRepository) Get(id uint64) (db.Section, error) {
	for _, section := range r.sections {
		if section.Id == id {
			return section, nil
		}
	}
	return db.Section{}, errors.New("Section not found")
}

func (r *sectionRepository) Create(
	number uint64, currentTemperature float32, minimumTemperature float32, currentCapacity uint32,
	minimumCapacity uint32, maximumCapacity uint32, warehouseId uint64, productTypeId uint64,
) (db.Section, error) {

	s := db.Section{
		Id:                 r.getNextId(),
		Number:             number,
		CurrentTemperature: currentTemperature,
		MinimumTemperature: minimumTemperature,
		CurrentCapacity:    currentCapacity,
		MinimumCapacity:    minimumCapacity,
		MaximumCapacity:    maximumCapacity,
		WarehouseId:        warehouseId,
		ProductTypeId:      productTypeId,
		Products:           []db.Product{},
	}

	r.sections = append(r.sections, s)
	return s, nil
}

func (r *sectionRepository) Update(id uint64, updatedSection db.Section) (db.Section, error) {
	for index, section := range r.sections {
		if section.Id == id {
			r.sections[index] = updatedSection
			return updatedSection, nil
		}
	}
	return db.Section{}, fmt.Errorf("Section not found")
}

func (r *sectionRepository) Delete(id uint64) error {
	for i := range r.sections {
		if r.sections[i].Id == id {
			r.sections = append(r.sections[:i], r.sections[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Section not found")
}

func (r *sectionRepository) ExistsSectionNumber(number uint64) bool {
	for _, section := range r.sections {
		if section.Number == number {
			return true
		}
	}
	return false
}

func (r *sectionRepository) getNextId() uint64 {

	sections, err := r.GetAll()
	if err != nil {
		return 1
	}

	if len(sections) == 0 {
		return 1
	}

	return sections[len(sections)-1].Id + 1
}
