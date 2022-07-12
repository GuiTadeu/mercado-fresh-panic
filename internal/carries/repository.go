package carries

import (
	"database/sql"

	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

type CarrierInfo struct {
	LocalityId   string `json:"locality_id"`
	CarriesCount uint64 `json:"carries_count"`
	LocalityName string `json:"locality_name"`
}

type CarrierRepository interface {
	Create(cid string, companyName string, address string, telephone string, localityId string) (database.Carrier, error)
	ExistsCarrierCid(cid string) (bool, error)
	GetAllCarrierInfo(id string) ([]CarrierInfo, error)
	FindLocalityId(localityId string) bool
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

	rows, err := r.db.Query("SELECT id FROM carriers WHERE cid = ?", cid)

	if err != nil {
		return false, err
	}

	for rows.Next() {
		err := rows.Scan(&carrier.Id)

		if err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}

func (r *carrierRepository) GetAllCarrierInfo(id string) ([]CarrierInfo, error) {
	var carrierInfo []CarrierInfo	

	stmt, err := r.db.Query(` SELECT localities.id, locality_name, COUNT(carriers.id)
	FROM localities
	LEFT JOIN carriers
	ON localities.id = carriers.locality_id
	WHERE localities.id = ?
	GROUP BY (localities.id);`, id)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	for stmt.Next() {
		var carrier CarrierInfo

		if err = stmt.Scan(
			&carrier.LocalityId,
			&carrier.LocalityName,
			&carrier.CarriesCount,						
		); err != nil {
			return []CarrierInfo{}, err
		}

		carrierInfo = append(carrierInfo, carrier)
	}

	return carrierInfo, nil
}

func (r *carrierRepository) FindLocalityId(localityId string) bool {

    var locality database.Locality

    err := r.db.QueryRow("SELECT id FROM localities WHERE id = ?", localityId).Scan(&locality.Id)

    return err == nil
}