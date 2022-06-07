package sellers

import (
	"fmt"
	"net/http"

	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/imdario/mergo"
)

type Service interface {
	FindAll() ([]database.Seller, int, error)
	Create(cid uint64, companyName string, address string, telephone string) ([]database.Seller, int, error)
	FindOne(id uint64) (database.Seller, int, error)
	Update(id uint64, cid uint64, companyName string, address string, telephone string) (database.Seller, int, error)
	Delete(id uint64) (int, error)
}

type service struct {
	repo Repository
}

func (s service) FindAll() ([]database.Seller, int, error) {
	db, err := s.repo.FindAll()

	if err != nil {
		return []database.Seller{}, http.StatusInternalServerError, fmt.Errorf("error on loading data")
	}

	return db, http.StatusOK, err
}

func (s service) FindOne(id uint64) (database.Seller, int, error) {
	db, err := s.repo.FindOne(id)

	if err != nil {
		return db, http.StatusNotFound, err
	}
	return db, http.StatusOK, err
}

func (s service) Create(cid uint64, companyName string, address string, telephone string) ([]database.Seller, int, error) {

	isUsedCid := s.repo.FindCid(cid)

	if isUsedCid {
		return []database.Seller{}, http.StatusConflict, fmt.Errorf("seller with this cid already exists")
	}

	sellersData, err := s.repo.Create(cid, companyName, address, telephone)

	if err != nil {
		return []database.Seller{}, http.StatusInternalServerError, fmt.Errorf("error on writing data")
	}
	return sellersData, http.StatusCreated, nil
}

func (s service) Update(id uint64, cid uint64, companyName string, address string, telephone string) (database.Seller, int, error) {
	foundSeller, err := s.repo.FindOne(id)
	if err != nil {
		return database.Seller{}, http.StatusNotFound, fmt.Errorf("error: seller not found")
	}

	isUsedCid := s.repo.FindCid(cid)

	if isUsedCid {
		return database.Seller{}, http.StatusConflict, fmt.Errorf("seller with this cid already exists")
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
		return database.Seller{}, http.StatusInternalServerError, fmt.Errorf("error: internal server error")
	}

	return newSeller, http.StatusOK, nil
}

func (s service) Delete(id uint64) (int, error) {
	err := s.repo.Delete(id)

	if err != nil {
		return http.StatusNotFound, err
	}
	return http.StatusNoContent, err
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}
