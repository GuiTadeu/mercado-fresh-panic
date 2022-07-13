package controller

import (
	"github.com/GuiTadeu/mercado-fresh-panic/internal/purchaseOrders"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PurchaseOrdersController struct {
	purchaseOrdesService purchaseOrders.PurchaseOrdersService
}

func NewPurchaseOrderController(s purchaseOrders.PurchaseOrdersService) *PurchaseOrdersController {
	return &PurchaseOrdersController{
		purchaseOrdesService: s,
	}
}

type CreatePurchaseOrderRequest struct {
	OrderNumber     string `json:"order_number" binding:"required"`
	OrderDate       string `json:"order_date" binding:"required"`
	TrackingCode    string `json:"tracking_code" binding:"required"`
	BuyerId         uint64 `json:"buyer_id" binding:"required"`
	OrderStatusId   uint64 `json:"order_status_id" binding:"required"`
	ProductRecordId uint64 `json:"product_record_id" binding:"required"`
}

func (c *PurchaseOrdersController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreatePurchaseOrderRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}
		purchaseOrders, err := c.purchaseOrdesService.Create(
			req.OrderNumber,
			req.OrderDate,
			req.TrackingCode,
			req.BuyerId,
			req.OrderStatusId,
			req.ProductRecordId,
		)

		if err != nil {
			status := purchaseOrderErrorHandler(err)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, purchaseOrders, ""))
	}
}

func purchaseOrderErrorHandler(err error) int {
	switch err {

	case purchaseOrders.ExistsIdError:
		return http.StatusConflict

	case purchaseOrders.PurchaseOrderNotFoundError:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
