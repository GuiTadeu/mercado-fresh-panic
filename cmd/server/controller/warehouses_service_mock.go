package controller

import (
	database "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"

)

type mockWarehouseService struct {
	result any
	err    error
}

func (m mockWarehouseService) GetAll() ([]database.Warehouse, error) {
	if m.err != nil {
		return []database.Warehouse{}, m.err
	}
	return m.result.([]database.Warehouse), nil
}

func (m mockWarehouseService) Get(id uint64) (database.Warehouse, error) {
	if m.err != nil {
		return database.Warehouse{}, m.err
	}
	return m.result.(database.Warehouse), nil
}

func (m mockWarehouseService) Delete(id uint64) error {
	if m.err != nil {
		return m.err
	}
	return nil

}

func (m mockWarehouseService) Create(Code string, address string, telephone string, minimunCapacity uint32, minimunTemperature float32, localityId string) (database.Warehouse, error) {
	if m.err != nil {
		return database.Warehouse{}, m.err
	}
	return m.result.(database.Warehouse), nil
}

func (m mockWarehouseService) Update(id uint64, code string, address string, telephone string, minimumCapacity uint32, minimumTemperature float32) (database.Warehouse, error) {
	if m.err != nil {
		return database.Warehouse{}, m.err
	}
	return m.result.(database.Warehouse), nil
}
