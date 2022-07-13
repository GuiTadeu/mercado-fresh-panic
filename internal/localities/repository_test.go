package localities

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"

	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	util "github.com/GuiTadeu/mercado-fresh-panic/pkg/util"
)

func Test_Repo_Create_Ok(t *testing.T) {

	expectedLocality := models.Locality{
		Id:         "11065001",
		Name:       "Santos",
		ProvinceId: 1,
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_LOCALITY_TABLE)

	repository := NewRepository(database)
	result, err := repository.Create(expectedLocality.Id, expectedLocality.Name, expectedLocality.ProvinceId)

	assert.Nil(t, err)
	assert.Equal(t, expectedLocality, result)

	util.DropDB(database)
}

func Test_Repo_Create_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_LOCALITY_TABLE)

	repository := NewRepository(database)

	database.Close()
	_, err := repository.Create("11065001", "Santos", 1)

	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_ExistsProvinceId(t *testing.T) {

	firstLocality := models.Locality{
		Id:         "11065001",
		Name:       "Santos",
		ProvinceId: 1,
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_LOCALITY_TABLE)
	util.QueryExec(database, CREATE_PROVINCE_TABLE)
	util.QueryExec(database, INSERT_PROVINCE)

	repository := NewRepository(database)
	_, err := repository.Create(firstLocality.Id, firstLocality.Name, firstLocality.ProvinceId)

	assert.Nil(t, err)

	isLocalityIdFound := repository.FindLocalityId("11065001")

	assert.True(t, isLocalityIdFound)

	isProvinceIdFound := repository.ExistsProvinceId(1)

	assert.True(t, isProvinceIdFound)

	util.DropDB(database)
}

func Test_Repo_GetLocalityInfo_Ok(t *testing.T) {

	localityInfo := []LocalityInfo{
		{
			LocalityId:   "11065001",
			LocalityName: "Santos",
			SellersCount: 1,
		},
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_LOCALITY_TABLE)
	util.QueryExec(database, CREATE_SELLERS_TABLE)
	util.QueryExec(database, INSERT_SELLER)

	repository := NewRepository(database)
	_, err := repository.Create("11065001", "Santos", 1)

	assert.Nil(t, err)

	result, err := repository.GetLocalityInfo("11065001")

	assert.Nil(t, err)
	assert.Equal(t, localityInfo, result)

	util.DropDB(database)
}

func Test_Repo_GetLocalityInfo_InternalError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_LOCALITY_TABLE)
	util.QueryExec(database, CREATE_SELLERS_TABLE)
	util.QueryExec(database, INSERT_SELLER)

	repository := NewRepository(database)
	_, err := repository.Create("11065001", "Santos", 1)

	assert.Nil(t, err)

	database.Close()
	_, err = repository.GetLocalityInfo("11065001")

	assert.NotNil(t, err)

	util.DropDB(database)
}

const CREATE_LOCALITY_TABLE = `
	CREATE TABLE "localities" (
		id TEXT NOT NULL,
		locality_name TEXT NOT NULL,
		province_id INTEGER NOT NULL,
		PRIMARY KEY ("id")
	);
	
`

const CREATE_PROVINCE_TABLE = `
	CREATE TABLE "provinces"(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		province_name TEXT NOT NULL,
		id_country_fk BIGINT  NOT NULL
	);	
`

const INSERT_PROVINCE = `
	INSERT INTO provinces(province_name, id_country_fk) VALUES ("São Paulo", 1);
`

const CREATE_SELLERS_TABLE = `
	CREATE TABLE "sellers" (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		cid BIGINT (64) UNIQUE NOT NULL,
		company_name TEXT NOT NULL,
		address TEXT NOT NULL,
		telephone TEXT NOT NULL,
		locality_id TEXT NOT NULL,
		FOREIGN KEY (locality_id) REFERENCES localities(id)
	);
`

const INSERT_SELLER = `
	INSERT INTO sellers(cid, company_name, address, telephone, locality_id)  
	VALUES  (1, "Nike", "Rua Pedro Américo, 212", "13990984533", "11065001");
`
