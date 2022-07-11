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
