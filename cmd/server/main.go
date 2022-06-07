package main

import (
	controller "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/controllers"
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	rp "github.com/GuiTadeu/mercado-fresh-panic/internal/sections"
	"github.com/gin-gonic/gin"
)

func main() {

	// sellers, warehouses, sections, products, employees, buyers
	var _, _, sections, _, _, _ = db.CreateDatabases()

	sectionRepository := rp.NewRepository(sections)
	sectionService := rp.NewService(sectionRepository)
	sectionHandler := controller.NewController(sectionService)

	r := gin.Default()

	sectionRoutes := r.Group("/api/v1/sections")

	sectionRoutes.GET("/", sectionHandler.GetAll())
	sectionRoutes.GET("/:id", sectionHandler.Get())
	sectionRoutes.POST("/", sectionHandler.Create())
	sectionRoutes.PATCH("/:id", sectionHandler.Update())
	sectionRoutes.DELETE("/:id", sectionHandler.Delete())

	r.Run()
}
