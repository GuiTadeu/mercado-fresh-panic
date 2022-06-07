package controller

import (
	"net/http"
	"strconv"

	"github.com/GuiTadeu/mercado-fresh-panic/internal/warehouse"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"	
)

type warehouseRequest struct {
	Id                  uint64  `json:"id"`
	Code      			string  `json:"code"`
	Address             string  `json:"endereco"`
	Telephone           string  `json:"number"`
	MinimunCapacity     uint32   `json:"minimun_capacity"`
	MinimunTemperature float32 `json:"minimun_temperature"`
}

type warehouseController struct {
	warehouseService warehouse.WarehouseService 
}

func NewWarehouseController(warehouse warehouse.WarehouseService) *warehouseController {
	return &warehouseController{
		warehouseService: warehouse,
	}
}

func (c *warehouseController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		warehouse, err := c.warehouseService.GetAll()

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H {
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, warehouse, ""))
	}
}

func (r *warehouseController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req warehouseRequest

		err := ctx.ShouldBindJSON(&req)

		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		addedWarehouse, err := r.warehouseService.Create(req.Id, req.Code, req.Address, req.Telephone, req.MinimunCapacity, req.MinimunTemperature)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, addedWarehouse, ""))
	}
}


func (c *warehouseController) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		warehouse, err := c.warehouseService.Get(uint64(id))

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, warehouse, ""))
	}
}

func (c *warehouseController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		_, err = c.warehouseService.Get(uint64(id))

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = c.warehouseService.Delete(uint64(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
	}
}