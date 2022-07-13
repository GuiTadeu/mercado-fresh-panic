package buyers

import (
	"errors"
	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/imdario/mergo"
)

var (
	ExistsBuyerCardNumberIdError = errors.New("buyer card_number_id already exists")
	BuyerNotFoundError           = errors.New("buyer not found")
)

type BuyerService interface {
	Create(cardNumberId, firstName, lastName string) (models.Buyer, error)
	Get(id uint64) (models.Buyer, error)
	GetAll() ([]models.Buyer, error)
	CountPurchaseOrdersByBuyer(id uint64) (models.CountBuyer, error)
	CountPurchaseOrdersByBuyers() ([]models.CountBuyer, error)
	Update(id uint64, cardNumberId, firstName, lastName string) (models.Buyer, error)
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

func (s *buyerService) Create(cardNumberId, firstName, lastName string) (models.Buyer, error) {
	existsBuyer, err := s.ExistsBuyerCardNumberId(cardNumberId)
	if err != nil {
		return models.Buyer{}, err
	}
	if existsBuyer {
		return models.Buyer{}, ExistsBuyerCardNumberIdError
	}

	return s.buyerRepository.Create(cardNumberId, firstName, lastName)
}

func (s *buyerService) Get(id uint64) (models.Buyer, error) {
	return s.buyerRepository.Get(id)
}

func (s *buyerService) GetAll() ([]models.Buyer, error) {
	return s.buyerRepository.GetAll()
}

func (s *buyerService) CountPurchaseOrdersByBuyer(id uint64) (models.CountBuyer, error) {
	return s.buyerRepository.CountPurchaseOrdersByBuyer(id)
}

func (s *buyerService) CountPurchaseOrdersByBuyers() ([]models.CountBuyer, error) {
	return s.buyerRepository.CountPurchaseOrdersByBuyers()
}

func (s *buyerService) Update(id uint64, cardNumberId, firstName, lastName string) (models.Buyer, error) {
	buyer, err := s.Get(id)
	if err != nil {
		return models.Buyer{}, BuyerNotFoundError
	}

	existsBuyer, err := s.ExistsBuyerCardNumberId(cardNumberId)
	if err != nil {
		return models.Buyer{}, err
	}
	if existsBuyer {
		return models.Buyer{}, ExistsBuyerCardNumberIdError
	}

	data := models.Buyer{Id: id, CardNumberId: cardNumberId, FirstName: firstName, LastName: lastName}

	err = mergo.Merge(&buyer, data, mergo.WithOverride)
	if err != nil {
		return models.Buyer{}, err
	}

	return s.buyerRepository.Update(buyer)
}

func (s *buyerService) Delete(id uint64) error {
	_, err := s.Get(id)
	if err != nil {
		return BuyerNotFoundError
	}
	return s.buyerRepository.Delete(id)
}

func (s *buyerService) ExistsBuyerCardNumberId(cardNumberId string) (bool, error) {
	return s.buyerRepository.ExistsBuyerCardNumberId(cardNumberId)
}
