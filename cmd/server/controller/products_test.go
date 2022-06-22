package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/products"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TODO: ADICIONAR CHECK DE MENSAGEM DE ERRO

func Test_Create_201(t *testing.T) {

	validProduct := db.Product{
		Id:                      1,
		Code:                    "ABC",
		Description:             "ABC",
		Width:                   1.0,
		Height:                  1.0,
		Length:                  1.0,
		NetWeight:               1.0,
		ExpirationRate:          1.0,
		RecommendedFreezingTemp: 1.0,
		FreezingRate:            1.0,
		ProductTypeId:           1,
		SellerId:                1,
	}

	jsonValue, _ := json.Marshal(validProduct)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockProductService{
		result: validProduct,
		err:    nil,
	}

	router := setupRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/products", requestBody)
	router.ServeHTTP(response, request)

	responseData := db.Product{}
	decodeWebResponse(response, &responseData)

	assert.Equal(t, 201, response.Code)
	assert.Equal(t, validProduct, responseData)
}

func Test_Create_422(t *testing.T) {

	invalidProduct := db.Product{}
	jsonValue, _ := json.Marshal(invalidProduct)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockProductService{}

	router := setupRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/products", requestBody)
	router.ServeHTTP(response, request)

	assert.Equal(t, 422, response.Code)
}

func Test_Create_409(t *testing.T) {

	validProduct := db.Product{
		Id:                      1,
		Code:                    "ABC",
		Description:             "ABC",
		Width:                   1.0,
		Height:                  1.0,
		Length:                  1.0,
		NetWeight:               1.0,
		ExpirationRate:          1.0,
		RecommendedFreezingTemp: 1.0,
		FreezingRate:            1.0,
		ProductTypeId:           1,
		SellerId:                1,
	}

	jsonValue, _ := json.Marshal(validProduct)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockProductService{
		result: db.Product{},
		err:    products.ExistsProductCodeError,
	}

	router := setupRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/products", requestBody)
	router.ServeHTTP(response, request)

	assert.Equal(t, 409, response.Code)
}

func Test_GetAll_200(t *testing.T) {

	productsList := []db.Product{
		{
			Id:                      1,
			Code:                    "ABC",
			Description:             "ABC",
			Width:                   1.0,
			Height:                  1.0,
			Length:                  1.0,
			NetWeight:               1.0,
			ExpirationRate:          1.0,
			RecommendedFreezingTemp: 1.0,
			FreezingRate:            1.0,
			ProductTypeId:           1,
			SellerId:                1,
		},
		{
			Id:                      2,
			Code:                    "DEF",
			Description:             "DEF",
			Width:                   2.0,
			Height:                  2.0,
			Length:                  2.0,
			NetWeight:               2.0,
			ExpirationRate:          2.0,
			RecommendedFreezingTemp: 2.0,
			FreezingRate:            2.0,
			ProductTypeId:           2,
			SellerId:                2,
		},
		{
			Id:                      3,
			Code:                    "GHI",
			Description:             "GHI",
			Width:                   3.0,
			Height:                  3.0,
			Length:                  3.0,
			NetWeight:               3.0,
			ExpirationRate:          3.0,
			RecommendedFreezingTemp: 3.0,
			FreezingRate:            3.0,
			ProductTypeId:           3,
			SellerId:                3,
		},
	}

	mockService := mockProductService{
		result: productsList,
		err:    nil,
	}

	router := setupRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/products", nil)
	router.ServeHTTP(response, request)

	responseData := []db.Product{}
	decodeWebResponse(response, &responseData)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, productsList, responseData)
}

func Test_Get_200(t *testing.T) {

	foundProduct := db.Product{
		Id:                      1,
		Code:                    "ABC",
		Description:             "ABC",
		Width:                   1.0,
		Height:                  1.0,
		Length:                  1.0,
		NetWeight:               1.0,
		ExpirationRate:          1.0,
		RecommendedFreezingTemp: 1.0,
		FreezingRate:            1.0,
		ProductTypeId:           1,
		SellerId:                1,
	}

	mockService := mockProductService{
		result: foundProduct,
		err:    nil,
	}

	router := setupRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/products/666", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
}

func Test_Get_404(t *testing.T) {

	mockService := mockProductService{
		result: db.Product{},
		err:    products.ProductNotFoundError,
	}

	router := setupRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/products/666", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 404, response.Code)
}

// func Test_Update_200(t *testing.T) {

// 	productToUpdate := db.Product{
// 		Id:                      1,
// 		Code:                    "ABC",
// 		Description:             "ABC",
// 		Width:                   1.0,
// 		Height:                  1.0,
// 		Length:                  1.0,
// 		NetWeight:               1.0,
// 		ExpirationRate:          1.0,
// 		RecommendedFreezingTemp: 1.0,
// 		FreezingRate:            1.0,
// 		ProductTypeId:           1,
// 		SellerId:                1,
// 	}

// 	jsonValue, _ := json.Marshal(productToUpdate)
// 	requestBody := bytes.NewBuffer(jsonValue)

// 	updatedProduct := db.Product{
// 		Id:                      2,
// 		Code:                    "DEF",
// 		Description:             "DEF",
// 		Width:                   2.0,
// 		Height:                  2.0,
// 		Length:                  2.0,
// 		NetWeight:               2.0,
// 		ExpirationRate:          2.0,
// 		RecommendedFreezingTemp: 2.0,
// 		FreezingRate:            2.0,
// 		ProductTypeId:           1,
// 		SellerId:                1,
// 	}

// 	mockService := mockProductService{
// 		result: updatedProduct,
// 		err:    nil,
// 	}

// 	router := setupRouter(mockService)

// 	response := httptest.NewRecorder()
// 	request, _ := http.NewRequest("PATCH", "/api/v1/products", requestBody)
// 	router.ServeHTTP(response, request)

// 	assert.Equal(t, 409, response.Code)
// }

func decodeWebResponse(response *httptest.ResponseRecorder, responseData any) {
	responseStruct := web.Response{}
	json.Unmarshal(response.Body.Bytes(), &responseStruct)
	jsonData, _ := json.Marshal(responseStruct.Data)
	json.Unmarshal(jsonData, &responseData)
}

func setupRouter(mockService mockProductService) *gin.Engine {
	controller := NewProductController(mockService)

	router := gin.Default()
	router.POST("/api/v1/products", controller.Create())
	router.GET("/api/v1/products", controller.GetAll())
	router.GET("/api/v1/products/:id", controller.Get())
	router.PATCH("/api/v1/products/:id", controller.Update())

	return router
}
