package batches

import (
	"testing"

	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/products"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/sections"
	"github.com/stretchr/testify/assert"
)

func Test_Create_Ok(t *testing.T) {

	expectedResult := models.ProductBatch{
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

	mockProductBatchesRepository := MockProductBatchesRepository{
		result:            expectedResult,
		err:               nil,
		existsBatchNumber: false,
	}

	mockSectionRepository := sections.MockSectionRepository{
		GetById: models.Section{
			Id:                 1,
			Number:             666,
			CurrentTemperature: 11.1,
			MinimumTemperature: 2.0,
			CurrentCapacity:    600,
			MinimumCapacity:    50,
			MaximumCapacity:    1000,
			WarehouseId:        1,
			ProductTypeId:      1,
		},
	}

	mockProductRepository := products.MockProductRepository{
		GetById: models.Product{
			Id:                      1,
			Code:                    "KKK",
			Description:             "NÃO AGUENTO MAIS",
			Width:                   200,
			Height:                  450,
			Length:                  300,
			NetWeight:               120,
			ExpirationRate:          302,
			RecommendedFreezingTemp: 45,
			FreezingRate:            67,
			ProductTypeId:           1,
			SellerId:                1,
		},
	}

	service := NewProductBatchesService(mockProductBatchesRepository, mockSectionRepository, mockProductRepository)
	result, err := service.Create(666, 666, 666, "2012", 666, "2012", "16:20", 666, 1, 1)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Create_ShouldReturnErrorWhenNumberAlreadyExists(t *testing.T) {

	expectedError := ExistsBatchNumberError

	mockProductBatchesRepository := MockProductBatchesRepository{
		existsBatchNumber: true,
	}

	mockSectionRepository := sections.MockSectionRepository{}
	mockProductRepository := products.MockProductRepository{}

	service := NewProductBatchesService(mockProductBatchesRepository, mockSectionRepository, mockProductRepository)
	_, err := service.Create(666, 666, 666, "2012", 666, "2012", "16:20", 666, 1, 1)

	assert.Equal(t, expectedError, err)
}

func Test_Create_ShouldReturnErrorWhenNotFoundProduct(t *testing.T) {

	expectedError := ProductNotFoundError

	mockProductBatchesRepository := MockProductBatchesRepository{
		err:           expectedError,
		result:        models.ProductBatch{},
	}

	mockSectionRepository := sections.MockSectionRepository{}
	mockProductRepository := products.MockProductRepository{}

	service := NewProductBatchesService(mockProductBatchesRepository, mockSectionRepository, mockProductRepository)
	_, err := service.Create(666, 666, 666, "2012", 666, "2012", "16:20", 666, 1, 1)

	assert.Equal(t, expectedError, err)
}

func Test_Create_ShouldReturnErrorWhenNotFoundSection(t *testing.T) {

	expectedError := SectionNotFoundError

	mockProductBatchesRepository := MockProductBatchesRepository{
		err:           expectedError,
		result:        models.ProductBatch{},
	}

	mockSectionRepository := sections.MockSectionRepository{}
	mockProductRepository := products.MockProductRepository{}

	service := NewProductBatchesService(mockProductBatchesRepository, mockSectionRepository, mockProductRepository)
	_, err := service.Create(666, 666, 666, "2012", 666, "2012", "16:20", 666, 1, 1)

	assert.Equal(t, expectedError, err)
}

func Test_CountProductsBySectionId_Ok(t *testing.T) {

	expectedResult := models.CountProductsBySectionIdReport{
		SectionId:     1,
		SectionNumber: 1,
		ProductsCount: 666,
	}

	mockProductBatchesRepository := MockProductBatchesRepository{
		result:            expectedResult,
		err:               nil,
		existsBatchNumber: false,
	}

	mockSectionRepository := sections.MockSectionRepository{}
	mockProductRepository := products.MockProductRepository{}

	service := NewProductBatchesService(mockProductBatchesRepository, mockSectionRepository, mockProductRepository)
	result, err := service.CountProductsBySectionId(1)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_CountProductsBySections_Ok(t *testing.T) {

	expectedResult := []models.CountProductsBySectionIdReport{
		{
			SectionId:     1,
			SectionNumber: 123,
			ProductsCount: 666,
		},

		{
			SectionId:     2,
			SectionNumber: 456,
			ProductsCount: 100,
		},
	}

	mockProductBatchesRepository := MockProductBatchesRepository{
		result:            expectedResult,
		err:               nil,
		existsBatchNumber: false,
	}

	mockSectionRepository := sections.MockSectionRepository{}
	mockProductRepository := products.MockProductRepository{}

	service := NewProductBatchesService(mockProductBatchesRepository, mockSectionRepository, mockProductRepository)
	result, err := service.CountProductsBySections()

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}
