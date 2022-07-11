package controller

import (
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/sections"
)

type mockSectionService struct {
	result any
	err    error
}

func (m mockSectionService) GetAll() ([]db.Section, error) {
	if m.err != nil {
		return []db.Section{}, m.err
	}
	return m.result.([]db.Section), nil
}

func (m mockSectionService) Get(id uint64) (db.Section, error) {
	if m.err != nil {
		return db.Section{}, m.err
	}

	return m.result.(db.Section), nil
}

func (m mockSectionService) Delete(id uint64) error {
	if m.err != nil {
		return m.err
	}
	return nil
}

func (m mockSectionService) ExistsSectionNumber(number uint64) (bool, error) {
	return m.err == sections.ErrExistsSectionNumberError, m.err
}

func (m mockSectionService) Create(
	number uint64, currentTemperature float32, minimumTemperature float32,
	currentCapacity uint32, minimumCapacity uint32, maximumCapacity uint32,
	warehouseId uint64, productTypeId uint64) (db.Section, error) {
	if m.err != nil {
		return db.Section{}, m.err
	}
	return m.result.(db.Section), nil
}

func (m mockSectionService) Update(id uint64, number uint64, currentTemperature float32, minimumTemperature float32,
	currentCapacity uint32, minimumCapacity uint32, maximumCapacity uint32) (db.Section, error) {
	if m.err != nil {
		return db.Section{}, m.err
	}
	return m.result.(db.Section), nil
}
