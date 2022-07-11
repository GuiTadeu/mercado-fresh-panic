package inboundorders

import (
	"database/sql"
	"log"

	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type InboundOrderRepository interface {
	Create(orderDate, orderNumber string, employeeId, productBatchId, warehouseId uint64) (database.InboundOrder, error)
	Get(id uint64) (database.InboundOrder, error)
}

type inboundOrderRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) InboundOrderRepository {
	return &inboundOrderRepository{
		db: db,
	}
}

func (r *inboundOrderRepository) Create(orderDate, orderNumber string, employeeId, productBatchId, warehouseId uint64) (database.InboundOrder, error) {

	stmt, err := r.db.Prepare("INSERT INTO inbound_orders(order_date, order_number, employee_id, product_batch_id, warehouse_id) VALUES(?,?,?,?,?)")
	if err != nil {
		return database.InboundOrder{}, err
	}

	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(orderDate, orderNumber, employeeId, productBatchId, warehouseId)
	if err != nil {
		return database.InboundOrder{}, err
	}
	insertedId, _ := result.LastInsertId()
	inboundOrder := database.InboundOrder{
		Id:             uint64(insertedId),
		OrderDate:      orderDate,
		OrderNumber:    orderNumber,
		EmployeeId:     employeeId,
		ProductBatchId: productBatchId,
		WarehouseId:    warehouseId}

	return inboundOrder, nil
}

func (r *inboundOrderRepository) Get(id uint64) (database.InboundOrder, error) {
	var inboundOrder database.InboundOrder
	rows, err := r.db.Query("SELECT * FROM inbound_orders WHERE id = ?", id)

	if err != nil {
		log.Println(err)
		return inboundOrder, err
	}

	for rows.Next() {

		err := rows.Scan(
			&inboundOrder.Id,
			&inboundOrder.OrderDate,
			&inboundOrder.OrderNumber,
			&inboundOrder.EmployeeId,
			&inboundOrder.ProductBatchId,
			&inboundOrder.WarehouseId,
		)
		if err != nil {
			log.Println(err.Error())
			return inboundOrder, nil
		}
	}

	return inboundOrder, nil
}
