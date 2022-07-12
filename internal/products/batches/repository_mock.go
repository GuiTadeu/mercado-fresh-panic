package batches

import (
	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type MockProductBatchesRepository struct {
	result            any
	err               error
	existsBatchNumber bool
	getById           models.ProductBatch
}

func (m MockProductBatchesRepository) Create(
	number uint64, currentQuantity uint64, currentTemperature float32,
	dueDate string, initialQuantity uint64, manufacturingDate string, manufacturingHour string,
	minimumTemperature float32, productId uint64, sectionId uint64,
) (models.ProductBatch, error) {
	if m.err != nil || m.existsBatchNumber {
		return models.ProductBatch{}, m.err
	}
	return m.result.(models.ProductBatch), nil
}

func (m MockProductBatchesRepository) CountProductsBySections() ([]models.CountProductsBySectionIdReport, error) {
	return m.result.([]models.CountProductsBySectionIdReport), m.err
}

func (m MockProductBatchesRepository) CountProductsBySectionId(sectionId uint64) (models.CountProductsBySectionIdReport, error) {
	return m.result.(models.CountProductsBySectionIdReport), m.err
}

func (m MockProductBatchesRepository) ExistsBatchNumber(number uint64) (bool, error) {
	return m.existsBatchNumber, m.err
}

func (m MockProductBatchesRepository) Get(id uint64) (models.ProductBatch, error) {
	if (m.getById == models.ProductBatch{} && m.err != nil) {
		return models.ProductBatch{}, m.err
	}

	return m.getById, nil
}
