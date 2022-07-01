package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/sections"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Section_Create_201(t *testing.T) {

	validSection := db.Section{
		Number:             4,
		CurrentTemperature: 1.0,
		MinimumTemperature: 1.0,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	jsonValue, _ := json.Marshal(validSection)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockSectionService{
		result: validSection,
		err:    nil,
	}

	router := setupSectionRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/sections", requestBody)
	router.ServeHTTP(response, request)

	responseData := db.Section{}
	decodeSectionWebResponse(response, &responseData)

	assert.Equal(t, 201, response.Code)
	assert.Equal(t, validSection, responseData)
}

func Test_Section_Create_422(t *testing.T) {

	invalidSection := db.Section{}
	jsonValue, _ := json.Marshal(invalidSection)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockSectionService{}

	router := setupSectionRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/sections", requestBody)
	router.ServeHTTP(response, request)

	assert.Equal(t, 422, response.Code)
}

func Test_Section_Create_409(t *testing.T) {

	validSection := db.Section{
		Id:                 1,
		Number:             4,
		CurrentTemperature: 1.0,
		MinimumTemperature: 1.0,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	jsonValue, _ := json.Marshal(validSection)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockSectionService{
		result: db.Section{},
		err:    sections.ExistsSectionNumberError,
	}

	router := setupSectionRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/sections", requestBody)
	router.ServeHTTP(response, request)

	assert.Equal(t, 409, response.Code)
}

func Test_Section_GetAll_200(t *testing.T) {

	sectionList := []db.Section{
		{
			Id:                 1,
			Number:             4,
			CurrentTemperature: 1.0,
			MinimumTemperature: 1.0,
			CurrentCapacity:    1,
			MinimumCapacity:    1,
			MaximumCapacity:    1,
			WarehouseId:        1,
			ProductTypeId:      1,
		},
		{
			Id:                 2,
			Number:             5,
			CurrentTemperature: 1.5,
			MinimumTemperature: 1.5,
			CurrentCapacity:    2,
			MinimumCapacity:    2,
			MaximumCapacity:    2,
			WarehouseId:        2,
			ProductTypeId:      2,
		},
		{
			Id:                 3,
			Number:             6,
			CurrentTemperature: 2.0,
			MinimumTemperature: 2.0,
			CurrentCapacity:    3,
			MinimumCapacity:    3,
			MaximumCapacity:    3,
			WarehouseId:        3,
			ProductTypeId:      3,
		},
	}

	mockService := mockSectionService{
		result: sectionList,
		err:    nil,
	}

	router := setupSectionRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/sections", nil)
	router.ServeHTTP(response, request)

	responseData := []db.Section{}
	decodeSectionWebResponse(response, &responseData)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, sectionList, responseData)
}

func Test_Section_Get_200(t *testing.T) {

	foundSection := db.Section{
		Id:                 1,
		Number:             4,
		CurrentTemperature: 1.0,
		MinimumTemperature: 1.0,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	mockService := mockSectionService{
		result: foundSection,
		err:    nil,
	}

	router := setupSectionRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/sections/666", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
}

func Test_Section_Get_404(t *testing.T) {

	mockService := mockSectionService{
		result: db.Section{},
		err:    sections.SectionNotFoundError,
	}

	router := setupSectionRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/sections/666", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 404, response.Code)
}

func Test_Section_Get_500(t *testing.T) {

	mockService := mockSectionService{
		result: db.Section{},
		err:    errors.New("internal server error"),
	}

	router := setupSectionRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/sections/666", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 500, response.Code)
}
func Test_Section_Update_200(t *testing.T) {

	sectionToUpdate := db.Section{
		Id:                 1,
		Number:             4,
		CurrentTemperature: 1.0,
		MinimumTemperature: 1.0,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	jsonValue, _ := json.Marshal(sectionToUpdate)
	requestBody := bytes.NewBuffer(jsonValue)

	updatedSection := db.Section{
		Id:                 2,
		Number:             5,
		CurrentTemperature: 1.5,
		MinimumTemperature: 1.5,
		CurrentCapacity:    2,
		MinimumCapacity:    2,
		MaximumCapacity:    2,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	mockService := mockSectionService{
		result: updatedSection,
		err:    nil,
	}

	router := setupSectionRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/v1/sections/1", requestBody)
	router.ServeHTTP(response, request)

	responseData := db.Section{}
	decodeSectionWebResponse(response, &responseData)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, updatedSection, responseData)
}

func Test_Section_Update_404(t *testing.T) {

	sectionToUpdate := db.Section{
		Id:                 1,
		Number:             4,
		CurrentTemperature: 1.0,
		MinimumTemperature: 1.0,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	jsonValue, _ := json.Marshal(sectionToUpdate)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockSectionService{
		result: db.Section{},
		err:    sections.SectionNotFoundError,
	}

	router := setupSectionRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/v1/sections/1", requestBody)
	router.ServeHTTP(response, request)

	responseData := db.Section{}
	decodeSectionWebResponse(response, &responseData)

	assert.Equal(t, 404, response.Code)
}

func Test_Section_Delete_204(t *testing.T) {

	mockService := mockSectionService{
		result: db.Section{},
		err:    nil,
	}

	router := setupSectionRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", "/api/v1/sections/1", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 204, response.Code)
}

func Test_Section_Delete_404(t *testing.T) {

	mockService := mockSectionService{
		result: db.Section{},
		err:    sections.SectionNotFoundError,
	}

	router := setupSectionRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", "/api/v1/sections/123", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 404, response.Code)
}

func decodeSectionWebResponse(response *httptest.ResponseRecorder, responseData any) {
	responseStruct := web.Response{}
	json.Unmarshal(response.Body.Bytes(), &responseStruct)

	jsonData, _ := json.Marshal(responseStruct.Data)
	json.Unmarshal(jsonData, &responseData)
}

func setupSectionRouter(mockService mockSectionService) *gin.Engine {
	controller := NewSectionController(mockService)

	router := gin.Default()
	router.POST("/api/v1/sections", controller.Create())
	router.GET("/api/v1/sections", controller.GetAll())
	router.GET("/api/v1/sections/:id", controller.Get())
	router.PATCH("/api/v1/sections/:id", controller.Update())
	router.DELETE("/api/v1/sections/:id", controller.Delete())

	return router
}
