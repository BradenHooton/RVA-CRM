# Models

## List of Models

- Customer
- Order
- Product
- OrderItem
- Payment
- Invoice
- InvoiceItem
- Project
- ProjectTask
- Milestone
- ProjectNote
- Task
- User
- Role
- Permission


## Model Definitions

### Base Model

```go

package models

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

/*
################  Base Model  #################
*/
type BaseModel struct {
    ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
    CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
    DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// BeforeCreate hook to generate UUID
func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
    if base.ID == uuid.Nil {
        base.ID = uuid.New()
    }
    return nil
}


/*
################  Customer Management  #################
*/
// Customer represents a customer in the CRM system
type Customer struct {
    BaseModel
    
    // Basic Information
    FirstName    string `json:"first_name" gorm:"not null;size:100" validate:"required,min=2,max=100"`
    LastName     string `json:"last_name" gorm:"not null;size:100" validate:"required,min=2,max=100"`
    Email        string `json:"email" gorm:"uniqueIndex;not null;size:255" validate:"required,email"`
    Phone        string `json:"phone" gorm:"size:20" validate:"omitempty,e164"`
    CompanyName  string `json:"company_name" gorm:"size:255"`
    JobTitle     string `json:"job_title" gorm:"size:100"`
    
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
    Notes     []CustomerNote `json:"notes" gorm:"foreignKey:CustomerID"`
    
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


/*
################  Product Management  #################
*/
// Product represents a product or service
type Product struct {
    BaseModel
    
    // Basic Information
    Name        string `json:"name" gorm:"not null;size:255" validate:"required"`
    SKU         string `json:"sku" gorm:"uniqueIndex;not null;size:100" validate:"required"`
    Description string `json:"description" gorm:"type:text"`
    Category    string `json:"category" gorm:"size:100"`
    
    // Pricing
    Price       float64 `json:"price" gorm:"not null" validate:"required,gt=0"`
    CostPrice   float64 `json:"cost_price" gorm:"default:0"`
    Currency    string  `json:"currency" gorm:"default:'USD';size:3"`
    
    // Inventory
    StockQuantity    int  `json:"stock_quantity" gorm:"default:0"`
    LowStockAlert    int  `json:"low_stock_alert" gorm:"default:10"`
    TrackInventory   bool `json:"track_inventory" gorm:"default:true"`
    
    // Status
    Status      ProductStatus `json:"status" gorm:"default:'active'"`
    IsDigital   bool         `json:"is_digital" gorm:"default:false"`
    
    // Relationships
    OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:ProductID"`
    
    // Metadata
    Tags        []string `json:"tags" gorm:"type:text[]"`
    Images      []string `json:"images" gorm:"type:text[]"`
    Attributes  map[string]interface{} `json:"attributes" gorm:"type:jsonb"`
}

type ProductStatus string
const (
    ProductStatusActive      ProductStatus = "active"
    ProductStatusInactive    ProductStatus = "inactive"
    ProductStatusDiscontinued ProductStatus = "discontinued"
)


/*
################  Order Management  #################
*/
// Order represents a customer order
type Order struct {
    BaseModel
    
    // Basic Information
    OrderNumber string      `json:"order_number" gorm:"uniqueIndex;not null"`
    CustomerID  uuid.UUID   `json:"customer_id" gorm:"not null"`
    Status      OrderStatus `json:"status" gorm:"default:'pending'"`
    
    // Financial Information
    SubTotal    float64 `json:"subtotal" gorm:"not null"`
    TaxAmount   float64 `json:"tax_amount" gorm:"default:0"`
    ShippingFee float64 `json:"shipping_fee" gorm:"default:0"`
    Discount    float64 `json:"discount" gorm:"default:0"`
    Total       float64 `json:"total" gorm:"not null"`
    Currency    string  `json:"currency" gorm:"default:'USD';size:3"`
    
    // Dates
    OrderDate    time.Time  `json:"order_date" gorm:"not null"`
    ShippedDate  *time.Time `json:"shipped_date"`
    DeliveredDate *time.Time `json:"delivered_date"`
    
    // Addresses
    BillingAddressID  *uuid.UUID `json:"billing_address_id"`
    ShippingAddressID *uuid.UUID `json:"shipping_address_id"`
    
    // Relationships
    Customer        Customer    `json:"customer" gorm:"foreignKey:CustomerID"`
    OrderItems      []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
    Payments        []Payment   `json:"payments" gorm:"foreignKey:OrderID"`
    BillingAddress  *Address    `json:"billing_address" gorm:"foreignKey:BillingAddressID"`
    ShippingAddress *Address    `json:"shipping_address" gorm:"foreignKey:ShippingAddressID"`
    
    // Metadata
    Notes       string                 `json:"notes" gorm:"type:text"`
    Metadata    map[string]interface{} `json:"metadata" gorm:"type:jsonb"`
}

type OrderStatus string
const (
    OrderStatusPending    OrderStatus = "pending"
    OrderStatusConfirmed  OrderStatus = "confirmed"
    OrderStatusProcessing OrderStatus = "processing"
    OrderStatusShipped    OrderStatus = "shipped"
    OrderStatusDelivered  OrderStatus = "delivered"
    OrderStatusCancelled  OrderStatus = "cancelled"
    OrderStatusRefunded   OrderStatus = "refunded"
)

// OrderItem represents individual items in an order
type OrderItem struct {
    BaseModel
    OrderID   uuid.UUID `json:"order_id" gorm:"not null"`
    ProductID uuid.UUID `json:"product_id" gorm:"not null"`
    Quantity  int       `json:"quantity" gorm:"not null" validate:"required,gt=0"`
    UnitPrice float64   `json:"unit_price" gorm:"not null"`
    Total     float64   `json:"total" gorm:"not null"`
    
    // Relationships
    Order   Order   `json:"order" gorm:"foreignKey:OrderID"`
    Product Product `json:"product" gorm:"foreignKey:ProductID"`
}



/*
################  Payment Management  #################
*/
// Payment represents a payment transaction
type Payment struct {
    BaseModel
    
    // Basic Information
    OrderID         uuid.UUID     `json:"order_id" gorm:"not null"`
    PaymentMethod   PaymentMethod `json:"payment_method" gorm:"not null"`
    Status          PaymentStatus `json:"status" gorm:"default:'pending'"`
    
    // Financial Information
    Amount          float64 `json:"amount" gorm:"not null"`
    Currency        string  `json:"currency" gorm:"default:'USD';size:3"`
    
    // Transaction Details
    TransactionID   string     `json:"transaction_id" gorm:"size:255"`
    ProcessorRef    string     `json:"processor_ref" gorm:"size:255"`
    ProcessedAt     *time.Time `json:"processed_at"`
    
    // Relationships
    Order Order `json:"order" gorm:"foreignKey:OrderID"`
    
    // Metadata
    ProcessorData map[string]interface{} `json:"processor_data" gorm:"type:jsonb"`
    Notes         string                 `json:"notes" gorm:"type:text"`
}

type PaymentMethod string
const (
    PaymentMethodCreditCard PaymentMethod = "credit_card"
    PaymentMethodDebitCard  PaymentMethod = "debit_card"
    PaymentMethodPayPal     PaymentMethod = "paypal"
    PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
    PaymentMethodCash       PaymentMethod = "cash"
    PaymentMethodCheck      PaymentMethod = "check"
)

type PaymentStatus string
const (
    PaymentStatusPending   PaymentStatus = "pending"
    PaymentStatusCompleted PaymentStatus = "completed"
    PaymentStatusFailed    PaymentStatus = "failed"
    PaymentStatusRefunded  PaymentStatus = "refunded"
    PaymentStatusCancelled PaymentStatus = "cancelled"
)


/*
################  Project Management  #################
*/
// Project represents a customer project
type Project struct {
    BaseModel
    
    // Basic Information
    Name        string        `json:"name" gorm:"not null;size:255" validate:"required"`
    Description string        `json:"description" gorm:"type:text"`
    CustomerID  uuid.UUID     `json:"customer_id" gorm:"not null"`
    Status      ProjectStatus `json:"status" gorm:"default:'planning'"`
    Priority    Priority      `json:"priority" gorm:"default:'medium'"`
    
    // Timeline
    StartDate    *time.Time `json:"start_date"`
    EndDate      *time.Time `json:"end_date"`
    DeadlineDate *time.Time `json:"deadline_date"`
    
    // Financial
    Budget       float64 `json:"budget" gorm:"default:0"`
    ActualCost   float64 `json:"actual_cost" gorm:"default:0"`
    Currency     string  `json:"currency" gorm:"default:'USD';size:3"`
    
    // Assignment
    AssignedToID *uuid.UUID `json:"assigned_to_id"`
    
    // Relationships
    Customer     Customer      `json:"customer" gorm:"foreignKey:CustomerID"`
    AssignedTo   *User         `json:"assigned_to" gorm:"foreignKey:AssignedToID"`
    Tasks        []ProjectTask `json:"tasks" gorm:"foreignKey:ProjectID"`
    Milestones   []Milestone   `json:"milestones" gorm:"foreignKey:ProjectID"`
    Notes        []ProjectNote `json:"notes" gorm:"foreignKey:ProjectID"`
    
    // Metadata
    Tags     []string               `json:"tags" gorm:"type:text[]"`
    Metadata map[string]interface{} `json:"metadata" gorm:"type:jsonb"`
}

type ProjectStatus string
const (
    ProjectStatusPlanning   ProjectStatus = "planning"
    ProjectStatusActive     ProjectStatus = "active"
    ProjectStatusOnHold     ProjectStatus = "on_hold"
    ProjectStatusCompleted  ProjectStatus = "completed"
    ProjectStatusCancelled  ProjectStatus = "cancelled"
)

type Priority string
const (
    PriorityLow    Priority = "low"
    PriorityMedium Priority = "medium"
    PriorityHigh   Priority = "high"
    PriorityUrgent Priority = "urgent"
)

// ProjectTask represents a task within a project
type ProjectTask struct {
    BaseModel
    
    ProjectID    uuid.UUID  `json:"project_id" gorm:"not null"`
    Title        string     `json:"title" gorm:"not null;size:255"`
    Description  string     `json:"description" gorm:"type:text"`
    Status       TaskStatus `json:"status" gorm:"default:'todo'"`
    Priority     Priority   `json:"priority" gorm:"default:'medium'"`
    
    // Timeline
    DueDate      *time.Time `json:"due_date"`
    CompletedAt  *time.Time `json:"completed_at"`
    
    // Assignment
    AssignedToID *uuid.UUID `json:"assigned_to_id"`
    
    // Relationships
    Project    Project `json:"project" gorm:"foreignKey:ProjectID"`
    AssignedTo *User   `json:"assigned_to" gorm:"foreignKey:AssignedToID"`
}

type TaskStatus string
const (
    TaskStatusTodo       TaskStatus = "todo"
    TaskStatusInProgress TaskStatus = "in_progress"
    TaskStatusReview     TaskStatus = "review"
    TaskStatusCompleted  TaskStatus = "completed"
    TaskStatusCancelled  TaskStatus = "cancelled"
)

// Milestone represents a project milestone
type Milestone struct {
    BaseModel
    
    ProjectID   uuid.UUID `json:"project_id" gorm:"not null"`
    Title       string    `json:"title" gorm:"not null;size:255"`
    Description string    `json:"description" gorm:"type:text"`
    DueDate     time.Time `json:"due_date" gorm:"not null"`
    IsCompleted bool      `json:"is_completed" gorm:"default:false"`
    CompletedAt *time.Time `json:"completed_at"`
    
    // Relationships
    Project Project `json:"project" gorm:"foreignKey:ProjectID"`
}



/*
################  User Management  #################
*/
// User represents a system user
type User struct {
    BaseModel
    
    // Basic Information
    FirstName string `json:"first_name" gorm:"not null;size:100" validate:"required"`
    LastName  string `json:"last_name" gorm:"not null;size:100" validate:"required"`
    Email     string `json:"email" gorm:"uniqueIndex;not null;size:255" validate:"required,email"`
    Username  string `json:"username" gorm:"uniqueIndex;not null;size:50" validate:"required"`
    
    // Authentication
    PasswordHash string `json:"-" gorm:"not null"`
    IsActive     bool   `json:"is_active" gorm:"default:true"`
    LastLoginAt  *time.Time `json:"last_login_at"`
    
    // Profile
    Avatar   string `json:"avatar" gorm:"size:500"`
    Phone    string `json:"phone" gorm:"size:20"`
    Timezone string `json:"timezone" gorm:"default:'UTC';size:50"`
    
    // Relationships
    Roles            []Role         `json:"roles" gorm:"many2many:user_roles;"`
    AssignedProjects []Project      `json:"assigned_projects" gorm:"foreignKey:AssignedToID"`
    AssignedTasks    []ProjectTask  `json:"assigned_tasks" gorm:"foreignKey:AssignedToID"`
    CreatedNotes     []CustomerNote `json:"created_notes" gorm:"foreignKey:CreatedByID"`
}

// Role represents a user role
type Role struct {
    BaseModel
    
    Name        string `json:"name" gorm:"uniqueIndex;not null;size:100"`
    Description string `json:"description" gorm:"type:text"`
    IsActive    bool   `json:"is_active" gorm:"default:true"`
    
    // Relationships
    Users       []User       `json:"users" gorm:"many2many:user_roles;"`
    Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions;"`
}

// Permission represents a system permission
type Permission struct {
    BaseModel
    
    Name        string `json:"name" gorm:"uniqueIndex;not null;size:100"`
    Description string `json:"description" gorm:"type:text"`
    Resource    string `json:"resource" gorm:"not null;size:100"` // e.g., "customers", "orders"
    Action      string `json:"action" gorm:"not null;size:50"`    // e.g., "create", "read", "update", "delete"
    
    // Relationships
    Roles []Role `json:"roles" gorm:"many2many:role_permissions;"`
}


/*
################  Notes and Communication  #################
*/
// CustomerNote represents notes about customers
type CustomerNote struct {
    BaseModel
    
    CustomerID  uuid.UUID `json:"customer_id" gorm:"not null"`
    CreatedByID uuid.UUID `json:"created_by_id" gorm:"not null"`
    Title       string    `json:"title" gorm:"size:255"`
    Content     string    `json:"content" gorm:"type:text;not null"`
    IsPrivate   bool      `json:"is_private" gorm:"default:false"`
    
    // Relationships
    Customer  Customer `json:"customer" gorm:"foreignKey:CustomerID"`
    CreatedBy User     `json:"created_by" gorm:"foreignKey:CreatedByID"`
}

// ProjectNote represents notes about projects
type ProjectNote struct {
    BaseModel
    
    ProjectID   uuid.UUID `json:"project_id" gorm:"not null"`
    CreatedByID uuid.UUID `json:"created_by_id" gorm:"not null"`
    Title       string    `json:"title" gorm:"size:255"`
    Content     string    `json:"content" gorm:"type:text;not null"`
    IsPrivate   bool      `json:"is_private" gorm:"default:false"`
    
    // Relationships
    Project   Project `json:"project" gorm:"foreignKey:ProjectID"`
    CreatedBy User    `json:"created_by" gorm:"foreignKey:CreatedByID"`
}



/*
################ Model Validation with Hooks #################
*/

// BeforeCreate hooks for generating order numbers, etc.
func (o *Order) BeforeCreate(tx *gorm.DB) error {
    if o.OrderNumber == "" {
        o.OrderNumber = generateOrderNumber()
    }
    return nil
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) error {
    oi.Total = oi.UnitPrice * float64(oi.Quantity)
    return nil
}

// Helper functions
func generateOrderNumber() string {
    return fmt.Sprintf("ORD-%d", time.Now().Unix())
}

```