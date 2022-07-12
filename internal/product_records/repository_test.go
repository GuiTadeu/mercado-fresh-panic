package productrecords

import (
	"testing"

	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/util"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func Test_Create_Ok(t *testing.T) {
	expectedProductRecords := models.ProductRecord{
		Id:             1,
		LastUpdateDate: "2021-04-04",
		PurchasePrice:  10,
		SalePrice:      15,
		ProductId:      1,
	}

	database := util.CreateDB()

	util.QueryExec(database, CREATE_PRODUCT_RECORDS_TABLE)

	repository := NewProductRecordsRepository(database)
	_, err := repository.Create("2021-04-04", 10, 15, 1)
	assert.Nil(t, err)

	productRecordFound, err := repository.Get(1)

	assert.Nil(t, err)
	assert.Equal(t, expectedProductRecords, productRecordFound)
	util.DropDB(database)
}

func Test_Create_ConnectionError(t *testing.T) {
	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCT_RECORDS_TABLE)

	repository := NewProductRecordsRepository(database)

	database.Close()
	_, err := repository.Create("", 0, 0, 0)
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_GetAll_OK(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCT_RECORDS_TABLE)

	repository := NewProductRecordsRepository(database)

	expectedCountRows := 2

	_, err := repository.Create("", 0, 0, 0)
	assert.Nil(t, err)

	_, err = repository.Create("", 0, 0, 0)
	assert.Nil(t, err)

	foundProducts, _ := repository.GetAll()
	assert.Equal(t, expectedCountRows, len(foundProducts))

	util.DropDB(database)
}

func Test_Repo_GetAll_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCT_RECORDS_TABLE)

	repository := NewProductRecordsRepository(database)

	database.Close()
	_, err := repository.GetAll()
	assert.NotNil(t, err)

	util.DropDB(database)
}

const CREATE_PRODUCT_RECORDS_TABLE = `
	CREATE TABLE "product_records"(
		id INTEGER PRIMARY KEY AUTOINCREMENT ,
		last_update_date TEXT NOT NULL,
		purchase_price TEXT NOT NULL,
		sale_price BIGINT  NOT NULL,
		product_id BIGINT NOT NULL,
		FOREIGN KEY (product_id) REFERENCES products(id)
	);
	`
