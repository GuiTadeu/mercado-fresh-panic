package controller

import (
	"github.com/GuiTadeu/mercado-fresh-panic/internal/employee"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EmployeeController struct {
	employeeService employee.EmployeeService
}

func NewEmployee(s employee.EmployeeService) *EmployeeController {
	return &EmployeeController{
		employeeService: s,
	}
}

type request struct {
	CardNumberId string `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	WarehouseId  uint64 `json:"warehouse_id" binding:"required"`
}

func (c *EmployeeController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity,
				gin.H{
					"error":   "Unprocessable Entity",
					"message": "Invalid inputs. Please check your inputs"})
			return
		}
		e, _ := c.employeeService.Create(req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)
		ctx.JSON(http.StatusCreated, gin.H{"data": e})
	}
}

func (c *EmployeeController) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		employee, err := c.employeeService.Get(uint64(id))

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(200, employee, ""))

	}
}

func (c *EmployeeController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		var req request
		ctx.ShouldBindJSON(&req)

		e, err := c.employeeService.Update(uint64(id), req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"error":   "VALIDATEERR-1",
					"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": e})
	}
}

func (c *EmployeeController) GetAll() gin.HandlerFunc {
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

func (c *EmployeeController) Delete() gin.HandlerFunc {
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
