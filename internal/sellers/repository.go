package sellers

import (
	"database/sql"
	"log"

	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	_ "github.com/go-sql-driver/mysql"
)

const (
	FindAllQuery = "SELECT id, cid, company_name, address, telephone, locality_id FROM sellers"
	FindOneQuery = "SELECT id, cid, company_name, address, telephone, locality_id FROM sellers WHERE id = ?"
	CreateQuery  = "INSERT INTO sellers(cid, company_name, address, telephone, locality_id) VALUES(?, ?, ?, ?, ?)"
	DeleteQuery  = "DELETE FROM sellers WHERE id = ?"
	FindCidQuery = "SELECT cid FROM sellers WHERE cid = ?"
	UpdateQuery  = `
		UPDATE 
			sellers 
		SET
			cid = ?,
			company_name = ?,
			address = ?,
			telephone = ?,
			locality_id = ?
		WHERE 
			id = ?
	`
)

type Repository interface {
	FindAll() ([]database.Seller, error)
	Create(cid uint64, companyName string, address string, telephone string, localityId string) (database.Seller, error)
	FindOne(id uint64) (database.Seller, error)
	Update(seller database.Seller) (database.Seller, error)
	Delete(id uint64) error
	FindCid(cid uint64) bool
}

type repository struct {
	db []database.Seller
}

func (r *repository) FindAll() ([]database.Seller, error) {
	var sellers []database.Seller

	db := database.StorageDB

	rows, err := db.Query(FindAllQuery)

	if err != nil {
		log.Fatal(err)
		return []database.Seller{}, err
	}

	for rows.Next() {
		var seller database.Seller
		if err := rows.Scan(&seller.Id, &seller.Cid, &seller.CompanyName, &seller.Address, &seller.Telephone, &seller.LocalityId); err != nil {
			log.Println(err.Error())
			return []database.Seller{}, err
		}

		sellers = append(sellers, seller)
	}

	return sellers, nil
}

func (r *repository) FindOne(id uint64) (database.Seller, error) {
	var seller database.Seller

	db := database.StorageDB

	err := db.QueryRow(FindOneQuery, id).Scan(&seller.Id, &seller.Cid, &seller.CompanyName, &seller.Address, &seller.Telephone, &seller.LocalityId)

	if err != nil {
		log.Println(err)
		return database.Seller{}, err
	}

	return seller, nil
}

func (r *repository) Create(cid uint64, companyName string, address string, telephone string, localityId string) (database.Seller, error) {
	db := database.StorageDB
	stmt, err := db.Prepare(CreateQuery)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	var result sql.Result

	result, err = stmt.Exec(cid, companyName, address, telephone, localityId)

	if err != nil {
		return database.Seller{}, err
	}

	insertedId, _ := result.LastInsertId()

	insertedSeller := createSeller(uint64(insertedId), cid, companyName, address, telephone, localityId)

	return insertedSeller, nil
}

func (r *repository) Update(seller database.Seller) (database.Seller, error) {
	db := database.StorageDB
	stmt, err := db.Prepare(UpdateQuery)

	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()
	_, err = stmt.Exec(
		seller.Cid,
		seller.CompanyName,
		seller.Address,
		seller.Telephone,
		seller.LocalityId,
		seller.Id,
	)

	if err != nil {
		return database.Seller{}, err
	}

	return seller, nil
}

func (r *repository) Delete(id uint64) error {
	db := database.StorageDB

	_, err := db.Exec(DeleteQuery, id)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *repository) FindCid(cid uint64) bool {

	var seller database.Seller

	db := database.StorageDB

	err := db.QueryRow(FindCidQuery, cid).Scan(&seller.Cid)

	return err == nil
}

func createSeller(id uint64, cid uint64, companyName string, address string, telephone string, localityId string) database.Seller {
	return database.Seller{
		Id:          id,
		Cid:         cid,
		CompanyName: companyName,
		Address:     address,
		Telephone:   telephone,
		LocalityId:  localityId,
	}
}

func NewRepository(db []database.Seller) Repository {
	return &repository{
		db: db,
	}
}
