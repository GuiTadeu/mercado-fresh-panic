package products

import (
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type MockProductRepository struct {
	Result             any
	Err                error
	ExistsProductsCode bool
	GetById            db.Product
}

func (m MockProductRepository) GetAll() ([]db.Product, error) {
	if m.Err != nil {
		return []db.Product{}, m.Err
	}
	return m.Result.([]db.Product), nil
}

func (m MockProductRepository) Get(id uint64) (db.Product, error) {
	if (m.GetById == db.Product{} && m.Err != nil) {
		return db.Product{}, m.Err
	}

	return m.GetById, nil
}

func (m MockProductRepository) Delete(id uint64) error {
	return m.Err
}

func (m MockProductRepository) ExistsProductCode(code string) (bool, error) {
	return m.ExistsProductsCode, m.Err
}

func (m MockProductRepository) Create(
	code string, description string, width float32, height float32, length float32, netWeight float32,
	expirationRate float32, recommendedFreezingTemp float32, freezingRate float32, productTypeId uint64,
	sellerId uint64) (db.Product, error) {
	if m.Err != nil || m.ExistsProductsCode {
		return db.Product{}, m.Err
	}
	return m.Result.(db.Product), nil
}

func (m MockProductRepository) GetReportRecords(id uint64) (db.ProductReportRecords, error) {
	if m.Err != nil {
		return db.ProductReportRecords{}, m.Err
	}
	return m.Result.(db.ProductReportRecords), nil
}

func (m MockProductRepository) GetAllReportRecords() ([]db.ProductReportRecords, error) {
	if m.Err != nil {
		return []db.ProductReportRecords{}, m.Err
	}
	return m.Result.([]db.ProductReportRecords), nil
}

func (m MockProductRepository) Update(updatedproduct db.Product) (db.Product, error) {
	if (m.Result.(db.Product) != db.Product{}) {
		return updatedproduct, nil
	}
	return db.Product{}, m.Err
}
