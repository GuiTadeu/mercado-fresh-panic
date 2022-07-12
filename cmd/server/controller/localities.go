package controller

import (
	"net/http"

	"github.com/GuiTadeu/mercado-fresh-panic/internal/localities"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
)

type LocalitiesController struct {
	service localities.Service
}

func NewLocality(s localities.Service) *LocalitiesController {
	return &LocalitiesController{
		service: s,
	}
}

func (control *LocalitiesController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req createLocalityRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			return
		}

		s, err := control.service.Create(req.LocalityId, req.LocalityName, req.ProvinceId)

		if err != nil {
			status := localityErrorHandler(err, ctx)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, s, ""))
	}
}

func (control *LocalitiesController) GetLocalityInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		localityId := ctx.Query("id")

		localityData, err := control.service.GetLocalityInfo(localityId)

		if err != nil {
			status := localityErrorHandler(err, ctx)
			ctx.JSON(status, web.NewResponse(status, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, localityData, ""))
	}
}

func localityErrorHandler(err error, ctx *gin.Context) int {
	switch err {

	case localities.LocalityNotFoundError:
		return http.StatusNotFound

	case localities.ExistsLocalityId:
		return http.StatusConflict

	case localities.ExistsProvinceIdError:
		return http.StatusConflict

	default:
		return http.StatusInternalServerError
	}
}

type createLocalityRequest struct {
	LocalityId   string `json:"id" binding:"required"`
	LocalityName string `json:"locality_name" binding:"required"`
	ProvinceId   uint64 `json:"province_id" binding:"required"`
}
