package products

import (
	"errors"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/imdario/mergo"
)

var (
	ErrExistsProductCodeError      = errors.New("product code already exists")
	ErrProductNotFoundError        = errors.New("product not found")
	ErrParameterNotAcceptableError = errors.New("parameter not accepted")
)

type ProductService interface {
	GetAll() ([]db.Product, error)
	Get(id uint64) (db.Product, error)
	Delete(id uint64) error
	ExistsProductCode(code string) (bool, error)

	GetReportRecords(id uint64) (db.ProductReportRecords, error)
	GetAllReportRecords() ([]db.ProductReportRecords, error)

	Create(code string, description string, width float32, height float32, length float32, netWeight float32, expirationRate float32,
		recommendedFreezingTemp float32, freezingRate float32, productTypeId uint64, sellerId uint64) (db.Product, error)

	Update(id uint64, newCode string, newDescription string, newWidth float32, newHeight float32, newLength float32,
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

func (s *productService) GetReportRecords(id uint64) (db.ProductReportRecords, error) {
	productFound, err := s.productRepository.Get(id)
	if err != nil {
		return db.ProductReportRecords{}, err
	}

	if productFound.Id != id {
		return db.ProductReportRecords{}, ErrProductNotFoundError
	}

	return s.productRepository.GetReportRecords(id)
}

func (s *productService) GetAllReportRecords() ([]db.ProductReportRecords, error) {
	return s.productRepository.GetAllReportRecords()
}

func (s *productService) Create(
	code string, description string, width float32,
	height float32, length float32, netWeight float32, expirationRate float32,
	recommendedFreezingTemp float32, freezingRate float32, productTypeId uint64, sellerId uint64,
) (db.Product, error) {

	existsProduct, err := s.ExistsProductCode(code)

	if err != nil {
		return db.Product{}, err
	}

	if existsProduct {
		return db.Product{}, ErrExistsProductCodeError
	}

	product, err := s.productRepository.Create(
		code, description, width, height, length, netWeight, expirationRate,
		recommendedFreezingTemp, freezingRate, productTypeId, sellerId,
	)

	if err != nil {
		return db.Product{}, err
	}

	return product, nil
}

func (s *productService) Update(
	id uint64, newCode string, newDescription string, newWidth float32, newHeight float32, newLength float32,
	newNetWeight float32, newExpirationRate float32, newRecommendedFreezingTemp float32, newFreezingRate float32,
) (db.Product, error) {

	foundProduct, err := s.Get(id)
	if err != nil {
		return db.Product{}, ErrProductNotFoundError
	}

	existsProduct, err := s.ExistsProductCode(newCode)

	if err != nil {
		return db.Product{}, err
	}

	if existsProduct {
		return db.Product{}, ErrExistsProductCodeError
	}

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

	err = mergo.Merge(&foundProduct, updatedProduct, mergo.WithOverride)
	if err != nil {
		return db.Product{}, err
	}

	return s.productRepository.Update(foundProduct)
}

func (s *productService) Delete(id uint64) error {

	_, err := s.Get(id)
	if err != nil {
		return ErrProductNotFoundError
	}

	return s.productRepository.Delete(id)
}

func (s *productService) ExistsProductCode(code string) (bool, error) {
	return s.productRepository.ExistsProductCode(code)
}
