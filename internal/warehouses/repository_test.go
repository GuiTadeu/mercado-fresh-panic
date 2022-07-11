package warehouses

import (
	"testing"

	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	util "github.com/GuiTadeu/mercado-fresh-panic/pkg/util"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func Test_Repo_Warehouse_Create_Ok(t *testing.T) {

	expectedWarehouse := models.Warehouse{
		Id:                 1,
		Code:               "CD",
		Address:            "spc",
		Telephone:          "11998723778",
		MinimunCapacity:    5,
		MinimumTemperature: 2.0,
		LocalityID:         "1",
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_WAREHOUSES_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create("CD", "spc", "11998723778", 5, 2.0, "1")
	assert.Nil(t, err)

	foundWarehouse, err := repository.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedWarehouse, foundWarehouse)

	util.DropDB(database)
}

func Test_Repo_Create_Warehouse_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_WAREHOUSES_TABLE)

	repository := NewRepository(database)

	database.Close()
	_, err := repository.Get(1)
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_GetAll_Warehouse_Ok(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_WAREHOUSES_TABLE)

	repository := NewRepository(database)

	expectedCountRows := 2

	_, err := repository.Create("CD", "spc", "11998723778", 5, 2.0, "1")
	assert.Nil(t, err)

	_, err = repository.Create("LP", "disco", "11976723778", 4, 3.0, "1")
	assert.Nil(t, err)

	foundWarehouses, _ := repository.GetAll()
	assert.Equal(t, expectedCountRows, len(foundWarehouses))

	util.DropDB(database)

}

func Test_Repo_GetAll_Warehouse_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_WAREHOUSES_TABLE)

	repository := NewRepository(database)

	database.Close()
	_, err := repository.GetAll()
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_Update_Warehouse_Ok(t *testing.T) {

	expectedOldWarehouse := models.Warehouse{
		Id:                 1,
		Code:               "CD",
		Address:            "sao paulo",
		Telephone:          "123456789",
		MinimunCapacity:    1,
		MinimumTemperature: 1.0,
		LocalityID:         "1",
	}

	expectedUpdateWarehouse := models.Warehouse{
		Id:                 1,
		Code:               "LP",
		Address:            "bernambuco",
		Telephone:          "123446789",
		MinimunCapacity:    1,
		MinimumTemperature: 1.0,
		LocalityID:         "1",
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_WAREHOUSES_TABLE)

	repository := NewRepository(database)

	_, err := repository.Create(
		expectedOldWarehouse.Code,
		expectedOldWarehouse.Address,
		expectedOldWarehouse.Telephone,
		expectedOldWarehouse.MinimunCapacity,
		expectedOldWarehouse.MinimumTemperature,
		expectedOldWarehouse.LocalityID,
	)
	assert.Nil(t, err)

	foundWarehouse, _ := repository.Get(1)
	assert.Equal(t, expectedOldWarehouse, foundWarehouse)

	_, err = repository.Update(expectedUpdateWarehouse)
	assert.Nil(t, err)

	foundWarehouse, _ = repository.Get(1)
	assert.Equal(t, expectedUpdateWarehouse, foundWarehouse)

	util.DropDB(database)
}

func Test_Repo_Update_Warehouse_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_WAREHOUSES_TABLE)

	repository := NewRepository(database)

	database.Close()
	_, err := repository.Update(models.Warehouse{})
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_Delete_Warehouse_Ok(t *testing.T) {

	expectedWarehouse := models.Warehouse{
		Id:                 1,
		Code:               "LP",
		Address:            "viela",
		Telephone:          "234567891",
		MinimunCapacity:    2,
		MinimumTemperature: 7.0,
		LocalityID:         "1",
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_WAREHOUSES_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create("LP", "viela", "234567891", 2, 7.0, "1")
	assert.Nil(t, err)

	foundWarehouse, err := repository.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedWarehouse, foundWarehouse)

	err = repository.Delete(1)
	assert.Nil(t, err)

	foundWarehouse, _ = repository.Get(1)
	assert.Empty(t, foundWarehouse)

	util.DropDB(database)
}

func Test_Repo_Delete_Warehouse_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_WAREHOUSES_TABLE)

	repository := NewRepository(database)

	database.Close()
	err := repository.Delete(1)
	assert.NotNil(t, err)

	util.DropDB(database)

}

func Test_Repo_ExistsWarehouseCode_Warehouse_Ok(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_WAREHOUSES_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create("LP", "viela", "234567891", 2, 7.0, "1")
	assert.Nil(t, err)

	existsWarehouse, err := repository.ExistsWarehouseCode("LP")
	assert.Nil(t, err)
	assert.True(t, existsWarehouse)

	util.DropDB(database)
}

func Test_Repo_ExistsWarehouseCode_Warehouse_NotFound(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_WAREHOUSES_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create("LP", "viela", "234567891", 2, 7.0, "1")
	assert.Nil(t, err)

	existsWarehouse, err := repository.ExistsWarehouseCode("MINI")
	assert.Nil(t, err)
	assert.False(t, existsWarehouse)

	util.DropDB(database)
}

func Test_Repo_ExistsWarehouseCode_Warehouse_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_WAREHOUSES_TABLE)

	repository := NewRepository(database)

	database.Close()
	_,err := repository.ExistsWarehouseCode("")
	assert.NotNil(t, err)

	util.DropDB(database)
}

const CREATE_WAREHOUSES_TABLE = `
	CREATE TABLE "warehouses" (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		address TEXT NOT NULL,
		telephone TEXT NOT NULL,
		warehouse_code TEXT UNIQUE NOT NULL,
		minimum_capacity BIGINT NOT NULL,
		minimum_temperature DECIMAL(19, 2) NOT NULL,
		locality_id TEXT NOT NULL,
		FOREIGN KEY (locality_id) REFERENCES localities(id)	
	);
`
