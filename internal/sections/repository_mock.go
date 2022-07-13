package sections

import (
	"reflect"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type MockSectionRepository struct {
	Result              any
	err                 error
	existsSectionNumber bool
	GetById             db.Section
}

func (m MockSectionRepository) GetAll() ([]db.Section, error) {
	if m.err != nil {
		return []db.Section{}, m.err
	}
	return m.Result.([]db.Section), nil
}

func (m MockSectionRepository) Get(id uint64) (db.Section, error) {
	if (reflect.DeepEqual(m.GetById, db.Section{}) && m.err != nil) {
		return db.Section{}, m.err
	}

	return m.GetById, nil
}

func (m MockSectionRepository) Delete(id uint64) error {
	if m.err != nil {
		return m.err
	}
	return nil
}

func (m MockSectionRepository) ExistsSectionNumber(number uint64) (bool, error) {
	return m.existsSectionNumber, m.err
}

func (m MockSectionRepository) Create(
	number uint64, currentTemperature float32, minimumTemperature float32,
	currentCapacity uint32, minimumCapacity uint32, maximumCapacity uint32,
	warehouseId uint64, productTypeId uint64) (db.Section, error) {
	if m.err != nil || m.existsSectionNumber {
		return db.Section{}, m.err
	}
	return m.Result.(db.Section), nil
}

func (m MockSectionRepository) Update(updatedSection db.Section) (db.Section, error) {
	if (reflect.DeepEqual(m.Result.(db.Section), db.Section{}) == false) {
		return updatedSection, nil
	}
	return db.Section{}, m.err
}
