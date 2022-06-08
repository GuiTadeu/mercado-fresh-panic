package products

import (
	"errors"
	"fmt"

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
	return r.products, nil
}

func (r *productRepository) Get(id uint64) (db.Product, error) {
	for _, product := range r.products {
		if product.Id == id {
			return product, nil
		}
	}
	return db.Product{}, errors.New("Product not found")
}

func (r *productRepository) Create(
	code string, description string, width float32, height float32, length float32,
	netWeight float32, expirationRate float32, recommendedFreezingTemp float32,
	freezingRate float32, productTypeId uint64, sellerId uint64,
) (db.Product, error) {

	s := db.Product{
		Id:                      r.getNextId(),
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
	r.products = append(r.products, s)
	return s, nil
}

func (r *productRepository) Update(id uint64, updatedProduct db.Product) (db.Product, error) {
	for index, product := range r.products {
		if product.Id == id {
			r.products[index] = updatedProduct
			return updatedProduct, nil
		}
	}
	return db.Product{}, fmt.Errorf("Product not found")
}

func (r *productRepository) Delete(id uint64) error {
	for i := range r.products {
		if r.products[i].Id == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Product not found")
}

func (r *productRepository) ExistsProductCode(code string) bool {
	for _, product := range r.products {
		if product.Code == code {
			return true
		}
	}
	return false
}

func (r *productRepository) getNextId() uint64 {

	products, err := r.GetAll()
	if err != nil {
		return 1
	}

	if len(products) == 0 {
		return 1
	}

	return products[len(products)-1].Id + 1
}