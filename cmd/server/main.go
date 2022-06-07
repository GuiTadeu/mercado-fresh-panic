package main

import (
	controller "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/controllers"
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	rp "github.com/GuiTadeu/mercado-fresh-panic/internal/sections"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/warehouse"
	"github.com/gin-gonic/gin"
)

func main() {

	// sellers, warehouses, sections, products, employees, buyers
	var _, warehouses, sections, _, _, _ = db.CreateDatabases()

	rwarehouses := warehouse.NewRepository(warehouses)

	swarehouses := warehouse.NewService(rwarehouses)

	cwarehouses := controller.NewWarehouseController(swarehouses)	

	rSections := rp.NewRepository(sections)

	sSections := rp.NewService(rSections)

	cSections := controller.NewController(sSections)


	r := gin.Default()

	sectionRoutes := r.Group("/api/v1/sections")

	sectionRoutes.GET("/", cSections.GetAll())
	sectionRoutes.GET("/:id", cSections.Get())
	sectionRoutes.POST("/", cSections.Create())
	sectionRoutes.PUT("/:id", cSections.Update())
	sectionRoutes.DELETE("/:id", cSections.Delete())

	warehouseRoutes := r.Group("/api/v1/warehouse")

	warehouseRoutes.GET("/", cwarehouses.GetAll())
	warehouseRoutes.GET("/:id",cwarehouses.Get())
	warehouseRoutes.POST("/", cwarehouses.Create())
	//warehouseRoutes.PUT("/", cwarehouses.Update())
	warehouseRoutes.DELETE("/", cwarehouses.Delete())

	r.Run()

}
