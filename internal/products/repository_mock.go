package products

import (
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type MockProductRepository struct {
	Result            any
	err               error
	existsProductCode bool
	GetById           db.Product
}

func (m MockProductRepository) GetAll() ([]db.Product, error) {
	if m.err != nil {
		return []db.Product{}, m.err
	}
	return m.Result.([]db.Product), nil
}

func (m MockProductRepository) Get(id uint64) (db.Product, error) {
	if (m.GetById == db.Product{} && m.err != nil) {
		return db.Product{}, m.err
	}

	return m.GetById, nil
}

func (m MockProductRepository) Delete(id uint64) error {
	return m.err
}

func (m MockProductRepository) ExistsProductCode(code string) (bool, error) {
	return m.existsProductCode, m.err
}

func (m MockProductRepository) Create(
	code string, description string, width float32, height float32, length float32, netWeight float32,
	expirationRate float32, recommendedFreezingTemp float32, freezingRate float32, productTypeId uint64,
	sellerId uint64) (db.Product, error) {
	if m.err != nil || m.existsProductCode {
		return db.Product{}, m.err
	}
	return m.Result.(db.Product), nil
}

func (m MockProductRepository) Update(updatedproduct db.Product) (db.Product, error) {
	if (m.Result.(db.Product) != db.Product{}) {
		return updatedproduct, nil
	}
	return db.Product{}, m.err
}
