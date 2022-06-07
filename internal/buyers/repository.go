package buyers

import (
	"errors"
	"fmt"
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type BuyerRepository interface {
	Create(cardNumberId, firstName, lastName string) (db.Buyer, error)
	Get(id uint64) (db.Buyer, error)
	GetAll() ([]db.Buyer, error)
	//Update()
	Delete(id uint64) error
	getNextId() uint64
	Update(id uint64, cardNumberId string, firstName string, lastName string) (db.Buyer, error)
}

func NewBuyerRepository(buyers []db.Buyer) BuyerRepository {
	return &buyerRepository{
		buyers,
	}

}

type buyerRepository struct {
	buyers []db.Buyer
}

func (r *buyerRepository) Create(cardNumberId, firstName, lastName string) (db.Buyer, error) {
	b := db.Buyer{
		Id:           r.getNextId(),
		CardNumberId: cardNumberId,
		FirstName:    firstName,
		LastName:     lastName,
	}
	r.buyers = append(r.buyers, b)
	return b, nil
}

func (r *buyerRepository) getNextId() uint64 {
	buyers, err := r.GetAll()
	if err != nil {
		return 1
	}

	if len(buyers) == 0 {
		return 1
	}

	return buyers[len(buyers)-1].Id + 1
}

func (r *buyerRepository) Get(id uint64) (db.Buyer, error) {
	for _, buyer := range r.buyers {
		if buyer.Id == id {
			return buyer, nil
		}
	}
	return db.Buyer{}, errors.New("buyer not found")
}

func (r *buyerRepository) GetAll() ([]db.Buyer, error) {
	return r.buyers, nil
}

func (r *buyerRepository) Delete(id uint64) error {
	for i := range r.buyers {
		if r.buyers[i].Id == id {
			r.buyers = append(r.buyers[:1], r.buyers[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Buyer not found")
}

func (r *buyerRepository) Update(id uint64, cardNumberId string, firstName string, lastName string) (db.Buyer, error) {
	uBuyer := db.Buyer{id, cardNumberId, firstName, lastName}
	for index, buyer := range r.buyers {
		if buyer.Id == id {
			r.buyers[index] = uBuyer
			return uBuyer, nil
		}
	}
	return db.Buyer{}, fmt.Errorf("Buyer not found")
}
