package controller

import (
	"net/http"
	"strconv"

	"github.com/GuiTadeu/mercado-fresh-panic/internal/sellers"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
)

type SellersController struct {
	service sellers.Service
}

func NewSeller(s sellers.Service) *SellersController {
	return &SellersController{
		service: s,
	}
}

func (control *SellersController) FindAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		s, err := control.service.FindAll()

		if err != nil {
			status := sellerErrorHandler(err, ctx)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, s, ""))
	}
}

func (control *SellersController) FindOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")

		id, err := strconv.ParseUint(idParam, 0, 64)

		if err != nil {
			status := sellerErrorHandler(err, ctx)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}

		s, err := control.service.FindOne(id)

		if err != nil {
			status := sellerErrorHandler(err, ctx)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, s, ""))
	}
}

func (control *SellersController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req createSellerRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			return
		}

		s, err := control.service.Create(req.Cid, req.CompanyName, req.Address, req.Telephone, req.LocalityId)

		if err != nil {
			status := sellerErrorHandler(err, ctx)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, s, ""))
	}
}

func (control *SellersController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req updateSellerRequest

		idParam := ctx.Param("id")

		id, err := strconv.ParseUint(idParam, 0, 64)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, "id in wrong format"))
			return
		}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			return
		}

		s, err := control.service.Update(id, req.Cid, req.CompanyName, req.Address, req.Telephone, req.LocalityId)

		if err != nil {
			status := sellerErrorHandler(err, ctx)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, s, ""))
	}
}

func (control *SellersController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")

		id, err := strconv.ParseUint(idParam, 0, 64)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "id in wrong format"))
			return
		}

		err = control.service.Delete(id)

		if err != nil {
			status := sellerErrorHandler(err, ctx)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, ""))
	}
}

func sellerErrorHandler(err error, ctx *gin.Context) int {
	switch err {

	case sellers.SellerNotFoundError:
		return http.StatusNotFound

	case sellers.ExistsSellerCodeError:
		return http.StatusConflict

	default:
		return http.StatusInternalServerError
	}
}

type createSellerRequest struct {
	Cid         uint64 `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
	LocalityId  string `json:"locality_id" binding:"required"`
}

type updateSellerRequest struct {
	Cid         uint64 `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	LocalityId  string `json:"locality_id"`
}
