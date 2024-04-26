package repositories

import (
	"context"

	"github.com/astrica1/order-delay-report/internal/models"
	"gorm.io/gorm"
)

type TripRepository interface {
	Get(ctx context.Context, id int) (*models.Trip, error)
	GetAll(ctx context.Context) ([]models.Trip, error)
	Create(ctx context.Context, trip *models.Trip) error
	Update(ctx context.Context, trip *models.Trip) error
	Delete(ctx context.Context, id int) error
}

type tripRepository struct {
	db *gorm.DB
}

func NewTripRepository(db *gorm.DB) TripRepository {
	return &tripRepository{
		db: db,
	}
}

func (r *tripRepository) Get(ctx context.Context, id int) (*models.Trip, error) {
	var trip models.Trip
	if err := r.db.WithContext(ctx).First(&trip, id).Error; err != nil {
		return nil, err
	}

	return &trip, nil
}

func (r *tripRepository) GetAll(ctx context.Context) ([]models.Trip, error) {
	var trips []models.Trip
	if err := r.db.WithContext(ctx).Find(&trips).Error; err != nil {
		return nil, err
	}

	return trips, nil
}

func (r *tripRepository) Create(ctx context.Context, trip *models.Trip) error {
	if err := r.db.WithContext(ctx).Create(trip).Error; err != nil {
		return err
	}

	return nil
}

func (r *tripRepository) Update(ctx context.Context, trip *models.Trip) error {
	if err := r.db.WithContext(ctx).Save(trip).Error; err != nil {
		return err
	}

	return nil
}

func (r *tripRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&models.Trip{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *tripRepository) GetTripByOrderID(ctx context.Context, orderID int) (*models.Trip, error) {
	var trip models.Trip
	if err := r.db.WithContext(ctx).
		Where("order_id = ?", orderID).
		Preload("Courier").
		Preload("Order").
		Last(&trip).Error; err != nil {
		return nil, err
	}

	return &trip, nil
}
