package main

import (
	"database/sql"
	"context"


	"github.com/google/uuid"
)

// Repository interface defines all database operations
type CustomerRepository interface {
	// Customer operations
	GetCustomer(ctx context.Context, id uuid.UUID) (*Customer, error)
	CreateCustomer(ctx context.Context, customer *Customer) error
	UpdateCustomer(ctx context.Context, customer *Customer) error
	DeleteCustomer(ctx context.Context, id uuid.UUID) error
}

type OrderRepository interface {
	// Order operations
	GetOrder(ctx context.Context, id uuid.UUID) (*Order, error)
	CreateOrder(ctx context.Context, order *Order) error
	UpdateOrder(ctx context.Context, order *Order) error
	DeleteOrder(ctx context.Context, id uuid.UUID) error
}

type OrderItemRepository interface {
	// OrderItem operations
	GetOrderItem(ctx context.Context, id uuid.UUID) (*OrderItem, error)
	CreateOrderItem(ctx context.Context, orderItem *OrderItem) error
	UpdateOrderItem(ctx context.Context, orderItem *OrderItem) error
	DeleteOrderItem(ctx context.Context, id uuid.UUID) error
}

type ProductRepository interface {
	// Product operations
	GetProduct(ctx context.Context, id uuid.UUID) (*Product, error)
	CreateProduct(ctx context.Context, product *Product) error
	UpdateProduct(ctx context.Context, product *Product) error
	DeleteProduct(ctx context.Context, id uuid.UUID) error
}

type ProjectRepository interface {
	// Project operations
	GetProject(ctx context.Context, id uuid.UUID) (*Project, error)
	CreateProject(ctx context.Context, project *Project) error
	UpdateProject(ctx context.Context, project *Project) error
	DeleteProject(ctx context.Context, id uuid.UUID) error
}

type NoteRepository interface {
	// Note operations
	GetNote(ctx context.Context, id uuid.UUID) (*Note, error)
	CreateNote(ctx context.Context, note *Note) error
	UpdateNote(ctx context.Context, note *Note) error
	DeleteNote(ctx context.Context, id uuid.UUID) error
}

type PaymentRepository interface {
	// Payment operations
	GetPayment(ctx context.Context, id uuid.UUID) (*Payment, error)
	CreatePayment(ctx context.Context, payment *Payment) error
	UpdatePayment(ctx context.Context, payment *Payment) error
	DeletePayment(ctx context.Context, id uuid.UUID) error
}

type TaskRepository interface {
	// Task operations
	GetTask(ctx context.Context, id uuid.UUID) (*ProjectTask, error)
	CreateTask(ctx context.Context, task *ProjectTask) error
	UpdateTask(ctx context.Context, task *ProjectTask) error
	DeleteTask(ctx context.Context, id uuid.UUID) error
}

// repository struct implements the Repository interface
type PostgresRepository struct {
	db *sql.DB
}

// NewRepository creates a new repository instance
func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

// Customer repository methods
func (r *PostgresRepository) GetCustomer(ctx context.Context, id uuid.UUID) (*Customer, error) {
	rows, err := r.db.Query("SELECT * FROM customers WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customer Customer
	if rows.Next() {
		err = rows.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Phone, &customer.CompanyName, &customer.JobTitle, &customer.Status, &customer.CustomerType, &customer.Source, &customer.CreditLimit, &customer.TotalSpent, &customer.LastPurchaseAt, &customer.Addresses, &customer.Orders, &customer.Projects, &customer.Notes, &customer.Tags, &customer.CustomFields)
		if err != nil {
			return nil, err
		}
	}
	return &customer, nil
}

func (r *PostgresRepository) CreateCustomer(ctx context.Context, customer *Customer) error {
	rows, err := r.db.Query("INSERT INTO customers (first_name, last_name, email, phone, company_name, job_title, status, customer_type, source, credit_limit, total_spent, last_purchase_at, addresses, orders, projects, notes, tags, custom_fields) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) RETURNING *", customer.FirstName, customer.LastName, customer.Email, customer.Phone, customer.CompanyName, customer.JobTitle, customer.Status, customer.CustomerType, customer.Source, customer.CreditLimit, customer.TotalSpent, customer.LastPurchaseAt, customer.Addresses, customer.Orders, customer.Projects, customer.Notes, customer.Tags, customer.CustomFields)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Phone, &customer.CompanyName, &customer.JobTitle, &customer.Status, &customer.CustomerType, &customer.Source, &customer.CreditLimit, &customer.TotalSpent, &customer.LastPurchaseAt, &customer.Addresses, &customer.Orders, &customer.Projects, &customer.Notes, &customer.Tags, &customer.CustomFields)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresRepository) UpdateCustomer(ctx context.Context, customer *Customer) error {
	rows, err := r.db.Query("UPDATE customers SET first_name = $2, last_name = $3, email = $4, phone = $5, company_name = $6, job_title = $7, status = $8, customer_type = $9, source = $10, credit_limit = $11, total_spent = $12, last_purchase_at = $13, addresses = $14, orders = $15, projects = $16, notes = $17, tags = $18, custom_fields = $19 WHERE id = $1 RETURNING *", customer.FirstName, customer.LastName, customer.Email, customer.Phone, customer.CompanyName, customer.JobTitle, customer.Status, customer.CustomerType, customer.Source, customer.CreditLimit, customer.TotalSpent, customer.LastPurchaseAt, customer.Addresses, customer.Orders, customer.Projects, customer.Notes, customer.Tags, customer.CustomFields)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Phone, &customer.CompanyName, &customer.JobTitle, &customer.Status, &customer.CustomerType, &customer.Source, &customer.CreditLimit, &customer.TotalSpent, &customer.LastPurchaseAt, &customer.Addresses, &customer.Orders, &customer.Projects, &customer.Notes, &customer.Tags, &customer.CustomFields)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresRepository) DeleteCustomer(id uuid.UUID) error {
	rows, err := r.db.Query("DELETE FROM customers WHERE id = $1 RETURNING *", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

// Order repository methods
func (r *PostgresRepository) GetOrder(ctx context.Context, id uuid.UUID) (*Order, error) {
	rows, err := r.db.Query("SELECT * FROM orders WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var order Order
	if rows.Next() {
		err = rows.Scan(&order.ID, &order.OrderNumber, &order.Customer, &order.Status, &order.SubTotal, &order.TaxAmount, &order.Discount, &order.Total, &order.OrderDate, &order.ShippedDate, &order.DeliveredDate, &order.BillingAddressID, &order.ShippingAddressID, &order.OrderCustomer, &order.OrderItems, &order.Payments, &order.BillingAddress, &order.ShippingAddress, &order.Notes, &order.Metadata)
		if err != nil {
			return nil, err
		}
	}
	return &order, nil
}

func (r *PostgresRepository) CreateOrder(ctx context.Context, order *Order) error {
	rows, err := r.db.Query("INSERT INTO orders (order_number, customer_id, status, sub_total, tax_amount, discount, total, order_date, shipped_date, delivered_date, billing_address_id, shipping_address_id, order_customer, order_items, payments, billing_address, shipping_address, notes, metadata) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19) RETURNING *", order.OrderNumber, order.Customer, order.Status, order.SubTotal, order.TaxAmount, order.Discount, order.Total, order.OrderDate, order.ShippedDate, order.DeliveredDate, order.BillingAddressID, order.ShippingAddressID, order.OrderCustomer, order.OrderItems, order.Payments, order.BillingAddress, order.ShippingAddress, order.Notes, order.Metadata)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&order.ID, &order.OrderNumber, &order.Customer, &order.Status, &order.SubTotal, &order.TaxAmount, &order.Discount, &order.Total, &order.OrderDate, &order.ShippedDate, &order.DeliveredDate, &order.BillingAddressID, &order.ShippingAddressID, &order.OrderCustomer, &order.OrderItems, &order.Payments, &order.BillingAddress, &order.ShippingAddress, &order.Notes, &order.Metadata)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresRepository) UpdateOrder(ctx context.Context, order *Order) error {
	rows, err := r.db.Query("UPDATE orders SET order_number = $2, customer_id = $3, status = $4, sub_total = $5, tax_amount = $6, discount = $7, total = $8, order_date = $9, shipped_date = $10, delivered_date = $11, billing_address_id = $12, shipping_address_id = $13, order_customer = $14, order_items = $15, payments = $16, billing_address = $17, shipping_address = $18, notes = $19, metadata = $20 WHERE id = $1 RETURNING *", order.OrderNumber, order.Customer, order.Status, order.SubTotal, order.TaxAmount, order.Discount, order.Total, order.OrderDate, order.ShippedDate, order.DeliveredDate, order.BillingAddressID, order.ShippingAddressID, order.OrderCustomer, order.OrderItems, order.Payments, order.BillingAddress, order.ShippingAddress, order.Notes, order.Metadata)
	if err != nil {
		return err
	}
	defer rows.Close()


	if rows.Next() {
		err = rows.Scan(&order.ID, &order.OrderNumber, &order.Customer, &order.Status, &order.SubTotal, &order.TaxAmount, &order.Discount, &order.Total, &order.OrderDate, &order.ShippedDate, &order.DeliveredDate, &order.BillingAddressID, &order.ShippingAddressID, &order.OrderCustomer, &order.OrderItems, &order.Payments, &order.BillingAddress, &order.ShippingAddress, &order.Notes, &order.Metadata)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresRepository) DeleteOrder(ctx context.Context, id uuid.UUID) error {
	rows, err := r.db.Query("DELETE FROM orders WHERE id = $1 RETURNING *", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil

}

// Product repository methods
func (r *PostgresRepository) GetProduct(ctx context.Context, id uuid.UUID) (*Product, error) {
	rows, err := r.db.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var product Product
	if rows.Next() {
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
	}
	return &product, nil
}

func (r *PostgresRepository) UpdateProduct(ctx context.Context, product *Product) error {
	rows, err := r.db.Query("UPDATE products SET name = $2, price = $3 WHERE id = $1 RETURNING *", product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresRepository) DeleteProduct(id uuid.UUID) error {
	rows, err := r.db.Query("DELETE FROM products WHERE id = $1 RETURNING *", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

// Project repository methods
func (r *PostgresRepository) GetProject(id uuid.UUID) (*Project, error) {
	rows, err := r.db.Query("SELECT * FROM projects WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	var project Project
	if rows.Next() {
		err = rows.Scan(&project.ID, &project.ProjectName, &project.ProjectDescription, &project.ProjectStatus, &project.ProjectStartDate, &project.ProjectEndDate, &project.ProjectBudget, &project.ProjectProgress, &project.ProjectNotes, &project.ProjectTasks)
		if err != nil {
			return nil, err
		}
	}
	return &project, nil
}

func (r *PostgresRepository) CreateProject(ctx context.Context, project *Project) error {
	rows, err := r.db.Query("INSERT INTO projects (project_name, project_description, project_status, project_start_date, project_end_date, project_budget, project_progress, project_notes, project_tasks) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *", project.ProjectName, project.ProjectDescription, project.ProjectStatus, project.ProjectStartDate, project.ProjectEndDate, project.ProjectBudget, project.ProjectProgress, project.ProjectNotes, project.ProjectTasks)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&project.ID, &project.ProjectName, &project.ProjectDescription, &project.ProjectStatus, &project.ProjectStartDate, &project.ProjectEndDate, &project.ProjectBudget, &project.ProjectProgress, &project.ProjectNotes, &project.ProjectTasks)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresRepository) UpdateProject(ctx context.Context, project *Project) error {
	rows, err := r.db.Query("UPDATE projects SET project_name = $2, project_description = $3, project_status = $4, project_start_date = $5, project_end_date = $6, project_budget = $7, project_progress = $8, project_notes = $9, project_tasks = $10 WHERE id = $1 RETURNING *", project.ProjectName, project.ProjectDescription, project.ProjectStatus, project.ProjectStartDate, project.ProjectEndDate, project.ProjectBudget, project.ProjectProgress, project.ProjectNotes, project.ProjectTasks)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&project.ID, &project.ProjectName, &project.ProjectDescription, &project.ProjectStatus, &project.ProjectStartDate, &project.ProjectEndDate, &project.ProjectBudget, &project.ProjectProgress, &project.ProjectNotes, &project.ProjectTasks)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresRepository) DeleteProject(id uuid.UUID) error {
	rows, err := r.db.Query("DELETE FROM projects WHERE id = $1 RETURNING *", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

// Note repository methods
func (r *PostgresRepository) GetNote(ctx context.Context, id uuid.UUID) (*Note, error) {
	rows, err := r.db.Query("SELECT * FROM notes WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var note Note
	if rows.Next() {
		err = rows.Scan(&note.ID, &note.Note, &note.NoteType, &note.NoteDate, &note.NoteAuthor, &note.NoteStatus, &note.NotePriority, &note.NoteCategory)
		if err != nil {
			return nil, err
		}
	}
	return &note, nil
}

func (r *PostgresRepository) CreateNote(ctx context.Context, note *Note) error {
	rows, err := r.db.Query("INSERT INTO notes (note, note_type, note_date, note_author, note_status, note_priority, note_category) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *", note.Note, note.NoteType, note.NoteDate, note.NoteAuthor, note.NoteStatus, note.NotePriority, note.NoteCategory)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&note.ID, &note.Note, &note.NoteType, &note.NoteDate, &note.NoteAuthor, &note.NoteStatus, &note.NotePriority, &note.NoteCategory)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresRepository) UpdateNote(ctx context.Context, note *Note) error {
	rows, err := r.db.Query("UPDATE notes SET note = $2, note_type = $3, note_date = $4, note_author = $5, note_status = $6, note_priority = $7, note_category = $8 WHERE id = $1 RETURNING *", note.Note, note.NoteType, note.NoteDate, note.NoteAuthor, note.NoteStatus, note.NotePriority, note.NoteCategory)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&note.ID, &note.Note, &note.NoteType, &note.NoteDate, &note.NoteAuthor, &note.NoteStatus, &note.NotePriority, &note.NoteCategory)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresRepository) DeleteNote(id uuid.UUID) error {
	rows, err := r.db.Query("DELETE FROM notes WHERE id = $1 RETURNING *", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil

}

// Payment repository methods
func (r *PostgresRepository) GetPayment(ctx context.Context, id uuid.UUID) (*Payment, error) {
	rows, err := r.db.Query("SELECT * FROM payments WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payment Payment
	if rows.Next() {
		err = rows.Scan(&payment.ID, &payment.Order, &payment.Amount, &payment.Status)
		if err != nil {
			return nil, err
		}
	}
	return &payment, nil
}

func (r *PostgresRepository) CreatePayment(ctx context.Context, payment *Payment) error {
	rows, err := r.db.Query("INSERT INTO payments (order_id, amount, status) VALUES ($1, $2, $3) RETURNING *", payment.Order, payment.Amount, payment.Status)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&payment.ID, &payment.Order, &payment.Amount, &payment.Status)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresRepository) UpdatePayment(ctx context.Context, payment *Payment) error {
	rows, err := r.db.Query("UPDATE payments SET order_id = $2, amount = $3, status = $4 WHERE id = $1 RETURNING *", payment.Order, payment.Amount, payment.Status)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&payment.ID, &payment.Order, &payment.Amount, &payment.Status)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresRepository) DeletePayment(id uuid.UUID) error {
	rows, err := r.db.Query("DELETE FROM payments WHERE id = $1 RETURNING *", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

// Task repository methods
func (r *PostgresRepository) GetTask(ctx context.Context, id uuid.UUID) (*ProjectTask, error) {
	rows, err := r.db.Query("SELECT * FROM tasks WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var task ProjectTask
	if rows.Next() {
		err = rows.Scan(&task.ID, &task.TaskName, &task.TaskDescription, &task.TaskStatus, &task.TaskStartDate, &task.TaskEndDate, &task.Assignee, &task.TaskType, &task.TaskPriority)
		if err != nil {
			return nil, err
		}
	}
	return &task, nil
}

func (r *PostgresRepository) CreateTask(ctx context.Context, task *ProjectTask) error {
	rows, err := r.db.Query("INSERT INTO tasks (task_name, task_description, task_status, task_start_date, task_end_date, assignee, task_type, task_priority) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *", task.TaskName, task.TaskDescription, task.TaskStatus, task.TaskStartDate, task.TaskEndDate, task.Assignee, task.TaskType, task.TaskPriority)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&task.ID, &task.TaskName, &task.TaskDescription, &task.TaskStatus, &task.TaskStartDate, &task.TaskEndDate, &task.Assignee, &task.TaskType, &task.TaskPriority)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresRepository) UpdateTask(ctx context.Context, task *ProjectTask) error {
	rows, err := r.db.Query("UPDATE tasks SET task_name = $2, task_description = $3, task_status = $4, task_start_date = $5, task_end_date = $6, assignee = $7, task_type = $8, task_priority = $9 WHERE id = $1 RETURNING *", task.TaskName, task.TaskDescription, task.TaskStatus, task.TaskStartDate, task.TaskEndDate, task.Assignee, task.TaskType, task.TaskPriority)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&task.ID, &task.TaskName, &task.TaskDescription, &task.TaskStatus, &task.TaskStartDate, &task.TaskEndDate, &task.Assignee, &task.TaskType, &task.TaskPriority)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresRepository) DeleteTask(id uuid.UUID) error {
	rows, err := r.db.Query("DELETE FROM tasks WHERE id = $1 RETURNING *", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return nil
	}
	return nil
}