package customers

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type customerRepository struct {
	db *sql.DB
}

type addressRepository struct {
	db *sql.DB
}

type opportunityRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func NewAddressRepository(db *sql.DB) AddressRepository {
	return &addressRepository{db: db}
}

func NewOpportunityRepository(db *sql.DB) OpportunityRepository {
	return &opportunityRepository{db: db}
}

func (r *customerRepository) GetCustomerByID(ctx context.Context, id uuid.UUID) (Customer, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM customers WHERE id = $1", id)
	if err != nil {
		return Customer{}, err
	}
	defer rows.Close()

	var customer Customer
	if rows.Next() {
		err = rows.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Phone, &customer.CompanyName, &customer.JobTitle, &customer.Status, &customer.CustomerType, &customer.Source, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			return Customer{}, err
		}
	}
	return customer, nil
}

func (r *customerRepository) ListCustomers(ctx context.Context) ([]Customer, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var customers []Customer
	for rows.Next() {
		var customer Customer
		err = rows.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Phone, &customer.CompanyName, &customer.JobTitle, &customer.Status, &customer.CustomerType, &customer.Source, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (r *customerRepository) CreateCustomer(ctx context.Context, customer Customer) (*Customer, error) {
	rows, err := r.db.QueryContext(ctx, "INSERT INTO customers (id, first_name, last_name, email, phone, company_name, job_title, status, customer_type, source) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING *", customer.ID, customer.FirstName, customer.LastName, customer.Email, customer.Phone, customer.CompanyName, customer.JobTitle, customer.Status, customer.CustomerType, customer.Source)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var createdCustomer Customer
	if rows.Next() {
		err = rows.Scan(&createdCustomer.ID, &createdCustomer.FirstName, &createdCustomer.LastName, &createdCustomer.Email, &createdCustomer.Phone, &createdCustomer.CompanyName, &createdCustomer.JobTitle, &createdCustomer.Status, &createdCustomer.CustomerType, &createdCustomer.Source, &createdCustomer.CreatedAt, &createdCustomer.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &createdCustomer, nil
}

func (r *customerRepository) UpdateCustomer(ctx context.Context, customer Customer) (*Customer, error) {
	rows, err := r.db.QueryContext(ctx, "UPDATE customers SET first_name = $1, last_name = $2, email = $3, phone = $4, company_name = $5, job_title = $6, status = $7, customer_type = $8, source = $9 WHERE id = $10 RETURNING *", customer.FirstName, customer.LastName, customer.Email, customer.Phone, customer.CompanyName, customer.JobTitle, customer.Status, customer.CustomerType, customer.Source, customer.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var updatedCustomer Customer
	if rows.Next() {
		err = rows.Scan(&updatedCustomer.ID, &updatedCustomer.FirstName, &updatedCustomer.LastName, &updatedCustomer.Email, &updatedCustomer.Phone, &updatedCustomer.CompanyName, &updatedCustomer.JobTitle, &updatedCustomer.Status, &updatedCustomer.CustomerType, &updatedCustomer.Source, &updatedCustomer.CreatedAt, &updatedCustomer.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &updatedCustomer, nil
}

func (r *customerRepository) DeleteCustomer(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM customers WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *addressRepository) GetAddressByID(ctx context.Context, id uuid.UUID) (*Address, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM addresses WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var address Address
	if rows.Next() {
		err = rows.Scan(&address.ID, &address.CustomerID, &address.Type, &address.Street1, &address.Street2, &address.City, &address.State, &address.PostalCode, &address.Country, &address.IsDefault, &address.CreatedAt, &address.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &address, nil
}

func (r *addressRepository) GetAddressesByCustomerID(ctx context.Context, customerID uuid.UUID) ([]*Address, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM addresses WHERE customer_id = $1", customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []*Address
	for rows.Next() {
		var address Address
		err = rows.Scan(&address.ID, &address.CustomerID, &address.Type, &address.Street1, &address.Street2, &address.City, &address.State, &address.PostalCode, &address.Country, &address.IsDefault, &address.CreatedAt, &address.UpdatedAt)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, &address)
	}
	return addresses, nil
}

func (r *addressRepository) CreateAddress(ctx context.Context, address Address) (*Address, error) {
	rows, err := r.db.QueryContext(ctx, "INSERT INTO addresses (id, customer_id, type, street1, street2, city, state, postal_code, country, is_default) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING *", address.ID, address.CustomerID, address.Type, address.Street1, address.Street2, address.City, address.State, address.PostalCode, address.Country, address.IsDefault)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var createdAddress Address
	if rows.Next() {
		err = rows.Scan(&createdAddress.ID, &createdAddress.CustomerID, &createdAddress.Type, &createdAddress.Street1, &createdAddress.Street2, &createdAddress.City, &createdAddress.State, &createdAddress.PostalCode, &createdAddress.Country, &createdAddress.IsDefault, &createdAddress.CreatedAt, &createdAddress.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &createdAddress, nil
}

func (r *addressRepository) UpdateAddress(ctx context.Context, address Address) (*Address, error) {
	rows, err := r.db.QueryContext(ctx, "UPDATE addresses SET customer_id = $1, type = $2, street1 = $3, street2 = $4, city = $5, state = $6, postal_code = $7, country = $8, is_default = $9 WHERE id = $10 RETURNING *", address.CustomerID, address.Type, address.Street1, address.Street2, address.City, address.State, address.PostalCode, address.Country, address.IsDefault, address.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var updatedAddress Address
	if rows.Next() {
		err = rows.Scan(&updatedAddress.ID, &updatedAddress.CustomerID, &updatedAddress.Type, &updatedAddress.Street1, &updatedAddress.Street2, &updatedAddress.City, &updatedAddress.State, &updatedAddress.PostalCode, &updatedAddress.Country, &updatedAddress.IsDefault, &updatedAddress.CreatedAt, &updatedAddress.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &updatedAddress, nil
}

func (r *addressRepository) DeleteAddress(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM addresses WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *opportunityRepository) GetOpportunityByID(ctx context.Context, id uuid.UUID) (*Opportunity, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM opportunities WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var opportunity Opportunity
	if rows.Next() {
		err = rows.Scan(&opportunity.ID, &opportunity.CustomerID, &opportunity.Name, &opportunity.Description, &opportunity.Value, &opportunity.Stage, &opportunity.Probability, &opportunity.ExpectedCloseDate, &opportunity.ActualCloseDate, &opportunity.Source, &opportunity.CreatedAt, &opportunity.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &opportunity, nil
}

func (r *opportunityRepository) GetOpportunitiesByCustomerID(ctx context.Context, customerID uuid.UUID) ([]*Opportunity, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM opportunities WHERE customer_id = $1", customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var opportunities []*Opportunity
	for rows.Next() {
		var opportunity Opportunity
		err = rows.Scan(&opportunity.ID, &opportunity.CustomerID, &opportunity.Name, &opportunity.Description, &opportunity.Value, &opportunity.Stage, &opportunity.Probability, &opportunity.ExpectedCloseDate, &opportunity.ActualCloseDate, &opportunity.Source, &opportunity.CreatedAt, &opportunity.UpdatedAt)
		if err != nil {
			return nil, err
		}
		opportunities = append(opportunities, &opportunity)
	}
	return opportunities, nil
}

func (r *opportunityRepository) CreateOpportunity(ctx context.Context, opportunity Opportunity) (*Opportunity, error) {
	rows, err := r.db.QueryContext(ctx, "INSERT INTO opportunities (id, customer_id, name, description, value, stage, probability, expected_close_date, actual_close_date, source) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING *", opportunity.ID, opportunity.CustomerID, opportunity.Name, opportunity.Description, opportunity.Value, opportunity.Stage, opportunity.Probability, opportunity.ExpectedCloseDate, opportunity.ActualCloseDate, opportunity.Source)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var createdOpportunity Opportunity
	if rows.Next() {
		err = rows.Scan(&createdOpportunity.ID, &createdOpportunity.CustomerID, &createdOpportunity.Name, &createdOpportunity.Description, &createdOpportunity.Value, &createdOpportunity.Stage, &createdOpportunity.Probability, &createdOpportunity.ExpectedCloseDate, &createdOpportunity.ActualCloseDate, &createdOpportunity.Source, &createdOpportunity.CreatedAt, &createdOpportunity.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &createdOpportunity, nil
}

func (r *opportunityRepository) UpdateOpportunity(ctx context.Context, opportunity Opportunity) (*Opportunity, error) {
	rows, err := r.db.QueryContext(ctx, "UPDATE opportunities SET customer_id = $1, name = $2, description = $3, value = $4, stage = $5, probability = $6, expected_close_date = $7, actual_close_date = $8, source = $9 WHERE id = $10 RETURNING *", opportunity.CustomerID, opportunity.Name, opportunity.Description, opportunity.Value, opportunity.Stage, opportunity.Probability, opportunity.ExpectedCloseDate, opportunity.ActualCloseDate, opportunity.Source, opportunity.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var updatedOpportunity Opportunity
	if rows.Next() {
		err = rows.Scan(&updatedOpportunity.ID, &updatedOpportunity.CustomerID, &updatedOpportunity.Name, &updatedOpportunity.Description, &updatedOpportunity.Value, &updatedOpportunity.Stage, &updatedOpportunity.Probability, &updatedOpportunity.ExpectedCloseDate, &updatedOpportunity.ActualCloseDate, &updatedOpportunity.Source, &updatedOpportunity.CreatedAt, &updatedOpportunity.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &updatedOpportunity, nil
}

func (r *opportunityRepository) DeleteOpportunity(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM opportunities WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}