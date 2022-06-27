package warehouses

import (
	database "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	
)

type mockWarehouseRepository struct {
	result any	
	err error
	findCode bool
	getById database.Warehouse	
}

func (m mockWarehouseRepository) GetAll() ([]database.Warehouse, error) {
	if m.err != nil {
		return []database.Warehouse{}, m.err
	}

	return m.result.([]database.Warehouse), nil
}

func (m mockWarehouseRepository) Get(id uint64) (database.Warehouse, error) {
	if (m.getById == database.Warehouse{} && m.err != nil) {
		return database.Warehouse{}, m.err
	}
	return m.getById, nil
}

func (m mockWarehouseRepository) Delete(id uint64) error {
	if m.err != nil {
		return m.err
	}
	return nil

}

func (m mockWarehouseRepository) FindCode(code string) bool {
	return m.findCode
}

func (m mockWarehouseRepository) Create(Code string, address string, telephone string, minimunCapacity uint32, minimunTemperature float32) (database.Warehouse, error) {
	if m.err != nil || m.findCode {
		return database.Warehouse{}, m.err
	}
	return m.result.(database.Warehouse), nil
}

func (m mockWarehouseRepository) Update(warehouse database.Warehouse) (database.Warehouse, error) {
	if m.err != nil {
		return database.Warehouse{}, m.err
	}
	return m.result.(database.Warehouse), nil
}
