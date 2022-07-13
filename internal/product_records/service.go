package productrecords

import (
	"errors"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/products"
)

var (
	ErrProductNotFoundError = errors.New("product not found")
)

type ProductRecordsService interface {
	Create(lastUpdateDate string, purchasePrice float32, salePrice float32, productId uint64) (db.ProductRecord, error)
}

type productRecordsService struct {
	productRecordsRepository ProductRecordsRepository
	productRepository        products.ProductRepository
}

func NewProductRecordsService(r ProductRecordsRepository, pr products.ProductRepository) ProductRecordsService {
	return &productRecordsService{
		productRecordsRepository: r,
		productRepository:        pr,
	}
}

func (s *productRecordsService) Create(
	lastUpdateDate string, purchasePrice float32, salePrice float32, productId uint64,
) (db.ProductRecord, error) {
	productFound, err := s.productRepository.Get(productId)
	if err != nil {
		return db.ProductRecord{}, err
	}

	if productFound.Id != productId {
		return db.ProductRecord{}, ErrProductNotFoundError
	}

	productRecords, err := s.productRecordsRepository.Create(
		lastUpdateDate, purchasePrice, salePrice, productId,
	)

	if err != nil {
		return db.ProductRecord{}, err
	}

	return productRecords, nil
}
