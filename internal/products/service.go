package products

import (
	"errors"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/imdario/mergo"
)

type ProductService interface {

	GetNextId() uint64
	GetAll() ([]db.Product, error)
	Get(id uint64) (db.Product, error)
	Delete(id uint64) error

	Create(id uint64, code string, description string, width float32, height float32, length float32, netWeight float32, expirationDate string,
		recommendedFreezingTemp float32, freezingRate float32, productTypeId uint64, sellerId uint64) (db.Product, error)

	Update(id uint64, newCode string, newDescription string, newWidth float32, newHeight float32, newLength float32, 
		newNetWeight float32, newExpirationDate string, newRecommendedFreezingTemp float32, newFreezingRate float32) (db.Product, error)
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
	id uint64, code string, description string, width float32, 
	height float32, length float32, netWeight float32, expirationDate string,
	recommendedFreezingTemp float32, freezingRate float32, productTypeId uint64, sellerId uint64,
) (db.Product, error) {
	return s.productRepository.Create(
		id, code, description, width, height, length, netWeight, expirationDate,
		recommendedFreezingTemp, freezingRate, productTypeId, sellerId,
	)
}

func (s *productService) Update(
	id uint64, newCode string, newDescription string, newWidth float32, newHeight float32, newLength float32, 
	newNetWeight float32, newExpirationDate string, newRecommendedFreezingTemp float32, newFreezingRate float32,
) (db.Product, error) {

	foundProduct, err := s.Get(id)
	if err != nil {
		return db.Product{}, errors.New("Product not found")
	}

	updatedProduct := db.Product{
		Id: id,
		Code: newCode,
		Description: newDescription,
		Width: newWidth,
		Height: newHeight,
		Length: newLength,
		NetWeight: newNetWeight,
		ExpirationDate: newExpirationDate,
		RecommendedFreezingTemp: newRecommendedFreezingTemp,
		FreezingRate: newFreezingRate,
	}

	mergo.Merge(&foundProduct, updatedProduct, mergo.WithOverride)
	return s.productRepository.Update(id, foundProduct)
}

func (s *productService) Delete(id uint64) error {
	return s.productRepository.Delete(id)
}

func (s *productService) GetNextId() uint64 {
	products, err := s.productRepository.GetAll()
	if err != nil {
		return 1
	}

	if len(products) == 0 {
		return 1
	}

	return products[len(products)-1].Id + 1
}
