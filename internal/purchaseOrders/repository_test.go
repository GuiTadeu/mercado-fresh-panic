package purchaseOrders

import (
	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/util"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Create_Ok(t *testing.T) {
	expectedPurchaseOrders := models.PurchaseOrder{
		Id:              1,
		OrderNumber:     "1",
		OrderDate:       "2022-07-12",
		TrackingCode:    "1",
		BuyerId:         1,
		OrderStatusId:   1,
		ProductRecordId: 1,
	}

	database := util.CreateDB()

	util.QueryExec(database, CREATE_PURCHASE_ORDERS_TABLE)

	repository := NewPurchaseOrdersRepository(database)
	_, err := repository.Create("1", "2022-07-12", "1", 1, 1, 1)
	assert.Nil(t, err)

	purchaseOrderFounded, err := repository.Get(1)

	assert.Nil(t, err)
	assert.Equal(t, expectedPurchaseOrders, purchaseOrderFounded)
	util.DropDB(database)
}

func Test_Create_ConnectionError(t *testing.T) {
	database := util.CreateDB()
	util.QueryExec(database, CREATE_PURCHASE_ORDERS_TABLE)

	repository := NewPurchaseOrdersRepository(database)

	database.Close()
	_, err := repository.Create("", "", "", 0, 0, 0)
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Get_OK(t *testing.T) {

	expectedPurchaseOrders := models.PurchaseOrder{
		Id:              1,
		OrderNumber:     "1",
		OrderDate:       "2022-07-12",
		TrackingCode:    "1",
		BuyerId:         1,
		OrderStatusId:   1,
		ProductRecordId: 1,
	}
	database := util.CreateDB()
	util.QueryExec(database, CREATE_PURCHASE_ORDERS_TABLE)

	repository := NewPurchaseOrdersRepository(database)

	_, err := repository.Create("1", "2022-07-12", "1", 1, 1, 1)
	assert.Nil(t, err)

	purchaseOrderFounded, err := repository.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedPurchaseOrders, purchaseOrderFounded)

	util.DropDB(database)
}

func Test_Repo_Get_ShouldReturnEmptyPurchaseOrderWhenIdNotExists(t *testing.T) {
	database := util.CreateDB()
	util.QueryExec(database, CREATE_PURCHASE_ORDERS_TABLE)

	repository := NewPurchaseOrdersRepository(database)
	_, err := repository.Create("1", "2022-07-12", "1", 1, 1, 1)
	assert.Nil(t, err)

	foundPurchaseOrder, _ := repository.Get(10)
	assert.Empty(t, foundPurchaseOrder)

	util.DropDB(database)
}

func Test_Repo_Get_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PURCHASE_ORDERS_TABLE)

	repository := NewPurchaseOrdersRepository(database)

	database.Close()
	_, err := repository.Get(1)
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_ExistsId_OK(t *testing.T) {
	database := util.CreateDB()
	util.QueryExec(database, CREATE_PURCHASE_ORDERS_TABLE)

	repository := NewPurchaseOrdersRepository(database)

	_, err := repository.Create("1", "2022-07-12", "1", 1, 1, 1)

	existId := repository.ExistsBuyerId(1)
	assert.False(t, existId)
	assert.Nil(t, err)

	util.DropDB(database)
}

func Test_Repo_ExistsId_ConnectionError(t *testing.T) {
	database := util.CreateDB()
	util.QueryExec(database, CREATE_PURCHASE_ORDERS_TABLE)

	repository := NewPurchaseOrdersRepository(database)

	database.Close()
	err := repository.ExistsBuyerId(1)
	assert.NotNil(t, err)

	util.DropDB(database)
}

const CREATE_PURCHASE_ORDERS_TABLE = `
	CREATE TABLE "purchase_orders"(
		id INTEGER PRIMARY KEY AUTOINCREMENT ,
		order_number TEXT NOT NULL,
		order_date TEXT NOT NULL,
		tracking_code TEXT  NOT NULL,
		buyer_id BIGINT  NOT NULL,
		order_status_id BIGINT  NOT NULL,
		product_record_id BIGINT  NOT NULL,
		FOREIGN KEY (buyer_id) REFERENCES buyers(id),
		FOREIGN KEY (order_status_id) REFERENCES order_status(id),
		FOREIGN KEY (product_record_id) REFERENCES product_records(id)
	);
`
