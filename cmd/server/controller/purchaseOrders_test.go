package controller

import (
	"bytes"
	"encoding/json"
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/purchaseOrders"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_PurchaseOrders_Create_201(t *testing.T) {

	validPurchaseOrder := db.PurchaseOrder{
		Id:              1,
		OrderNumber:     "777",
		OrderDate:       "2022-07-12",
		TrackingCode:    "777",
		BuyerId:         1,
		OrderStatusId:   1,
		ProductRecordId: 1,
	}

	jsonValue, _ := json.Marshal(validPurchaseOrder)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockPurchaseOrdersService{
		result: validPurchaseOrder,
		err:    nil,
	}

	router := setupPurchaseOrdersRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/purchaseOrders", requestBody)
	router.ServeHTTP(response, request)

	responseData := db.PurchaseOrder{}
	decodePurchaseOrdersWebResponse(response, &responseData)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, validPurchaseOrder, responseData)
}

func Test_PurchaseOrders_Create_422(t *testing.T) {

	invalidPurchaseOrder := db.PurchaseOrder{}
	jsonValue, _ := json.Marshal(invalidPurchaseOrder)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockPurchaseOrdersService{}

	router := setupPurchaseOrdersRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/purchaseOrders", requestBody)
	router.ServeHTTP(response, request)

	assert.Equal(t, 422, response.Code)
}

func Test_PurchaseOrders_Create_409(t *testing.T) {
	validPurchaseOrder := db.PurchaseOrder{
		Id:              1,
		OrderNumber:     "777",
		OrderDate:       "2022-07-12",
		TrackingCode:    "777",
		BuyerId:         1,
		OrderStatusId:   1,
		ProductRecordId: 1,
	}

	jsonValue, _ := json.Marshal(validPurchaseOrder)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockPurchaseOrdersService{
		result: db.PurchaseOrder{},
		err:    purchaseOrders.ExistsIdError,
	}

	router := setupPurchaseOrdersRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/purchaseOrders", requestBody)
	router.ServeHTTP(response, request)

	assert.Equal(t, 409, response.Code)
}

func decodePurchaseOrdersWebResponse(response *httptest.ResponseRecorder, responseData any) {
	responseStruct := web.Response{}
	json.Unmarshal(response.Body.Bytes(), &responseStruct)

	jsonData, _ := json.Marshal(responseStruct.Data)
	json.Unmarshal(jsonData, &responseData)
}

func setupPurchaseOrdersRouter(mockService mockPurchaseOrdersService) *gin.Engine {
	controller := NewPurchaseOrderController(mockService)

	router := gin.Default()
	router.POST("/api/v1/purchaseOrders", controller.Create())

	return router
}
