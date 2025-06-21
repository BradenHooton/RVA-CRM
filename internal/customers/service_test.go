package customers

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) GetCustomerByID(ctx context.Context, id uuid.UUID) (Customer, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(Customer), args.Error(1)
}

func (m *MockCustomerRepository) ListCustomers(ctx context.Context) ([]Customer, error) {
	args := m.Called(ctx)
	return args.Get(0).([]Customer), args.Error(1)
}

func (m *MockCustomerRepository) CreateCustomer(ctx context.Context, customer Customer) (*Customer, error) {
	args := m.Called(ctx, customer)
	return args.Get(0).(*Customer), args.Error(1)
}

func (m *MockCustomerRepository) UpdateCustomer(ctx context.Context, customer Customer) (*Customer, error) {
	args := m.Called(ctx, customer)
	return args.Get(0).(*Customer), args.Error(1)
}

func (m *MockCustomerRepository) DeleteCustomer(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type CustomerServiceTestSuite struct {
	suite.Suite
	mockRepo *MockCustomerRepository
	service CustomerService
}

func (s *CustomerServiceTestSuite) SetupTest() {
	s.mockRepo = new(MockCustomerRepository)
	s.service = NewCustomerService(s.mockRepo)
}

func (s *CustomerServiceTestSuite) TearDownTest() {
	s.mockRepo.AssertExpectations(s.T())
}

func TestCustomerServiceSuite(t *testing.T) {
	suite.Run(t, new(CustomerServiceTestSuite))
}

func (s *CustomerServiceTestSuite) TestCreateCustomer_Success() {
	// Arrange
	ctx := context.Background()
	customerID := uuid.New()
	inputCustomer := Customer{
		FirstName: "Luke",
		LastName: "Skywalker",
		Email: "luke@jedi.com"
	}

	expectedCustomer := &Customer{
		ID: customerID,
		FirstName: "Luke",
		LastName: "Skywalker",
		Email: "luke@jedi.com",
	}

	s.mockRepo.On("CreateCustomer", ctx, inputCustomer).Return(expectedCustomer, nil)

	// Act
	result, err := s.service.CreateCustomer(ctx, inputCustomer)

	// Asset
	s.NoError(err)
	s.Equal(expectedCustomer, result)
}

func (s *CustomerServiceTestSuite) TestCreateCustomerByID_Success() {
	// Arrange
	ctx := context.Background()
	customerID := uuid.New()
	expectedCustomer := Customer{
		ID: customerID,
		FirstName: "Luke",
		LastName: "Skywalker",
		Email: "luke@jedi.com",
	}

	s.mockRepo.On("GetCustomerByID", ctx, customerID).Return(expectedCustomer, nil)

	// Act
	result, err := s.service.GetCustomerByID(ctx, customerID)

	// Assert
	s.NoError(err)
	s.Equal(expectedCustomer, result)
}

func (s *CustomerServiceTestSuite) TestCreateCustomer_Error() {
	// Arrange
	ctx := context.Background()
	customerID := uuid.New()

	s.mockRepo.On("GetCustomerByID", ctx, customerID).Return(Customer{}, errors.New("customer not found"))

	// Act
	result, err := s.service.GetCustomerByID(ctx, customerID)

	// Assert
	s.Error(err)
	s.Equal(Customer{}, result)
	s.Contains(err.Error(), "customer not found")
}
