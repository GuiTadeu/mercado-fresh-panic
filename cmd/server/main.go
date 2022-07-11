package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/GuiTadeu/mercado-fresh-panic/cmd/server/controller"
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/buyers"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/employees"
	inboundorders "github.com/GuiTadeu/mercado-fresh-panic/internal/inboundOrders"
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
	var sellersDB, warehousesDB, sectionsDB, employeeDB, buyersDB = db.CreateDatabases()

	productRepository, inboundOrderRepository := buildRepositories(storageDB)

	sellersHandlers(sellersDB, server)
	warehousesHandlers(warehousesDB, server)
	sectionHandlers(sectionsDB, server)
	productHandlers(productRepository, server)
	buyerHandlers(buyersDB, server)
	employeeHandlers(employeeDB, server)
	inboundOrderHandlers(inboundOrderRepository, employeeDB, warehousesDB, server)

	port := os.Getenv("MERCADO_FRESH_HOST_PORT")
	server.Run(port)
}

func sellersHandlers(sellersDB []db.Seller, server *gin.Engine) {
	sellerRepository := sellers.NewRepository(sellersDB)
	sellerService := sellers.NewService(sellerRepository)
	sellerController := controller.NewSeller(sellerService)

	sellerGroup := server.Group("/api/v1/sellers")
	sellerGroup.GET("/", sellerController.FindAll())
	sellerGroup.GET("/:id", sellerController.FindOne())
	sellerGroup.POST("/", sellerController.Create())
	sellerGroup.PATCH("/:id", sellerController.Update())
	sellerGroup.DELETE("/:id", sellerController.Delete())
}

func warehousesHandlers(warehousesDB []db.Warehouse, server *gin.Engine) {
	warehouseRepository := warehouses.NewRepository(warehousesDB)
	warehouseService := warehouses.NewService(warehouseRepository)
	warehouseController := controller.NewWarehouseController(warehouseService)

	warehouseGroup := server.Group("/api/v1/warehouses")
	warehouseGroup.GET("/", warehouseController.GetAll())
	warehouseGroup.GET("/:id", warehouseController.Get())
	warehouseGroup.POST("/", warehouseController.Create())
	warehouseGroup.PATCH("/:id", warehouseController.Update())
	warehouseGroup.DELETE("/:id", warehouseController.Delete())
}

func productHandlers(productRepository products.ProductRepository, server *gin.Engine) {

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

func employeeHandlers(employeeDB []db.Employee, server *gin.Engine) {

	employeeRepository := employees.NewRepository(employeeDB)
	employeeService := employees.NewEmployeeService(employeeRepository)
	employeeHandler := controller.NewEmployeeController(employeeService)

	employeeRoutes := server.Group("/api/v1/employees")

	employeeRoutes.GET("/", employeeHandler.GetAll())
	employeeRoutes.POST("/", employeeHandler.Create())
	employeeRoutes.DELETE("/:id", employeeHandler.Delete())
	employeeRoutes.GET("/:id", employeeHandler.Get())
	employeeRoutes.PATCH("/:id", employeeHandler.Update())
	employeeRoutes.GET("/reportInboundOrders", employeeHandler.ReportInboundOrders())
}

func inboundOrderHandlers(inboundOrderRepository inboundorders.InboundOrderRepository, employeeDB []db.Employee, warehousesDB []db.Warehouse, server *gin.Engine) {

	employeeRepository := employees.NewRepository(employeeDB)
	warehouseRepository := warehouses.NewRepository(warehousesDB)

	inboundOrderService := inboundorders.NewInboundOrderService(employeeRepository, warehouseRepository, inboundOrderRepository)

	cInboundOrders := controller.NewInboundOrderController(inboundOrderService)

	inboundOrderRoutes := server.Group("/api/v1/inboundOrders")

	inboundOrderRoutes.POST("/", cInboundOrders.Create())
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

func buildRepositories(storageDB *sql.DB) (
	products.ProductRepository, 
	inboundorders.InboundOrderRepository) {
		
	productsRepository := products.NewProductRepository(storageDB)
	inboundOrderRepository := inboundorders.NewRepository(storageDB)
	return productsRepository, inboundOrderRepository
}

