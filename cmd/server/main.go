package main

import db "github.com/GuiTadeu/mercado-fresh-panic/cmd/server/database"

func main() {

	// sellers, warehouses, sections, products, employees, buyers
	var _, _, _, _, _, _ = db.CreateDatabases()
}