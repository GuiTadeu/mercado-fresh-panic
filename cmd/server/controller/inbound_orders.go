package controller

import (
	"net/http"

	inboundorders "github.com/GuiTadeu/mercado-fresh-panic/internal/inboundOrders"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
)

type createInboundOrdersRequest struct {
	OrderDate      string `json:"order_date" binding:"required"`
	OrderNumber    string `json:"order_number" binding:"required"`
	EmployeeId     uint64 `json:"employee_id" binding:"required"`
	ProductBatchId uint64 `json:"product_batch_id" binding:"required"`
	WarehouseId    uint64 `json:"warehouse_id" binding:"required"`
}

type inboundOrderController struct {
	inboundOrderService inboundorders.InboundOrderService
}

func NewInboundOrderController(s inboundorders.InboundOrderService) *inboundOrderController {
	return &inboundOrderController{
		inboundOrderService: s,
	}
}

func (c inboundOrderController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createInboundOrdersRequest

		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			return
		}

		inboundOrder, err := c.inboundOrderService.Create(req.OrderDate, req.OrderNumber, req.EmployeeId, req.ProductBatchId, req.WarehouseId)

		if err != nil {
			status := inboundOrderErrorHandler(err)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, inboundOrder, ""))

	}

}

func inboundOrderErrorHandler(err error) int {
	switch err {
	case inboundorders.EmployeeNotFoundError:
		return http.StatusConflict
	case inboundorders.WarehouseNotFoundError:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
