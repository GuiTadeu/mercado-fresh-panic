package inboundorders

import (
	"errors"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/employees"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/warehouses"
)

var (
	WarehouseNotFoundError   = errors.New("warehouse not found")
	EmployeeNotFoundError   = errors.New("employee not found")

)

type InboundOrderService interface {
	Create(orderDate, orderNumber string, employeeId, productBatchId, warehouseId uint64) (db.InboundOrder, error)
}

type inboundOrderService struct {
	employeeRepository employees.EmployeeRepository
	warehouseRepository warehouses.WarehouseRepository
	inboundOrderRepository InboundOrderRepository
}

func NewInboundOrderService(employeeRepository employees.EmployeeRepository, warehouseRepository warehouses.WarehouseRepository, inboundOrderRepository InboundOrderRepository) InboundOrderService {
	return &inboundOrderService{
		employeeRepository,
		warehouseRepository,
		inboundOrderRepository,
	}
}

func (s *inboundOrderService) Create(orderDate, orderNumber string, employeeId, productBatchId, warehouseId uint64) (db.InboundOrder, error) {
	// TODO adicionar validaçao se productBatchId existe
	 if !s.employeeRepository.ExistsEmployee(employeeId) {
	 	return db.InboundOrder{}, EmployeeNotFoundError
	 }

	 if _, err := s.warehouseRepository.Get(warehouseId); err != nil {
	 	return db.InboundOrder{}, WarehouseNotFoundError
	 }

	inboundOrder, err := s.inboundOrderRepository.Create(orderDate, orderNumber, employeeId, productBatchId, warehouseId)
	if err != nil {
		return db.InboundOrder{}, err
	}

	return inboundOrder, nil
}