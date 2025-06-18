package main

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Customer struct {
	BaseModel


	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	CompanyName    string    `json:"company_name"`
	JobTitle       string    `json:"job_title"`

	// Status and Classification
    Status       CustomerStatus `json:"status" gorm:"default:'active'"`
    CustomerType CustomerType   `json:"customer_type" gorm:"default:'prospect'"`
    Source       string         `json:"source" gorm:"size:100"` // How they found us
    
    // Financial Information
    CreditLimit     float64 `json:"credit_limit" gorm:"default:0"`
    TotalSpent      float64 `json:"total_spent" gorm:"default:0"`
    LastPurchaseAt  *time.Time `json:"last_purchase_at"`
    
    // Relationships
    Addresses []Address `json:"addresses" gorm:"foreignKey:CustomerID"`
    Orders    []Order   `json:"orders" gorm:"foreignKey:CustomerID"`
    Projects  []Project `json:"projects" gorm:"foreignKey:CustomerID"`
    Notes     []Note `json:"notes" gorm:"foreignKey:CustomerID"`
    
    // Metadata
    Tags        []string `json:"tags" gorm:"type:text[]"`
    CustomFields map[string]interface{} `json:"custom_fields" gorm:"type:jsonb"`
}

type CustomerStatus string
const (
    CustomerStatusActive    CustomerStatus = "active"
    CustomerStatusInactive  CustomerStatus = "inactive"
    CustomerStatusBlocked   CustomerStatus = "blocked"
)

type CustomerType string
const (
    CustomerTypeProspect CustomerType = "prospect"
    CustomerTypeLead     CustomerType = "lead"
    CustomerTypeActive   CustomerType = "active"
    CustomerTypeChurned  CustomerType = "churned"
)

// Address represents customer addresses
type Address struct {
    BaseModel
    CustomerID   uuid.UUID   `json:"customer_id" gorm:"not null"`
    Type         AddressType `json:"type" gorm:"default:'billing'"`
    Street1      string      `json:"street1" gorm:"not null;size:255"`
    Street2      string      `json:"street2" gorm:"size:255"`
    City         string      `json:"city" gorm:"not null;size:100"`
    State        string      `json:"state" gorm:"not null;size:100"`
    PostalCode   string      `json:"postal_code" gorm:"not null;size:20"`
    Country      string      `json:"country" gorm:"not null;size:100"`
    IsDefault    bool        `json:"is_default" gorm:"default:false"`
}

type AddressType string
const (
    AddressTypeBilling  AddressType = "billing"
    AddressTypeShipping AddressType = "shipping"
)

type Contact struct {
	BaseModel

	Customer  []Customer
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	JobTitle  string `json:"job_title"`
}

type Order struct {
	BaseModel

	OrderNumber string
	Customer  Customer
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
    OrderCustomer   Customer    `json:"customer"`
    OrderItems      []OrderItem `json:"order_items"`
    Payments        []Payment   `json:"payments"`
    BillingAddress  *Address    `json:"billing_address"`
    ShippingAddress *Address    `json:"shipping_address"`
    
    // Metadata
    Notes       string                 `json:"notes"`
    Metadata    map[string]interface{} `json:"metadata"`
}


type Product struct {
	BaseModel
	Name  string
	Price float64
}

type OrderItem struct {
	BaseModel
	Order     Order
	Product   Product
	Quantity  int
	Total     float64
}

type Payment struct {
	BaseModel
	Order     Order
	Amount    float64
	Status    string
}

type Project struct {
	BaseModel

	Customer Customer
	ProjectName string
	ProjectDescription string
	ProjectStatus string
	ProjectStartDate time.Time
	ProjectEndDate time.Time
	ProjectBudget float64
	ProjectProgress float64
	ProjectNotes string
	ProjectTasks []ProjectTask
}

type ProjectTask struct {
	BaseModel

	Project Project
	TaskName string
	TaskDescription string
	TaskStatus string
	TaskStartDate time.Time
	TaskEndDate time.Time
    Assignee string
    TaskType string
    TaskPriority string
}

type Note struct {
	BaseModel
	Customer Customer
	Note string
	NoteType string
	NoteDate time.Time
	NoteAuthor string
	NoteStatus string
	NotePriority string
	NoteCategory string
}

