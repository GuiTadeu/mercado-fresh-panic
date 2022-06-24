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

	assert.Equal(t, 201, response.Code)
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

}

func Test_Buyer_Get_200(t *testing.T) {

}

func Test_Buyer_Get_404(t *testing.T) {

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

	assert.Equal(t, 200, response.Code)
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

	assert.Equal(t, 404, response.Code)

}

func Test_Buyer_Delete_204(t *testing.T) {

}

func Test_Buyer_Delete_404(t *testing.T) {

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
