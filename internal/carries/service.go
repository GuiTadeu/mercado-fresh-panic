package carries

import (
	"errors"

	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"	
)

var (
	ExistsCarrierCidError = errors.New("carriers cid already exists")
	CarrierNotFoundError = errors.New("carriers not found")
)

type CarrierService interface {
	Create(Cid string, Company_Name string, Address string, Telephone string, localityId string) (database.Carrier, error)
	GetAllCarrierInfo() ([]database.Carrier, error)
}

type carrierService struct {
	carrierRepo CarrierRepository
}

func NewCarrierService(r CarrierRepository) CarrierService {
	return &carrierService {
		carrierRepo: r,
	}
}

func (s *carrierService) Create(cid string, companyName string, address string, telephone string, localityId string) (database.Carrier, error) {
	isUsedCid, err := s.carrierRepo.ExistsCarrierCid(cid)
	if err != nil {
		return database.Carrier{}, err
	}

	if isUsedCid {
		return database.Carrier{}, ExistsCarrierCidError
	}
	return s.carrierRepo.Create(cid, companyName, address, telephone, localityId)
}

func (s *carrierService) GetAllCarrierInfo() ([]database.Carrier, error) {
	return s.carrierRepo.GetAllCarrierInfo()
}