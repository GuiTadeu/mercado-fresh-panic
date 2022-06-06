package sellers

import (
	"fmt"

	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

const initialId uint64 = 0

type Repository interface {
	FindAll() ([]database.Seller, error)
	Create(cid uint64, companyName string, address string, telephone string) ([]database.Seller, error)
	FindOne(id uint64) (database.Seller, error)
	UpdateAddress(id uint64, address string) (database.Seller, error)
	Delete(id uint64) error
}

type repository struct {
	db []database.Seller
}

func (r *repository) FindAll() ([]database.Seller, error) {
	return r.db, nil
}

func (r *repository) FindOne(id uint64) (database.Seller, error) {
	for _, seller := range r.db {
		if seller.Id == id {
			return seller, nil
		}
	}
	return database.Seller{}, fmt.Errorf("error: seller with id %d not found", id)
}

func (r *repository) Create(cid uint64, companyName string, address string, telephone string) ([]database.Seller, error) {
	id := r.generateId()
	data := createSeller(id, cid, companyName, address, telephone)
	r.db = append(r.db, data)
	return r.db, nil
}

func (r *repository) UpdateAddress(id uint64, address string) (database.Seller, error) {
	for idx, seller := range r.db {
		if seller.Id == id {
			r.db[idx].Address = address
			return r.db[idx], nil
		}
	}
	return database.Seller{}, fmt.Errorf("error: seller with id %d not found", id)
}

func (r *repository) Delete(id uint64) error {
	for idx, seller := range r.db {
		if seller.Id == id {
			r.db = append(r.db[:idx], r.db[idx+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error: seller with id %d not found", id)
}

func createSeller(id uint64, cid uint64, companyName string, address string, telephone string) database.Seller {
	return database.Seller{
		Id:          id,
		Cid:         cid,
		CompanyName: companyName,
		Address:     address,
		Telephone:   telephone,
	}
}

func (r repository) generateId() uint64 {
	if len(r.db) == 0 {
		return initialId
	}
	return r.db[len(r.db)-1].Id + 1
}

func NewRepository(db []database.Seller) Repository {
	return &repository{
		db: db,
	}
}
