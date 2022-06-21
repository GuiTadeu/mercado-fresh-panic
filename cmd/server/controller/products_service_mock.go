package controller

import (
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type mockProductService struct {
	result any
	err    error
	getById db.Product
	existsProductCode bool
}

func (m mockProductService) GetAll() ([]db.Product, error) {
	if m.err != nil {
		return []db.Product{}, m.err
	}
	return m.result.([]db.Product), nil
}

func (m mockProductService) Get(id uint64) (db.Product, error) {
	if (m.getById == db.Product{} && m.err != nil) {
		return db.Product{}, m.err
	}

	return m.getById, nil
}

func (m mockProductService) Delete(id uint64) error {
	if m.err != nil {
		return m.err
	}
	return nil
}

func (m mockProductService) ExistsProductCode(code string) bool {
	return m.existsProductCode
}

func (m mockProductService) Create(
	code string, description string, width float32, height float32, length float32, netWeight float32,
	expirationRate float32, recommendedFreezingTemp float32, freezingRate float32, productTypeId uint64,
	sellerId uint64) (db.Product, error) {
	if m.err != nil || m.existsProductCode {
		return db.Product{}, m.err
	}
	return m.result.(db.Product), nil
}

func (m mockProductService) Update(id uint64, newCode string, newDescription string, newWidth float32, newHeight float32, newLength float32,
	newNetWeight float32, newExpirationRate float32, newRecommendedFreezingTemp float32, newFreezingRate float32) (db.Product, error) {
	// if (m.result.(db.Product) != db.Product{}) {
	// 	return updatedproduct, nil
	// }
	// return db.Product{}, m.err
	return db.Product{}, m.err
}