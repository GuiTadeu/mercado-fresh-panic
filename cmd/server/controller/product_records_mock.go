package controller

import (
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type mockProductRecordsService struct {
	result any
	err    error
}

func (m mockProductRecordsService) Create(
	lastUpdateDate string, purchasePrice float32, salePrice float32, productId uint64,
) (db.ProductRecord, error) {
	if m.err != nil {
		return db.ProductRecord{}, m.err
	}
	return m.result.(db.ProductRecord), nil
}
