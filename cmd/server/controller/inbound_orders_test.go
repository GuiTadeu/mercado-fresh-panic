package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
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