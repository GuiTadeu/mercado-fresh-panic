package sellers

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"

	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	util "github.com/GuiTadeu/mercado-fresh-panic/pkg/util"
)

func Test_Repo_Create_Ok(t *testing.T) {

	expectedSeller := models.Seller{
		Id:          1,
		Cid:         1,
		CompanyName: "Microsoft",
		Address:     "Rua Pedro Américo, 123",
		Telephone:   "1233265466",
		LocalityId:  "11065001",
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SELLERS_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create(1, "Microsoft", "Rua Pedro Américo, 123", "1233265466", "11065001")
	assert.Nil(t, err)

	foundSeller, err := repository.FindOne(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedSeller, foundSeller)

	util.DropDB(database)
}

func Test_Repo_Create_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SELLERS_TABLE)

	repository := NewRepository(database)

	database.Close()
	_, err := repository.Create(1, "Microsoft", "Rua Pedro Américo, 123", "1233265466", "11065001")
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_FindOne_WhenSellerIdNotExists(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SELLERS_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create(1, "Microsoft", "Rua Pedro Américo, 123", "1233265466", "11065001")
	assert.Nil(t, err)

	foundSeller, _ := repository.FindOne(2)
	assert.Empty(t, foundSeller)

	util.DropDB(database)
}

func Test_Repo_FindOne_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SELLERS_TABLE)

	repository := NewRepository(database)

	database.Close()
	_, err := repository.FindOne(1)
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_FindAll_OK(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SELLERS_TABLE)

	repository := NewRepository(database)

	expectedCountRows := 2

	_, err := repository.Create(1, "Microsoft", "Rua Pedro Américo, 123", "1233265466", "11065001")
	assert.Nil(t, err)

	_, err = repository.Create(2, "Apple", "Rua Goiás, 1233", "1233265412", "11092001")
	assert.Nil(t, err)

	foundSellers, _ := repository.FindAll()
	assert.Equal(t, expectedCountRows, len(foundSellers))

	util.DropDB(database)
}

func Test_Repo_FindAll_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SELLERS_TABLE)

	repository := NewRepository(database)

	database.Close()
	_, err := repository.FindAll()
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_Update_Ok(t *testing.T) {

	updatedSeller := models.Seller{
		Id:          2,
		Cid:         1,
		CompanyName: "Apple",
		Address:     "Rua Goiás, 73",
		Telephone:   "1233265466",
		LocalityId:  "11065001",
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SELLERS_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create(1, "Microsoft", "Rua Pedro Américo, 123", "1233265466", "11065001")
	assert.Nil(t, err)

	foundSeller, err := repository.Update(updatedSeller)
	assert.Nil(t, err)
	assert.Equal(t, updatedSeller, foundSeller)

	util.DropDB(database)
}

func Test_Repo_Update_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SELLERS_TABLE)

	repository := NewRepository(database)

	database.Close()
	_, err := repository.Update(models.Seller{})
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_Delete_Ok(t *testing.T) {

	expectedSeller := models.Seller{
		Id:          1,
		Cid:         1,
		CompanyName: "Microsoft",
		Address:     "Rua Pedro Américo, 123",
		Telephone:   "1233265466",
		LocalityId:  "11065001",
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SELLERS_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create(1, "Microsoft", "Rua Pedro Américo, 123", "1233265466", "11065001")
	assert.Nil(t, err)

	foundSeller, err := repository.FindOne(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedSeller, foundSeller)

	err = repository.Delete(1)
	assert.Nil(t, err)

	foundSeller, _ = repository.FindOne(1)
	assert.Empty(t, foundSeller)

	util.DropDB(database)
}

func Test_Repo_Delete_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SELLERS_TABLE)

	repository := NewRepository(database)

	database.Close()
	err := repository.Delete(1)
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_FindCid_OK(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SELLERS_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create(1, "Microsoft", "Rua Pedro Américo, 123", "1233265466", "11065001")
	assert.Nil(t, err)

	existsSellerCid := repository.FindCid(1)
	assert.True(t, existsSellerCid)

	util.DropDB(database)
}

func Test_Repo_FindCid_NotFound(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SELLERS_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create(1, "Microsoft", "Rua Pedro Américo, 123", "1233265466", "11065001")
	assert.Nil(t, err)

	existsSellerCid := repository.FindCid(2)
	assert.False(t, existsSellerCid)

	util.DropDB(database)
}

func Test_Repo_FindCid_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SELLERS_TABLE)

	repository := NewRepository(database)

	database.Close()
	existsSellerCid := repository.FindCid(1)
	assert.NotNil(t, existsSellerCid)

	util.DropDB(database)
}

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
