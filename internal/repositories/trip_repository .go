package repositories

import (
	"context"

	"github.com/astrica1/order-delay-report/internal/models"
	"gorm.io/gorm"
)

type TripRepository struct {
	db *gorm.DB
}

func NewTripRepository(db *gorm.DB) *TripRepository {
	return &TripRepository{
		db: db,
	}
}

func (r *TripRepository) Get(ctx context.Context, id int) (*models.Trip, error) {
	var trip models.Trip
	if err := r.db.WithContext(ctx).First(&trip, id).Error; err != nil {
		return nil, err
	}

	return &trip, nil
}

func (r *TripRepository) GetAll(ctx context.Context) ([]models.Trip, error) {
	var trips []models.Trip
	if err := r.db.WithContext(ctx).Find(&trips).Error; err != nil {
		return nil, err
	}

	return trips, nil
}

func (r *TripRepository) Create(ctx context.Context, trip *models.Trip) error {
	if err := r.db.WithContext(ctx).Create(trip).Error; err != nil {
		return err
	}

	return nil
}

func (r *TripRepository) Update(ctx context.Context, trip *models.Trip) error {
	if err := r.db.WithContext(ctx).Save(trip).Error; err != nil {
		return err
	}

	return nil
}

func (r *TripRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&models.Trip{}, id).Error; err != nil {
		return err
	}

	return nil
}
