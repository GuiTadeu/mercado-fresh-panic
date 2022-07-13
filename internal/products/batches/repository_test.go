package batches

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"

	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/sections"
	util "github.com/GuiTadeu/mercado-fresh-panic/pkg/util"
)

func Test_Repo_Create_Ok(t *testing.T) {

	expectedBatch := models.ProductBatch{
		Id:                 1,
		Number:             666,
		CurrentQuantity:    666,
		CurrentTemperature: 666,
		DueDate:            "2012",
		InitialQuantity:    666,
		ManufacturingDate:  "2012",
		ManufacturingHour:  "16:20",
		MinimumTemperature: 666,
		ProductId:          1,
		SectionId:          1,
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCT_BATCHES_TABLE)

	repository := NewProductBatchRepository(database)
	_, err := repository.Create(666, 666, 666, "2012", 666, "2012", "16:20", 666, 1, 1)
	assert.Nil(t, err)

	foundBatch, err := repository.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedBatch, foundBatch)

	util.DropDB(database)
}

func Test_Repo_Create_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCT_BATCHES_TABLE)

	repository := NewProductBatchRepository(database)

	database.Close()
	_, err := repository.Create(1, 1, 1, "", 1, "", "", 1, 1, 1)
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_Get_OK(t *testing.T) {

	expectedBatch := models.ProductBatch{
		Id:                 1,
		Number:             666,
		CurrentQuantity:    666,
		CurrentTemperature: 666,
		DueDate:            "2012",
		InitialQuantity:    666,
		ManufacturingDate:  "2012",
		ManufacturingHour:  "16:20",
		MinimumTemperature: 666,
		ProductId:          1,
		SectionId:          1,
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCT_BATCHES_TABLE)

	repository := NewProductBatchRepository(database)
	_, err := repository.Create(666, 666, 666, "2012", 666, "2012", "16:20", 666, 1, 1)
	assert.Nil(t, err)

	foundBatch, err := repository.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedBatch, foundBatch)

	util.DropDB(database)
}

func Test_Repo_Get_ShouldReturnEmptyProductBatchWhenIdNotExists(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCT_BATCHES_TABLE)

	repository := NewProductBatchRepository(database)
	_, err := repository.Create(666, 666, 666, "2012", 666, "2012", "16:20", 666, 1, 1)
	assert.Nil(t, err)

	foundProduct, err := repository.Get(2)
	assert.Empty(t, foundProduct)

	util.DropDB(database)
}

func Test_Repo_Get_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCT_BATCHES_TABLE)

	repository := NewProductBatchRepository(database)

	database.Close()
	_, err := repository.Get(1)
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_CountProductsBySections_Ok(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SECTIONS_TABLE)
	util.QueryExec(database, CREATE_PRODUCT_BATCHES_TABLE)

	batchRepository := NewProductBatchRepository(database)
	sectionRepository := sections.NewRepository(database)

	_, err := sectionRepository.Create(444, 44.4, 4.0, 400, 40, 400, 4, 4) // id: 1
	_, err = sectionRepository.Create(999, 99.9, 9.0, 900, 90, 900, 9, 9) // id: 2

	_, err = batchRepository.Create(666, 666, 666, "2012", 666, "2012", "16:20", 666, 1, 1)
	_, err = batchRepository.Create(777, 777, 777, "2013", 777, "2013", "17:20", 777, 1, 2)

	report, err := batchRepository.CountProductsBySections()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(report))

	util.DropDB(database)
}

func Test_Repo_CountProductsBySections_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCT_BATCHES_TABLE)

	repository := NewProductBatchRepository(database)

	database.Close()
	_, err := repository.CountProductsBySections()
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_CountProductsBySectionId_Ok(t *testing.T) {

	expectedReport := models.CountProductsBySectionIdReport{
		SectionId: 1,
		SectionNumber: 444,
		ProductsCount: 2,
	}

	database := util.CreateDB()
	util.QueryExec(database, CREATE_SECTIONS_TABLE)
	util.QueryExec(database, CREATE_PRODUCT_BATCHES_TABLE)

	batchRepository := NewProductBatchRepository(database)
	sectionRepository := sections.NewRepository(database)

	_, err := sectionRepository.Create(444, 44.4, 4.0, 400, 40, 400, 4, 4) // id: 1
	_, err = sectionRepository.Create(999, 99.9, 9.0, 900, 90, 900, 9, 9) // id: 2

	_, err = batchRepository.Create(666, 666, 666, "2012", 666, "2012", "16:20", 666, 1, 1)
	_, err = batchRepository.Create(777, 777, 777, "2013", 777, "2013", "17:20", 777, 2, 1)
	_, err = batchRepository.Create(777, 777, 777, "2013", 777, "2013", "17:20", 777, 1, 3)

	report, err := batchRepository.CountProductsBySectionId(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedReport, report)

	util.DropDB(database)
}

func Test_Repo_CountProductsBySectionId_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCT_BATCHES_TABLE)

	repository := NewProductBatchRepository(database)

	database.Close()
	_, err := repository.CountProductsBySectionId(1)
	assert.NotNil(t, err)

	util.DropDB(database)
}

func Test_Repo_ExistsBatchNumber_ShouldReturnTrue(t *testing.T) {

	expectedBatchNumber := uint64(666)
	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCT_BATCHES_TABLE)

	repository := NewProductBatchRepository(database)
	_, err := repository.Create(expectedBatchNumber, 666, 666, "2012", 666, "2012", "16:20", 666, 1, 1)
	assert.Nil(t, err)

	existsBatchNumber, err := repository.ExistsBatchNumber(expectedBatchNumber)
	assert.Nil(t, err)
	assert.True(t, existsBatchNumber)

	util.DropDB(database)
}

func Test_Repo_ExistsBatchNumber_ShouldReturnFalse(t *testing.T) {

	expectedBatchNumber := uint64(666)
	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCT_BATCHES_TABLE)

	repository := NewProductBatchRepository(database)
	_, err := repository.Create(expectedBatchNumber, 666, 666, "2012", 666, "2012", "16:20", 666, 1, 1)
	assert.Nil(t, err)

	existsBatchNumber, err := repository.ExistsBatchNumber(2345678)
	assert.False(t, existsBatchNumber)

	util.DropDB(database)
}

func Test_Repo_ExistsBatchNumber_ConnectionError(t *testing.T) {

	database := util.CreateDB()
	util.QueryExec(database, CREATE_PRODUCT_BATCHES_TABLE)

	repository := NewProductBatchRepository(database)

	database.Close()
	_, err := repository.ExistsBatchNumber(1)
	assert.NotNil(t, err)

	util.DropDB(database)
}

const CREATE_PRODUCT_BATCHES_TABLE = `
	CREATE TABLE "product_batches"(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		batch_number BIGINT NOT NULL,
		current_quantity BIGINT NOT NULL,
		current_temperature DECIMAL(19, 2) NOT NULL,
		due_date TEXT NOT NULL,
		initial_quantity BIGINT NOT NULL,
		manufacturing_date TEXT NOT NULL,
		manufacturing_hour TEXT NOT NULL,
		minimum_temperature DECIMAL(19, 2) NOT NULL,
		product_id BIGINT NOT NULL,
		section_id BIGINT NOT NULL,
		FOREIGN KEY (product_id) REFERENCES products(id),
		FOREIGN KEY (section_id) REFERENCES sections(id)
	);
`

const CREATE_SECTIONS_TABLE = `
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
