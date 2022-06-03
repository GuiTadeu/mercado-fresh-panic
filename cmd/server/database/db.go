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
	Id                 uint64
	Number             string
	CurrentTemperature float32
	MinimumTemperature float32
	CurrentCapacity    uint32
	MinimumCapacity    uint32
	MaximumCapacity    uint32
	WarehouseId        int64
	ProductTypeId      uint64
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
