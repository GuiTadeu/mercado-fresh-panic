package controller

import (
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/sellers"
)

type mockSellerService struct {
	result any
	err    error
}

func (m mockSellerService) FindAll() ([]db.Seller, error) {
	if m.err != nil {
		return []db.Seller{}, m.err
	}
	return m.result.([]db.Seller), nil
}

func (m mockSellerService) FindOne(id uint64) (db.Seller, error) {
	if m.err != nil {
		return db.Seller{}, m.err
	}
	return m.result.(db.Seller), nil
}

func (m mockSellerService) Delete(id uint64) error {
	if m.err != nil {
		return m.err
	}
	return nil
}

func (m mockSellerService) ExistsSellerCode(code string) bool {
	return m.err == sellers.ExistsSellerCodeError
}

func (m mockSellerService) Create(cid uint64, companyName string, address string, telephone string, localityId string) (db.Seller, error) {
	if m.err != nil {
		return db.Seller{}, m.err
	}
	return m.result.(db.Seller), nil
}

func (m mockSellerService) Update(id uint64, cid uint64, companyName string, address string, telephone string, localityId string) (db.Seller, error) {
	if m.err != nil {
		return db.Seller{}, m.err
	}
	return m.result.(db.Seller), nil
}
