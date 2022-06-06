package controllers

import (
	"net/http"
	"strconv"

	"github.com/GuiTadeu/mercado-fresh-panic/internal/sellers"
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
		s, statusCode, err := control.service.FindAll()

		if err != nil {
			ctx.JSON(statusCode, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(statusCode, gin.H{
			"data": s,
		})
	}
}

func (control *SellersController) FindOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")

		id, err := strconv.ParseUint(idParam, 0, 64)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "ID in wrong format",
			})
			return
		}

		s, statusCode, err := control.service.FindOne(id)

		if err != nil {
			ctx.JSON(statusCode, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(statusCode, gin.H{
			"data": s,
		})
	}
}

func (control *SellersController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req createRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": err.Error(),
			})
			return
		}

		s, statusCode, err := control.service.Create(req.Cid, req.CompanyName, req.Address, req.Telephone)

		if err != nil {
			ctx.JSON(statusCode, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(statusCode, gin.H{
			"data": s,
		})
	}
}

func (control *SellersController) UpdateAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req updateAddressRequest

		idParam := ctx.Param("id")

		id, err := strconv.ParseUint(idParam, 0, 64)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "ID in wrong format",
			})
			return
		}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		s, statusCode, err := control.service.UpdateAddress(id, req.Address)

		if err != nil {
			ctx.JSON(statusCode, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(statusCode, gin.H{
			"data": s,
		})
	}
}

type createRequest struct {
	Cid         uint64 `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

type updateAddressRequest struct {
	Address string `json:"address"`
}
