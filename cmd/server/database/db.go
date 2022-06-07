package database

import "fmt"

type Seller struct {
	Id          uint64
	Cid         string
	CompanyName string
	Address     string
	Telephone   string
}

type Warehouse struct {
	Id                 uint64
	Code               string
	Address            string
	Telephone          string
	MinimunCapacity    uint32
	MinimumTemperature float32
}

type Section struct {
	Id                 uint64  `json:"id"`
	Number             uint64  `json:"number" binding:"required"`
	CurrentTemperature float32 `json:"current_temperature" binding:"required"`
	MinimumTemperature float32 `json:"minimum_temperature" binding:"required"`
	CurrentCapacity    uint32  `json:"current_capacity" binding:"required"`
	MinimumCapacity    uint32  `json:"minimum_capacity" binding:"required"`
	MaximumCapacity    uint32  `json:"maximum_capacity" binding:"required"`
	WarehouseId        uint64  `json:"warehouse_id" binding:"required"`
	ProductTypeId      uint64  `json:"product_type_id" binding:"required"`
	Products           []Product
}

type Product struct {
	Id                      uint64
	Code                    string
	Description             string
	Width                   float32
	Height                  float32
	Length                  float32
	NetWeight               float32
	ExpirationDate          string
	RecommendedFreezingTemp float32
	FreezingRate            float32
	ProductTypeId           uint64
	SellerId                uint64
}

type Employee struct {
	Id           uint64 `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  uint64 `json:"warehouse_id"`
}

type Buyer struct {
	Id           uint64
	CardNumberId uint64
	FirstName    string
	LastName     string
}

func CreateDatabases() (
	sellers []Seller,
	warehouses []Warehouse,
	sections []Section,
	products []Product,
	employees []Employee,
	buyers []Buyer,
) {

	fmt.Println("Create Databases - Starting...")

	sellers = []Seller{}
	warehouses = []Warehouse{}
	sections = []Section{}
	products = []Product{}
	employees = []Employee{}
	buyers = []Buyer{}

	fmt.Printf("\n sellers:%v", sellers)
	fmt.Printf("\n warehouses:%v", warehouses)
	fmt.Printf("\n sections:%v", sections)
	fmt.Printf("\n products:%v", products)
	fmt.Printf("\n employees:%v", employees)
	fmt.Printf("\n buyers:%v", buyers)

	fmt.Println("\n Create Databases - Done!")
	return
}
