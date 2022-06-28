package controller

import (
	"bytes"
	"encoding/json"
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/buyers"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Buyer_Create_201(t *testing.T) {

	validBuyer := db.Buyer{
		Id:           1,
		CardNumberId: "ABC",
		FirstName:    "Melii",
		LastName:     "Developers",
	}

	jsonValue, _ := json.Marshal(validBuyer)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockBuyerService{
		result: validBuyer,
		err:    nil,
	}

	router := setupBuyerRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/buyers", requestBody)
	router.ServeHTTP(response, request)

	responseData := db.Buyer{}
	decodeBuyerWebResponse(response, &responseData)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, validBuyer, responseData)
}

func Test_Buyer_Create_422(t *testing.T) {

	invalidBuyer := db.Buyer{}
	jsonValue, _ := json.Marshal(invalidBuyer)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockBuyerService{}

	router := setupBuyerRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/buyers", requestBody)
	router.ServeHTTP(response, request)

	assert.Equal(t, 422, response.Code)
}

func Test_Buyer_Create_409(t *testing.T) {
	validBuyer := db.Buyer{
		Id:           1,
		CardNumberId: "ABC",
		FirstName:    "Melii",
		LastName:     "Developers",
	}

	jsonValue, _ := json.Marshal(validBuyer)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockBuyerService{
		result: db.Buyer{},
		err:    buyers.ExistsBuyerCardNumberIdError,
	}

	router := setupBuyerRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/buyers", requestBody)
	router.ServeHTTP(response, request)

	assert.Equal(t, 409, response.Code)
}

func Test_Buyer_GetAll_200(t *testing.T) {
	buyers := []db.Buyer{
		{
			Id:           1,
			CardNumberId: "5",
			FirstName:    "NameT1",
			LastName:     "LastNameT1",
		},
		{
			Id:           2,
			CardNumberId: "10",
			FirstName:    "NameT2",
			LastName:     "LastNameT2",
		},
		{
			Id:           3,
			CardNumberId: "15",
			FirstName:    "NameT3",
			LastName:     "LastNameT3",
		},
	}

	mockService := mockBuyerService{
		result: buyers,
		err:    nil,
	}

	router := setupBuyerRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/buyers", nil)
	router.ServeHTTP(response, request)

	responseData := []db.Buyer{}
	decodeBuyerWebResponse(response, &responseData)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, buyers, responseData)
}

func Test_Buyer_Get_200(t *testing.T) {
	buyer := db.Buyer{
		Id:           1,
		CardNumberId: "5",
		FirstName:    "NameT1",
		LastName:     "LastNameT1",
	}

	mockService := mockBuyerService{
		result: buyer,
		err:    nil,
	}

	router := setupBuyerRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/buyers/1", nil)
	router.ServeHTTP(response, request)

	responseData := db.Buyer{}
	decodeBuyerWebResponse(response, &responseData)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, buyer, responseData)
}

func Test_Buyer_Get_404(t *testing.T) {

	mockService := mockBuyerService{
		result: db.Buyer{},
		err:    buyers.BuyerNotFoundError,
	}

	expectedError := buyers.BuyerNotFoundError.Error()

	router := setupBuyerRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/buyers/1", nil)
	router.ServeHTTP(response, request)

	responseStruct := web.Response{}
	json.Unmarshal(response.Body.Bytes(), &responseStruct)

	assert.Equal(t, http.StatusNotFound, response.Code)
	assert.Equal(t, expectedError, responseStruct.Error)
}

func Test_Buyer_Update_200(t *testing.T) {

	buyerToUpdate := db.Buyer{
		Id:           1,
		CardNumberId: "ABC",
		FirstName:    "Melii",
		LastName:     "Developers",
	}

	jsonValue, _ := json.Marshal(buyerToUpdate)
	requestBody := bytes.NewBuffer(jsonValue)

	updatedBuyer := db.Buyer{
		Id:           2,
		CardNumberId: "BCD",
		FirstName:    "Todo ",
		LastName:     "Mundo em Panic",
	}

	mockService := mockBuyerService{
		result: updatedBuyer,
		err:    nil,
	}
	router := setupBuyerRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/v1/buyers/1", requestBody)
	router.ServeHTTP(response, request)

	responseData := db.Buyer{}
	decodeBuyerWebResponse(response, &responseData)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, updatedBuyer, responseData)
}

func Test_Buyer_Update_404(t *testing.T) {
	buyerToUpdate := db.Buyer{
		Id:           1,
		CardNumberId: "ABC",
		FirstName:    "Melii",
		LastName:     "Developers",
	}

	jsonValue, _ := json.Marshal(buyerToUpdate)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockBuyerService{
		result: db.Buyer{},
		err:    buyers.BuyerNotFoundError,
	}
	router := setupBuyerRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/v1/buyers/1", requestBody)
	router.ServeHTTP(response, request)

	responseData := db.Buyer{}
	decodeBuyerWebResponse(response, &responseData)

	assert.Equal(t, http.StatusNotFound, response.Code)

}

func Test_Buyer_Delete_204(t *testing.T) {
	mockService := mockBuyerService{
		result: db.Buyer{},
		err:    nil,
	}

	router := setupBuyerRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", "/api/v1/buyers/1", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusNoContent, response.Code)
}

func Test_Buyer_Delete_404(t *testing.T) {
	mockService := mockBuyerService{
		result: db.Product{},
		err:    buyers.BuyerNotFoundError,
	}

	router := setupBuyerRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", "/api/v1/buyers/1", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusNotFound, response.Code)
}

func decodeBuyerWebResponse(response *httptest.ResponseRecorder, responseData any) {
	responseStruct := web.Response{}
	json.Unmarshal(response.Body.Bytes(), &responseStruct)

	jsonData, _ := json.Marshal(responseStruct.Data)
	json.Unmarshal(jsonData, &responseData)
}

func setupBuyerRouter(mockService mockBuyerService) *gin.Engine {
	controller := NewBuyerController(mockService)

	router := gin.Default()
	router.POST("/api/v1/buyers", controller.Create())
	router.GET("/api/v1/buyers", controller.GetAll())
	router.GET("/api/v1/buyers/:id", controller.Get())
	router.PATCH("/api/v1/buyers/:id", controller.Update())
	router.DELETE("/api/v1/buyers/:id", controller.Delete())

	return router
}
