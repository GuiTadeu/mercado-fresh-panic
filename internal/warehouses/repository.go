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
	ExistsWarehouseCode(code string) (bool, error)
}

func NewRepository(db *sql.DB) WarehouseRepository {
	return &warehouseRepository{
		db: db,
	}
}

type warehouseRepository struct {
	db *sql.DB
}

func (r *warehouseRepository) GetAll() ([]database.Warehouse, error) {
	stmt, err := r.db.Query("SELECT id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id FROM warehouses")
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
			return nil, err
		}

		warehouses = append(warehouses, warehouse)
	}

	return warehouses, nil
}

func (r *warehouseRepository) Create(code string, address string, telephone string, minimumCapacity uint32, minimumTemperature float32, localityId string) (database.Warehouse, error) {
	stmt, err := r.db.Prepare("INSERT INTO warehouses(warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return database.Warehouse{}, err
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
	err := r.db.QueryRow("SELECT id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id FROM warehouses WHERE id = ?",
		id).Scan(&warehouse.Id, &warehouse.Code, &warehouse.Address, &warehouse.Telephone, &warehouse.MinimunCapacity, &warehouse.MinimumTemperature, &warehouse.LocalityID)
	if err != nil {
		log.Println(err)
		return database.Warehouse{}, err
	}
	return warehouse, nil
}

func (r *warehouseRepository) Delete(id uint64) error {
	stmt, err := r.db.Prepare("DELETE FROM warehouses WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (r *warehouseRepository) ExistsWarehouseCode(code string) (bool, error) {

	var warehouse database.Warehouse

	rows, err := r.db.Query("SELECT * FROM warehouses WHERE warehouse_code = ?", code)

	if err != nil {
		return false, err
	}

	for rows.Next() {
		err := rows.Scan(&warehouse.Id, &warehouse.Code, &warehouse.Address, &warehouse.Telephone, &warehouse.MinimunCapacity, &warehouse.MinimumTemperature, &warehouse.LocalityID)

		if err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil	
}

func (r *warehouseRepository) Update(warehouse database.Warehouse) (database.Warehouse, error) {
	stmt, err := r.db.Prepare(`
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
		return database.Warehouse{}, err
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
