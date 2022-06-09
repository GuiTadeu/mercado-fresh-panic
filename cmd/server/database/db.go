package database

import "fmt"

type Seller struct {
	Id          uint64 `json:"id"`
	Cid         uint64 `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
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
	Id                      uint64  `json:"id"`
	Code                    string  `json:"product_code" binding:"required"`
	Description             string  `json:"description" binding:"required"`
	Width                   float32 `json:"width" binding:"required"`
	Height                  float32 `json:"height" binding:"required"`
	Length                  float32 `json:"length" binding:"required"`
	NetWeight               float32 `json:"net_weight" binding:"required"`
	ExpirationRate          float32 `json:"expiration_rate" binding:"required"`
	RecommendedFreezingTemp float32 `json:"recommended_freezing_temperature" binding:"required"`
	FreezingRate            float32 `json:"freezing_rate" binding:"required"`
	ProductTypeId           uint64  `json:"product_type_id" binding:"required"`
	SellerId                uint64  `json:"seller_id" binding:"required"`
}

type Employee struct {
	Id           uint64
	CardNumberId uint64
	FirstName    string
	LastName     string
	WarehouseId  uint64
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
