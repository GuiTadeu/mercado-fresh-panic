package controller

import (
	"encoding/json"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func Test_Employee_Create_201(t *testing.T) {
}

func Test_Employee_Create_422(t *testing.T) {

}

func Test_Employee_Create_409(t *testing.T) {

}

func Test_Employee_GetAll_200(t *testing.T) {

}

func Test_Employee_Get_200(t *testing.T) {

}

func Test_Employee_Get_404(t *testing.T) {

}

func Test_Employee_Update_200(t *testing.T) {

}

func Test_Employee_Update_404(t *testing.T) {

}

func Test_Employee_Delete_204(t *testing.T) {

}

func Test_Employee_Delete_404(t *testing.T) {

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

