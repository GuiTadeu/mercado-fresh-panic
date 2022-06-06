package controller

import (
	employee2 "github.com/GuiTadeu/mercado-fresh-panic/internal/employee"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmployeeController struct {
	service employee2.Service
}

func NewEmployee(e employee2.Service) *EmployeeController {
	return &EmployeeController{
		service: e,
	}
}

type employee struct {
	CardNumberId uint64
	FirstName    string
	LastName     string
	WarehouseId  uint64
}

func (e *EmployeeController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requestEmployee
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"error":   "VALIDATEERR-1",
					"message": "Invalid inputs. Please check your inputs"})
			return
		}
		e, _ := e.service.Create(req.Id, req.CardNumberId, req.FirstName, req.LastName, req.WarehouseId)
		c.JSON(http.StatusOK, gin.H{"criado": e})
	}
}

type requestEmployee struct {
	Id           uint64 `json:"id"`
	CardNumberId uint64 `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  uint64 `json:"warehouse_id"`
}
