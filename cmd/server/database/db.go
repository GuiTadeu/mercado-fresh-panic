package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	StorageDB *sql.DB
)

func Init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("failed to load .env")
	}

	datasource := os.Getenv("CONNECT_MYSQL")

	StorageDB, err = sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database configured")
}

type Seller struct {
	Id          uint64 `json:"id"`
	Cid         uint64 `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
	LocalityID  string `json:"locality_id" binding:"required"`
}

type Warehouse struct {
	Id                 uint64  `json:"id"`
	Code               string  `json:"warehouse_code" binding:"required"`
	Address            string  `json:"address" binding:"required"`
	Telephone          string  `json:"telephone" binding:"required"`
	MinimunCapacity    uint32  `json:"minimum_capacity" binding:"required"`
	MinimumTemperature float32 `json:"minimum_temperature" binding:"required"`
	LocalityID         string  `json:"locality_id" binding:"required"`
}

type Section struct {
	Id                 uint64  `json:"id"`
	Number             uint64  `json:"section_number" binding:"required"`
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
	Id           uint64 `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseId  uint64 `json:"warehouse_id"`
}

type Buyer struct {
	Id           uint64 `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type Country struct {
	Id   uint64 `json:"id"`
	Name string `json:"country_name"`
}

type Province struct {
	Id        uint64 `json:"id"`
	Name      string `json:"province_name"`
	CountryID uint64 `json:"id_country_fk"`
}

type Locality struct {
	Id         string `json:"id"`
	Name       string `json:"locality_name"`
	ProvinceId uint64 `json:"province_id"`
}

type Carrier struct {
	Id          uint64 `json:"id"`
	Cid         uint64 `json:"cid" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Telephone   string `json:"telephone" binding:"required"`
	LocalityID  string `json:"locality_id" binding:"required"`
}

type ProductType struct {
	Id          uint64 `json:"id"`
	Description string `json:"description"`
}

type ProductBatch struct {
	Id                 uint64    `json:"id"`
	Number             uint64    `json:"batch_number"`
	CurrentQuantity    uint64    `json:"current_quantity"`
	CurrentTemperature float32   `json:"current_temperature"`
	DueDate            time.Time `json:"due_date"`
	InitialQuantity    uint64    `json:"initial_quantity"`
	ManufacturingDate  time.Time `json:"manufacturing_date"`
	ManufacturingHour  time.Time `json:"manufacturing_hour"`
	MinimumTemperature float32   `json:"minimum_temperature"`
	ProductId          uint64    `json:"product_id"`
	SectionId          uint64    `json:"section_id"`
}

type ProductRecord struct {
	Id             uint64    `json:"id"`
	LastUpdateDate time.Time `json:"last_update_date"`
	PurchasePrice  float32   `json:"purchase_price"`
	SalePrice      float32   `json:"sale_price"`
	ProductId      uint64    `json:"product_id"`
}

type InboundOrder struct {
	Id             uint64    `json:"id"`
	OrderDate      time.Time `json:"order_date"`
	OrderNumber    string    `json:"order_number"`
	EmployeeId     uint64    `json:"employee_id"`
	ProductBatchId uint64    `json:"product_batch_id"`
	WarehouseId    uint64    `json:"warehouse_id"`
}

type OrderStatus struct {
	Id          uint64 `json:"id"`
	Description string `json:"description"`
}

type PurchaseOrder struct {
	Id              uint64    `json:"id"`
	OrderNumber     string    `json:"order_number"`
	OrderDate       time.Time `json:"order_date"`
	TrackingCode    string    `json:"tracking_code"`
	BuyerId         uint64    `json:"buyer_id"`
	OrderStatusId   uint64    `json:"order_status_id"`
	ProductRecordId uint64    `json:"product_record_id"`
}

type OrderDetails struct {
	Id                uint64  `json:"id"`
	CleanLinessStatus string  `json:"clean_liness_status"`
	Quantity          uint64  `json:"quantity"`
	Temperature       float32 `json:"temperature"`
	ProductRecordId   uint64  `json:"product_record_id"`
	PurchaseOrderId   uint64  `json:"purchase_order_id"`
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
