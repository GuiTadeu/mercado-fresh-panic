package controller

import (
	"github.com/GuiTadeu/mercado-fresh-panic/internal/employee"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EmployeeController struct {
	service employee.EmployeeService
}

func NewEmployee(s employee.EmployeeService) *EmployeeController {
	return &EmployeeController{
		service: s,
	}
}

func (c *EmployeeController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"error":   "VALIDATEERR-1",
					"message": "Invalid inputs. Please check your inputs"})
			return
		}
		e, _ := c.service.Create(req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)
		ctx.JSON(http.StatusOK, gin.H{"criado": e})
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

		employee, err := c.service.Get(uint64(id))

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
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"error":   "VALIDATEERR-1",
					"message": "Invalid inputs. Please check your inputs"})
			return
		}
		e, err := c.service.Update(uint64(id), req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)

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

type request struct {
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  uint64 `json:"warehouse_id"`
}
