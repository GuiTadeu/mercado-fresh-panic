package warehouse

import (
	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	
)

type WarehouseService interface {
	GetAll() ([]database.Warehouse, error)
	Create(id uint64, Code string, address string, telephone string, minimunCapacity uint32, minimunTemperature float32) (database.Warehouse, error)
	Get(id uint64) (database.Warehouse, error)
	Delete(id uint64) error
	
}

func NewService(warehouseRepo WarehouseRepository) WarehouseService {
	return &warehouseService{
		warehouseRepo: warehouseRepo,
	}
}

type warehouseService struct {
	warehouseRepo WarehouseRepository
}

func (s *warehouseService) GetAll() ([]database.Warehouse, error) {
	return s.warehouseRepo.GetAll()
}

func (s *warehouseService) Create(id uint64, Code string, address string, telephone string, minimumCapacity uint32, minimumTemperature float32) (database.Warehouse, error) {	
	return s.warehouseRepo.Create(id, Code, address, telephone, minimumCapacity, minimumTemperature)
}

func (s *warehouseService) Get(id uint64) (database.Warehouse, error) {
	return s.warehouseRepo.Get(id)
}

func (s *warehouseService) Delete(id uint64) error {
	return s.warehouseRepo.Delete(id)
}