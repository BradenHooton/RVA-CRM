# Service Layer

## Overview
The service layer acts as an intermediary between the application's business logic and the repository layer. It encapsulates the business logic of the application, making it easier to manage complex operations and interactions between different parts of the system.

## Service Interface
```go
type Service interface {
    GetCustomer(id uuid.UUID) (*Customer, error)
    CreateCustomer(customer *Customer) error
    UpdateCustomer(customer *Customer) error
    DeleteCustomer(id uuid.UUID) error

    GetOrder(id uuid.UUID) (*Order, error)
    CreateOrder(order *Order) error
    UpdateOrder(order *Order) error
    DeleteOrder(id uuid.UUID) error

    // Add more methods for other entities
}
```

## Service Struct
```go
type service struct {
    repository Repository
}
```

## Service Methods
```go
func (s *service) GetCustomer(id uuid.UUID) (*Customer, error) {
    // Example:
    customer, err := s.repository.GetCustomer(id)
    if err != nil {
        return nil, err
    }
    return customer, nil
}

func (s *service) CreateCustomer(customer *Customer) error {
    // Example:
    err := s.repository.CreateCustomer(customer)
    if err != nil {
        return err
    }
    return nil
}

func (s *service) UpdateCustomer(customer *Customer) error {
    // Example:
    err := s.repository.UpdateCustomer(customer)
    if err != nil {
        return err
    }
    return nil
}

func (s *service) DeleteCustomer(id uuid.UUID) error {
    // Example:
    err := s.repository.DeleteCustomer(id)
    if err != nil {
        return err
    }
    return nil
}

func (s *service) GetOrder(id uuid.UUID) (*Order, error) {
    // Example:
    order, err := s.repository.GetOrder(id)
    if err != nil {
        return nil, err
    }
    return order, nil
}

func (s *service) CreateOrder(order *Order) error {
    // Example:
    err := s.repository.CreateOrder(order)
    if err != nil {
        return err
    }
    return nil
}

func (s *service) UpdateOrder(order *Order) error {
    // Example:
    err := s.repository.UpdateOrder(order)
    if err != nil {
        return err
    }
    return nil
}

func (s *service) DeleteOrder(id uuid.UUID) error {
    // Example:
    err := s.repository.DeleteOrder(id)
    if err != nil {
        return err
    }
    return nil
}

// Advanced Example:
func (s *service) ProcessOrderWithPayment(customerID uuid.UUID, orderItems []OrderItem, paymentAmount float64) (*Order, *Payment, error) {
    // Validate customer exists
    customer, err := s.repository.GetCustomer(customerID)
    if err != nil {
        return nil, nil, fmt.Errorf("customer not found: %w", err)
    }

    // Calculate total from order items
    var total float64
    for _, item := range orderItems {
        total += item.Total
    }

    // Validate payment amount matches order total
    if paymentAmount != total {
        return nil, nil, fmt.Errorf("payment amount %.2f does not match order total %.2f", paymentAmount, total)
    }

    // Create order
    order := &Order{
        BaseModel: BaseModel{
            ID:        uuid.New(),
            CreatedAt: time.Now(),
            UpdatedAt: time.Now(),
        },
        Customer: *customer,
        Total:    total,
    }

    err = s.repository.CreateOrder(order)
    if err != nil {
        return nil, nil, fmt.Errorf("failed to create order: %w", err)
    }

    // Create order items
    for _, item := range orderItems {
        item.Order = *order
        item.ID = uuid.New()
        item.CreatedAt = time.Now()
        item.UpdatedAt = time.Now()
        
        err = s.repository.CreateOrderItem(&item)
        if err != nil {
            // Rollback order creation if order items fail
            s.repository.DeleteOrder(order.ID)
            return nil, nil, fmt.Errorf("failed to create order item: %w", err)
        }
    }

    // Process payment
    payment := &Payment{
        BaseModel: BaseModel{
            ID:        uuid.New(),
            CreatedAt: time.Now(),
            UpdatedAt: time.Now(),
        },
        Order:  *order,
        Amount: paymentAmount,
        Status: "pending",
    }

    err = s.repository.CreatePayment(payment)
    if err != nil {
        // Rollback order and items if payment fails
        s.repository.DeleteOrder(order.ID)
        return nil, nil, fmt.Errorf("failed to create payment: %w", err)
    }

    // Update payment status to completed
    payment.Status = "completed"
    payment.UpdatedAt = time.Now()
    
    err = s.repository.UpdatePayment(payment)
    if err != nil {
        return nil, nil, fmt.Errorf("failed to update payment status: %w", err)
    }

    return order, payment, nil
}

func (s *service) GetCustomerOrderHistory(customerID uuid.UUID, limit int, offset int) ([]Order, error) {
    // Validate customer exists
    // __ is a placeholder for the actual implementation. Hint: Repository layer. 
    _, err := s.repository.GetCustomer(customerID)
    if err != nil {
        return nil, fmt.Errorf("customer not found: %w", err)
    }

    // Get paginated order history
    orders, err := s.repository.GetOrdersByCustomerID(customerID, limit, offset)
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve order history: %w", err)
    }

    return orders, nil
}

func (s *service) CalculateCustomerLifetimeValue(customerID uuid.UUID) (float64, error) {
    // Validate customer exists
    _, err := s.repository.GetCustomer(customerID)
    if err != nil {
        return 0, fmt.Errorf("customer not found: %w", err)
    }

    // Get all completed orders for customer
    orders, err := s.repository.GetOrdersByCustomerID(customerID, 0, 0) // 0,0 means no pagination
    if err != nil {
        return 0, fmt.Errorf("failed to retrieve customer orders: %w", err)
    }

    var totalValue float64
    for _, order := range orders {
        // Only count orders with completed payments
        payment, err := s.repository.GetPaymentByOrderID(order.ID)
        if err == nil && payment.Status == "completed" {
            totalValue += order.Total
        }
    }

    return totalValue, nil
}

```

## Initialize the Service
```go
func NewService(repository Repository) Service {
    return &service{repository: repository}
}

