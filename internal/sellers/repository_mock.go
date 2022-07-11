package sellers

import (
	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type mockSellerRepository struct {
	result          any
	err             error
	existsSellerCid bool
	getByID         database.Seller
}

func (m mockSellerRepository) FindAll() ([]database.Seller, error) {
	if m.err != nil {
		return []database.Seller{}, m.err
	}
	return m.result.([]database.Seller), nil
}

func (m mockSellerRepository) FindOne(id uint64) (database.Seller, error) {
	if (m.getByID == database.Seller{} && m.err != nil) {
		return database.Seller{}, m.err
	}
	return m.getByID, nil
}

func (m mockSellerRepository) Delete(id uint64) error {
	return m.err
}

func (m mockSellerRepository) FindCid(cid uint64) bool {
	return m.err != nil
}

func (m mockSellerRepository) Create(cid uint64, companyName string, address string, telephone string, localityId string) (database.Seller, error) {
	if m.err != nil || m.existsSellerCid {
		return database.Seller{}, m.err
	}
	return m.result.(database.Seller), nil
}

func (m mockSellerRepository) Update(seller database.Seller) (database.Seller, error) {
	if (m.result.(database.Seller) != database.Seller{}) {
		return seller, nil
	}
	return database.Seller{}, m.err
}
