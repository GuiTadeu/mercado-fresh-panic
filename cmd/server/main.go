package main

import (
	controller "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/controllers"
	db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"
	"github.com/GuiTadeu/mercado-fresh-panic/internal/employee"
	rp "github.com/GuiTadeu/mercado-fresh-panic/internal/sections"
	"github.com/gin-gonic/gin"
)

func main() {

	// sellers, warehouses, sections, products, employees, buyers
	var _, _, sections, _, employees, _ = db.CreateDatabases()

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

	rEmployee := employee.NewRepository(employees)
	sEmployee := employee.NewService(rEmployee)
	pEmployee := controller.NewEmployee(sEmployee)

	pr := r.Group("/api/v1/employees")
	pr.GET("/", pEmployee.GetAll())
	pr.POST("/", pEmployee.Create())
	pr.DELETE("/:id", pEmployee.Delete())

	r.Run()

}
