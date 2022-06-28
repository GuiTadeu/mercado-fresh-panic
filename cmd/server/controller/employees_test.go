package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/employees"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Employee_Create_201(t *testing.T) {

	validEmployee := db.Employee{
		Id:           1,
		FirstName:    "criando",
		LastName:     "teste",
		CardNumberId: "1",
		WarehouseId:  1,
	}

	jsonValue, _ := json.Marshal(validEmployee)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockEmployeeService{
		result: validEmployee,
		err:    nil,
	}

	router := setupEmployeeRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/employees", requestBody)
	router.ServeHTTP(response, request)

	responseData := db.Employee{}
	decodeWebResponse(response, &responseData)

	assert.Equal(t, 201, response.Code)
	assert.Equal(t, validEmployee, responseData)
}

func Test_Employee_Create_422(t *testing.T) {

	invalidEmployee := db.Employee{}
	jsonValue, _ := json.Marshal(invalidEmployee)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockEmployeeService{}

	router := setupEmployeeRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/employees", requestBody)
	router.ServeHTTP(response, request)

	assert.Equal(t, 422, response.Code)
}

func Test_Employee_Create_409(t *testing.T) {

	validEmployee := db.Employee{
		Id:           1,
		FirstName:    "criando",
		LastName:     "teste",
		CardNumberId: "1",
		WarehouseId:  1,
	}

	jsonValue, _ := json.Marshal(validEmployee)
	requestBody := bytes.NewBuffer(jsonValue)

	mockService := mockEmployeeService{
		result: db.Employee{},
		err:    employees.ExistsCardNumberIdError,
	}

	router := setupEmployeeRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/api/v1/employees", requestBody)
	router.ServeHTTP(response, request)

	assert.Equal(t, 409, response.Code)
}

func Test_Employee_GetAll_200(t *testing.T) {

	employeeList := []db.Employee{
		{
			Id:           1,
			FirstName:    "criando",
			LastName:     "teste",
			CardNumberId: "1",
			WarehouseId:  1,
		},
		{
			Id:           2,
			FirstName:    "criando",
			LastName:     "teste",
			CardNumberId: "1",
			WarehouseId:  1,
		},
		{
			Id:           3,
			FirstName:    "criando",
			LastName:     "teste",
			CardNumberId: "1",
			WarehouseId:  1,
		},
	}

	mockService := mockEmployeeService{
		result: employeeList,
		err:    nil,
	}

	router := setupEmployeeRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/employees", nil)
	router.ServeHTTP(response, request)

	responseData := []db.Employee{}
	decodeEmployeeWebResponse(response, &responseData)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, employeeList, responseData)
}

func Test_Employee_Get_200(t *testing.T) {

	foundEmployee := db.Employee{
		Id:           1,
		FirstName:    "criando",
		LastName:     "teste",
		CardNumberId: "1",
		WarehouseId:  1,
	}

	mockService := mockEmployeeService{
		result: foundEmployee,
		err:    nil,
	}

	router := setupEmployeeRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/employees/100", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
}

func Test_Employee_Get_404(t *testing.T) {

	mockService := mockEmployeeService{
		result: db.Employee{},
		err:    employees.EmployeeNotFoundError,
	}

	router := setupEmployeeRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/employees/100", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 404, response.Code)
}

func Test_Employee_Update_200(t *testing.T) {
	employeeToUpdate := db.Employee{
		Id: 10,
		CardNumberId: "1",
		FirstName: "Nelson",
		LastName: "Nerd",
		WarehouseId: 1,
	}

	jsonValue, _ := json.Marshal(employeeToUpdate)
	requestBody := bytes.NewBuffer(jsonValue)

	updatedEmployee := db.Employee{
		Id: 10,
		CardNumberId: "5",
		FirstName: "Nelson",
		LastName: "Lord",
		WarehouseId: 3,
	}

	mockEmployeeService := mockEmployeeService{
		result: updatedEmployee,
		err: nil,
	}

	router := setupEmployeeRouter(mockEmployeeService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/v1/employees/10", requestBody)
	router.ServeHTTP(response, request)

	responseData := db.Employee{}
	decodeEmployeeWebResponse(response, &responseData)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, updatedEmployee, responseData)
}

func Test_Employee_Update_404(t *testing.T) {
	employeeToUpdate := db.Employee{
		Id: 10,
		CardNumberId: "1",
		FirstName: "Nelson",
		LastName: "Nerd",
		WarehouseId: 1,
	}

	expectedError := employees.EmployeeNotFoundError.Error()

	jsonValue, _ := json.Marshal(employeeToUpdate)
	requestBody := bytes.NewBuffer(jsonValue)

	mockEmployeeService := mockEmployeeService{
		result: db.Product{},
		err: employees.EmployeeNotFoundError,
	}

	router := setupEmployeeRouter(mockEmployeeService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("PATCH", "/api/v1/employees/10", requestBody)
	router.ServeHTTP(response, request)

	responseStruct := web.Response{}
	json.Unmarshal(response.Body.Bytes(), &responseStruct)

	assert.Equal(t, http.StatusNotFound, response.Code)
	assert.Equal(t, expectedError, responseStruct.Error)
}

func Test_Employee_Delete_204(t *testing.T) {

	mockService := mockEmployeeService{
		result: db.Employee{},
		err:    nil,
	}

	router := setupEmployeeRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", "/api/v1/employees/1", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 204, response.Code)
}

func Test_Employee_Delete_404(t *testing.T) {

	mockService := mockEmployeeService{
		result: db.Employee{},
		err:    employees.EmployeeNotFoundError,
	}

	router := setupEmployeeRouter(mockService)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", "/api/v1/employees/1", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 404, response.Code)
}

func decodeEmployeeWebResponse(response *httptest.ResponseRecorder, responseData any) {
	responseStruct := web.Response{}
	json.Unmarshal(response.Body.Bytes(), &responseStruct)

	jsonData, _ := json.Marshal(responseStruct.Data)
	json.Unmarshal(jsonData, &responseData)
}

func setupEmployeeRouter(mockService mockEmployeeService) *gin.Engine {
	controller := NewEmployeeController(mockService)

	router := gin.Default()
	router.POST("/api/v1/employees", controller.Create())
	router.GET("/api/v1/employees", controller.GetAll())
	router.GET("/api/v1/employees/:id", controller.Get())
	router.PATCH("/api/v1/employees/:id", controller.Update())
	router.DELETE("/api/v1/employees/:id", controller.Delete())

	return router
}
