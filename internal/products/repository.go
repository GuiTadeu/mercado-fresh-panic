package products

import (
	"database/sql"
	"log"

	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type ProductRepository interface {
	GetAll() ([]db.Product, error)
	Get(id uint64) (db.Product, error)
	Update(id uint64, updatedproduct db.Product) (db.Product, error)
	Delete(id uint64) error
	ExistsProductCode(code string) bool

	Create(code string, description string, width float32, height float32, length float32, netWeight float32, expirationRate float32,
		recommendedFreezingTemp float32, freezingRate float32, productTypeId uint64, sellerId uint64) (db.Product, error)
}

func NewProductRepository(products []db.Product) ProductRepository {
	return &productRepository{
		products: products,
	}
}

type productRepository struct {
	products []db.Product
}

func (r *productRepository) GetAll() ([]db.Product, error) {

	db := db.StorageDB
	rows, err := db.Query("SELECT * FROM products")
	
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var products []database.Product
	for rows.Next() {
		
		var product database.Product
		err := rows.Scan(
			&product.Id,
			&product.Code,
			&product.Description,
			&product.Width,
			&product.Height,
			&product.Length,
			&product.Length,
			&product.NetWeight,
			&product.ExpirationRate,
			&product.RecommendedFreezingTemp,
			&product.FreezingRate,
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

func (r *productRepository) Get(id uint64) (db.Product, error) {

	var product database.Product
	db := db.StorageDB
	rows, err := db.Query("SELECT * FROM products WHERE id = ?", id)
	
	if err != nil {
		log.Println(err)
		return product, err
	}

	for rows.Next() {
		err := rows.Scan(
			&product.Id,
			&product.Code,
			&product.Description,
			&product.Width,
			&product.Height,
			&product.Length,
			&product.Length,
			&product.NetWeight,
			&product.ExpirationRate,
			&product.RecommendedFreezingTemp,
			&product.FreezingRate,
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
) (db.Product, error) {

	db := database.StorageDB
	stmt, err := db.Prepare(`
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
		log.Fatal(err)
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
		return database.Product{}, err
	}

	insertedId, _ := result.LastInsertId()
	product := database.Product{
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

func (r *productRepository) Update(id uint64, updatedProduct db.Product) (db.Product, error) {
	
	db := db.StorageDB
	stmt, err := db.Prepare(`
		UPDATE products SET
		product_code = ?
		description = ?
		width = ?
		height = ?
		length = ?
		net_weight = ? 
		expiration_rate = ? 
		recommended_freezing_temperature = ? 
		freezing_rate = ?
		product_type = ?
		seller_id = ?
		WHERE id = ?
	`)

	if err != nil {
		log.Fatal(err)
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
		return database.Product{}, err
	}

	return updatedProduct, nil
}

func (r *productRepository) Delete(id uint64) error {

	db := db.StorageDB
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepository) ExistsProductCode(code string) bool {
	
	var product database.Product
	db := db.StorageDB
	rows, err := db.Query("SELECT * FROM products WHERE product_code = ?", code)
	
	if err != nil {
		log.Println(err)
		return false
	}

	for rows.Next() {
		err := rows.Scan(
			&product.Id,
			&product.Code,
			&product.Description,
			&product.Width,
			&product.Height,
			&product.Length,
			&product.Length,
			&product.NetWeight,
			&product.ExpirationRate,
			&product.RecommendedFreezingTemp,
			&product.FreezingRate,
			&product.ProductTypeId,
			&product.SellerId,
		)
		
		if err != nil {
			log.Println(err.Error())
			return true
		}
	}

	return false
}