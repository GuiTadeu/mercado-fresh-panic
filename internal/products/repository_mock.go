package products

import (
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type mockProductRepository struct {
	result any
	err    error
}

func (m mockProductRepository) GetAll() ([]db.Product, error) {
	if m.err != nil {
		return []db.Product{}, m.err
	}
	return m.result.([]db.Product), nil
}

func (m mockProductRepository) Get(id uint64) (db.Product, error) {
	if m.err != nil {
		return db.Product{}, m.err
	}
	return m.result.(db.Product), nil
}

func (m mockProductRepository) Delete(id uint64) error {
	if m.err != nil {
		return m.err
	}
	return nil
}

func (m mockProductRepository) ExistsProductCode(code string) bool {
	if m.err != nil {
		return false
	}
	return true
}

func (m mockProductRepository) Create(
	code string, description string, width float32, height float32, length float32, netWeight float32,
	expirationRate float32, recommendedFreezingTemp float32, freezingRate float32, productTypeId uint64,
	sellerId uint64) (db.Product, error) {
	if m.err != nil {
		return db.Product{}, m.err
	}
	return m.result.(db.Product), nil
}

func (m mockProductRepository) Update(id uint64, updatedproduct db.Product) (db.Product, error) {
	if m.err != nil {
		return db.Product{}, m.err
	}
	return m.result.(db.Product), nil
}
