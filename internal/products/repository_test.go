package products

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"

	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	util "github.com/GuiTadeu/mercado-fresh-panic/pkg/util"
)

func Test_Repo_Create_Ok(t *testing.T) {

	expectedProduct := models.Product{
		Id: 1,
		Code: "dvd",
		Description: "Pirata",
		Width: 1,
		Height: 1,
		Length: 1,
		NetWeight: 1,
		ExpirationRate: 1,
		RecommendedFreezingTemp: 1,
		FreezingRate: 1,
		ProductTypeId: 1,
		SellerId: 1,
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCTS_TABLE)
	
	repository := NewProductRepository(database)
	_, err := repository.Create("dvd", "Pirata", 1, 1, 1, 1, 1, 1, 1, 1, 1)
	assert.Nil(t, err)

	foundProduct, err := repository.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedProduct, foundProduct)

	util.DropDB(database)
}

func Test_Repo_Create_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCTS_TABLE)
	
	repository := NewProductRepository(database)

	database.Close()
	_, err := repository.Create("", "", 0, 0, 0, 0, 0, 0, 0, 0, 0)
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_Get_OK(t *testing.T) {

	expectedProduct := models.Product{
		Id: 1,
		Code: "dvd",
		Description: "Pirata",
		Width: 1,
		Height: 1,
		Length: 1,
		NetWeight: 1,
		ExpirationRate: 1,
		RecommendedFreezingTemp: 1,
		FreezingRate: 1,
		ProductTypeId: 1,
		SellerId: 1,
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCTS_TABLE)
	
	repository := NewProductRepository(database)
	_, err := repository.Create("dvd", "Pirata", 1, 1, 1, 1, 1, 1, 1, 1, 1)
	assert.Nil(t, err)

	foundProduct, err := repository.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedProduct, foundProduct)

	util.DropDB(database)
}

func Test_Repo_Get_ShouldReturnEmptyProductWhenIdNotExists(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCTS_TABLE)
	
	repository := NewProductRepository(database)
	_, err := repository.Create("dvd", "Pirata", 1, 1, 1, 1, 1, 1, 1, 1, 1)
	assert.Nil(t, err)

	foundProduct, err := repository.Get(2)
	assert.Empty(t, foundProduct)

	util.DropDB(database)
}

func Test_Repo_Get_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCTS_TABLE)
	
	repository := NewProductRepository(database)

	database.Close()
	_, err := repository.Get(1)
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_GetAll_OK(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCTS_TABLE)
	
	repository := NewProductRepository(database)

	expectedCountRows := 2

	_, err := repository.Create("dvd", "Pirata", 1, 1, 1, 1, 1, 1, 1, 1, 1)
	assert.Nil(t, err)

	_, err = repository.Create("cd", "Original", 1, 1, 1, 1, 1, 1, 1, 1, 1)
	assert.Nil(t, err)

	foundProducts, err := repository.GetAll()
	assert.Equal(t, expectedCountRows, len(foundProducts))

	util.DropDB(database)
}

func Test_Repo_GetAll_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCTS_TABLE)
	
	repository := NewProductRepository(database)

	database.Close()
	_, err := repository.GetAll()
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_Update_OK(t *testing.T) {

	expectedOldProduct := models.Product{
		Id: 1,
		Code: "dvd",
		Description: "Pirata",
		Width: 1,
		Height: 1,
		Length: 1,
		NetWeight: 1,
		ExpirationRate: 1,
		RecommendedFreezingTemp: 1,
		FreezingRate: 1,
		ProductTypeId: 1,
		SellerId: 1,
	}

	expectedUpdatedProduct := models.Product{
		Id: 1,
		Code: "almanaque",
		Description: "Deep Web",
		Width: 2,
		Height: 2,
		Length: 2,
		NetWeight: 2,
		ExpirationRate: 2,
		RecommendedFreezingTemp: 2,
		FreezingRate: 2,
		ProductTypeId: 1,
		SellerId: 1,
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCTS_TABLE)
	
	repository := NewProductRepository(database)

	_, err := repository.Create(
		expectedOldProduct.Code,
		expectedOldProduct.Description,
		expectedOldProduct.Width,
		expectedOldProduct.Height,
		expectedOldProduct.Length,
		expectedOldProduct.NetWeight,
		expectedOldProduct.ExpirationRate,
		expectedOldProduct.RecommendedFreezingTemp,
		expectedOldProduct.FreezingRate,
		expectedOldProduct.ProductTypeId,
		expectedOldProduct.SellerId,
	)
	assert.Nil(t, err)

	foundProduct, _ := repository.Get(1)
	assert.Equal(t, expectedOldProduct, foundProduct)

	_, err = repository.Update(expectedUpdatedProduct)
	assert.Nil(t, err)

	foundProduct, _ = repository.Get(1)
	assert.Equal(t, expectedUpdatedProduct, foundProduct)

	util.DropDB(database)
}

func Test_Repo_Update_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCTS_TABLE)
	
	repository := NewProductRepository(database)

	database.Close()
	_, err := repository.Update(models.Product{})
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_Delete_Ok(t *testing.T) {

	expectedProduct := models.Product{
		Id: 1,
		Code: "dvd",
		Description: "Pirata",
		Width: 1,
		Height: 1,
		Length: 1,
		NetWeight: 1,
		ExpirationRate: 1,
		RecommendedFreezingTemp: 1,
		FreezingRate: 1,
		ProductTypeId: 1,
		SellerId: 1,
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCTS_TABLE)
	
	repository := NewProductRepository(database)
	_, err := repository.Create("dvd", "Pirata", 1, 1, 1, 1, 1, 1, 1, 1, 1)
	assert.Nil(t, err)

	foundProduct, err := repository.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedProduct, foundProduct)

	err = repository.Delete(1)
	assert.Nil(t, err)

	foundProduct, err = repository.Get(1)
	assert.Empty(t, foundProduct)

	util.DropDB(database)
}

func Test_Repo_Delete_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCTS_TABLE)
	
	repository := NewProductRepository(database)

	database.Close()
	err := repository.Delete(1)
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_ExistsProductCode_OK(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCTS_TABLE)
	
	repository := NewProductRepository(database)
	_, err := repository.Create("dvd", "Pirata", 1, 1, 1, 1, 1, 1, 1, 1, 1)
	assert.Nil(t, err)

	existsProduct, err := repository.ExistsProductCode("dvd")
	assert.Nil(t, err)
	assert.True(t, existsProduct)

	util.DropDB(database)
}

func Test_Repo_ExistsProductCode_NotFound(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCTS_TABLE)
	
	repository := NewProductRepository(database)
	_, err := repository.Create("dvd", "Pirata", 1, 1, 1, 1, 1, 1, 1, 1, 1)
	assert.Nil(t, err)

	existsProduct, err := repository.ExistsProductCode("goiaba")
	assert.Nil(t, err)
	assert.False(t, existsProduct)

	util.DropDB(database)
}

func Test_Repo_ExistsProductCode_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCTS_TABLE)
	
	repository := NewProductRepository(database)

	database.Close()
	_, err := repository.ExistsProductCode("")
	assert.NotNil(t, err)

	util.DropDB(database)
}

const CREATE_PRODUCTS_TABLE = `
	CREATE TABLE "products" (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		expiration_rate DECIMAL(19, 2) NOT NULL,
		freezing_rate DECIMAL(19, 2) NOT NULL,
		height DECIMAL(19, 2) NOT NULL,
		length DECIMAL(19, 2) NOT NULL,
		net_weight DECIMAL(19, 2) NOT NULL,
		product_code TEXT NOT NULL,
		recommended_freezing_temperature DECIMAL(19, 2) NOT NULL,
		width DECIMAL(19, 2) NOT NULL,
		product_type BIGINT  NOT NULL,
		seller_id BIGINT  NOT NULL,
		FOREIGN KEY (product_type) REFERENCES products_types(id),
		FOREIGN KEY (seller_id) REFERENCES sellers(id)
	);
`