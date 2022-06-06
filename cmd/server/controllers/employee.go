package controller

import (
	"github.com/GuiTadeu/mercado-fresh-panic/internal/employee"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type requestEmployee struct {
	Id           uint64 `json:"id"`
	CardNumberId uint64 `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  uint64 `json:"warehouse_id"`
}

type employeeController struct {
	employeeService employee.EmployeeService
}

func NewEmployee(e employee.EmployeeService) *employeeController {
	return &employeeController{
		employeeService: e,
	}
}

func (c *employeeController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		employees, err := c.employeeService.GetAll()

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, web.NewResponse(200, employees, ""))
	}
}

func (e *employeeController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requestEmployee

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity,
				gin.H{
					"error":   "Unprocessable Entity",
					"message": "Invalid inputs. Please check your inputs"})
			return
		}
		e, _ := e.employeeService.Create(e.employeeService.GetNextId(), req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)
		c.JSON(http.StatusCreated, gin.H{"data": e})
	}
}

func (c *employeeController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		_, err = c.employeeService.Get(uint64(id))

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = c.employeeService.Delete(uint64(id))
		ctx.JSON(204, web.NewResponse(204, nil, ""))
	}
}
