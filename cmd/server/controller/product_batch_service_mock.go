package controller

import (
	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type mockProductBatchService struct {
	result any
	err    error
}

func (m mockProductBatchService) Create(number uint64, currentQuantity uint64, currentTemperature float32,
	dueDate string, initialQuantity uint64, manufacturingDate string, manufacturingHour string,
	minimumTemperature float32, productId uint64, sectionId uint64) (models.ProductBatch, error) {
	if m.err != nil {
		return models.ProductBatch{}, m.err
	}
	return m.result.(models.ProductBatch), nil
}

func (m mockProductBatchService) CountProductsBySections() ([]models.CountProductsBySectionIdReport, error) {
	if m.err != nil {
		return []models.CountProductsBySectionIdReport{}, m.err
	}
	return m.result.([]models.CountProductsBySectionIdReport), nil
}

func (m mockProductBatchService) CountProductsBySectionId(sectionId uint64) (models.CountProductsBySectionIdReport, error) {
	if m.err != nil {
		return models.CountProductsBySectionIdReport{}, m.err
	}
	return m.result.(models.CountProductsBySectionIdReport), nil
}