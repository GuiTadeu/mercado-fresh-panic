package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	models "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/products/batches"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func Test_Create_Batch_201(t *testing.T) {
	validProductBatch := models.ProductBatch{
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

	jsonValue, _ := json.Marshal(validProductBatch)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockProductBatchService{
		result: validProductBatch,
		err:    nil,
	}

	router := setupBatchRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/productBatches", requestBody)
	router.ServeHTTP(response, request)

	responseData := models.ProductBatch{}
	decodeInboundOrderWebResponse(response, &responseData)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, validProductBatch, responseData)
}

func Test_Create_Batch_409_NotFound_ProductId(t *testing.T) {
	expectedError := batches.ProductNotFoundError

	validProductBatch := models.ProductBatch{
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

	jsonValue, _ := json.Marshal(validProductBatch)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockProductBatchService{
		err: expectedError,
	}

	router := setupBatchRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/productBatches", requestBody)
	router.ServeHTTP(response, request)

	responseData := web.Response{}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	assert.Equal(t, http.StatusConflict, response.Code)
	assert.Equal(t, expectedError.Error(), responseData.Error)
}

func Test_Create_Batch_409_NotFound_SectionId(t *testing.T) {
	expectedError := batches.SectionNotFoundError

	validProductBatch := models.ProductBatch{
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

	jsonValue, _ := json.Marshal(validProductBatch)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockProductBatchService{
		err: expectedError,
	}

	router := setupBatchRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/productBatches", requestBody)
	router.ServeHTTP(response, request)

	responseData := web.Response{}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	assert.Equal(t, http.StatusConflict, response.Code)
	assert.Equal(t, expectedError.Error(), responseData.Error)
}

func Test_CountProductsBySections(t *testing.T) {
	
	report := []models.CountProductsBySectionIdReport{
		{
			SectionId: 1,
			SectionNumber: 200,
			ProductsCount: 350,
		},
		{
			SectionId: 2,
			SectionNumber: 900,
			ProductsCount: 140,
		},
	}

	mockService := mockProductBatchService{
		result: report,
		err:    nil,
	}

	router := setupBatchRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/sections/reportProducts", nil)
	router.ServeHTTP(response, request)

	responseData := []models.CountProductsBySectionIdReport{}
	decodeBatchWebResponse(response, &responseData)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, report, responseData)
}

func Test_CountProductsBySections_With_SectionId(t *testing.T) {
	
	report := models.CountProductsBySectionIdReport{
		SectionId: 1,
		SectionNumber: 200,
		ProductsCount: 350,
	}

	mockService := mockProductBatchService{
		result: report,
		err:    nil,
	}

	router := setupBatchRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/sections/reportProducts?id=1", nil)
	router.ServeHTTP(response, request)

	responseData := models.CountProductsBySectionIdReport{}
	decodeBatchWebResponse(response, &responseData)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, report, responseData)
}

func decodeBatchWebResponse(response *httptest.ResponseRecorder, responseData any) {
	responseStruct := web.Response{}
	json.Unmarshal(response.Body.Bytes(), &responseStruct)

	jsonData, _ := json.Marshal(responseStruct.Data)
	json.Unmarshal(jsonData, &responseData)
}

func setupBatchRouter(mockService mockProductBatchService) *gin.Engine {
	controller := NewProductBatchController(mockService)

	router := gin.Default()
	router.POST("/api/v1/productBatches", controller.Create())
	router.GET("/api/v1/sections/reportProducts", controller.CountProductsBySections())

	return router
}
