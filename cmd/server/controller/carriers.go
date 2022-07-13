package controller

import (
	"net/http"	

	"github.com/GuiTadeu/mercado-fresh-panic/internal/carries"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
)

type createCarrierRequest struct {
	Cid         string `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
	LocalityID  string `json:"locality_id" binding:"required"`
}

type carrierController struct {
	carrierService carries.CarrierService
}

func NewCarrierController(carrier carries.CarrierService) *carrierController {
	return &carrierController{
		carrierService: carrier,
	}
}

func (c *carrierController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var required createCarrierRequest

		err := ctx.ShouldBindJSON(&required)

		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			return
		}

		addedCarrier, err := c.carrierService.Create(required.Cid, required.CompanyName, required.Address, required.Telephone, required.LocalityID)

		if err != nil {
			status := carrierErrorHandler(err)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, addedCarrier, " "))
	}
}

func (c *carrierController) GetAllCarrierInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		localityId := ctx.Query("id")
		
		carrier, err := c.carrierService.GetAllCarrierInfo(localityId)
		if err != nil {
			status := carrierErrorHandler(err)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, carrier, ""))
	}
}

func carrierErrorHandler(err error) int {
	switch err {
	case carries.CarrierNotFoundError:
		return http.StatusNotFound
	case carries.ExistsCarrierCidError:
		return http.StatusConflict
	case carries.LocalityIdNotExistsError:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}