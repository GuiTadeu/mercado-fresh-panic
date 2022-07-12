package batches

import (
	"errors"

	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/products"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/sections"
)

var (
	ProductNotFoundError   = errors.New("product not found")
	SectionNotFoundError   = errors.New("section not found")
	ExistsBatchNumberError = errors.New("number already exists")
)

const (
	NOT_FOUND_ID = 0
)

type ProductBatchService interface {
	Create(number uint64, currentQuantity uint64, currentTemperature float32,
		dueDate string, initialQuantity uint64, manufacturingDate string, manufacturingHour string,
		minimumTemperature float32, productId uint64, sectionId uint64) (models.ProductBatch, error)

	CountProductsBySections() ([]models.CountProductsBySectionIdReport, error)
	CountProductsBySectionId(sectionId uint64) (models.CountProductsBySectionIdReport, error)
}

type productBatchService struct {
	productBatchRepository ProductBatchRepository
	sectionRepository      sections.SectionRepository
	productRepository      products.ProductRepository
}

func NewProductBatchesService(
	pbr ProductBatchRepository,
	sr sections.SectionRepository,
	pr products.ProductRepository,
) ProductBatchService {
	return &productBatchService{
		productBatchRepository: pbr,
		sectionRepository:      sr,
		productRepository:      pr,
	}
}

func (s *productBatchService) CountProductsBySections() ([]models.CountProductsBySectionIdReport, error) {
	return s.productBatchRepository.CountProductsBySections()
}

func (s *productBatchService) CountProductsBySectionId(sectionId uint64) (models.CountProductsBySectionIdReport, error) {
	return s.productBatchRepository.CountProductsBySectionId(sectionId)
}

func (s *productBatchService) Create(
	number uint64, currentQuantity uint64, currentTemperature float32,
	dueDate string, initialQuantity uint64, manufacturingDate string, manufacturingHour string,
	minimumTemperature float32, productId uint64, sectionId uint64,
) (models.ProductBatch, error) {

	existsNumber, err := s.ExistsBatchNumber(number)

	if err != nil {
		return models.ProductBatch{}, err
	}

	if existsNumber {
		return models.ProductBatch{}, ExistsBatchNumberError
	}

	foundProduct, err := s.productRepository.Get(productId)
	if err != nil {
		return models.ProductBatch{}, err
	}

	if (foundProduct == models.Product{}) {
		return models.ProductBatch{}, ProductNotFoundError
	}

	_, err = s.sectionRepository.Get(sectionId)
	if err != nil {
		return models.ProductBatch{}, SectionNotFoundError
	}

	product, err := s.productBatchRepository.Create(
		number, currentQuantity, currentTemperature, dueDate,
		initialQuantity, manufacturingDate, manufacturingHour, minimumTemperature, productId, sectionId,
	)

	if err != nil {
		return models.ProductBatch{}, err
	}

	return product, nil
}

func (s *productBatchService) ExistsBatchNumber(number uint64) (bool, error) {
	return s.productBatchRepository.ExistsBatchNumber(number)
}
