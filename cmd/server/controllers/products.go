package controller

import (
	"strconv"

	"github.com/GuiTadeu/mercado-fresh-panic/internal/products"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
)

type CreateProductRequest struct {
	Code string `json:"product_code" binding:"required"`
	Description string `json:"description" binding:"required"`
	Width float32 `json:"width" binding:"required"`
	Height float32 `json:"height" binding:"required"`
	Length float32 `json:"length" binding:"required"`
	NetWeight float32 `json:"netWeight" binding:"required"`
	ExpirationDate string `json:"expiration_date" binding:"required"`
	RecommendedFreezingTemp float32 `json:"recommended_freezing_temperature" binding:"required"`
	FreezingRate float32 `json:"freezing_rate" binding:"required"`
	ProductTypeId uint64 `json:"product_type_id" binding:"required"`
	SellerId uint64 `json:"seller_id" binding:"required"`
}

type UpdateProductRequest struct {
	Code string `json:"code"`
	Description string `json:"description"`
	Width float32 `json:"width"`
	Height float32 `json:"height"`
	Length float32 `json:"length"`
	NetWeight float32 `json:"netWeight"`
	ExpirationDate string `json:"expiration_date"`
	RecommendedFreezingTemp float32 `json:"recommended_freezing_temp"`
	FreezingRate float32 `json:"freezing_rate"`
}

type productController struct {
	productService products.ProductService
}

func NewProductController(s products.ProductService) *productController {
	return &productController{
		productService: s,
	}
}

func (c *productController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		products, err := c.productService.GetAll()

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, web.NewResponse(200, products, ""))
	}
}

func (c *productController) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		product, err := c.productService.Get(uint64(id))

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, web.NewResponse(200, product, ""))
	}
}

func (c *productController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request CreateProductRequest
		err := ctx.ShouldBindJSON(&request)

		if err != nil {
			ctx.JSON(422, gin.H{
				"error": err.Error(),
			})
			return
		}

		products, err := c.productService.GetAll()
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		for _, product := range products {
			if product.Code == request.Code {
				ctx.JSON(409, web.NewResponse(409, nil, "Product code already existis"))
				return
			}
		}

		addedProduct, err := c.productService.Create(
			c.productService.GetNextId(),
			request.Code,
			request.Description,
			request.Width,
			request.Height,
			request.Length,
			request.NetWeight,
			request.ExpirationDate,
			request.RecommendedFreezingTemp,
			request.FreezingRate,
			request.ProductTypeId,
			request.SellerId,
		)

		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, web.NewResponse(200, addedProduct, ""))
	}
}

func (c *productController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request UpdateProductRequest
		err := ctx.ShouldBindJSON(&request)

		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "product id binding error"))
			return
		}

		updatedproduct, err := c.productService.Update(
			id,
			request.Code,
			request.Description,
			request.Width,
			request.Height,
			request.Length,
			request.NetWeight,
			request.ExpirationDate,
			request.RecommendedFreezingTemp,
			request.FreezingRate,
		)

		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, updatedproduct, ""))
	}
}

func (c *productController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		_, err = c.productService.Get(uint64(id))

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = c.productService.Delete(uint64(id))
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(204, web.NewResponse(204, nil, ""))
	}
}
