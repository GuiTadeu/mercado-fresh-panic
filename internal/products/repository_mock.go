package products

import (
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type mockProductRepository struct {
	result            any
	err               error
	existsProductCode bool
	getById           db.Product
}

func (m mockProductRepository) GetAll() ([]db.Product, error) {
	if m.err != nil {
		return []db.Product{}, m.err
	}
	return m.result.([]db.Product), nil
}

func (m mockProductRepository) Get(id uint64) (db.Product, error) {
	if (m.getById == db.Product{} && m.err != nil) {
		return db.Product{}, m.err
	}

	return m.getById, nil
}

func (m mockProductRepository) Delete(id uint64) error {
	return m.err
}

func (m mockProductRepository) ExistsProductCode(code string) (bool, error) {
	return m.existsProductCode, m.err
}

func (m mockProductRepository) Create(
	code string, description string, width float32, height float32, length float32, netWeight float32,
	expirationRate float32, recommendedFreezingTemp float32, freezingRate float32, productTypeId uint64,
	sellerId uint64) (db.Product, error) {
	if m.err != nil || m.existsProductCode {
		return db.Product{}, m.err
	}
	return m.result.(db.Product), nil
}

func (m mockProductRepository) GetReportRecords(id uint64) (db.ProductReportRecords, error) {
	if m.err != nil {
		return db.ProductReportRecords{}, m.err
	}
	return m.result.(db.ProductReportRecords), nil
}

func (m mockProductRepository) GetAllReportRecords() ([]db.ProductReportRecords, error) {
	if m.err != nil {
		return []db.ProductReportRecords{}, m.err
	}
	return m.result.([]db.ProductReportRecords), nil
}

func (m mockProductRepository) Update(updatedproduct db.Product) (db.Product, error) {
	if (m.result.(db.Product) != db.Product{}) {
		return updatedproduct, nil
	}
	return db.Product{}, m.err
}
