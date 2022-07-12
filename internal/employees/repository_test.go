package employees

import (
	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	util "github.com/GuiTadeu/mercado-fresh-panic/pkg/util"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Repo_Create_Ok(t *testing.T) {

	expectedEmployee := models.Employee{
		Id:           1,
		CardNumberId: "1",
		FirstName:    "Jack",
		LastName:     "Sparrow",
		WarehouseId:  1,
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_EMPLOYEE_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create("1", "Jack", "Sparrow", 1)
	assert.Nil(t, err)

	foundEmployee, err := repository.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedEmployee, foundEmployee)

	util.DropDB(database)
}

func Test_Repo_Create_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_EMPLOYEE_TABLE)

	repository := NewRepository(database)

	database.Close()
	_, err := repository.Create("", "", "", 0)
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_Get_OK(t *testing.T) {

	expectedEmployee := models.Employee{
		Id:           1,
		CardNumberId: "1",
		FirstName:    "Jack",
		LastName:     "Sparrow",
		WarehouseId:  1,
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_EMPLOYEE_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create("1", "Jack", "Sparrow", 1)
	assert.Nil(t, err)

	foundEmployee, err := repository.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedEmployee, foundEmployee)

	util.DropDB(database)
}

func Test_Repo_Get_ShouldReturnEmptyEmployeeWhenIdNotExists(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_EMPLOYEE_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create("1", "Jack", "Sparrow", 1)
	assert.Nil(t, err)

	foundEmployee, err := repository.Get(2)
	assert.Empty(t, foundEmployee)

	util.DropDB(database)
}

func Test_Repo_Get_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_EMPLOYEE_TABLE)

	repository := NewRepository(database)

	database.Close()
	_, err := repository.Get(1)
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_GetAll_OK(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_EMPLOYEE_TABLE)

	repository := NewRepository(database)

	expectedCountRows := 2

	_, err := repository.Create("1", "Jack", "Sparrow", 1)
	assert.Nil(t, err)

	_, err = repository.Create("2", "Will", "Turner", 1)
	assert.Nil(t, err)

	foundEmployees, err := repository.GetAll()
	assert.Equal(t, expectedCountRows, len(foundEmployees))

	util.DropDB(database)
}

func Test_Repo_GetAll_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_EMPLOYEE_TABLE)

	repository := NewRepository(database)

	database.Close()
	_, err := repository.GetAll()
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_Update_OK(t *testing.T) {

	expectedOldEmployee := models.Employee{
		Id:           1,
		CardNumberId: "1",
		FirstName:    "Jack",
		LastName:     "Sparrow",
		WarehouseId:  1,
	}

	expectedNewEmployee := models.Employee{
		Id:           1,
		CardNumberId: "1",
		FirstName:    "Charles",
		LastName:     "do Bronx",
		WarehouseId:  1,
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_EMPLOYEE_TABLE)

	repository := NewRepository(database)

	_, err := repository.Create(
		expectedOldEmployee.CardNumberId,
		expectedOldEmployee.FirstName,
		expectedOldEmployee.LastName,
		expectedOldEmployee.WarehouseId,
	)
	assert.Nil(t, err)

	foundEmployee, _ := repository.Get(1)
	assert.Equal(t, expectedOldEmployee, foundEmployee)

	_, err = repository.Update(expectedNewEmployee)
	assert.Nil(t, err)

	foundEmployee, _ = repository.Get(1)
	assert.Equal(t, expectedNewEmployee, foundEmployee)

	util.DropDB(database)
}

func Test_Repo_Update_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_EMPLOYEE_TABLE)

	repository := NewRepository(database)

	database.Close()
	_, err := repository.Update(models.Employee{})
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_Delete_Ok(t *testing.T) {

	expectedEmployee := models.Employee{
		Id:           1,
		CardNumberId: "1",
		FirstName:    "Jack",
		LastName:     "Sparrow",
		WarehouseId:  1,
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_EMPLOYEE_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create("1", "Jack", "Sparrow", 1)
	assert.Nil(t, err)

	foundEmployee, err := repository.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedEmployee, foundEmployee)

	err = repository.Delete(1)
	assert.Nil(t, err)

	foundEmployee, err = repository.Get(1)
	assert.Empty(t, foundEmployee)

	util.DropDB(database)
}

func Test_Repo_Delete_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_EMPLOYEE_TABLE)

	repository := NewRepository(database)

	database.Close()
	err := repository.Delete(1)
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_ExistsEmployeeCode_OK(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_EMPLOYEE_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create("1", "Jack", "Sparrow", 1)
	assert.Nil(t, err)

	existEmployee, err := repository.ExistsEmployeeCardNumberId("1")
	assert.Nil(t, err)
	assert.True(t, existEmployee)

	util.DropDB(database)
}

func Test_Repo_ExistsCardNumber_NotFound(t *testing.T) {
	database := util.CreateDB()
	util.QueryExec(database, CREATE_EMPLOYEE_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create("1", "Jack", "Sparrow", 1)
	assert.Nil(t, err)

	existEmployee, err := repository.ExistsEmployeeCardNumberId("234")
	assert.Nil(t, err)
	assert.False(t, existEmployee)

	util.DropDB(database)
}

func Test_Repo_ExistsProductCode_ConnectionError(t *testing.T) {
	database := util.CreateDB()
	util.QueryExec(database, CREATE_EMPLOYEE_TABLE)

	repository := NewRepository(database)

	database.Close()
	_, err := repository.ExistsEmployeeCardNumberId("")
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_Count_Inbound_Orders_By_Employee_OK(t *testing.T) {

	expectedReport := models.ReportInboundOrders{
		Id:           1,
		CardNumberId: "1111222233334444",
		FirstName:    "José",
		LastName:     "Neto",
		WarehouseId:  1,
		InboundOrdersCount: 2,
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_EMPLOYEE_TABLE)
	util.QueryExec(database, CREATE_INBOUND_ORDERS_TABLE)
	util.QueryExec(database, INSERT_EMPLOYEE)
	util.QueryExec(database, INSERT_INBOUND_ORDERS)

	repository := NewRepository(database)
	result, err := repository.CountInboundOrdersByEmployeeId(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedReport, result)

	util.DropDB(database)
}

func Test_Repo_Count_Inbound_Orders_OK(t *testing.T) {

	expectedReport := []models.ReportInboundOrders{
		{
			Id:           1,
			CardNumberId: "1111222233334444",
			FirstName:    "José",
			LastName:     "Neto",
			WarehouseId:  1,
			InboundOrdersCount: 2,
		},
		{
			Id:           2,
			CardNumberId: "1456542642455555",
			FirstName:    "Fernando",
			LastName:     "Diniz",
			WarehouseId:  1,
			InboundOrdersCount: 2,
		},
		{
			Id:           3,
			CardNumberId: "2543542532354543",
			FirstName:    "Paulo",
			LastName:     "Souza",
			WarehouseId:  2,
			InboundOrdersCount: 0,
		},
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_EMPLOYEE_TABLE)
	util.QueryExec(database, CREATE_INBOUND_ORDERS_TABLE)
	util.QueryExec(database, INSERT_EMPLOYEE)
	util.QueryExec(database, INSERT_INBOUND_ORDERS)

	repository := NewRepository(database)
	result, err := repository.CountInboundOrders()
	assert.Nil(t, err)
	assert.Equal(t, expectedReport, result)

	util.DropDB(database)
}



const CREATE_EMPLOYEE_TABLE = `
	CREATE TABLE "employees"(
	  id INTEGER PRIMARY KEY AUTOINCREMENT,
	  id_card_number TEXT NOT NULL,
	  first_name TEXT NOT NULL,
	  last_name TEXT NOT NULL,
	  warehouse_id BIGINT NOT NULL,
	  FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
	);
`

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

const INSERT_INBOUND_ORDERS = `
	INSERT INTO inbound_orders(order_date, order_number, employee_id, product_batch_id, warehouse_id)
	VALUES	("2022-03-21 12:11:21", "1234", 1, 1, 1),
		("2022-04-21 13:11:21", "2134", 1, 2, 2),
		("2022-05-21 14:11:21", "3543", 2, 2, 2),
		("2022-06-21 15:11:21", "3561", 2, 1, 1);
`

const INSERT_EMPLOYEE = `
	INSERT INTO employees(id_card_number, first_name, last_name, warehouse_id)
	VALUES	("1111222233334444", "José", "Neto", 1),
			("1456542642455555", "Fernando", "Diniz", 1),
			("2543542532354543", "Paulo", "Souza", 2);
`
