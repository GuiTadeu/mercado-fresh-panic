package controller

import (
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type mockInboundOrderService struct {
	result any
	err    error
}

func (m mockInboundOrderService) Create(orderDate, orderNumber string, employeeId, productBatchId, warehouseId uint64) (db.InboundOrder, error) {
	if m.err != nil {
		return db.InboundOrder{}, m.err
	}
	return m.result.(db.InboundOrder), nil
}
