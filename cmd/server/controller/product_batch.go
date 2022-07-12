package controller

import (
	"net/http"
	"strconv"

	"github.com/GuiTadeu/mercado-fresh-panic/internal/products/batches"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
)

type CreateProductBatchRequest struct {
	Id                 uint64  `json:"id"`
	Number             uint64  `json:"batch_number"`
	CurrentQuantity    uint64  `json:"current_quantity"`
	CurrentTemperature float32 `json:"current_temperature"`
	DueDate            string  `json:"due_date"`
	InitialQuantity    uint64  `json:"initial_quantity"`
	ManufacturingDate  string  `json:"manufacturing_date"`
	ManufacturingHour  string  `json:"manufacturing_hour"`
	MinimumTemperature float32 `json:"minimum_temperature"`
	ProductId          uint64  `json:"product_id"`
	SectionId          uint64  `json:"section_id"`
}

type productBatchController struct {
	productBatchService batches.ProductBatchService
}

func NewProductBatchController(s batches.ProductBatchService) *productBatchController {
	return &productBatchController{
		productBatchService: s,
	}
}

func (c *productBatchController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var service = c.productBatchService
		var request CreateProductBatchRequest

		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			ctx.JSON(
				http.StatusUnprocessableEntity,
				web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()),
			)
			return
		}

		addedProductBatch, err := service.Create(
			request.Number,
			request.CurrentQuantity,
			request.CurrentTemperature,
			request.DueDate,
			request.CurrentQuantity,
			request.ManufacturingDate,
			request.ManufacturingHour,
			request.MinimumTemperature,
			request.ProductId,
			request.SectionId,
		)

		if err != nil {
			status := productBatchErrorHandler(err)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, addedProductBatch, ""))
	}
}

func (c *productBatchController) CountProductsBySections() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		var report any
		var err error
		var sectionId uint64

		sectionIdParam := ctx.Query("id")

		if(sectionIdParam != "") {
			sectionId, err = strconv.ParseUint(sectionIdParam, 10, 64)
			report, err = c.productBatchService.CountProductsBySectionId(sectionId)
		}

		if(sectionIdParam == "") {
			report, err = c.productBatchService.CountProductsBySections()
		}

		if err != nil {
			status := productErrorHandler(err)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, report, ""))
	}
}

func productBatchErrorHandler(err error) int {
	switch err {

	case batches.ProductNotFoundError:
		return http.StatusConflict

	case batches.SectionNotFoundError:
		return http.StatusConflict

	case batches.ExistsBatchNumberError:
		return http.StatusConflict

	default:
		return http.StatusInternalServerError
	}
}
