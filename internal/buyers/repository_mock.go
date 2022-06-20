package buyers

import db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"

type mockBuyerRepository struct {
	result                  any
	err                     error
	existsBuyerCardNumberId bool
	getById                 db.Buyer
}

func (m mockBuyerRepository) GetAll() ([]db.Buyer, error) {
	if m.err != nil {
		return []db.Buyer{}, m.err
	}
	return m.result.([]db.Buyer), nil
}

func (m mockBuyerRepository) Get(id uint64) (db.Buyer, error) {
	if (m.getById == db.Buyer{} && m.err != nil) {
		return db.Buyer{}, m.err
	}
	return m.result.(db.Buyer), nil
}

func (m mockBuyerRepository) Delete(id uint64) error {
	if m.err != nil {
		return m.err
	}
	return nil
}

func (m mockBuyerRepository) Create(cardNumberId, firstName, lastName string) (db.Buyer, error) {
	if m.err != nil || m.existsBuyerCardNumberId {
		return db.Buyer{}, m.err
	}
	return m.result.(db.Buyer), nil
}

func (m mockBuyerRepository) getNextId() uint64 {
	return m.result.(uint64)
}

func (m mockBuyerRepository) Update(id uint64, cardNumberId string, firstName string, lastName string) (db.Buyer, error) {
	if (m.result.(db.Buyer) != db.Buyer{}) {
		return db.Buyer{id, cardNumberId, firstName, lastName}, m.err
	}
	return db.Buyer{}, nil
}

func (m mockBuyerRepository) ExistsBuyerCardNumberId(cardNumberId string) bool {
	return m.existsBuyerCardNumberId
}
