package main

import (
	controller "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/controllers"
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	rpb "github.com/GuiTadeu/mercado-fresh-panic/internal/buyers"
	rp "github.com/GuiTadeu/mercado-fresh-panic/internal/sections"
	"github.com/gin-gonic/gin"
)

func main() {

	// sellers, warehouses, sections, products, employees, buyers
	var _, _, sections, _, _, buyers = db.CreateDatabases()

	rSections := rp.NewRepository(sections)

	sSections := rp.NewService(rSections)

	cSections := controller.NewController(sSections)

	rBuyers := rpb.NewBuyerRepository(buyers)

	sBuyers := rpb.NewBuyerService(rBuyers)

	cBuyers := controller.NewBuyerController(sBuyers)

	r := gin.Default()

	sectionRoutes := r.Group("/api/v1/sections")

	sectionRoutes.GET("/", cSections.GetAll())
	sectionRoutes.GET("/:id", cSections.Get())
	sectionRoutes.POST("/", cSections.Create())
	sectionRoutes.PUT("/:id", cSections.Update())
	sectionRoutes.DELETE("/:id", cSections.Delete())

	buyerRoutes := r.Group("/api/v1/buyers")

	buyerRoutes.GET("/", cBuyers.GetAll())
	buyerRoutes.GET("/:id", cBuyers.Get())
	buyerRoutes.POST("/", cBuyers.Create())
	//buyerRoutes.PUT("/:id", cBuyers.Update())
	buyerRoutes.DELETE("/:id", cBuyers.Delete())

	r.Run()

}
