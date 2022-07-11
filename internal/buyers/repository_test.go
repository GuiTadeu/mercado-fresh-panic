package buyers

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"

	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	util "github.com/GuiTadeu/mercado-fresh-panic/pkg/util"
)

func Test_Repo_Create_OK(t *testing.T) {

	expectedBuyer := models.Buyer{
		Id:           1,
		CardNumberId: "11",
		FirstName:    "Willy",
		LastName:     "Passos",
	}
	database := util.CreateDB()
	util.QueryExec(database, CREATE_BUYERS_TABLE)

	repository := NewBuyerRepository(database)
	_, err := repository.Create("11", "Willy", "Passos")
	assert.Nil(t, err)

	foundBuyer, err := repository.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedBuyer, foundBuyer)

	util.DropDB(database)

}

func Test_Repo_Create_Buyer_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_BUYERS_TABLE)

	repository := NewBuyerRepository(database)

	database.Close()
	_, err := repository.Get(1)
	assert.NotNil(t, err)

	util.DropDB(database)

}

func Test_Repo_GetAll_Buyer_OK(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_BUYERS_TABLE)

	repository := NewBuyerRepository(database)

	expectedCountRouws := 2

	_, err := repository.Create("11", "Joao", "Silva")

	_, err = repository.Create("22", "Jose", "Santos")
	assert.Nil(t, err)

	foundBuyers, _ := repository.GetAll()
	assert.Equal(t, expectedCountRouws, len(foundBuyers))

	util.DropDB(database)

}

func Test_Repo_GetAll_Buyer_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_BUYERS_TABLE)

	repository := NewBuyerRepository(database)

	database.Close()
	_, err := repository.GetAll()
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_Update_Buyer_Ok(t *testing.T) {

	expectedOldBuyer := models.Buyer{
		Id:           1,
		CardNumberId: "10",
		FirstName:    "Joao",
		LastName:     "Silva",
	}

	expectedUpdateBuyer := models.Buyer{
		Id:           1,
		CardNumberId: "10",
		FirstName:    "Joao",
		LastName:     "Silva",
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_BUYERS_TABLE)

	repository := NewBuyerRepository(database)

	_, err := repository.Create(
		expectedOldBuyer.CardNumberId,
		expectedOldBuyer.FirstName,
		expectedOldBuyer.LastName,
	)
	assert.Nil(t, err)

	foundBuyer, _ := repository.Get(1)
	assert.Equal(t, expectedOldBuyer, foundBuyer)

	_, err = repository.Update(expectedUpdateBuyer)
	assert.Nil(t, err)

	foundBuyer, _ = repository.Get(1)
	assert.Equal(t, expectedUpdateBuyer, foundBuyer)

	util.DropDB(database)
}

func Test_Repo_Update_Buyer_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_BUYERS_TABLE)

	repository := NewBuyerRepository(database)

	database.Close()
	_, err := repository.Update(models.Buyer{})
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_Delete_Buyer_Ok(t *testing.T) {

	expectedBuyer := models.Buyer{
		Id:           1,
		CardNumberId: "11",
		FirstName:    "Levi",
		LastName:     "Passos",
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_BUYERS_TABLE)

	repository := NewBuyerRepository(database)
	_, err := repository.Create("11", "Levi", "Passos")
	assert.Nil(t, err)

	foundBuyer, err := repository.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedBuyer, foundBuyer)

	err = repository.Delete(1)
	assert.Nil(t, err)

	foundBuyer, _ = repository.Get(1)
	assert.Empty(t, foundBuyer)

	util.DropDB(database)
}

func Test_Repo_Delete_Buyer_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_BUYERS_TABLE)

	repository := NewBuyerRepository(database)

	database.Close()
	err := repository.Delete(1)
	assert.NotNil(t, err)

	util.DropDB(database)

}

func Test_Repo_ExistsBuyerCardNumberId_Buyer_Ok(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_BUYERS_TABLE)

	repository := NewBuyerRepository(database)
	_, err := repository.Create("20", "Antoaio", "Souza")
	assert.Nil(t, err)

	existsBuyer, err := repository.ExistsBuyerCardNumberId("20")
	assert.Nil(t, err)
	assert.True(t, existsBuyer)

	util.DropDB(database)
}

func Test_Repo_ExistsBuyerCardNumberId_Buyer_NotFound(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_BUYERS_TABLE)

	repository := NewBuyerRepository(database)
	_, err := repository.Create("13", "Paulo", "Pereira")
	assert.Nil(t, err)

	existsBuyer, err := repository.ExistsBuyerCardNumberId("16")
	assert.Nil(t, err)
	assert.False(t, existsBuyer)

	util.DropDB(database)
}

func Test_Repo_ExistsBuyerCardNumberId_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_BUYERS_TABLE)

	repository := NewBuyerRepository(database)

	database.Close()
	_, err := repository.ExistsBuyerCardNumberId("45")
	assert.NotNil(t, err)

	util.DropDB(database)
}

const CREATE_BUYERS_TABLE = `CREATE TABLE  "buyers"(
id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL ,
id_card_number TEXT NOT NULL,
first_name TEXT NOT NULL,
last_name TEXT NOT NULL
);
`
