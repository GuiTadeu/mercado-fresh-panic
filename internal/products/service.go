package products

import (
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/imdario/mergo"
)

type ProductService interface {
	GetAll() ([]db.Product, error)
	Get(id uint64) (db.Product, error)
	Delete(id uint64) error
	ExistsProductCode(code string) bool

	Create(code string, description string, width float32, height float32, length float32, netWeight float32, expirationRate float32,
		recommendedFreezingTemp float32, freezingRate float32, productTypeId uint64, sellerId uint64) (db.Product, error)

	Update(oldProduct db.Product, newCode string, newDescription string, newWidth float32, newHeight float32, newLength float32,
		newNetWeight float32, newExpirationRate float32, newRecommendedFreezingTemp float32, newFreezingRate float32) (db.Product, error)
}

type productService struct {
	productRepository ProductRepository
}

func NewProductService(r ProductRepository) ProductService {
	return &productService{
		productRepository: r,
	}
}

func (s *productService) GetAll() ([]db.Product, error) {
	return s.productRepository.GetAll()
}

func (s *productService) Get(id uint64) (db.Product, error) {
	return s.productRepository.Get(id)
}

func (s *productService) Create(
	code string, description string, width float32,
	height float32, length float32, netWeight float32, expirationRate float32,
	recommendedFreezingTemp float32, freezingRate float32, productTypeId uint64, sellerId uint64,
) (db.Product, error) {
	return s.productRepository.Create(
		code, description, width, height, length, netWeight, expirationRate,
		recommendedFreezingTemp, freezingRate, productTypeId, sellerId,
	)
}

func (s *productService) Update(
	foundProduct db.Product, newCode string, newDescription string, newWidth float32, newHeight float32, newLength float32,
	newNetWeight float32, newExpirationRate float32, newRecommendedFreezingTemp float32, newFreezingRate float32,
) (db.Product, error) {

	id := foundProduct.Id
	updatedProduct := db.Product{
		Id:                      id,
		Code:                    newCode,
		Description:             newDescription,
		Width:                   newWidth,
		Height:                  newHeight,
		Length:                  newLength,
		NetWeight:               newNetWeight,
		ExpirationRate:          newExpirationRate,
		RecommendedFreezingTemp: newRecommendedFreezingTemp,
		FreezingRate:            newFreezingRate,
	}

	mergo.Merge(&foundProduct, updatedProduct, mergo.WithOverride)
	return s.productRepository.Update(id, foundProduct)
}

func (s *productService) Delete(id uint64) error {
	return s.productRepository.Delete(id)
}

func (s *productService) ExistsProductCode(code string) bool {
	return s.productRepository.ExistsProductCode(code)
}
