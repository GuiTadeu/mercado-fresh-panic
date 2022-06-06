package main

import (
	"github.com/GuiTadeu/mercado-fresh-panic/cmd/controllers"
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/sellers"
	"github.com/gin-gonic/gin"
)

func main() {

	// sellers, warehouses, sections, products, employees, buyers
	var dataSellers, _, _, _, _, _ = db.CreateDatabases()

	r := gin.Default()

	sellerRepository := sellers.NewRepository(dataSellers)
	sellerService := sellers.NewService(sellerRepository)
	sellerController := controllers.NewSeller(sellerService)

	sellerGroup := r.Group("/api/v1/seller")
	sellerGroup.GET("/", sellerController.FindAll())
	sellerGroup.GET("/:id", sellerController.FindOne())
	sellerGroup.POST("/", sellerController.Create())
	sellerGroup.PATCH("/:id", sellerController.UpdateAddress())

	r.Run()
}
