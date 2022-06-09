package controller

import (
	"net/http"
	"strconv"

	"github.com/GuiTadeu/mercado-fresh-panic/internal/products"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
)

type CreateProductRequest struct {
	Code                    string  `json:"product_code" binding:"required"`
	Description             string  `json:"description" binding:"required"`
	Width                   float32 `json:"width" binding:"required"`
	Height                  float32 `json:"height" binding:"required"`
	Length                  float32 `json:"length" binding:"required"`
	NetWeight               float32 `json:"net_weight" binding:"required"`
	ExpirationRate          float32 `json:"expiration_rate" binding:"required"`
	RecommendedFreezingTemp float32 `json:"recommended_freezing_temperature" binding:"required"`
	FreezingRate            float32 `json:"freezing_rate" binding:"required"`
	ProductTypeId           uint64  `json:"product_type_id" binding:"required"`
	SellerId                uint64  `json:"seller_id" binding:"required"`
}

type UpdateProductRequest struct {
	Code                    string  `json:"product_code"`
	Description             string  `json:"description"`
	Width                   float32 `json:"width"`
	Height                  float32 `json:"height"`
	Length                  float32 `json:"length"`
	NetWeight               float32 `json:"net_weight"`
	ExpirationRate          float32 `json:"expiration_rate"`
	RecommendedFreezingTemp float32 `json:"recommended_freezing_temperature"`
	FreezingRate            float32 `json:"freezing_rate"`
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
			status, header := productErrorHandler(err, ctx)
			ctx.JSON(status, header)
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, products, ""))
	}
}

func (c *productController) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		product, err := c.productService.Get(id)

		if err != nil {
			status, header := productErrorHandler(err, ctx)
			ctx.JSON(status, header)
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, product, ""))
	}
}

//TODO Adicionar verificação de ProductTypeId e SellerId (ambos precisam existir)
func (c *productController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var service = c.productService
		var request CreateProductRequest

		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		addedProduct, err := service.Create(
			request.Code,
			request.Description,
			request.Width,
			request.Height,
			request.Length,
			request.NetWeight,
			request.ExpirationRate,
			request.RecommendedFreezingTemp,
			request.FreezingRate,
			request.ProductTypeId,
			request.SellerId,
		)

		if err != nil {
			status, header := productErrorHandler(err, ctx)
			ctx.JSON(status, header)
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, addedProduct, ""))
	}
}

func (c *productController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request UpdateProductRequest

		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(http.StatusBadRequest, nil, "Product id binding error"),
			)
			return
		}

		updatedProduct, err := c.productService.Update(
			id,
			request.Code,
			request.Description,
			request.Width,
			request.Height,
			request.Length,
			request.NetWeight,
			request.ExpirationRate,
			request.RecommendedFreezingTemp,
			request.FreezingRate,
		)

		if err != nil {
			status, header := productErrorHandler(err, ctx)
			ctx.JSON(status, header)
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, updatedProduct, ""))
	}
}

func (c *productController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = c.productService.Delete(id)

		if err != nil {
			status, header := productErrorHandler(err, ctx)
			ctx.JSON(status, header)
			return
		}

		ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, ""))
	}
}

func productErrorHandler(err error, ctx *gin.Context) (int, gin.H) {
	switch err {

	case products.ProductNotFoundError:
		return http.StatusNotFound, gin.H{"error": err.Error()}

	case products.ExistsProductCodeError:
		return http.StatusConflict, gin.H{"error": err.Error()}

	default:
		return http.StatusInternalServerError, gin.H{"error": err.Error()}
	}
}
