package buyers

import db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"

type BuyerService interface {
	Create(cardNumberId, firstName, lastName string) (db.Buyer, error)
	Get(id uint64) (db.Buyer, error)
	GetAll() ([]db.Buyer, error)
	//Update()
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
	return s.buyerRepository.Create(cardNumberId, firstName, lastName)
}

func (s *buyerService) Get(id uint64) (db.Buyer, error) {
	return s.buyerRepository.Get(id)
}

func (s *buyerService) GetAll() ([]db.Buyer, error) {
	return s.buyerRepository.GetAll()
}

func (s *buyerService) Delete(id uint64) error {
	return s.buyerRepository.Delete(id)
}
