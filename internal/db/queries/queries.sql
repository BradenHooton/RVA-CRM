-- name: GetCustomer :one
SELECT * FROM customers WHERE id = $1;

-- name: CreateCustomer :one
INSERT INTO customers (first_name, last_name, email, phone, company_name, job_title, status, customer_type, source, credit_limit, total_spent, last_purchase_at, addresses, orders, projects, notes, tags, custom_fields) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) RETURNING *;

-- name: UpdateCustomer :one
UPDATE customers SET first_name = $2, last_name = $3, email = $4, phone = $5, company_name = $6, job_title = $7, status = $8, customer_type = $9, source = $10, credit_limit = $11, total_spent = $12, last_purchase_at = $13, addresses = $14, orders = $15, projects = $16, notes = $17, tags = $18, custom_fields = $19 WHERE id = $1 RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM customers WHERE id = $1;

-- name: GetOrder :one
SELECT * FROM orders WHERE id = $1;

-- name: CreateOrder :one
INSERT INTO orders (customer_id, total_price, payment_status, order_status, order_type, order_date, delivery_date, delivery_address, delivery_status, delivery_date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING *;

-- name: UpdateOrder :one
UPDATE orders SET customer_id = $2, total_price = $3, payment_status = $4, order_status = $5, order_type = $6, order_date = $7, delivery_date = $8, delivery_address = $9, delivery_status = $10, delivery_date = $11, created_at = $12, updated_at = $13 WHERE id = $1 RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders WHERE id = $1;

-- name: GetOrderItems :many
SELECT * FROM order_items WHERE order_id = $1;

-- name: GetOrderItem :one
SELECT * FROM order_items WHERE id = $1;

-- name: CreateOrderItem :one
INSERT INTO order_items (order_id, product_id, quantity, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: UpdateOrderItem :one
UPDATE order_items SET order_id = $2, product_id = $3, quantity = $4, price = $5, created_at = $6, updated_at = $7 WHERE id = $1 RETURNING *;

-- name: DeleteOrderItem :exec
DELETE FROM order_items WHERE id = $1;

-- name: GetProduct :one
SELECT * FROM products WHERE id = $1;

-- name: CreateProduct :one
INSERT INTO products (name, description, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: UpdateProduct :one
UPDATE products SET name = $2, description = $3, price = $4, created_at = $5, updated_at = $6 WHERE id = $1 RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;

-- name: GetProject :one
SELECT * FROM projects WHERE id = $1;

-- name: CreateProject :one
INSERT INTO projects (name, description, customer_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: UpdateProject :one
UPDATE projects SET name = $2, description = $3, customer_id = $4, created_at = $5, updated_at = $6 WHERE id = $1 RETURNING *;

-- name: DeleteProject :exec
DELETE FROM projects WHERE id = $1;

-- name: GetNote :one
SELECT * FROM notes WHERE id = $1;

-- name: CreateNote :one
INSERT INTO notes (project_id, content, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: DeleteNote :exec
DELETE FROM notes WHERE id = $1;

-- name: UpdateNote :one
UPDATE notes SET project_id = $2, content = $3, created_at = $4, updated_at = $5 WHERE id = $1 RETURNING *;

-- name: GetPayment :one
SELECT * FROM payments WHERE id = $1;

-- name: DeletePayment :exec
DELETE FROM payments WHERE id = $1;

-- name: CreatePayment :one
INSERT INTO payments (order_id, amount, payment_method, payment_status, payment_date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: UpdatePayment :one
UPDATE payments SET order_id = $2, amount = $3, payment_method = $4, payment_status = $5, payment_date = $6, created_at = $7, updated_at = $8 WHERE id = $1 RETURNING *;

