package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	database "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/warehouses"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Warehouse_Create_201(t *testing.T) {

	validWarehouse := database.Warehouse{
		Id:                 2,
		Code:               "HSK",
		Address:            "Rua",
		Telephone:          "1145674458",
		MinimunCapacity:    4,
		MinimumTemperature: 12,
	}

	jsonValue, _ := json.Marshal(validWarehouse)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockWarehouseService{
		result: validWarehouse,
		err:    nil,
	}

	router := setupWarehouseRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/warehouses", requestBody)
	router.ServeHTTP(response, request)

	responseData := database.Warehouse{}
	decodeWarehouseWebResponse(response, &responseData)

	assert.Equal(t, 201, response.Code)
	assert.Equal(t, validWarehouse, responseData)

}

func Test_Warehouse_Create_422(t *testing.T) {

	invalidWarehouse := database.Warehouse{}
	jsonValue, _ := json.Marshal(invalidWarehouse)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockWarehouseService{}

	router := setupWarehouseRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/warehouses", requestBody)
	router.ServeHTTP(response, request)

	assert.Equal(t, 422, response.Code)

}

func Test_Warehouse_Create_409(t *testing.T) {

	validWarehouse := database.Warehouse{
		Id:                 2,
		Code:               "HSC",
		Address:            "Rua",
		Telephone:          "1145674458",
		MinimunCapacity:    4,
		MinimumTemperature: 12,
	}

	jsonValue, _ := json.Marshal(validWarehouse)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockWarehouseService{
		result: database.Warehouse{},
		err:    warehouses.ExistsWarehouseCodeError,
	}

	router := setupWarehouseRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/warehouses", requestBody)
	router.ServeHTTP(response, request)

	assert.Equal(t, 409, response.Code)

}

func Test_Warehouse_GetAll_200(t *testing.T) {

	warehousesList := []database.Warehouse{
		{
			Id:                 2,
			Code:               "HSC",
			Address:            "Rua",
			Telephone:          "1145674458",
			MinimunCapacity:    4,
			MinimumTemperature: 12,
		},
		{
			Id:                 1,
			Code:               "MARDUK",
			Address:            "MDK",
			Telephone:          "333888777",
			MinimunCapacity:    6,
			MinimumTemperature: 2.7,
		},
		{
			Id:                 3,
			Code:               "MAIDEN",
			Address:            "MDN",
			Telephone:          "111222444",
			MinimunCapacity:    2,
			MinimumTemperature: 8.9,
		},
	}

	mockService := mockWarehouseService{
		result: warehousesList,
		err:    nil,
	}

	router := setupWarehouseRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/warehouses", nil)
	router.ServeHTTP(response, request)

	responseData := []database.Warehouse{}
	decodeWarehouseWebResponse(response, &responseData)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, warehousesList, responseData)
}

func Test_Get_OneOk(t *testing.T) {

	foundWarehouse := database.Warehouse{
		Id:                 2,
		Code:               "HSC",
		Address:            "Rua",
		Telephone:          "1145674458",
		MinimunCapacity:    4,
		MinimumTemperature: 12,
	}

	mockService := mockWarehouseService{
		result: foundWarehouse,
		err:    nil,
	}

	router := setupWarehouseRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/warehouses/66", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)

}

func Test_Warehouse_Get_200(t *testing.T) {

	mockService := mockWarehouseService{
		result: database.Warehouse{},
		err:    warehouses.WarehouseNotFoundError,
	}

	router := setupWarehouseRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/warehouses/66", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 404, response.Code)
}

func Test_Warehouse_Update_200(t *testing.T) {

	warehouseToUpdate := database.Warehouse{
		Id:                 2,
		Code:               "HSC",
		Address:            "Rua",
		Telephone:          "1145674458",
		MinimunCapacity:    4,
		MinimumTemperature: 12,
	}

	jsonValue, _ := json.Marshal(warehouseToUpdate)
	requestBody := bytes.NewBuffer(jsonValue)

	updatedWarehouse := database.Warehouse{
		Id:                 1,
		Code:               "MARDUK",
		Address:            "MDK",
		Telephone:          "333888777",
		MinimunCapacity:    6,
		MinimumTemperature: 27,
	}

	mockService := mockWarehouseService{
		result: updatedWarehouse,
		err:    nil,
	}

	router := setupWarehouseRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/v1/warehouses/1", requestBody)
	router.ServeHTTP(response, request)

	responseData := database.Warehouse{}
	decodeWarehouseWebResponse(response, &responseData)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, updatedWarehouse, responseData)

}

func Test_Warehouse_Update_404(t *testing.T) {

	warehouseToUpdate := database.Warehouse{
		Id:                 2,
		Code:               "HSC",
		Address:            "hpsn",
		Telephone:          "1145674458",
		MinimunCapacity:    4,
		MinimumTemperature: 12,
	}

	jsonValue, _ := json.Marshal(warehouseToUpdate)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockWarehouseService{
		result: database.Warehouse{},
		err:    warehouses.WarehouseNotFoundError,
	}

	router := setupWarehouseRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/v1/warehouses/1", requestBody)
	router.ServeHTTP(response, request)

	responseData := []database.Warehouse{}
	decodeWarehouseWebResponse(response, &responseData)

	assert.Equal(t, 404, response.Code)
}

func Test_Warehouse_Delete_204(t *testing.T) {

	mockService := mockWarehouseService{
		result: database.Warehouse{},
		err:    nil,
	}

	router := setupWarehouseRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", "/api/v1/warehouses/1", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 204, response.Code)
}

func Test_Warehouse_Delete_404(t *testing.T) {

	mockService := mockWarehouseService{
		result: database.Warehouse{},
		err:    warehouses.WarehouseNotFoundError,
	}

	router := setupWarehouseRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", "/api/v1/products/1", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 404, response.Code)
}

func decodeWarehouseWebResponse(response *httptest.ResponseRecorder, responseData any) {
	responseStruct := web.Response{}
	json.Unmarshal(response.Body.Bytes(), &responseStruct)

	jsonData, _ := json.Marshal(responseStruct.Data)
	json.Unmarshal(jsonData, &responseData)
}

func setupWarehouseRouter(mockService mockWarehouseService) *gin.Engine {
	controller := NewWarehouseController(mockService)

	router := gin.Default()
	router.POST("/api/v1/warehouses", controller.Create())
	router.GET("/api/v1/warehouses", controller.GetAll())
	router.GET("/api/v1/warehouses/:id", controller.Get())
	router.PATCH("/api/v1/warehouses/:id", controller.Update())
	router.DELETE("/api/v1/warehouses/:id", controller.Delete())

	return router
}