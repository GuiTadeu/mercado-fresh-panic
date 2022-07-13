package purchaseOrders

import (
	"database/sql"
	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"log"
)

type PurchaseOrdersRepository interface {
	Create(
		orderNumber string,
		orderDate string,
		trackingCode string,
		buyerId uint64,
		orderStatusId uint64,
		productRecordId uint64,
	) (models.PurchaseOrder, error)
	Get(id uint64) (models.PurchaseOrder, error)
	ExistsBuyerId(buyerId uint64) bool
}

type purchaseOrdersRepository struct {
	db *sql.DB
}

func NewPurchaseOrdersRepository(purchaseOrders *sql.DB) PurchaseOrdersRepository {
	return &purchaseOrdersRepository{
		db: purchaseOrders,
	}
}

func (r *purchaseOrdersRepository) Create(
	orderNumber string, orderDate string, trackingCode string, buyerId uint64, orderStatusId uint64, productRecordId uint64,
) (models.PurchaseOrder, error) {

	stmt, err := r.db.Prepare(`
		INSERT INTO purchase_orders(
		    order_number,
		    order_date,
		    tracking_code,
		    buyer_id,
		    order_status_id,
		    product_record_id
		) VALUES (?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return models.PurchaseOrder{}, err
	}

	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(
		orderNumber,
		orderDate,
		trackingCode,
		buyerId,
		orderStatusId,
		productRecordId,
	)

	if err != nil {
		return models.PurchaseOrder{}, err
	}

	insertId, _ := result.LastInsertId()
	purchaseOrders := models.PurchaseOrder{
		Id:              uint64(insertId),
		OrderNumber:     orderNumber,
		OrderDate:       orderDate,
		TrackingCode:    trackingCode,
		BuyerId:         buyerId,
		OrderStatusId:   orderStatusId,
		ProductRecordId: productRecordId,
	}

	return purchaseOrders, nil
}

func (r *purchaseOrdersRepository) Get(id uint64) (models.PurchaseOrder, error) {
	var purchaseOrder models.PurchaseOrder
	rows, err := r.db.Query("SELECT * FROM purchase_orders WHERE id = ?", id)

	if err != nil {
		log.Println(err)
		return purchaseOrder, err
	}

	for rows.Next() {

		err := rows.Scan(
			&purchaseOrder.Id,
			&purchaseOrder.OrderNumber,
			&purchaseOrder.OrderDate,
			&purchaseOrder.TrackingCode,
			&purchaseOrder.BuyerId,
			&purchaseOrder.OrderStatusId,
			&purchaseOrder.ProductRecordId,
		)
		if err != nil {
			log.Println(err.Error())
			return purchaseOrder, nil
		}
	}

	return purchaseOrder, nil
}

func (r *purchaseOrdersRepository) ExistsBuyerId(buyerId uint64) bool {
	var myPurchaseOrder models.PurchaseOrder
	err := r.db.QueryRow(`
			SELECT 
		    id,
		FROM buyers
		WHERE id = ?
		`, buyerId).Scan(&myPurchaseOrder.BuyerId)

	return err == nil

}
