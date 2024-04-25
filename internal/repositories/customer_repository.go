package repositories

import (
	"context"

	"github.com/astrica1/order-delay-report/internal/models"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (r *CustomerRepository) Get(ctx context.Context, id int) (*models.Customer, error) {
	var customer models.Customer
	if err := r.db.First(&customer, id).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *CustomerRepository) GetAll(ctx context.Context) ([]models.Customer, error) {
	var customers []models.Customer
	if err := r.db.Find(&customers).Error; err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *CustomerRepository) Create(ctx context.Context, customer *models.Customer) error {
	if err := r.db.Create(customer).Error; err != nil {
		return err
	}

	return nil
}

func (r *CustomerRepository) Update(ctx context.Context, customer *models.Customer) error {
	if err := r.db.Save(customer).Error; err != nil {
		return err
	}

	return nil
}

func (r *CustomerRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.Delete(&models.Customer{}, id).Error; err != nil {
		return err
	}

	return nil
}
