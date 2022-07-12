package batches

import (
	"database/sql"
	"log"

	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type ProductBatchRepository interface {
	Create(number uint64, currentQuantity uint64, currentTemperature float32,
		dueDate string, initialQuantity uint64, manufacturingDate string, manufacturingHour string,
		minimumTemperature float32, productId uint64, sectionId uint64) (models.ProductBatch, error)

	CountProductsBySections() ([]models.CountProductsBySectionIdReport, error)
	CountProductsBySectionId(sectionId uint64) (models.CountProductsBySectionIdReport, error)

	ExistsBatchNumber(number uint64) (bool, error)

	Get(id uint64) (models.ProductBatch, error)
}

type productBatchRepository struct {
	db *sql.DB
}

func NewProductBatchRepository(db *sql.DB) ProductBatchRepository {
	return &productBatchRepository{
		db: db,
	}
}

func (r *productBatchRepository) Create(
	number uint64, currentQuantity uint64, currentTemperature float32,
	dueDate string, initialQuantity uint64, manufacturingDate string, manufacturingHour string,
	minimumTemperature float32, productId uint64, sectionId uint64,
) (models.ProductBatch, error) {

	stmt, err := r.db.Prepare(`
		INSERT INTO product_batches(
			batch_number,
			current_quantity,
			current_temperature,
			due_date,
			initial_quantity,
			manufacturing_date,
			manufacturing_hour,
			minimum_temperature,
			product_id,
			section_id
		) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return models.ProductBatch{}, err
	}

	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(
		number,
		currentQuantity,
		currentTemperature,
		dueDate,
		initialQuantity,
		manufacturingDate,
		manufacturingHour,
		minimumTemperature,
		productId,
		sectionId,
	)

	if err != nil {
		return models.ProductBatch{}, err
	}

	insertedId, _ := result.LastInsertId()
	productBatch := models.ProductBatch{
		Id:                 uint64(insertedId),
		Number:             number,
		CurrentQuantity:    currentQuantity,
		CurrentTemperature: currentTemperature,
		DueDate:            dueDate,
		InitialQuantity:    initialQuantity,
		ManufacturingDate:  manufacturingDate,
		ManufacturingHour:  manufacturingHour,
		MinimumTemperature: minimumTemperature,
		ProductId:          productId,
		SectionId:          sectionId,
	}

	return productBatch, nil
}

func (r *productBatchRepository) CountProductsBySections() ([]models.CountProductsBySectionIdReport, error) {

	var report []models.CountProductsBySectionIdReport

	rows, err := r.db.Query(`
		SELECT sc.id, sc.section_number, COUNT(pb.product_id) AS products_count
		FROM product_batches pb JOIN sections sc ON sc.id = pb.section_id GROUP BY (sc.id)`)

	if err != nil {
		log.Println(err)
		return report, err
	}

	var foundReport models.CountProductsBySectionIdReport
	for rows.Next() {

		err := rows.Scan(
			&foundReport.SectionId,
			&foundReport.SectionNumber,
			&foundReport.ProductsCount,
		)

		if err != nil {
			log.Println(err.Error())
			return report, nil
		}

		report = append(report, foundReport)
	}

	return report, nil
}

func (r *productBatchRepository) CountProductsBySectionId(sectionId uint64) (models.CountProductsBySectionIdReport, error) {

	var report models.CountProductsBySectionIdReport

	rows, err := r.db.Query(`
		SELECT sc.id, sc.section_number, COUNT(pb.product_id) AS products_count
		FROM product_batches pb JOIN sections sc ON sc.id = pb.section_id WHERE sc.id = ? GROUP BY (sc.id)`, sectionId,
	)

	if err != nil {
		log.Println(err)
		return report, err
	}

	for rows.Next() {

		err := rows.Scan(
			&report.SectionId,
			&report.SectionNumber,
			&report.ProductsCount,
		)

		if err != nil {
			log.Println(err.Error())
			return report, nil
		}
	}

	return report, nil
}

func (r *productBatchRepository) ExistsBatchNumber(number uint64) (bool, error) {
	var productBatch models.ProductBatch
	rows, err := r.db.Query("SELECT * FROM product_batches WHERE batch_number = ?", number)

	if err != nil {
		return false, err
	}

	for rows.Next() {

		// Fields must be in the same order as in the database
		err := rows.Scan(
			&productBatch.Number,
			&productBatch.CurrentQuantity,
			&productBatch.CurrentTemperature,
			&productBatch.DueDate,
			&productBatch.InitialQuantity,
			&productBatch.ManufacturingDate,
			&productBatch.ManufacturingHour,
			&productBatch.MinimumTemperature,
			&productBatch.ProductId,
			&productBatch.SectionId,
		)

		if err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}

func (r *productBatchRepository) Get(id uint64) (models.ProductBatch, error) {

	var productBatch models.ProductBatch
	rows, err := r.db.Query("SELECT * FROM product_batches WHERE id = ?", id)

	if err != nil {
		return models.ProductBatch{}, err
	}

	for rows.Next() {

		// Fields must be in the same order as in the database
		err := rows.Scan(
			&productBatch.Id,
			&productBatch.Number,
			&productBatch.CurrentQuantity,
			&productBatch.CurrentTemperature,
			&productBatch.DueDate,
			&productBatch.InitialQuantity,
			&productBatch.ManufacturingDate,
			&productBatch.ManufacturingHour,
			&productBatch.MinimumTemperature,
			&productBatch.ProductId,
			&productBatch.SectionId,
		)

		if err != nil {
			return productBatch, nil
		}
	}

	return productBatch, nil
}
