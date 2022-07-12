package inboundorders

import db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"

type MockInboundOrdersRepository struct {
	result any
	err    error
	getById db.InboundOrder
}

func (m MockInboundOrdersRepository) Create(orderDate, orderNumber string, employeeId, productBatchId, warehouseId uint64) (db.InboundOrder, error)  {
	if m.err != nil {
		return db.InboundOrder{}, m.err
	}
	return m.result.(db.InboundOrder), nil
}

func (m MockInboundOrdersRepository) Get(id uint64) (db.InboundOrder, error) {
	if (m.getById == db.InboundOrder{} && m.err != nil) {
		return db.InboundOrder{}, m.err
	}
	return m.getById, nil
}
