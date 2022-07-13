package controller

import (
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type mockPurchaseOrdersService struct {
	result any
	err    error
}

func (m mockPurchaseOrdersService) Create(
	orderNumber string, orderDate string, trackingCode string, buyerId uint64, orderStatusId uint64, productRecordId uint64,
) (db.PurchaseOrder, error) {
	if m.err != nil {
		return db.PurchaseOrder{}, m.err
	}
	return m.result.(db.PurchaseOrder), nil
}
