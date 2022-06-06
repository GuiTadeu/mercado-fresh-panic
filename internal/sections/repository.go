package sections

import (
	"errors"
	"fmt"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type SectionRepository interface {
	GetAll() ([]db.Section, error)
	Get(id uint64) (db.Section, error)
	Create(id uint64, number uint64, currentTemperature float32, minimumTemperature float32, currentCapacity uint32, minimumCapacity uint32, maximumCapacity uint32, warehouseId uint64, productTypeId uint64) (db.Section, error)
	Update(id uint64, number uint64, currentTemperature float32, minimumTemperature float32, currentCapacity uint32, minimumCapacity uint32, maximumCapacity uint32, warehouseId uint64, productTypeId uint64) (db.Section, error)
	Delete(id uint64) error
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
	return db.Section{}, errors.New("section not found")
}

func (r *sectionRepository) Create(
	id uint64,
	number uint64,
	currentTemperature float32,
	minimumTemperature float32,
	currentCapacity uint32,
	minimumCapacity uint32,
	maximumCapacity uint32,
	warehouseId uint64,
	productTypeId uint64) (db.Section, error) {

	s := db.Section{
		Id:                 id,
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

func (r *sectionRepository) Update(
	id uint64,
	number uint64,
	currentTemperature float32,
	minimumTemperature float32,
	currentCapacity uint32,
	minimumCapacity uint32,
	maximumCapacity uint32,
	warehouseId uint64,
	productTypeId uint64) (db.Section, error) {

	for i := range r.sections {
		if r.sections[i].Id == id {
			r.sections[i].Number = number
			r.sections[i].CurrentTemperature = currentTemperature
			r.sections[i].MinimumTemperature = minimumTemperature
			r.sections[i].CurrentCapacity = currentCapacity
			r.sections[i].MinimumCapacity = minimumCapacity
			r.sections[i].MaximumCapacity = maximumCapacity
			r.sections[i].WarehouseId = warehouseId
			r.sections[i].ProductTypeId = productTypeId
			return r.sections[i], nil
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
