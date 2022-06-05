package controller

import (
	"strconv"

	"github.com/GuiTadeu/mercado-fresh-panic/internal/sections"
	"github.com/GuiTadeu/mercado-fresh-panic/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Id                 uint64  `json:"id"`
	Number             uint64  `json:"number"`
	CurrentTemperature float32 `json:"currentTemperature"`
	MinimumTemperature float32 `json:"minimumTemperature"`
	CurrentCapacity    uint32  `json:"currentCapacity"`
	MinimumCapacity    uint32  `json:"minimumCapacity"`
	MaximumCapacity    uint32  `json:"maximumCapacity"`
	WarehouseId        uint64  `json:"warehouseId"`
	ProductTypeId      uint64  `json:"productTypeId"`
}

type sectionController struct {
	sectionService sections.SectionService
}

func NewController(s sections.SectionService) *sectionController {
	return &sectionController{
		sectionService: s,
	}
}

func (c *sectionController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		sections, err := c.sectionService.GetAll()

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, web.NewResponse(200, sections, ""))
	}
}

func (c *sectionController) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		section, err := c.sectionService.Get(uint64(id))

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, web.NewResponse(200, section, ""))
	}
}

func (c *sectionController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request

		err := ctx.ShouldBindJSON(&req)

		if err != nil {
			ctx.JSON(422, gin.H{
				"error": err.Error(),
			})
			return
		}

		section, err := c.sectionService.Get(req.Id)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		if section.Number == req.Number {
			ctx.JSON(409, gin.H{
				"error": err.Error(),
			})
			return
		}

		addedSection, err := c.sectionService.Create(c.sectionService.GetNextId(), req.Number, req.CurrentTemperature, req.MinimumTemperature, req.CurrentCapacity, req.MinimumCapacity, req.MaximumCapacity, req.WarehouseId, req.ProductTypeId)

		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, web.NewResponse(200, addedSection, ""))
	}

}

func (c *sectionController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request

		err := ctx.ShouldBindJSON(&req)

		if err != nil {
			ctx.JSON(422, gin.H{
				"error": err.Error(),
			})
			return
		}

		_, err = c.sectionService.Get(req.Id)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		updatedSection, err := c.sectionService.Update(req.Id, req.Number, req.CurrentTemperature, req.MinimumTemperature, req.CurrentCapacity, req.MinimumCapacity, req.MaximumCapacity, req.WarehouseId, req.ProductTypeId)

		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, web.NewResponse(200, updatedSection, ""))
	}

}

func (c *sectionController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		_, err = c.sectionService.Get(uint64(id))

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = c.sectionService.Delete(uint64(id))
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(204, web.NewResponse(204, nil, ""))
	}
}
