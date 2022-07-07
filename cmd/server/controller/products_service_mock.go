package controller

import (
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/products"
)

type mockProductService struct {
	result any
	err    error
}

func (m mockProductService) GetAll() ([]db.Product, error) {
	if m.err != nil {
		return []db.Product{}, m.err
	}
	return m.result.([]db.Product), nil
}

func (m mockProductService) Get(id uint64) (db.Product, error) {
	if m.err != nil {
		return db.Product{}, m.err
	}
	return m.result.(db.Product), nil
}

func (m mockProductService) Delete(id uint64) error {
	if m.err != nil {
		return m.err
	}
	return nil
}

func (m mockProductService) ExistsProductCode(code string) (bool, error) {
	return m.err == products.ExistsProductCodeError, m.err
}

func (m mockProductService) Create(
	code string, description string, width float32, height float32, length float32, netWeight float32,
	expirationRate float32, recommendedFreezingTemp float32, freezingRate float32, productTypeId uint64, sellerId uint64,
) (db.Product, error) {
	if m.err != nil {
		return db.Product{}, m.err
	}
	return m.result.(db.Product), nil
}

func (m mockProductService) Update(
	id uint64, newCode string, newDescription string, newWidth float32, newHeight float32, newLength float32,
	newNetWeight float32, newExpirationRate float32, newRecommendedFreezingTemp float32, newFreezingRate float32,
) (db.Product, error) {
	if m.err != nil {
		return db.Product{}, m.err
	}
	return m.result.(db.Product), nil
}
