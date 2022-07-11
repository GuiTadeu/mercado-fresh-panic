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

}

const CREATE_BUYERS_TABLE = `CREATE TABLE  "buyers"(
id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL ,
id_card_number TEXT NOT NULL,
first_name TEXT NOT NULL,
last_name TEXT NOT NULL
);
`
