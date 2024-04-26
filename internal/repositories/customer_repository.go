package repositories

import (
	"context"

	"github.com/astrica1/order-delay-report/internal/models"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Get(ctx context.Context, id int) (*models.Customer, error)
	GetAll(ctx context.Context) ([]models.Customer, error)
	Create(ctx context.Context, customer *models.Customer) error
	Update(ctx context.Context, customer *models.Customer) error
	Delete(ctx context.Context, id int) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (r *customerRepository) Get(ctx context.Context, id int) (*models.Customer, error) {
	var customer models.Customer
	if err := r.db.WithContext(ctx).First(&customer, id).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *customerRepository) GetAll(ctx context.Context) ([]models.Customer, error) {
	var customers []models.Customer
	if err := r.db.WithContext(ctx).Find(&customers).Error; err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *customerRepository) Create(ctx context.Context, customer *models.Customer) error {
	if err := r.db.WithContext(ctx).Create(customer).Error; err != nil {
		return err
	}

	return nil
}

func (r *customerRepository) Update(ctx context.Context, customer *models.Customer) error {
	if err := r.db.WithContext(ctx).Save(customer).Error; err != nil {
		return err
	}

	return nil
}

func (r *customerRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&models.Customer{}, id).Error; err != nil {
		return err
	}

	return nil
}
