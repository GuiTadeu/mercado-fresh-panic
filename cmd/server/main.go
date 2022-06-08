package main

import (
	controller "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/controllers"
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	products "github.com/GuiTadeu/mercado-fresh-panic/internal/products"
	sections "github.com/GuiTadeu/mercado-fresh-panic/internal/sections"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	// sellers, warehouses, sections, products, employees, buyers
	var _, _, sectionsDB, productsDB, _, _ = db.CreateDatabases()

	sectionHandlers(sectionsDB, server)
	productHandlers(productsDB, server)

	server.Run()
}

func productHandlers(productsDB []db.Product, server *gin.Engine) {

	productRepository := products.NewProductRepository(productsDB)
	productService := products.NewProductService(productRepository)
	productHandler := controller.NewProductController(productService)

	productRoutes := server.Group("/api/v1/products")

	productRoutes.GET("/", productHandler.GetAll())
	productRoutes.GET("/:id", productHandler.Get())
	productRoutes.POST("/", productHandler.Create())
	productRoutes.PATCH("/:id", productHandler.Update())
	productRoutes.DELETE("/:id", productHandler.Delete())
}

func sectionHandlers(sectionsDB []db.Section, server *gin.Engine) {

	sectionRepository := sections.NewRepository(sectionsDB)
	sectionService := sections.NewService(sectionRepository)
	sectionHandler := controller.NewSectionController(sectionService)

	sectionRoutes := server.Group("/api/v1/sections")

	sectionRoutes.GET("/", sectionHandler.GetAll())
	sectionRoutes.GET("/:id", sectionHandler.Get())
	sectionRoutes.POST("/", sectionHandler.Create())
	sectionRoutes.PATCH("/:id", sectionHandler.Update())
	sectionRoutes.DELETE("/:id", sectionHandler.Delete())
}
