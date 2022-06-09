package buyers

import (
	"errors"
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/imdario/mergo"
)

var (
	ExistsBuyerCardNumberIdError = errors.New("buyer card_number_id already exists")
	BuyerNotFoundError           = errors.New("buyer not found")
)

type BuyerService interface {
	Create(cardNumberId, firstName, lastName string) (db.Buyer, error)
	Get(id uint64) (db.Buyer, error)
	GetAll() ([]db.Buyer, error)
	Update(id uint64, cardNumberId, firstName, lastName string) (db.Buyer, error)
	Delete(id uint64) error
}

type buyerService struct {
	buyerRepository BuyerRepository
}

func NewBuyerService(r BuyerRepository) BuyerService {
	return &buyerService{
		buyerRepository: r,
	}

}

func (s *buyerService) Create(cardNumberId, firstName, lastName string) (db.Buyer, error) {
	if s.existsBuyerCardNumberId(cardNumberId) {
		return db.Buyer{}, ExistsBuyerCardNumberIdError
	}
	return s.buyerRepository.Create(cardNumberId, firstName, lastName)
}

func (s *buyerService) Get(id uint64) (db.Buyer, error) {
	return s.buyerRepository.Get(id)
}

func (s *buyerService) GetAll() ([]db.Buyer, error) {
	return s.buyerRepository.GetAll()
}

func (s *buyerService) Update(id uint64, cardNumberId, firstName, lastName string) (db.Buyer, error) {
	buyer, err := s.Get(id)
	if err != nil {
		return db.Buyer{}, BuyerNotFoundError
	}

	if s.existsBuyerCardNumberId(cardNumberId) {
		return db.Buyer{}, ExistsBuyerCardNumberIdError
	}

	data := db.Buyer{id, cardNumberId, firstName, lastName}

	err = mergo.Merge(&buyer, data, mergo.WithOverride)
	if err != nil {
		return db.Buyer{}, err
	}

	return s.buyerRepository.Update(buyer.Id, buyer.CardNumberId, buyer.FirstName, buyer.LastName)
}

func (s *buyerService) Delete(id uint64) error {
	_, err := s.Get(id)
	if err != nil {
		return BuyerNotFoundError
	}
	return s.buyerRepository.Delete(id)
}

func (s *buyerService) existsBuyerCardNumberId(cardNumberId string) bool {
	return s.buyerRepository.ExistsBuyerCardNumberId(cardNumberId)
}
