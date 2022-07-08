package sections

import (
	"database/sql"
	"log"

	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type SectionRepository interface {
	GetAll() ([]database.Section, error)
	Get(id uint64) (database.Section, error)
	Update(updatedSection database.Section) (database.Section, error)
	Delete(id uint64) error
	ExistsSectionNumber(number uint64) (bool, error)

	Create(number uint64, currentTemperature float32, minimumTemperature float32, currentCapacity uint32,
		minimumCapacity uint32, maximumCapacity uint32, warehouseId uint64, productTypeId uint64) (database.Section, error)
}

type sectionRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) SectionRepository {
	return &sectionRepository{
		db: db,
	}
}

func (r *sectionRepository) GetAll() ([]database.Section, error) {

	rows, err := r.db.Query("SELECT * FROM sections")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	var sections []database.Section
	for rows.Next() {

		var section database.Section

		err := rows.Scan(
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

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		sections = append(sections, section)
	}

	return sections, nil
}

func (r *sectionRepository) Get(id uint64) (database.Section, error) {

	var section database.Section
	rows, err := r.db.Query("SELECT * FROM sections WHERE ID = ?", id)

	if err != nil {
		log.Println(err)
		return section, err
	}

	for rows.Next() {

		err := rows.Scan(
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

		if err != nil {
			log.Println(err.Error())
			return section, err
		}

		return section, nil
	}

	return section, ErrSectionNotFoundError

}

func (r *sectionRepository) Create(number uint64, currentTemperature float32, minimumTemperature float32, currentCapacity uint32, minimumCapacity uint32, maximumCapacity uint32, warehouseId uint64, productTypeId uint64,
) (database.Section, error) {

	stmt, err := r.db.Prepare(`
	INSERT INTO sections(
		section_number, 
		current_temperature, 
		minimum_temperature, 
		current_capacity, 
		minimum_capacity, 
		maximum_capacity, 
		warehouse_id, product_type
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		`)

	if err != nil {
		return database.Section{}, err
	}

	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(
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
		return database.Section{}, err
	}

	insertedId, _ := result.LastInsertId()
	section := database.Section{
		Id:                 uint64(insertedId),
		Number:             number,
		CurrentTemperature: currentTemperature,
		MinimumTemperature: minimumTemperature,
		CurrentCapacity:    currentCapacity,
		MinimumCapacity:    minimumCapacity,
		MaximumCapacity:    maximumCapacity,
		WarehouseId:        warehouseId,
		ProductTypeId:      productTypeId,
	}

	return section, nil
}

func (r *sectionRepository) Update(updatedSection database.Section) (database.Section, error) {

	stmt, err := r.db.Prepare(`
	UPDATE sections SET 
	section_number=?, 
	current_temperature=?, 
	minimum_temperature=?, 
	current_capacity=?, 
	minimum_capacity=?, 
	maximum_capacity=?,
	warehouse_id=?, 
	product_type=? 
	WHERE id=?
	 `)

	if err != nil {
		return database.Section{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		updatedSection.Number,
		updatedSection.CurrentTemperature,
		updatedSection.MinimumTemperature,
		updatedSection.CurrentCapacity,
		updatedSection.MinimumCapacity,
		updatedSection.MaximumCapacity,
		updatedSection.WarehouseId,
		updatedSection.ProductTypeId,
		updatedSection.Id,
	)

	if err != nil {
		return database.Section{}, err
	}

	return updatedSection, nil
}

func (r *sectionRepository) Delete(id uint64) error {

	stmt, err := r.db.Prepare("DELETE FROM sections WHERE id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (r *sectionRepository) ExistsSectionNumber(number uint64) (bool, error) {

	var section database.Section

	rows, err := r.db.Query("SELECT section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type from Sections where section_number = ?", number)

	if err != nil {
		return false, err
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
			return false, err
		}

		return true, nil
	}

	return false, nil
}
