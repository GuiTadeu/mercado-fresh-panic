package main

import (
	controller "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/controllers"
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	rp "github.com/GuiTadeu/mercado-fresh-panic/internal/sections"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/sellers"
	"github.com/gin-gonic/gin"
)

func main() {

	// sellers, warehouses, sections, products, employees, buyers
	var dataSellers, _, sections, _, _, _ = db.CreateDatabases()

	r := gin.Default()

	rSections := rp.NewRepository(sections)

	sSections := rp.NewService(rSections)

	cSections := controller.NewController(sSections)

	sectionRoutes := r.Group("/api/v1/sections")

	sectionRoutes.GET("/", cSections.GetAll())
	sectionRoutes.GET("/:id", cSections.Get())
	sectionRoutes.POST("/", cSections.Create())
	sectionRoutes.PUT("/:id", cSections.Update())
	sectionRoutes.DELETE("/:id", cSections.Delete())

	sellerRepository := sellers.NewRepository(dataSellers)
	sellerService := sellers.NewService(sellerRepository)
	sellerController := controller.NewSeller(sellerService)

	sellerGroup := r.Group("/api/v1/seller")
	sellerGroup.GET("/", sellerController.FindAll())
	sellerGroup.GET("/:id", sellerController.FindOne())
	sellerGroup.POST("/", sellerController.Create())
	sellerGroup.PATCH("/:id", sellerController.UpdateAddress())
	sellerGroup.DELETE("/:id", sellerController.Delete())

	r.Run()
}
