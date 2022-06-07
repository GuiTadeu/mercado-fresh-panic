package controller

import (
	"github.com/GuiTadeu/mercado-fresh-panic/internal/buyers"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type requestBuyers struct {
	CardNumberId string `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
}

type buyerController struct {
	buyerService buyers.BuyerService
}

func NewBuyerController(s buyers.BuyerService) *buyerController {
	return &buyerController{
		buyerService: s,
	}
}

func (c buyerController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req requestBuyers

		err := ctx.ShouldBindJSON(&req)

		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}
		addedBuyer, err := c.buyerService.Create(req.CardNumberId, req.FirstName, req.LastName)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, addedBuyer, ""))
	}

}

func (c *buyerController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		buyers, err := c.buyerService.GetAll()

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, buyers, ""))
	}

}

func (c *buyerController) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		buyer, err := c.buyerService.Get(uint64(id))

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, buyer, ""))
	}
}

func (c *buyerController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		_, err = c.buyerService.Get(uint64(id))

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = c.buyerService.Delete(uint64(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
	}
}

