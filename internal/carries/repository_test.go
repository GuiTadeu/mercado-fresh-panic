package carries

import (
    "testing"

    _ "github.com/mattn/go-sqlite3"
    "github.com/stretchr/testify/assert"

    models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
    util "github.com/GuiTadeu/mercado-fresh-panic/pkg/util"
)

func Test_Repo_Create_Ok(t *testing.T) {

    expectedLocality := models.Carrier{
    	Id:          1,
    	Cid:         "SDX",
    	CompanyName: "CTX",
    	Address:     "Rua Marselha",
    	Telephone:   "1234561238",
    	LocalityID:  "2",
    }

    database := util.CreateDB()
    util.QueryExec(database, CREATE_CARRIERS_TABLE)

    repository := NewCarrierRepository(database)
    result, err := repository.Create(expectedLocality.Cid, expectedLocality.CompanyName, expectedLocality.Address, expectedLocality.Telephone, expectedLocality.LocalityID)

    assert.Nil(t, err)
    assert.Equal(t, expectedLocality, result)

    util.DropDB(database)
}

func Test_Repo_Create_ConnectionError(t *testing.T) {

    database := util.CreateDB()
    util.QueryExec(database, CREATE_CARRIERS_TABLE)

    repository := NewCarrierRepository(database)

    database.Close()
    _, err := repository.Create("SDX", "CTX", "Rua Marselha", "1234561238", "2")

    assert.NotNil(t, err)

    util.DropDB(database)
}

func Test_Repo_ExistsCarrierCid(t *testing.T) {

    firstLocality := models.Carrier{
    	Id:          1,
    	Cid:         "SDX",
    	CompanyName: "CTX",
    	Address:     "Rua Marselha",
    	Telephone:   "1234561238",
    	LocalityID:  "2",
    }

    database := util.CreateDB()
    util.QueryExec(database, CREATE_CARRIERS_TABLE)    

    repository := NewCarrierRepository(database)
    _, err := repository.Create(firstLocality.Cid, firstLocality.CompanyName, firstLocality.Address, firstLocality.Telephone, firstLocality.LocalityID)

    assert.Nil(t, err)

    isLocalityIdFound, _ := repository.ExistsCarrierCid("SDX")

    assert.True(t, isLocalityIdFound)

    isProvinceIdFound, _ := repository.ExistsCarrierCid("dfg")

    assert.False(t, isProvinceIdFound)

    util.DropDB(database)
}

func Test_Repo_FindLocalityId_Ok(t *testing.T) {

    database := util.CreateDB()
    util.QueryExec(database, CREATE_LOCALITY_TABLE)    
    util.QueryExec(database, INSERT_LOCALITY)

    repository := NewCarrierRepository(database)   

    result := repository.FindLocalityId("11065001")
  
    assert.True(t, result )

	result = repository.FindLocalityId("11065001TE")
  
    assert.False(t, result )

    util.DropDB(database)
}

func Test_Repo_GetCarrierInfo_Ok(t *testing.T) {
    localityInfo := []CarrierInfo{
        {
            LocalityId:   "11065001",           
            CarriesCount: 1,
			LocalityName: "Santos",
        },
    }

    database := util.CreateDB()
    util.QueryExec(database, CREATE_LOCALITY_TABLE)
    util.QueryExec(database, CREATE_CARRIERS_TABLE)
    util.QueryExec(database, INSERT_LOCALITY)
    repository := NewCarrierRepository(database)

    _, err := repository.Create("SDX", "CTX", "Rua Marselha", "1234561238", "11065001")

    assert.Nil(t, err)

    result, err := repository.GetAllCarrierInfo("11065001")
    assert.Nil(t, err)
    assert.Equal(t, localityInfo, result)

    util.DropDB(database)
}

func Test_Repo_GetCarrierInfo_InternalError(t *testing.T) {

    database := util.CreateDB()
    util.QueryExec(database, CREATE_LOCALITY_TABLE)
    util.QueryExec(database, CREATE_CARRIERS_TABLE)
    util.QueryExec(database, INSERT_LOCALITY)
    repository := NewCarrierRepository(database)
	
	database.Close()

    _, err := repository.GetAllCarrierInfo("11065001")
    assert.Error(t, err)

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

const CREATE_CARRIERS_TABLE = `
    CREATE TABLE "carriers" (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        cid BIGINT (64) UNIQUE NOT NULL,
        company_name TEXT NOT NULL,
        address TEXT NOT NULL,
        telephone TEXT NOT NULL,
        locality_id TEXT NOT NULL,
        FOREIGN KEY (locality_id) REFERENCES localities(id)
    );
`

const INSERT_LOCALITY = `
	INSERT INTO localities(id, locality_name, province_id)
	VALUES  ("11065001", "Santos", 1);
`