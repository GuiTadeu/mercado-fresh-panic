package sections

import (
	"reflect"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type mockSectionRepository struct {
	result              any
	err                 error
	existsSectionNumber bool
	getById             db.Section
}

func (m mockSectionRepository) GetAll() ([]db.Section, error) {
	if m.err != nil {
		return []db.Section{}, m.err
	}
	return m.result.([]db.Section), nil
}

func (m mockSectionRepository) Get(id uint64) (db.Section, error) {
	if (reflect.DeepEqual(m.getById, db.Section{}) && m.err != nil) {
		return db.Section{}, m.err
	}

	return m.getById, nil
}

func (m mockSectionRepository) Delete(id uint64) error {
	if m.err != nil {
		return m.err
	}
	return nil
}

func (m mockSectionRepository) ExistsSectionNumber(number uint64) (bool, error) {
	return m.existsSectionNumber, m.err
}

func (m mockSectionRepository) Create(
	number uint64, currentTemperature float32, minimumTemperature float32,
	currentCapacity uint32, minimumCapacity uint32, maximumCapacity uint32,
	warehouseId uint64, productTypeId uint64) (db.Section, error) {
	if m.err != nil || m.existsSectionNumber {
		return db.Section{}, m.err
	}
	return m.result.(db.Section), nil
}

func (m mockSectionRepository) Update(updatedSection db.Section) (db.Section, error) {
	if (reflect.DeepEqual(m.result.(db.Section), db.Section{}) == false) {
		return updatedSection, nil
	}
	return db.Section{}, m.err
}
