package warehouses

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"	
)

type WarehouseRepository interface {
	GetAll() ([]database.Warehouse, error)
	Create(Code string, address string, telephone string, minimunCapacity uint32, minimunTemperature float32) (database.Warehouse, error)
	Get(id uint64) (database.Warehouse, error)
	Delete(id uint64) error
	Update(warehouse database.Warehouse) (database.Warehouse, error)
	FindByCode(code string) bool
}

func NewRepository(warehouse []database.Warehouse) WarehouseRepository {
	return &warehouseRepository{
		warehouses: warehouse,
	}
}

type warehouseRepository struct {
	warehouses []database.Warehouse
}

func (r *warehouseRepository) GetAll() ([]database.Warehouse, error) {
	return r.warehouses, nil
}

func (r *warehouseRepository) Create(code string, address string, telephone string, minimumCapacity uint32, minimumTemperature float32) (database.Warehouse, error) {
	db := database.StorageDB

	stmt,err := db.Prepare("INSERT INTO warehouses(warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()	
	var result sql.Result
	result, err = stmt.Exec(code, address, telephone, minimumCapacity, minimumTemperature, 1)
	if err != nil {
		return database.Warehouse{}, err		
	}

	insertedId, _ := result.LastInsertId()
	warehouse := database.Warehouse{
		Id: uint64(insertedId),
		Code: code,
		Address: address,
		Telephone: telephone,
		MinimunCapacity: minimumCapacity,
		MinimumTemperature: minimumTemperature,
	}

	return warehouse, nil
}

func (r *warehouseRepository) getNextId() uint64 {
	warehouse, err := r.GetAll()
	if err != nil {
		return 1
	}
	if len(warehouse) == 0 {
		return 1
	}

	return warehouse[len(warehouse)-1].Id + 1
}

func (r *warehouseRepository) Get(id uint64) (database.Warehouse, error) {
	for _, warehouse := range r.warehouses {
		if warehouse.Id == id {
			return warehouse, nil
		}
	}

	return database.Warehouse{}, WarehouseNotFoundError
}

func (r *warehouseRepository) Delete(id uint64) error {
	for i := range r.warehouses {
		if r.warehouses[i].Id == id {
			r.warehouses = append(r.warehouses[:i], r.warehouses[i+1:]...)
			return nil
		}
	}
	return WarehouseNotFoundError
}

func (r *warehouseRepository) FindByCode(code string) bool {
	for _, warehouse := range r.warehouses {
		if warehouse.Code == code {
			return true
		}
	}
	return false
}

func (r *warehouseRepository) Update(warehouse database.Warehouse) (database.Warehouse, error) {
	for idx, w := range r.warehouses {
		if w.Id == warehouse.Id {
			r.warehouses[idx] = warehouse
			return r.warehouses[idx], nil
		}
	}
	return database.Warehouse{}, fmt.Errorf("error: warehoue with id %d not found", warehouse.Id)
}
