package inboundorders

import (
	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/util"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Repo_Create_Ok(t *testing.T) {
	expectedInboundOrders := models.InboundOrder{
		Id:             1,
		OrderDate:      "2021-04-04",
		OrderNumber:    "order#1",
		EmployeeId:     1,
		ProductBatchId: 1,
		WarehouseId:    1,
	}

	database := util.CreateDB()

	util.QueryExec(database, CREATE_INBOUND_ORDERS_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create("2021-04-04", "order#1", 1, 1, 1)
	assert.Nil(t, err)

	inboundOrderFounded, err := repository.Get(1)

	assert.Nil(t, err)
	assert.Equal(t, expectedInboundOrders, inboundOrderFounded)
	util.DropDB(database)
}

func Test_Repo_Create_ConnectionError(t *testing.T) {
	database := util.CreateDB()
	util.QueryExec(database, CREATE_INBOUND_ORDERS_TABLE)
	
	repository := NewRepository(database)

	database.Close()
	_, err := repository.Create("", "", 0, 0, 0)
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_Get_ConnectionError(t *testing.T) {
	database := util.CreateDB()
	util.QueryExec(database, CREATE_INBOUND_ORDERS_TABLE)
	
	repository := NewRepository(database)

	database.Close()
	_, err := repository.Get(1)
	assert.NotNil(t, err)

	util.DropDB(database)
}



const CREATE_INBOUND_ORDERS_TABLE = `
	CREATE TABLE "inbound_orders"(
		id INTEGER PRIMARY KEY AUTOINCREMENT ,
		order_date TEXT NOT NULL,
		order_number TEXT NOT NULL,
		employee_id BIGINT  NOT NULL,
		product_batch_id BIGINT  NOT NULL,
		warehouse_id BIGINT  NOT NULL,
		FOREIGN KEY (employee_id) REFERENCES employees(id),
		FOREIGN KEY (product_batch_id) REFERENCES product_batches(id),
		FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
	);
`
