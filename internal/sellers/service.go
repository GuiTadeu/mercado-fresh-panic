package sellers

import (
	"errors"

	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/imdario/mergo"
)

var (
	ExistsSellerCodeError = errors.New("seller code already exists")
	SellerNotFoundError   = errors.New("seller not found")
)

type Service interface {
	FindAll() ([]database.Seller, error)
	Create(cid uint64, companyName string, address string, telephone string) (database.Seller, error)
	FindOne(id uint64) (database.Seller, error)
	Update(id uint64, cid uint64, companyName string, address string, telephone string) (database.Seller, error)
	Delete(id uint64) error
}

type service struct {
	repo Repository
}

func (s service) FindAll() ([]database.Seller, error) {
	db, err := s.repo.FindAll()

	if err != nil {
		return []database.Seller{}, err
	}

	return db, err
}

func (s service) FindOne(id uint64) (database.Seller, error) {
	db, err := s.repo.FindOne(id)

	if err != nil {
		return db, SellerNotFoundError
	}
	return db, err
}

func (s service) Create(cid uint64, companyName string, address string, telephone string) (database.Seller, error) {

	isUsedCid := s.repo.FindCid(cid)

	if isUsedCid {
		return database.Seller{}, ExistsSellerCodeError
	}

	sellerData, err := s.repo.Create(cid, companyName, address, telephone)

	if err != nil {
		return database.Seller{}, err
	}
	return sellerData, nil
}

func (s service) Update(id uint64, cid uint64, companyName string, address string, telephone string) (database.Seller, error) {
	foundSeller, err := s.repo.FindOne(id)
	if err != nil {
		return database.Seller{}, SellerNotFoundError
	}

	isUsedCid := s.repo.FindCid(cid)

	if isUsedCid {
		return database.Seller{}, ExistsSellerCodeError
	}

	updatedSeller := database.Seller{
		Id:          id,
		Cid:         cid,
		CompanyName: companyName,
		Telephone:   telephone,
		Address:     address,
	}

	mergo.Merge(&foundSeller, updatedSeller, mergo.WithOverride)
	newSeller, err := s.repo.Update(foundSeller)

	if err != nil {
		return database.Seller{}, err
	}

	return newSeller, nil
}

func (s service) Delete(id uint64) error {
	err := s.repo.Delete(id)

	if err != nil {
		return SellerNotFoundError
	}
	return err
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}
