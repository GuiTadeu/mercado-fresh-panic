package localities

import (
	"database/sql"
	"log"

	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
)

const (
	CreateQuery          = "INSERT INTO localities(id, locality_name, province_id) VALUES(?, ?, ?)"
	FindLocalityId       = "SELECT id FROM localities WHERE id = ?"
	ExistsProvinceId     = "SELECT id FROM provinces WHERE id = ?"
	GetLocalityInfoQuery = `
		SELECT localities.id, localities.locality_name, COUNT(sellers.id) AS sellers_count 
		FROM localities
		LEFT JOIN sellers
		ON localities.id = sellers.locality_id
		WHERE localities.id = ?
		GROUP BY localities.id
	`
)

type LocalityInfo struct {
	LocalityId   string `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	SellersCount uint64 `json:"sellers_count"`
}

type Repository interface {
	Create(localityId string, localityName string, provinceId uint64) (database.Locality, error)
	FindLocalityId(localityId string) bool
	GetLocalityInfo(localityId string) ([]LocalityInfo, error)
	ExistsProvinceId(provinceId uint64) bool
}

type repository struct {
	db *sql.DB
}

func (r *repository) Create(localityId string, localityName string, provinceId uint64) (database.Locality, error) {
	stmt, err := r.db.Prepare(CreateQuery)

	if err != nil {
		return database.Locality{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(localityId, localityName, provinceId)

	if err != nil {
		return database.Locality{}, err
	}

	insertedLocality := createLocality(localityId, localityName, provinceId)

	return insertedLocality, nil
}

func (r *repository) FindLocalityId(localityId string) bool {

	var locality database.Locality

	err := r.db.QueryRow(FindLocalityId, localityId).Scan(&locality.Id)

	return err == nil
}

func (r *repository) ExistsProvinceId(provinceId uint64) bool {

	var locality database.Locality

	err := r.db.QueryRow(ExistsProvinceId, provinceId).Scan(&locality.ProvinceId)

	return err == nil
}

func (r *repository) GetLocalityInfo(localityId string) ([]LocalityInfo, error) {
	var localityInfos []LocalityInfo

	rows, err := r.db.Query(GetLocalityInfoQuery, localityId)

	if err != nil {
		return []LocalityInfo{}, err
	}

	for rows.Next() {
		var localityInfo LocalityInfo
		if err := rows.Scan(&localityInfo.LocalityId, &localityInfo.LocalityName, &localityInfo.SellersCount); err != nil {
			log.Println(err.Error())
			return []LocalityInfo{}, err
		}

		localityInfos = append(localityInfos, localityInfo)
	}

	return localityInfos, nil
}

func createLocality(localityId string, localityName string, provinceId uint64) database.Locality {
	return database.Locality{
		Id:         localityId,
		Name:       localityName,
		ProvinceId: provinceId,
	}
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
