package controller

import (
	"net/http"

	productrecords "github.com/GuiTadeu/mercado-fresh-panic/internal/product_records"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
)

type CreateProductRecordsRequest struct {
	LastUpdateDate string  `json:"last_update_date" binding:"required"`
	PurchasePrice  float32 `json:"purchase_price" binding:"required"`
	SalePrice      float32 `json:"sale_price" binding:"required"`
	ProductId      uint64  `json:"product_id" binding:"required"`
}

type productRecordsController struct {
	productRecordsService productrecords.ProductRecordsService
}

func NewProductRecordsController(s productrecords.ProductRecordsService) *productRecordsController {
	return &productRecordsController{
		productRecordsService: s,
	}
}

// TODO Adicionar verificação de ProductTypeId e SellerId (ambos precisam existir)
func (c *productRecordsController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var service = c.productRecordsService
		var request CreateProductRecordsRequest

		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			ctx.JSON(
				http.StatusUnprocessableEntity,
				web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()),
			)
			return
		}

		addedProductRecords, err := service.Create(
			request.LastUpdateDate,
			request.PurchasePrice,
			request.SalePrice,
			request.ProductId,
		)

		if err != nil {
			status := productRecordsErrorHandler(err)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, addedProductRecords, ""))
	}
}

func productRecordsErrorHandler(err error) int {
	switch err {

	case productrecords.ErrProductNotFoundError:
		return http.StatusConflict

	default:
		return http.StatusInternalServerError
	}
}
