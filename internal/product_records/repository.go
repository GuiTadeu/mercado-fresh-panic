package productrecords

import (
	"database/sql"
	"log"

	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type ProductRecordsRepository interface {
	Create(lastUpdateDate string, purchasePrice float32, salePrice float32, productId uint64) (models.ProductRecord, error)
	Get(id uint64) (models.ProductRecord, error)
	GetAll() ([]models.ProductRecord, error)
}

type productRecordsRepository struct {
	db *sql.DB
}

func NewProductRecordsRepository(db *sql.DB) ProductRecordsRepository {
	return &productRecordsRepository{
		db: db,
	}
}

func (r *productRecordsRepository) Create(
	lastUpdateDate string, purchasePrice float32, salePrice float32, productId uint64,
) (models.ProductRecord, error) {

	stmt, err := r.db.Prepare(`
		INSERT INTO product_records(
			last_update_date, 
			purchase_price, 
			sale_price, 
			product_id
		) VALUES(?, ?, ?, ?)
	`)

	if err != nil {
		return models.ProductRecord{}, err
	}

	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(
		lastUpdateDate,
		purchasePrice,
		salePrice,
		productId,
	)

	if err != nil {
		return models.ProductRecord{}, err
	}

	insertedId, _ := result.LastInsertId()
	product := models.ProductRecord{
		Id:             uint64(insertedId),
		LastUpdateDate: lastUpdateDate,
		PurchasePrice:  purchasePrice,
		SalePrice:      salePrice,
		ProductId:      productId,
	}

	return product, nil
}

func (r *productRecordsRepository) Get(id uint64) (models.ProductRecord, error) {
	var productRecords models.ProductRecord
	rows, err := r.db.Query("SELECT * FROM product_records WHERE id = ?", id)

	if err != nil {
		log.Println(err)
		return productRecords, err
	}

	for rows.Next() {

		err := rows.Scan(
			&productRecords.Id,
			&productRecords.LastUpdateDate,
			&productRecords.PurchasePrice,
			&productRecords.SalePrice,
			&productRecords.ProductId,
		)
		if err != nil {
			log.Println(err.Error())
			return productRecords, nil
		}
	}

	return productRecords, nil
}

func (r *productRecordsRepository) GetAll() ([]models.ProductRecord, error) {
	var productRecords []models.ProductRecord
	rows, err := r.db.Query("SELECT * FROM product_records")

	if err != nil {
		log.Println(err)
		return productRecords, err
	}

	for rows.Next() {
		var productRec models.ProductRecord

		err := rows.Scan(
			&productRec.Id,
			&productRec.LastUpdateDate,
			&productRec.PurchasePrice,
			&productRec.SalePrice,
			&productRec.ProductId,
		)
		if err != nil {
			log.Println(err.Error())
			return productRecords, nil
		}

		productRecords = append(productRecords, productRec)
	}

	return productRecords, nil
}
