package sections

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"

	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	util "github.com/GuiTadeu/mercado-fresh-panic/pkg/util"
)

func Test_Repo_Create_Ok(t *testing.T) {

	expectedSection := models.Section{
		Id:                 1,
		Number:             1,
		CurrentTemperature: 11.1,
		MinimumTemperature: 1.0,
		CurrentCapacity:    100,
		MinimumCapacity:    10,
		MaximumCapacity:    100,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SECTION_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create(1, 11.1, 1.0, 100, 10, 100, 1, 1)
	assert.Nil(t, err)

	foundSection, err := repository.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedSection, foundSection)

	util.DropDB(database)
}

func Test_Repo_Create_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SECTION_TABLE)

	repository := NewRepository(database)

	database.Close()
	_, err := repository.Create(1, 11.1, 1.0, 100, 10, 100, 1, 1)
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_Get_OK(t *testing.T) {

	expectedSection := models.Section{
		Id:                 1,
		Number:             1,
		CurrentTemperature: 11.1,
		MinimumTemperature: 1.0,
		CurrentCapacity:    100,
		MinimumCapacity:    10,
		MaximumCapacity:    100,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SECTION_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create(1, 11.1, 1.0, 100, 10, 100, 1, 1)
	assert.Nil(t, err)

	foundSection, err := repository.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedSection, foundSection)

	util.DropDB(database)
}

func Test_Repo_Get_ShouldReturnEmptySectionWhenIdNotExists(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SECTION_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create(1, 11.1, 1.0, 100, 10, 100, 1, 1)
	assert.Nil(t, err)

	foundSection, _ := repository.Get(2)
	assert.Empty(t, foundSection)

	util.DropDB(database)
}

func Test_Repo_Get_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SECTION_TABLE)

	repository := NewRepository(database)

	database.Close()
	_, err := repository.Get(1)
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_GetAll_OK(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SECTION_TABLE)

	repository := NewRepository(database)

	expectedCountRows := 2

	_, err := repository.Create(1, 11.1, 1.0, 100, 10, 100, 1, 1)
	assert.Nil(t, err)

	_, err = repository.Create(2, 22.2, 1.0, 100, 10, 100, 1, 1)
	assert.Nil(t, err)

	foundSections, _ := repository.GetAll()
	assert.Equal(t, expectedCountRows, len(foundSections))

	util.DropDB(database)
}

func Test_Repo_GetAll_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SECTION_TABLE)

	repository := NewRepository(database)

	database.Close()
	_, err := repository.GetAll()
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_Update_OK(t *testing.T) {

	expectedOldSection := models.Section{
		Id:                 1,
		Number:             1,
		CurrentTemperature: 11.1,
		MinimumTemperature: 1.0,
		CurrentCapacity:    100,
		MinimumCapacity:    10,
		MaximumCapacity:    100,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	expectedUpdatedSection := models.Section{
		Id:                 1,
		Number:             2,
		CurrentTemperature: 22.2,
		MinimumTemperature: 2.0,
		CurrentCapacity:    200,
		MinimumCapacity:    20,
		MaximumCapacity:    100,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SECTION_TABLE)

	repository := NewRepository(database)

	_, err := repository.Create(
		expectedOldSection.Number,
		expectedOldSection.CurrentTemperature,
		expectedOldSection.MinimumTemperature,
		expectedOldSection.CurrentCapacity,
		expectedOldSection.MinimumCapacity,
		expectedOldSection.MaximumCapacity,
		expectedOldSection.WarehouseId,
		expectedOldSection.ProductTypeId,
	)
	assert.Nil(t, err)

	_, err = repository.Update(expectedUpdatedSection)
	assert.Nil(t, err)

	foundSection, _ := repository.Get(1)
	assert.Equal(t, expectedUpdatedSection, foundSection)

	util.DropDB(database)
}

func Test_Repo_Update_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SECTION_TABLE)

	repository := NewRepository(database)

	database.Close()
	_, err := repository.Update(models.Section{})
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_Delete_Ok(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SECTION_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create(1, 11.1, 1.0, 100, 10, 100, 1, 1)
	assert.Nil(t, err)

	err = repository.Delete(1)
	assert.Nil(t, err)

	foundSection, _ := repository.Get(1)
	assert.Empty(t, foundSection)

	util.DropDB(database)
}

func Test_Repo_Delete_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SECTION_TABLE)

	repository := NewRepository(database)

	database.Close()
	err := repository.Delete(1)
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_ExistsNumberSection_OK(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SECTION_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create(1, 22.2, 1.0, 100, 10, 100, 1, 1)
	assert.Nil(t, err)

	existsProduct, _ := repository.ExistsSectionNumber(1)
	assert.Nil(t, err)
	assert.True(t, existsProduct)

	util.DropDB(database)
}

func Test_Repo_ExistsSectionNumber_NotFound(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SECTION_TABLE)

	repository := NewRepository(database)
	_, err := repository.Create(2, 22.2, 1.0, 100, 10, 100, 1, 1)
	assert.Nil(t, err)

	existsProduct, _ := repository.ExistsSectionNumber(1)
	assert.Nil(t, err)
	assert.False(t, existsProduct)

	util.DropDB(database)
}

func Test_Repo_ExistsSectionNumber_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SECTION_TABLE)

	repository := NewRepository(database)

	database.Close()
	_, err := repository.ExistsSectionNumber(0)
	assert.NotNil(t, err)

	util.DropDB(database)
}

const CREATE_SECTION_TABLE = `
	CREATE TABLE "sections" (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		section_number BIGINT NOT NULL,
		current_capacity  BIGINT NOT NULL,
		current_temperature DECIMAL(19, 2) NOT NULL,
		maximum_capacity  BIGINT NOT NULL,
		minimum_capacity BIGINT NOT NULL,
		minimum_temperature DECIMAL(19, 2) NOT NULL,
		product_type BIGINT NOT NULL,
		warehouse_id BIGINT NOT NULL,
		FOREIGN KEY (product_type) REFERENCES products_types(id),
		FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
	);
`
