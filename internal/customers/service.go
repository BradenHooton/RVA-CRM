package customers

import (
	"context"

	"github.com/google/uuid"
)

type CustomerService interface {
	CustomerManager
}

type AddressService interface {
	AddressManager
}

type OpportunityService interface {
	OpportunityManager
}

type CustomerRepository interface {
	CustomerManager
}

type AddressRepository interface {
	AddressManager
}

type OpportunityRepository interface {
	OpportunityManager
}

type customerService struct {
	repo CustomerRepository
}

type addressService struct {
	repo AddressRepository
}

type opportunityService struct {
	repo OpportunityRepository
}

func NewCustomerService(repo CustomerRepository) CustomerService {
	return &customerService{repo: repo}
}

func NewAddressService(repo AddressRepository) AddressService {
	return &addressService{repo: repo}
}

func NewOpportunityService(repo OpportunityRepository) OpportunityService {
	return &opportunityService{repo: repo}
}

type CustomerManager interface {
	CustomerReader
	CustomerWriter
}

type CustomerReader interface {
	CustomerRetriever
	CustomerLister
}

type CustomerWriter interface {
	CustomerCreator
	CustomerUpdater
	CustomerDeleter
}

type CustomerRetriever interface {
	GetCustomerByID(ctx context.Context, id uuid.UUID) (Customer, error)
}

type CustomerLister interface {
	ListCustomers(ctx context.Context) ([]Customer, error)
}

type CustomerCreator interface {
	CreateCustomer(ctx context.Context, customer Customer) (*Customer, error)
}

type CustomerUpdater interface {
	UpdateCustomer(ctx context.Context, customer Customer) (*Customer, error)
}

type CustomerDeleter interface {
	DeleteCustomer(ctx context.Context, id uuid.UUID) error
}

type AddressManager interface {
	AddressReader
	AddressWriter
}

type AddressReader interface {
	AddressRetriever
	AddressLister
}

type AddressWriter interface {
	AddressCreator
	AddressUpdater
	AddressDeleter
}

type AddressRetriever interface {
	GetAddressByID(ctx context.Context, id uuid.UUID) (*Address, error)
}

type AddressLister interface {
	GetAddressesByCustomerID(ctx context.Context, customerID uuid.UUID) ([]*Address, error)
}

type AddressCreator interface {
	CreateAddress(ctx context.Context, address Address) (*Address, error)
}

type AddressUpdater interface {
	UpdateAddress(ctx context.Context, address Address) (*Address, error)
}

type AddressDeleter interface {
	DeleteAddress(ctx context.Context, id uuid.UUID) error
}

type OpportunityManager interface {
	OpportunityReader
	OpportunityWriter
}

type OpportunityReader interface {
	OpportunityRetriever
	OpportunityLister
}

type OpportunityWriter interface {
	OpportunityCreator
	OpportunityUpdater
	OpportunityDeleter
}

type OpportunityCreator interface {
	CreateOpportunity(ctx context.Context, opportunity Opportunity) (*Opportunity, error)
}

type OpportunityUpdater interface {
	UpdateOpportunity(ctx context.Context, opportunity Opportunity) (*Opportunity, error)
}

type OpportunityDeleter interface {
	DeleteOpportunity(ctx context.Context, id uuid.UUID) error
}

type OpportunityRetriever interface {
	GetOpportunityByID(ctx context.Context, id uuid.UUID) (*Opportunity, error)
}

type OpportunityLister interface {
	GetOpportunitiesByCustomerID(ctx context.Context, customerID uuid.UUID) ([]*Opportunity, error)
}

func (s *customerService) GetCustomerByID(ctx context.Context, id uuid.UUID) (Customer, error) {
	return s.repo.GetCustomerByID(ctx, id)
}

func (s *customerService) ListCustomers(ctx context.Context) ([]Customer, error) {
	return s.repo.ListCustomers(ctx)
}

func (s *customerService) CreateCustomer(ctx context.Context, customer Customer) (*Customer, error) {
	return s.repo.CreateCustomer(ctx, customer)
}

func (s *customerService) UpdateCustomer(ctx context.Context, customer Customer) (*Customer, error) {
	return s.repo.UpdateCustomer(ctx, customer)
}

func (s *customerService) DeleteCustomer(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteCustomer(ctx, id)
}

func (s *addressService) GetAddressByID(ctx context.Context, id uuid.UUID) (*Address, error) {
	return s.repo.GetAddressByID(ctx, id)
}

func (s *addressService) GetAddressesByCustomerID(ctx context.Context, customerID uuid.UUID) ([]*Address, error) {
	return s.repo.GetAddressesByCustomerID(ctx, customerID)
}

func (s *addressService) CreateAddress(ctx context.Context, address Address) (*Address, error) {
	return s.repo.CreateAddress(ctx, address)
}

func (s *addressService) UpdateAddress(ctx context.Context, address Address) (*Address, error) {
	return s.repo.UpdateAddress(ctx, address)
}

func (s *addressService) DeleteAddress(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteAddress(ctx, id)
}

func (s *opportunityService) GetOpportunityByID(ctx context.Context, id uuid.UUID) (*Opportunity, error) {
	return s.repo.GetOpportunityByID(ctx, id)
}

func (s *opportunityService) GetOpportunitiesByCustomerID(ctx context.Context, customerID uuid.UUID) ([]*Opportunity, error) {
	return s.repo.GetOpportunitiesByCustomerID(ctx, customerID)
}

func (s *opportunityService) CreateOpportunity(ctx context.Context, opportunity Opportunity) (*Opportunity, error) {
	return s.repo.CreateOpportunity(ctx, opportunity)
}

func (s *opportunityService) UpdateOpportunity(ctx context.Context, opportunity Opportunity) (*Opportunity, error) {
	return s.repo.UpdateOpportunity(ctx, opportunity)
}

func (s *opportunityService) DeleteOpportunity(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteOpportunity(ctx, id)
}
