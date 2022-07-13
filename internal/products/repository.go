package products

import (
	"database/sql"
	"log"

	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	Get(id uint64) (models.Product, error)
	Update(updatedproduct models.Product) (models.Product, error)
	Delete(id uint64) error
	ExistsProductCode(code string) (bool, error)

	GetReportRecords(id uint64) (models.ProductReportRecords, error)
	GetAllReportRecords() ([]models.ProductReportRecords, error)

	Create(code string, description string, width float32, height float32, length float32, netWeight float32, expirationRate float32,
		recommendedFreezingTemp float32, freezingRate float32, productTypeId uint64, sellerId uint64) (models.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) GetAllReportRecords() ([]models.ProductReportRecords, error) {
	rows, err := r.db.Query(`SELECT p.id product_id, p.description, count(pr.product_id) records_count FROM products p
		left join product_records pr on pr.product_id = p.id
		group by p.id, p.description`)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var products []models.ProductReportRecords
	for rows.Next() {

		var product models.ProductReportRecords

		// Fields must be in the same order as in the database
		err := rows.Scan(
			&product.Id,
			&product.Description,
			&product.RecordsCount,
		)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil

}

func (r *productRepository) GetReportRecords(id uint64) (models.ProductReportRecords, error) {
	rows, err := r.db.Query(`SELECT p.id product_id, p.description, count(pr.product_id) records_count FROM products p
		left join product_records pr on pr.product_id = p.id
		where p.id = ?
		group by p.id, p.description`, id)

	var product models.ProductReportRecords

	if err != nil {
		log.Println(err)
		return product, err
	}

	for rows.Next() {

		// Fields must be in the same order as in the database
		err := rows.Scan(
			&product.Id,
			&product.Description,
			&product.RecordsCount,
		)

		if err != nil {
			log.Fatal(err)
			return product, err
		}

	}

	return product, nil

}

func (r *productRepository) GetAll() ([]models.Product, error) {

	rows, err := r.db.Query("SELECT * FROM products")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var products []models.Product
	for rows.Next() {

		var product models.Product

		// Fields must be in the same order as in the database
		err := rows.Scan(
			&product.Id,
			&product.Description,
			&product.ExpirationRate,
			&product.FreezingRate,
			&product.Height,
			&product.Length,
			&product.NetWeight,
			&product.Code,
			&product.RecommendedFreezingTemp,
			&product.Width,
			&product.ProductTypeId,
			&product.SellerId,
		)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *productRepository) Get(id uint64) (models.Product, error) {

	var product models.Product
	rows, err := r.db.Query("SELECT * FROM products WHERE id = ?", id)

	if err != nil {
		log.Println(err)
		return product, err
	}

	for rows.Next() {

		// Fields must be in the same order as in the database
		err := rows.Scan(
			&product.Id,
			&product.Description,
			&product.ExpirationRate,
			&product.FreezingRate,
			&product.Height,
			&product.Length,
			&product.NetWeight,
			&product.Code,
			&product.RecommendedFreezingTemp,
			&product.Width,
			&product.ProductTypeId,
			&product.SellerId,
		)

		if err != nil {
			log.Println(err.Error())
			return product, nil
		}
	}

	return product, nil
}

func (r *productRepository) Create(
	code string, description string, width float32, height float32, length float32,
	netWeight float32, expirationRate float32, recommendedFreezingTemp float32,
	freezingRate float32, productTypeId uint64, sellerId uint64,
) (models.Product, error) {

	stmt, err := r.db.Prepare(`
		INSERT INTO products(
			product_code, 
			description, 
			width, 
			height, 
			length, 
			net_weight, 
			expiration_rate, 
			recommended_freezing_temperature, 
			freezing_rate, 
			product_type, 
			seller_id
		) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return models.Product{}, err
	}

	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(
		code,
		description,
		width,
		height,
		length,
		netWeight,
		expirationRate,
		recommendedFreezingTemp,
		freezingRate,
		productTypeId,
		sellerId,
	)

	if err != nil {
		return models.Product{}, err
	}

	insertedId, _ := result.LastInsertId()
	product := models.Product{
		Id:                      uint64(insertedId),
		Code:                    code,
		Description:             description,
		Width:                   width,
		Height:                  height,
		Length:                  length,
		NetWeight:               netWeight,
		ExpirationRate:          expirationRate,
		RecommendedFreezingTemp: recommendedFreezingTemp,
		FreezingRate:            freezingRate,
		ProductTypeId:           productTypeId,
		SellerId:                sellerId,
	}

	return product, nil
}

func (r *productRepository) Update(updatedProduct models.Product) (models.Product, error) {

	stmt, err := r.db.Prepare(`
		UPDATE products SET
		product_code = ?,
		description = ?,
		width = ?,
		height = ?,
		length = ?,
		net_weight = ? ,
		expiration_rate = ? ,
		recommended_freezing_temperature = ?,
		freezing_rate = ?,
		product_type = ?,
		seller_id = ?
		WHERE id = ?
	`)

	if err != nil {
		return models.Product{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		updatedProduct.Code,
		updatedProduct.Description,
		updatedProduct.Width,
		updatedProduct.Height,
		updatedProduct.Length,
		updatedProduct.NetWeight,
		updatedProduct.ExpirationRate,
		updatedProduct.RecommendedFreezingTemp,
		updatedProduct.FreezingRate,
		updatedProduct.ProductTypeId,
		updatedProduct.SellerId,
		updatedProduct.Id,
	)

	if err != nil {
		return models.Product{}, err
	}

	return updatedProduct, nil
}

func (r *productRepository) Delete(id uint64) error {

	stmt, err := r.db.Prepare("DELETE FROM products WHERE id = ?")
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

func (r *productRepository) ExistsProductCode(code string) (bool, error) {

	var product models.Product
	rows, err := r.db.Query("SELECT * FROM products WHERE product_code = ?", code)

	if err != nil {
		return false, err
	}

	for rows.Next() {

		// Fields must be in the same order as in the database
		err := rows.Scan(
			&product.Id,
			&product.Description,
			&product.ExpirationRate,
			&product.FreezingRate,
			&product.Height,
			&product.Length,
			&product.NetWeight,
			&product.Code,
			&product.RecommendedFreezingTemp,
			&product.Width,
			&product.ProductTypeId,
			&product.SellerId,
		)

		if err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}
