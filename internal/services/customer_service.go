package services

import (
	"context"

	"github.com/astrica1/order-delay-report/internal/models"
	"github.com/astrica1/order-delay-report/internal/repositories"
)

type CustomerService interface {
	GetCustomerByID(ctx context.Context, id int) (*models.Customer, error)
	GetAllCustomers(ctx context.Context) ([]models.Customer, error)
	CreateNewCustomer(ctx context.Context, customer *models.Customer) error
	UpdateExistingCustomer(ctx context.Context, customer *models.Customer) error
	DeleteExistingCustomer(ctx context.Context, id int) error
}

type customerService struct {
	customerRepository repositories.CustomerRepository
}

func NewCustomerService(customerRepository repositories.CustomerRepository) CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}

func (s *customerService) GetCustomerByID(ctx context.Context, id int) (*models.Customer, error) {
	return s.customerRepository.Get(ctx, id)
}

func (s *customerService) GetAllCustomers(ctx context.Context) ([]models.Customer, error) {
	return s.customerRepository.GetAll(ctx)
}

func (s *customerService) CreateNewCustomer(ctx context.Context, customer *models.Customer) error {
	return s.customerRepository.Create(ctx, customer)
}

func (s *customerService) UpdateExistingCustomer(ctx context.Context, customer *models.Customer) error {
	return s.customerRepository.Update(ctx, customer)
}

func (s *customerService) DeleteExistingCustomer(ctx context.Context, id int) error {
	return s.customerRepository.Delete(ctx, id)
}
