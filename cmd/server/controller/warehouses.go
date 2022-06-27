package controller

import (
	"net/http"
	"strconv"

	"github.com/GuiTadeu/mercado-fresh-panic/internal/warehouses"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
)

type updateWarehouseRequest struct {
	Code               string  `json:"warehouse_code"`
	Address            string  `json:"address"`
	Telephone          string  `json:"telephone"`
	MinimunCapacity    uint32  `json:"minimun_capacity"`
	MinimunTemperature float32 `json:"minimun_temperature"`
}

type createWarehouseRequest struct {
	Code               string  `json:"warehouse_code" binding:"required"`
	Address            string  `json:"address" binding:"required"`
	Telephone          string  `json:"telephone" binding:"required"`
	MinimunCapacity    uint32  `json:"minimun_capacity" binding:"required"`
	MinimunTemperature float32 `json:"minimun_temperature" binding:"required"`
}

type warehouseController struct {
	warehouseService warehouses.WarehouseService
}

func NewWarehouseController(warehouse warehouses.WarehouseService) *warehouseController {
	return &warehouseController{
		warehouseService: warehouse,
	}
}

func (c *warehouseController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		warehouse, err := c.warehouseService.GetAll()

		if err != nil {
			status := warehouseErrorHandler(err, ctx)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, warehouse, ""))
	}
}

func (c *warehouseController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createWarehouseRequest

		err := ctx.ShouldBindJSON(&req)

		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			return
		}

		addedWarehouse, err := c.warehouseService.Create(req.Code, req.Address, req.Telephone, req.MinimunCapacity, req.MinimunTemperature)

		if err != nil {
			status := warehouseErrorHandler(err, ctx)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, addedWarehouse, ""))
	}
}

func (c *warehouseController) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			status := warehouseErrorHandler(err, ctx)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}
		warehouse, err := c.warehouseService.Get(uint64(id))

		if err != nil {
			status := warehouseErrorHandler(err, ctx)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, warehouse, ""))
	}
}

func (c *warehouseController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			status := warehouseErrorHandler(err, ctx)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}
		_, err = c.warehouseService.Get(uint64(id))

		if err != nil {
			status := warehouseErrorHandler(err, ctx)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}

		err = c.warehouseService.Delete(uint64(id))
		if err != nil {
			status := warehouseErrorHandler(err, ctx)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, ""))

	}
}

func (c *warehouseController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request updateWarehouseRequest
		err := ctx.ShouldBindJSON(&request)
		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "warehouses id binding error"))
			return
		}
		updatedWarehouse, err := c.warehouseService.Update(
			id,
			request.Code,
			request.Address,
			request.Telephone,
			request.MinimunCapacity,
			request.MinimunTemperature,
		)
		if err != nil {
			status := warehouseErrorHandler(err, ctx)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, updatedWarehouse, ""))
	}
}

func warehouseErrorHandler(err error, ctx *gin.Context) int {
	switch err {
	case warehouses.WarehouseNotFoundError:
		return http.StatusNotFound
	case warehouses.ExistsWarehouseCodeError:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
