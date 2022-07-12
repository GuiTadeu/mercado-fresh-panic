package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	inboundorders "github.com/GuiTadeu/mercado-fresh-panic/internal/inboundOrders"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)
func Test_Inbound_Order_Create_201(t *testing.T) {
	validInboundOrder := db.InboundOrder{
		Id: 1,
		OrderDate: "2021-04-04",
		OrderNumber: "order#1",
		EmployeeId: 1,
		ProductBatchId: 1,
		WarehouseId: 1,
	}

	jsonValue, _ := json.Marshal(validInboundOrder)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockInboundOrderService{
		result: validInboundOrder,
		err: nil,
	}

	router := setupInboundOrderRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/inboundOrders", requestBody)
	router.ServeHTTP(response, request)

	responseData := db.InboundOrder{}
	decodeInboundOrderWebResponse(response, &responseData)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, validInboundOrder, responseData)

}

func Test_Inbound_Order_Create_Employee_Not_Found_409(t *testing.T) {

	expectedError := inboundorders.EmployeeNotFoundError

	validInboundOrder := db.InboundOrder{
		Id: 1,
		OrderDate: "2021-04-04",
		OrderNumber: "order#1",
		EmployeeId: 1,
		ProductBatchId: 1,
		WarehouseId: 1,
	}

	jsonValue, _ := json.Marshal(validInboundOrder)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockInboundOrderService{
		err: expectedError,
	}

	router := setupInboundOrderRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/inboundOrders", requestBody)
	router.ServeHTTP(response, request)

	responseData := web.Response{}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	assert.Equal(t, http.StatusConflict, response.Code)
	assert.Equal(t, expectedError.Error(), responseData.Error)
}

func Test_Inbound_Order_Create_Warehouse_Not_Found_409(t *testing.T) {

	expectedError := inboundorders.WarehouseNotFoundError

	validInboundOrder := db.InboundOrder{
		Id: 1,
		OrderDate: "2021-04-04",
		OrderNumber: "order#1",
		EmployeeId: 1,
		ProductBatchId: 1,
		WarehouseId: 1,
	}

	jsonValue, _ := json.Marshal(validInboundOrder)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockInboundOrderService{
		err: expectedError,
	}

	router := setupInboundOrderRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/inboundOrders", requestBody)
	router.ServeHTTP(response, request)

	responseData := web.Response{}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	assert.Equal(t, http.StatusConflict, response.Code)
	assert.Equal(t, expectedError.Error(), responseData.Error)
}

func Test_Inbound_Order_Create_500(t *testing.T) {

	expectedError := errors.New("Internal Server Error")

	validInboundOrder := db.InboundOrder{
		Id: 1,
		OrderDate: "2021-04-04",
		OrderNumber: "order#1",
		EmployeeId: 1,
		ProductBatchId: 1,
		WarehouseId: 1,
	}

	jsonValue, _ := json.Marshal(validInboundOrder)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockInboundOrderService{
		err: expectedError,
	}

	router := setupInboundOrderRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/inboundOrders", requestBody)
	router.ServeHTTP(response, request)

	responseData := web.Response{}
	json.Unmarshal(response.Body.Bytes(), &responseData)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, expectedError.Error(), responseData.Error)
}

func decodeInboundOrderWebResponse(response *httptest.ResponseRecorder, responseData any) {
	responseStruct := web.Response{}
	json.Unmarshal(response.Body.Bytes(), &responseStruct)

	jsonData, _ := json.Marshal(responseStruct.Data)
	json.Unmarshal(jsonData, &responseData)
}

func setupInboundOrderRouter(mockService mockInboundOrderService) *gin.Engine {
	controller := NewInboundOrderController(mockService)

	router := gin.Default()
	router.POST("/api/v1/inboundOrders", controller.Create())

	return router
}