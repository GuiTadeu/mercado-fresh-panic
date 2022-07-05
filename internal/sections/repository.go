package sections

import (
	"database/sql"
	"errors"
	"log"

	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type SectionRepository interface {
	GetAll() ([]database.Section, error)
	Get(id uint64) (database.Section, error)
	Update(id uint64, updatedSection database.Section) (database.Section, error)
	Delete(id uint64) error
	ExistsSectionNumber(number uint64) bool

	Create(number uint64, currentTemperature float32, minimumTemperature float32, currentCapacity uint32,
		minimumCapacity uint32, maximumCapacity uint32, warehouseId uint64, productTypeId uint64) (database.Section, error)
}

func NewRepository(sections []database.Section) SectionRepository {
	return &sectionRepository{
		sections: sections,
	}
}

type sectionRepository struct {
	sections []database.Section
}

func (r *sectionRepository) GetAll() ([]database.Section, error) {
	sections := []database.Section{}
	db := database.StorageDB

	rows, err := db.Query("SELECT * FROM sections")
	if err != nil {
		return sections, err
	}

	defer rows.Close()

	for rows.Next() {
		var section database.Section

		if err := rows.Scan(
			&section.Id,
			&section.Number,
			&section.CurrentCapacity,
			&section.CurrentTemperature,
			&section.MaximumCapacity,
			&section.MinimumCapacity,
			&section.MinimumTemperature,
			&section.ProductTypeId,
			&section.WarehouseId,
		); err != nil {
			return sections, err
		}

		sections = append(sections, section)
	}

	return sections, nil
}

func (r *sectionRepository) Get(id uint64) (database.Section, error) {
	row := database.StorageDB.QueryRow("SELECT * FROM sections WHERE ID = ?", id)

	section := database.Section{}

	err := row.Scan(
		&section.Id,
		&section.Number,
		&section.CurrentCapacity,
		&section.CurrentTemperature,
		&section.MaximumCapacity,
		&section.MinimumCapacity,
		&section.MinimumTemperature,
		&section.ProductTypeId,
		&section.WarehouseId,
	)
	// ID not found
	if errors.Is(err, sql.ErrNoRows) {
		return section, ErrSectionNotFoundError
	}

	// Other errors
	if err != nil {
		return section, err
	}

	return section, nil

}

func (r *sectionRepository) Create(number uint64, currentTemperature float32, minimumTemperature float32, currentCapacity uint32, minimumCapacity uint32, maximumCapacity uint32, warehouseId uint64, productTypeId uint64,
) (database.Section, error) {

	section := database.Section{}
	db := database.StorageDB

	query := `INSERT INTO sections 
	(section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, 
		maximum_capacity, warehouse_id, product_type) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := db.Exec(
		query,
		number,
		currentTemperature,
		minimumTemperature,
		currentCapacity,
		minimumCapacity,
		maximumCapacity,
		warehouseId,
		productTypeId,
	)
	if err != nil {
		return section, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return section, err
	}

	section.Id = uint64(lastID)
	section.Number = number
	section.CurrentTemperature = currentTemperature
	section.MinimumTemperature = minimumTemperature
	section.CurrentCapacity = currentCapacity
	section.MinimumCapacity = minimumCapacity
	section.MaximumCapacity = maximumCapacity
	section.WarehouseId = warehouseId
	section.ProductTypeId = productTypeId

	return section, nil
}

func (r *sectionRepository) Update(id uint64, updatedSection database.Section) (database.Section, error) {
	db := database.StorageDB

	query := `UPDATE sections SET 
	section_number=?, current_temperature=?, minimum_temperature=?, current_capacity=?, 
	minimum_capacity=?, maximum_capacity=?, warehouse_id=?, product_type=? WHERE id=?`

	result, err := db.Exec(
		query,
		&updatedSection.Number,
		&updatedSection.CurrentTemperature,
		&updatedSection.MinimumTemperature,
		&updatedSection.CurrentCapacity,
		&updatedSection.MinimumCapacity,
		&updatedSection.MaximumCapacity,
		&updatedSection.WarehouseId,
		&updatedSection.ProductTypeId,
		&updatedSection.Id,
	)
	if err != nil {
		return updatedSection, err
	}

	affectedRows, err := result.RowsAffected()
	// ID not found
	if affectedRows == 0 {
		return updatedSection, ErrSectionNotFoundError
	}

	// Other errors
	if err != nil {
		return updatedSection, err
	}

	return updatedSection, nil
}

func (r *sectionRepository) Delete(id uint64) error {

	result, err := database.StorageDB.Exec("DELETE FROM sections WHERE id=?", id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()

	// ID not found
	if affectedRows == 0 {
		return ErrSectionNotFoundError
	}

	// Other errors
	if err != nil {
		return err
	}

	return nil
}

func (r *sectionRepository) ExistsSectionNumber(number uint64) bool {
	var section database.Section
	db := database.StorageDB
	rows, err := db.Query("SELECT * FROM sections WHERE section_number = ?", number)

	if err != nil {
		log.Println(err)
		return false
	}

	for rows.Next() {

		err := rows.Scan(
			&section.Number,
			&section.CurrentTemperature,
			&section.MinimumTemperature,
			&section.CurrentCapacity,
			&section.MinimumCapacity,
			&section.MaximumCapacity,
			&section.WarehouseId,
			&section.ProductTypeId,
		)

		if err != nil {
			log.Println(err.Error())
			return true
		}
	}
	return false
}
