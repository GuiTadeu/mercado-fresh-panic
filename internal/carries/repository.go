package carries

import (
	"database/sql"

	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type CarrierRepository interface {
	Create(cid string, companyName string, address string, telephone string, localityId string) (database.Carrier, error)
	ExistsCarrierCid(cid string) (bool, error)	
	GetAllCarrierInfo() ([]database.Carrier, error)
}

type carrierRepository struct {
	db *sql.DB
}

func NewCarrierRepository(db *sql.DB) CarrierRepository {
	return &carrierRepository{
		db: db,
	}
}

func (r *carrierRepository) Create(cid string, companyName string, address string, telephone string, localityId string) (database.Carrier, error) {
	stmt, err := r.db.Prepare("INSERT INTO carriers(cid, company_name, address, telephone, locality_id) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return database.Carrier{}, err
	}

	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(cid, companyName, address, telephone, localityId)
	if err != nil {
		return database.Carrier{}, err
	}

	insertedId, _ := result.LastInsertId()
	carrier := database.Carrier{
		Id:          uint64(insertedId),
		Cid:         cid,
		CompanyName: companyName,
		Address:     address,
		Telephone:   telephone,
		LocalityID:  localityId,
	}

	return carrier, nil
}

func (r *carrierRepository) ExistsCarrierCid(cid string) (bool, error) {

	var carrier database.Carrier

	rows, err := r.db.Query("SELECT * FROM carriers WHERE cid = ?", )

	if err != nil {
		return false, err
	}

	for rows.Next() {
		err := rows.Scan(&carrier.Id, &carrier.Cid, &carrier.CompanyName, &carrier.Address, &carrier.Telephone, &carrier.LocalityID)

		if err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil	
}

func (r *carrierRepository) GetAllCarrierInfo() ([]database.Carrier, error) {
	stmt, err := r.db.Query("SELECT id, cid, company_name, address, telephone, locality_id FROM carriers")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var carriers []database.Carrier

	for stmt.Next() {
		var carrier database.Carrier

		if err = stmt.Scan(
			&carrier.Id,
			&carrier.Cid,
			&carrier.CompanyName,
			&carrier.Address,
			&carrier.Telephone,
			&carrier.LocalityID,
		); err != nil {
			return nil, err
		}

		carriers = append(carriers, carrier)
	}

	return carriers, nil
}