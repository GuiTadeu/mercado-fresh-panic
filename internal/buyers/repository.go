package buyers

import (
	"database/sql"
	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"log"
)

type BuyerRepository interface {
	Create(cardNumberId, firstName, lastName string) (models.Buyer, error)
	Get(id uint64) (models.Buyer, error)
	GetAll() ([]models.Buyer, error)
	Delete(id uint64) error
	getNextId() uint64
	Update(updatedBuyer models.Buyer) (models.Buyer, error)
	ExistsBuyerCardNumberId(cardNumberId string) (bool, error)
}

type buyerRepository struct {
	db *sql.DB
}

func NewBuyerRepository(db *sql.DB) BuyerRepository {
	return &buyerRepository{
		db: db,
	}

}

func (r *buyerRepository) Create(cardNumberId, firstName, lastName string) (models.Buyer, error) {
	database := models.StorageDB

	stmt, err := database.Prepare("INSERT INTO buyer(buyer_code,card_number_d, first_name, last_name) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(cardNumberId, firstName, lastName)
	if err != nil {
		return models.Buyer{}, err
	}

	insertedId, _ := result.LastInsertId()
	buyer := models.Buyer{
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

func (r *buyerRepository) Get(id uint64) (models.Buyer, error) {
	var buyer models.Buyer
	database := models.StorageDB
	err := database.QueryRow("SELECT id,buyer_code,card_number_d, first_name, last_name FROM buyers WHERE id = ?",
		id).Scan(&buyer.Id, &buyer.CardNumberId, &buyer.FirstName, &buyer.LastName)
	if err != nil {
		log.Println(err)
		return models.Buyer{}, err
	}
	return buyer, nil
}

func (r *buyerRepository) GetAll() ([]models.Buyer, error) {

	database := models.StorageDB
	stmt, err := database.Query("SELECT * FROM buyers")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	var buyers []models.Buyer

	for stmt.Next() {
		var buyer models.Buyer

		if err = stmt.Scan(
			&buyer.Id,
			&buyer.CardNumberId,
			&buyer.FirstName,
			&buyer.LastName,
		); err != nil {
			log.Println(err)
			return nil, err
		}
		buyers = append(buyers, buyer)
	}
	return buyers, nil
}

func (r *buyerRepository) Delete(id uint64) error {
	database := models.StorageDB
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

func (r *buyerRepository) Update(updatedBuyer models.Buyer) (models.Buyer, error) {

	stmt, err := r.db.Prepare(`
	UPDATE buyers SET 
	buyers_code =?,
	card_number_id =?,
	first_name =?,
	last_name =?,
	where id =?
`)
	if err != nil {
		return models.Buyer{}, err
	}

	defer stmt.Close()
	_, err = stmt.Exec(
		updatedBuyer.Id,
		updatedBuyer.CardNumberId,
		updatedBuyer.FirstName,
		updatedBuyer.LastName,
	)

	if err != nil {
		return models.Buyer{}, err
	}
	return updatedBuyer, nil
}

func (r *buyerRepository) ExistsBuyerCardNumberId(cardNumberId string) (bool, error) {
	var buyer models.Buyer

	database := models.StorageDB

	stmt, err := database.Query("SELECT * FROM buyers WHERE card_number_id = ?", cardNumberId)

	if err != nil {
		return false, err
	}

	for stmt.Next() {

		err = stmt.Scan(
			&buyer.Id,
			&buyer.CardNumberId,
			&buyer.FirstName,
			&buyer.LastName,
		)
	}
	if err != nil {
		return false, err
	}
	return true, err
}
