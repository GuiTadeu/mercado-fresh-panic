package inboundorders

import (
	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"log"
)

type InboundOrderRepository interface {
	Create(orderDate, orderNumber string, employeeId, productBatchId, warehouseId uint64) (database.InboundOrder, error)
}

type inboundOrderRepository struct{}

func NewRepository() InboundOrderRepository {
	return &inboundOrderRepository{}
}

func (r *inboundOrderRepository) Create(orderDate, orderNumber string, employeeId, productBatchId, warehouseId uint64) (database.InboundOrder, error) {
	db := database.StorageDB

	stmt, err := db.Prepare("INSERT INTO inbound_orders(order_date, order_number, employee_id, product_batch_id, warehouse_id) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(orderDate, orderNumber, employeeId, productBatchId, warehouseId)
	if err != nil {
		return database.InboundOrder{}, err
	}
	insertedId, _ := result.LastInsertId()
	inboundOrder := database.InboundOrder{
		Id: uint64(insertedId),
		OrderDate:      orderDate,
		OrderNumber:    orderNumber,
		EmployeeId:     employeeId,
		ProductBatchId: productBatchId,
		WarehouseId:    warehouseId}

	return inboundOrder, nil
}
