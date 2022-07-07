package warehouses

import (
	"database/sql"
	"log"

	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type WarehouseRepository interface {
	GetAll() ([]database.Warehouse, error)
	Create(Code string, address string, telephone string, minimunCapacity uint32, minimunTemperature float32, localityId string) (database.Warehouse, error)
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
	db := database.StorageDB
	stmt, err := db.Query("SELECT id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id FROM warehouses")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var warehouses []database.Warehouse

	for stmt.Next() {
		var warehouse database.Warehouse

		if err = stmt.Scan(
			&warehouse.Id,
			&warehouse.Code,
			&warehouse.Address,
			&warehouse.Telephone,
			&warehouse.MinimunCapacity,
			&warehouse.MinimumTemperature,
			&warehouse.LocalityID,
		); err != nil {
			return r.warehouses, err
		}

		warehouses = append(warehouses, warehouse)
	}

	return warehouses, nil
}

func (r *warehouseRepository) Create(code string, address string, telephone string, minimumCapacity uint32, minimumTemperature float32, localityId string) (database.Warehouse, error) {
	db := database.StorageDB

	stmt, err := db.Prepare("INSERT INTO warehouses(warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(code, address, telephone, minimumCapacity, minimumTemperature, localityId)
	if err != nil {
		return database.Warehouse{}, err
	}

	insertedId, _ := result.LastInsertId()
	warehouse := database.Warehouse{
		Id:                 uint64(insertedId),
		Code:               code,
		Address:            address,
		Telephone:          telephone,
		MinimunCapacity:    minimumCapacity,
		MinimumTemperature: minimumTemperature,
		LocalityID:         localityId,
	}

	return warehouse, nil
}

func (r *warehouseRepository) Get(id uint64) (database.Warehouse, error) {
	var warehouse database.Warehouse
	db := database.StorageDB
	err := db.QueryRow("SELECT id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id FROM warehouses WHERE id = ?",
		id).Scan(&warehouse.Id, &warehouse.Code, &warehouse.Address, &warehouse.Telephone, &warehouse.MinimunCapacity, &warehouse.MinimumTemperature, &warehouse.LocalityID)
	if err != nil {
		log.Println(err)
		return database.Warehouse{}, err
	}
	return warehouse, nil
}

func (r *warehouseRepository) Delete(id uint64) error {
	db := database.StorageDB
	stmt, err := db.Prepare("DELETE FROM warehouses WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (r *warehouseRepository) FindByCode(code string) bool {

	var warehouse database.Warehouse

	db := database.StorageDB

	err := db.QueryRow("SELECT warehouse_code FROM warehouses WHERE warehouse_code = ?", code).Scan(&warehouse.Code)

	return err == nil
}

func (r *warehouseRepository) Update(warehouse database.Warehouse) (database.Warehouse, error) {
	db := database.StorageDB
    stmt, err := db.Prepare(`
        UPDATE
            warehouses
        SET
            warehouse_code = ?,
            address = ?,
            telephone = ?,
            minimum_capacity = ?,
            minimum_temperature = ?,
			locality_id = ?
        WHERE
            id = ?
    `)
    if err != nil {
        log.Println(err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(
        warehouse.Code,
        warehouse.Address,
        warehouse.Telephone,
        warehouse.MinimunCapacity,
        warehouse.MinimumTemperature,
        warehouse.LocalityID,
		warehouse.Id,
    )
    if err != nil {
        return database.Warehouse{}, err
    }
    return warehouse, nil
}