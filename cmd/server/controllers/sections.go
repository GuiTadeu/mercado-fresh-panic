package controller

import (
	"net/http"
	"strconv"

	"github.com/GuiTadeu/mercado-fresh-panic/internal/sections"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
)

type CreateSectionRequest struct {
	Number             uint64  `json:"number" binding:"required"`
	CurrentTemperature float32 `json:"current_temperature" binding:"required"`
	MinimumTemperature float32 `json:"minimum_temperature" binding:"required"`
	CurrentCapacity    uint32  `json:"current_capacity" binding:"required"`
	MinimumCapacity    uint32  `json:"minimum_capacity" binding:"required"`
	MaximumCapacity    uint32  `json:"maximum_capacity" binding:"required"`
	WarehouseId        uint64  `json:"warehouse_id" binding:"required"`
	ProductTypeId      uint64  `json:"product_type_id" binding:"required"`
}

type UpdateSectionRequest struct {
	Number             uint64  `json:"number"`
	CurrentTemperature float32 `json:"current_temperature"`
	MinimumTemperature float32 `json:"minimum_temperature"`
	CurrentCapacity    uint32  `json:"current_capacity"`
	MinimumCapacity    uint32  `json:"minimum_capacity"`
	MaximumCapacity    uint32  `json:"maximum_capacity"`
}

type sectionController struct {
	sectionService sections.SectionService
}

func NewSectionController(s sections.SectionService) *sectionController {
	return &sectionController{
		sectionService: s,
	}
}

func (c *sectionController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		sections, err := c.sectionService.GetAll()

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, sections, ""))
	}
}

func (c *sectionController) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		section, err := c.sectionService.Get(uint64(id))

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, section, ""))
	}
}

func (c *sectionController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req CreateSectionRequest
		err := ctx.ShouldBindJSON(&req)

		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		addedSection, err := c.sectionService.Create(req.Number, req.CurrentTemperature, req.MinimumTemperature, req.CurrentCapacity, req.MinimumCapacity, req.MaximumCapacity, req.WarehouseId, req.ProductTypeId)

		if err != nil {
			ctx.JSON(err.(*web.CustomError).Status, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, addedSection, ""))
	}
}

func (c *sectionController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request UpdateSectionRequest
		err := ctx.ShouldBindJSON(&request)

		id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {

			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "section id binding error"))
			return
		}

		updatedSection, err := c.sectionService.Update(
			id,
			request.Number,
			request.CurrentTemperature,
			request.MinimumTemperature,
			request.CurrentCapacity,
			request.MinimumCapacity,
			request.MaximumCapacity,
		)

		if err != nil {
			ctx.JSON(err.(*web.CustomError).Status, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, updatedSection, ""))
	}
}

func (c *sectionController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = c.sectionService.Delete(uint64(id))
		if err != nil {
			ctx.JSON(err.(*web.CustomError).Status, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, ""))
	}
}
