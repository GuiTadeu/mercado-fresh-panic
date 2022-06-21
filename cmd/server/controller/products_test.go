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
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

// func Test_Create_422(t *testing.T) {

// 	mockService := mockProductService{
// 		err:               nil,
// 		existsProductCode: false,
// 	}

// 	controller := NewProductController(mockService)
// 	create := controller.Create()

// 	response := httptest.NewRecorder()
// 	request, _ := gin.CreateTestContext(response)

// 	create(request)

// 	assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
// }

func Test_Create_422(t *testing.T) {

	mockService := mockProductService{
		err:               nil,
		existsProductCode: false,
	}

	controller := NewProductController(mockService)
	create := controller.Create()

	r := SetupRouter()
	r.POST("/api/v1/products", create)

	product := db.Product{}

	jsonValue, _ := json.Marshal(product)
	req, _ := http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(jsonValue))

	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
}

func Test_Create_201(t *testing.T) {

	expectedProduct := db.Product{
		Id:                      1,
		Code:                    "ABC",
		Description:             "ABC",
		Width:                   1,
		Height:                  1,
		Length:                  1,
		NetWeight:               1,
		ExpirationRate:          1,
		RecommendedFreezingTemp: 1,
		FreezingRate:            1,
		ProductTypeId:           1,
		SellerId:                1,
	}

	mockService := mockProductService{
		result:            expectedProduct,
		err:               nil,
		existsProductCode: false,
	}

	controller := NewProductController(mockService)
	create := controller.Create()

	r := SetupRouter()
	r.POST("/api/v1/products", create)

	productRequest := CreateProductRequest{
		Code:                    "ABC",
		Description:             "ABC",
		Width:                   1,
		Height:                  1,
		Length:                  1,
		NetWeight:               1,
		ExpirationRate:          1,
		RecommendedFreezingTemp: 1,
		FreezingRate:            1,
		ProductTypeId:           1,
		SellerId:                1,
	}

	jsonValue, _ := json.Marshal(productRequest)
	request, _ := http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(jsonValue))

	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	responseStruct := web.Response{}
	json.Unmarshal(response.Body.Bytes(), &responseStruct)

	responseData := db.Product{}
	mapstructure.Decode(responseStruct.Data, &responseData)

	assert.Equal(t, expectedProduct, responseData)
	assert.Equal(t, http.StatusCreated, response.Code)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}
