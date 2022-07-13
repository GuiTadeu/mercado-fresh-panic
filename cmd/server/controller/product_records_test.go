package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	productrecords "github.com/GuiTadeu/mercado-fresh-panic/internal/product_records"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Product_Records_Create_201(t *testing.T) {
	validProductRecords := db.ProductRecord{
		Id:             1,
		LastUpdateDate: "2021-04-04",
		PurchasePrice:  10,
		SalePrice:      15,
		ProductId:      1,
	}

	jsonValue, _ := json.Marshal(validProductRecords)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockProductRecordsService{
		result: validProductRecords,
		err:    nil,
	}

	router := setupProductRecordsRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/productRecords", requestBody)
	router.ServeHTTP(response, request)

	responseData := db.ProductRecord{}
	decodeProductRecords(response, &responseData)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, validProductRecords, responseData)

}

func Test_Product_Records_Create_409(t *testing.T) {
	validProductRecords := db.ProductRecord{
		Id:             1,
		LastUpdateDate: "2021-04-04",
		PurchasePrice:  10,
		SalePrice:      15,
		ProductId:      999,
	}

	jsonValue, _ := json.Marshal(validProductRecords)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockProductRecordsService{
		result: validProductRecords,
		err:    productrecords.ErrProductNotFoundError,
	}

	router := setupProductRecordsRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/productRecords", requestBody)
	router.ServeHTTP(response, request)

	responseData := db.ProductRecord{}
	decodeProductRecords(response, &responseData)

	assert.Equal(t, http.StatusConflict, response.Code)

}

func Test_Product_Records_Create_422(t *testing.T) {
	validProductRecords := "nil"

	jsonValue, _ := json.Marshal(validProductRecords)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockProductRecordsService{
		result: validProductRecords,
		err:    productrecords.ErrProductNotFoundError,
	}

	router := setupProductRecordsRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/productRecords", requestBody)
	router.ServeHTTP(response, request)

	responseData := db.ProductRecord{}
	decodeProductRecords(response, &responseData)

	assert.Equal(t, http.StatusUnprocessableEntity, response.Code)

}

func decodeProductRecords(response *httptest.ResponseRecorder, responseData any) {
	responseStruct := web.Response{}
	json.Unmarshal(response.Body.Bytes(), &responseStruct)

	jsonData, _ := json.Marshal(responseStruct.Data)
	json.Unmarshal(jsonData, &responseData)
}

func setupProductRecordsRouter(mockService mockProductRecordsService) *gin.Engine {
	controller := NewProductRecordsController(mockService)

	router := gin.Default()
	router.POST("/api/v1/productRecords", controller.Create())

	return router
}
