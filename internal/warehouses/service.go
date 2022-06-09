package warehouses

import (
	"errors"
	"fmt"

	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/imdario/mergo"
)

var (
	ExistsWarehouseCodeError = errors.New("warehouses code already exists")
	WarehouseNotFoundError   = errors.New("warehouses not found")
)

type WarehouseService interface {
	GetAll() ([]database.Warehouse, error)
	Create(Code string, address string, telephone string, minimunCapacity uint32, minimunTemperature float32) (database.Warehouse, error)
	Get(id uint64) (database.Warehouse, error)
	Delete(id uint64) error
	Update(id uint64, code string, address string, telephone string, minimumCapacity uint32, minimumTemperature float32) (database.Warehouse, error)
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

func (s *warehouseService) Create(code string, address string, telephone string, minimumCapacity uint32, minimumTemperature float32) (database.Warehouse, error) {
	warehouses, err := s.GetAll()
	if err != nil {
		return database.Warehouse{}, err
	}
	for _, v := range warehouses {
		if v.Code == code {
			return database.Warehouse{}, ExistsWarehouseCodeError
		}
	}
	return s.warehouseRepo.Create(code, address, telephone, minimumCapacity, minimumTemperature)
}

func (s *warehouseService) Get(id uint64) (database.Warehouse, error) {
	return s.warehouseRepo.Get(id)
}

func (s *warehouseService) Delete(id uint64) error {
	return s.warehouseRepo.Delete(id)
}

func (s *warehouseService) GetNextId() uint64 {
	warehouses, err := s.warehouseRepo.GetAll()
	if err != nil {
		return 1
	}

	if len(warehouses) == 0 {
		return 1
	}

	return warehouses[len(warehouses)-1].Id + 1
}

func (s *warehouseService) Update(id uint64, code string, address string, telephone string, minimumCapacity uint32, minimumTemperature float32) (database.Warehouse, error) {
	foundWarehouse, err := s.warehouseRepo.Get(id)
	if err != nil {
		return database.Warehouse{}, WarehouseNotFoundError
	}
	isUsedCid := s.warehouseRepo.FindCode(code)
	if isUsedCid {
		return database.Warehouse{}, ExistsWarehouseCodeError
	}
	updatedWarehouse := database.Warehouse{
		Id:                 id,
		Code:               code,
		Address:            address,
		Telephone:          telephone,
		MinimunCapacity:    minimumCapacity,
		MinimumTemperature: minimumTemperature,
	}
	mergo.Merge(&foundWarehouse, updatedWarehouse, mergo.WithOverride)
	newWarehouse, err := s.warehouseRepo.Update(foundWarehouse)
	if err != nil {
		return database.Warehouse{}, fmt.Errorf("error: internal server error")
	}
	return newWarehouse, nil
}
