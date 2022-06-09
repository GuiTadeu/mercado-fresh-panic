package controller

import (
	"github.com/GuiTadeu/mercado-fresh-panic/internal/employees"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EmployeeController struct {
	employeeService employees.EmployeeService
}

func NewEmployeeController(s employees.EmployeeService) *EmployeeController {
	return &EmployeeController{
		employeeService: s,
	}
}

type CreateEmployeeRequest struct {
	CardNumberId string `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	WarehouseId  uint64 `json:"warehouse_id" binding:"required"`
}

type UpdateEmployeeRequest struct {
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  uint64 `json:"warehouse_id"`
}

func (c *EmployeeController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateEmployeeRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}
		employee, err := c.employeeService.Create(req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)

		if err != nil {
			status, header := employeeErrorHandler(err, ctx)
			ctx.JSON(status, header)
			return
		}

		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, employee, ""))
	}
}

func (c *EmployeeController) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		employee, err := c.employeeService.Get(id)

		if err != nil {
			status, header := employeeErrorHandler(err, ctx)
			ctx.JSON(status, header)
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, employee, ""))

	}
}

func (c *EmployeeController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var req UpdateEmployeeRequest

		err = ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		employee, err := c.employeeService.Update(id, req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)

		if err != nil {
			status, header := employeeErrorHandler(err, ctx)
			ctx.JSON(status, header)
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, employee, ""))
	}
}

func (c *EmployeeController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		employees, err := c.employeeService.GetAll()

		if err != nil {
			status, header := employeeErrorHandler(err, ctx)
			ctx.JSON(status, header)
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, employees, ""))
	}
}

func (c *EmployeeController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = c.employeeService.Delete(id)
		if err != nil {
			status, header := employeeErrorHandler(err, ctx)
			ctx.JSON(status, header)
			return
		}
		ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, ""))
	}
}

func employeeErrorHandler(err error, ctx *gin.Context) (int, gin.H) {
	switch err {

	case employees.ExistsCardNumberIdError:
		return http.StatusConflict, gin.H{"error": err.Error()}

	case employees.EmployeeNotFoundError:
		return http.StatusNotFound, gin.H{"error": err.Error()}
	default:
		return http.StatusInternalServerError, gin.H{"error": err.Error()}
	}
}
