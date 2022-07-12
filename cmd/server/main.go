package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/controller"
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/buyers"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/carries"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/employees"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/products"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/sections"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/sellers"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/warehouses"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"	
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load .env")
	}

	storageDB := db.Init()
	server := gin.Default()

	// sellers, warehouses, sections, products, employees, buyers
	var buyersDB = db.CreateDatabases()

	sellersHandlers(storageDB, server)
	warehousesHandlers(storageDB, server)
	sectionHandlers(storageDB, server)
	productHandlers(storageDB, server)
	buyerHandlers(buyersDB, server)
	employeeHandlers(storageDB, server)
	carriersHandlers(storageDB, server)

	port := os.Getenv("MERCADO_FRESH_HOST_PORT")
	server.Run(port)
}

func carriersHandlers(storageDB *sql.DB, server *gin.Engine) {
	carrierRepository := carries.NewCarrierRepository(storageDB)
	carrierService := carries.NewCarrierService(carrierRepository)
	carrierController := controller.NewCarrierController(carrierService)

	CarrierGroup := server.Group("/api/v1/carries")
	CarrierGroup.GET("/reportCarries", carrierController.GetAllCarrierInfo())
//	CarrierGroup.GET("/:id", carrierController.FindOne())
	CarrierGroup.POST("/", carrierController.Create())
}

func sellersHandlers(storageDB *sql.DB, server *gin.Engine) {
	sellerRepository := sellers.NewRepository(storageDB)
	sellerService := sellers.NewService(sellerRepository)
	sellerController := controller.NewSeller(sellerService)

	sellerGroup := server.Group("/api/v1/sellers")
	sellerGroup.GET("/", sellerController.FindAll())
	sellerGroup.GET("/:id", sellerController.FindOne())
	sellerGroup.POST("/", sellerController.Create())
	sellerGroup.PATCH("/:id", sellerController.Update())
	sellerGroup.DELETE("/:id", sellerController.Delete())
}

func warehousesHandlers(storageDB *sql.DB, server *gin.Engine) {
	warehouseRepository := warehouses.NewRepository(storageDB)
	warehouseService := warehouses.NewService(warehouseRepository)
	warehouseController := controller.NewWarehouseController(warehouseService)

	warehouseGroup := server.Group("/api/v1/warehouses")

	warehouseGroup.GET("/", warehouseController.GetAll())
	warehouseGroup.GET("/:id", warehouseController.Get())
	warehouseGroup.POST("/", warehouseController.Create())
	warehouseGroup.PATCH("/:id", warehouseController.Update())
	warehouseGroup.DELETE("/:id", warehouseController.Delete())
}

func productHandlers(storageDB *sql.DB, server *gin.Engine) {

	productRepository := products.NewProductRepository(storageDB)
	productService := products.NewProductService(productRepository)
	productHandler := controller.NewProductController(productService)

	productRoutes := server.Group("/api/v1/products")

	productRoutes.GET("/", productHandler.GetAll())
	productRoutes.GET("/:id", productHandler.Get())
	productRoutes.POST("/", productHandler.Create())
	productRoutes.PATCH("/:id", productHandler.Update())
	productRoutes.DELETE("/:id", productHandler.Delete())
}

func sectionHandlers(storageDB *sql.DB, server *gin.Engine) {

	sectionRepository := sections.NewRepository(storageDB)
	sectionService := sections.NewService(sectionRepository)
	sectionHandler := controller.NewSectionController(sectionService)

	sectionRoutes := server.Group("/api/v1/sections")

	sectionRoutes.GET("/", sectionHandler.GetAll())
	sectionRoutes.GET("/:id", sectionHandler.Get())
	sectionRoutes.POST("/", sectionHandler.Create())
	sectionRoutes.PATCH("/:id", sectionHandler.Update())
	sectionRoutes.DELETE("/:id", sectionHandler.Delete())

}

func employeeHandlers(storageDB *sql.DB, server *gin.Engine) {

	employeeRepository := employees.NewRepository(storageDB)
	employeeService := employees.NewEmployeeService(employeeRepository)
	employeeHandler := controller.NewEmployeeController(employeeService)

	employeeRoutes := server.Group("/api/v1/employees")

	employeeRoutes.GET("/", employeeHandler.GetAll())
	employeeRoutes.POST("/", employeeHandler.Create())
	employeeRoutes.DELETE("/:id", employeeHandler.Delete())
	employeeRoutes.GET("/:id", employeeHandler.Get())
	employeeRoutes.PATCH("/:id", employeeHandler.Update())
}

func buyerHandlers(buyersDB []db.Buyer, server *gin.Engine) {

	rBuyers := buyers.NewBuyerRepository(buyersDB)
	sBuyers := buyers.NewBuyerService(rBuyers)
	cBuyers := controller.NewBuyerController(sBuyers)

	buyerRoutes := server.Group("/api/v1/buyers")

	buyerRoutes.GET("/", cBuyers.GetAll())
	buyerRoutes.GET("/:id", cBuyers.Get())
	buyerRoutes.POST("/", cBuyers.Create())
	buyerRoutes.PATCH("/:id", cBuyers.Update())
	buyerRoutes.DELETE("/:id", cBuyers.Delete())
}
