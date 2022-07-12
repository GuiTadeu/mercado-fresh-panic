package warehouses

import (
	database "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type MockWarehouseRepository struct {
	Result     any
	Err        error
	FindByCode bool
	GetById    database.Warehouse
}

func (m MockWarehouseRepository) GetAll() ([]database.Warehouse, error) {
	if m.Err != nil {
		return []database.Warehouse{}, m.Err
	}

	return m.Result.([]database.Warehouse), nil
}

func (m MockWarehouseRepository) Get(id uint64) (database.Warehouse, error) {
	if (m.GetById == database.Warehouse{} && m.Err != nil) {
		return database.Warehouse{}, m.Err
	}
	return m.GetById, nil
}

func (m MockWarehouseRepository) Delete(id uint64) error {
	if m.Err != nil {
		return m.Err
	}
	return nil

}

func (m MockWarehouseRepository) ExistsWarehouseCode(code string) (bool, error) {
	return m.FindByCode, m.Err
}

func (m MockWarehouseRepository) Create(Code string, address string, telephone string, minimunCapacity uint32, minimunTemperature float32, localityId string) (database.Warehouse, error) {
	if m.Err != nil || m.FindByCode {
		return database.Warehouse{}, m.Err
	}
	return m.Result.(database.Warehouse), nil
}

func (m MockWarehouseRepository) Update(warehouse database.Warehouse) (database.Warehouse, error) {
	if m.Err != nil {
		return database.Warehouse{}, m.Err
	}
	return m.Result.(database.Warehouse), nil
}
