package purchaseOrders

import (
	"errors"
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

var (
	ExistsIdError              = errors.New("id already exists")
	PurchaseOrderNotFoundError = errors.New(" purchase orders not found")
)

type PurchaseOrdersService interface {
	Create(
		orderNumber string,
		orderDate string,
		trackingCode string,
		buyerId uint64,
		orderStatusId uint64,
		productRecordId uint64,
	) (db.PurchaseOrder, error)
}

type purchaseOrdersService struct {
	purchaseOrdersRepository PurchaseOrdersRepository
}

func NewPurchaseOrdersService(r PurchaseOrdersRepository) PurchaseOrdersService {
	return &purchaseOrdersService{
		purchaseOrdersRepository: r,
	}
}

func (s *purchaseOrdersService) Create(
	orderNumber string, orderDate string, trackingCode string, buyerId uint64, orderStatusId uint64, productRecordId uint64,
) (db.PurchaseOrder, error) {
	purchaseOrder, err := s.purchaseOrdersRepository.Create(orderNumber, orderDate, trackingCode, buyerId, orderStatusId, productRecordId)

	if err != nil {
		return db.PurchaseOrder{}, err
	}
	existPurchaseOrderId := s.ExistsBuyerId(purchaseOrder.BuyerId)
	if existPurchaseOrderId {
		return db.PurchaseOrder{}, ExistsIdError
	}

	return purchaseOrder, nil
}

func (s *purchaseOrdersService) ExistsBuyerId(buyerId uint64) bool {
	return s.purchaseOrdersRepository.ExistsBuyerId(buyerId)
}
