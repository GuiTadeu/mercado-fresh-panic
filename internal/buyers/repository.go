package buyers

import (
	"database/sql"
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"log"
)

type BuyerRepository interface {
	Create(cardNumberId, firstName, lastName string) (db.Buyer, error)
	Get(id uint64) (db.Buyer, error)
	GetAll() ([]db.Buyer, error)
	Delete(id uint64) error
	getNextId() uint64
	Update(id uint64, cardNumberId string, firstName string, lastName string) (db.Buyer, error)
	ExistsBuyerCardNumberId(cardNumberId string) bool
}

func NewBuyerRepository(buyers []db.Buyer) BuyerRepository {
	return &buyerRepository{
		buyers: buyers,
	}

}

type buyerRepository struct {
	buyers []db.Buyer
}

func (r *buyerRepository) Create(cardNumberId, firstName, lastName string) (db.Buyer, error) {
	database := db.StorageDB

	stmt, err := database.Prepare("INSERT INTO buyer(buyer_code,cardNumberId, firstName, lastName) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(cardNumberId, firstName, lastName)
	if err != nil {
		return db.Buyer{}, err
	}

	insertedId, _ := result.LastInsertId()
	buyer := db.Buyer{
		Id:           uint64(insertedId),
		CardNumberId: cardNumberId,
		FirstName:    firstName,
		LastName:     lastName,
	}
	return buyer, nil
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
	var buyer db.Buyer
	database := db.StorageDB
	err := database.QueryRow("SELECT id,buyer_code,cardNumberId, firstName, lastName FROM buyers WHERE id = ?",
		id).Scan(&buyer.Id, &buyer.CardNumberId, &buyer.FirstName, &buyer.LastName)
	if err != nil {
		log.Println(err)
		return db.Buyer{}, err
	}
	return buyer, nil
}

func (r *buyerRepository) GetAll() ([]db.Buyer, error) {

	database := db.StorageDB
	stmt, err := database.Query("SELECT id,buyer_code,cardNumberId, firstName, lastName FROM buyers")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var buyers []db.Buyer

	for stmt.Next() {
		var buyer db.Buyer

		if err = stmt.Scan(
			&buyer.Id,
			&buyer.CardNumberId,
			&buyer.FirstName,
			&buyer.LastName,
		); err != nil {
			return r.buyers, err
		}
		buyers = append(buyers, buyer)
	}
	return buyers, nil
}

func (r *buyerRepository) Delete(id uint64) error {
	database := db.StorageDB
	stmt, err := database.Prepare("DELETE FROM buyers WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (r *buyerRepository) Update(id uint64, cardNumberId string, firstName string, lastName string) (db.Buyer, error) {
	uBuyer := db.Buyer{id, cardNumberId, firstName, lastName}
	for index, buyer := range r.buyers {
		if buyer.Id == id {
			r.buyers[index] = uBuyer
			return uBuyer, nil
		}
	}
	return db.Buyer{}, BuyerNotFoundError
}

func (r *buyerRepository) ExistsBuyerCardNumberId(cardNumberId string) bool {
	for _, buyer := range r.buyers {
		if buyer.CardNumberId == cardNumberId {
			return true
		}
	}
	return false
}
