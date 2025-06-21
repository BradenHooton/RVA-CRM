package billing

import (
	"time"

	"rva_crm/internal/core"
	"rva_crm/internal/customers"
	"github.com/google/uuid"
)

type Order struct {
	core.BaseModel

	OrderNumber string
	Customer  customers.Customer
	Status    string

	// Financial Information
	SubTotal    float64
	TaxAmount   float64
	Discount    float64
	Total       float64
	
	// Dates
    OrderDate    time.Time  `json:"order_date"`
    ShippedDate  *time.Time `json:"shipped_date"`
    DeliveredDate *time.Time `json:"delivered_date"`

	// Addresses
    BillingAddressID  *uuid.UUID `json:"billing_address_id"`
    ShippingAddressID *uuid.UUID `json:"shipping_address_id"`
    
    // Relationships
    OrderCustomer   customers.Customer    `json:"customer"`
    OrderItems      []OrderItem `json:"order_items"`
    Payments        []Payment   `json:"payments"`
    BillingAddress  *customers.Address    `json:"billing_address"`
    ShippingAddress *customers.Address    `json:"shipping_address"`
    
    // Metadata
    Notes       string                 `json:"notes"`
    Metadata    map[string]interface{} `json:"metadata"`
}


type Product struct {
	core.BaseModel
	Name  string
	Price float64
}

type OrderItem struct {
	core.BaseModel
	Order     Order
	Product   Product
	Quantity  int
	Total     float64
}

type Payment struct {
	core.BaseModel
	Order     Order
	Amount    float64
	Status    string
}