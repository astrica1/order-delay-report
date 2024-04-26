package repositories

import (
	"context"

	"github.com/astrica1/order-delay-report/internal/models"
	"gorm.io/gorm"
)

type CourierRepository struct {
	db *gorm.DB
}

func NewCourierRepository(db *gorm.DB) *CourierRepository {
	return &CourierRepository{
		db: db,
	}
}

func (r *CourierRepository) Get(ctx context.Context, id int) (*models.Courier, error) {
	var courier models.Courier
	if err := r.db.WithContext(ctx).First(&courier, id).Error; err != nil {
		return nil, err
	}

	return &courier, nil
}

func (r *CourierRepository) GetAll(ctx context.Context) ([]models.Courier, error) {
	var couriers []models.Courier
	if err := r.db.WithContext(ctx).Find(&couriers).Error; err != nil {
		return nil, err
	}

	return couriers, nil
}

func (r *CourierRepository) Create(ctx context.Context, courier *models.Courier) error {
	if err := r.db.WithContext(ctx).Create(courier).Error; err != nil {
		return err
	}

	return nil
}

func (r *CourierRepository) Update(ctx context.Context, courier *models.Courier) error {
	if err := r.db.WithContext(ctx).Save(courier).Error; err != nil {
		return err
	}

	return nil
}

func (r *CourierRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&models.Courier{}, id).Error; err != nil {
		return err
	}

	return nil
}
