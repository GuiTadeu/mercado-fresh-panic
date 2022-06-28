package controller

import db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"

type mockBuyerService struct {
	result any
	err    error
}

func (m mockBuyerService) Create(
	cardNumberId, firstName, lastName string,
) (db.Buyer, error) {
	if m.err != nil {
		return db.Buyer{}, m.err
	}
	return m.result.(db.Buyer), nil
}

func (m mockBuyerService) Get(id uint64) (db.Buyer, error) {
	if m.err != nil {
		return db.Buyer{}, m.err
	}
	return m.result.(db.Buyer), nil
}

func (m mockBuyerService) GetAll() ([]db.Buyer, error) {
	if m.err != nil {
		return []db.Buyer{}, m.err
	}
	return m.result.([]db.Buyer), nil
}

func (m mockBuyerService) Update(
	id uint64, cardNumberId, firstName, lastName string,
) (db.Buyer, error) {
	if m.err != nil {
		return db.Buyer{}, m.err
	}
	return m.result.(db.Buyer), nil
}

func (m mockBuyerService) Delete(id uint64) error {
	if m.err != nil {
		return m.err
	}
	return nil
}
