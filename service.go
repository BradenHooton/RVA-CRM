package main

import (
	"github.com/google/uuid"
)

type CustomerService interface {
	GetCustomer(id uuid.UUID) (*Customer, error)
	CreateCustomer(customer *Customer) error
	UpdateCustomer(customer *Customer) error
	DeleteCustomer(id uuid.UUID) error
}

type OrderService interface {
	GetOrder(id uuid.UUID) (*Order, error)
	CreateOrder(order *Order) error
	UpdateOrder(order *Order) error
	DeleteOrder(id uuid.UUID) error
}

type ProjectService interface {
	GetProject(id uuid.UUID) (*Project, error)
	CreateProject(project *Project) error
	UpdateProject(project *Project) error
	DeleteProject(id uuid.UUID) error
}

type NoteService interface {
	GetNote(id uuid.UUID) (*Note, error)
	CreateNote(note *Note) error
	UpdateNote(note *Note) error
	DeleteNote(id uuid.UUID) error
}

type OrderPaymentService interface {
	ProcessOrderWithPayment(order *Order, payment *Payment) error
}

type ServiceContainer struct {
	CustomerService CustomerService
	OrderService OrderService
	ProjectService ProjectService
	NoteService NoteService
	OrderPaymentService OrderPaymentService
}


type customerServiceImpl struct {
	repo CustomerRepository
}

