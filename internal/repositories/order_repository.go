package repositories

import (
	"context"

	"github.com/astrica1/order-delay-report/internal/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Get(ctx context.Context, id int) (*models.Order, error) {
	var order models.Order
	if err := r.db.WithContext(ctx).First(&order, id).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) GetAll(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.WithContext(ctx).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) Create(ctx context.Context, order *models.Order) error {
	if err := r.db.WithContext(ctx).Create(order).Error; err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) Update(ctx context.Context, order *models.Order) error {
	if err := r.db.WithContext(ctx).Save(order).Error; err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&models.Order{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetTripStatusByOrderID(ctx context.Context, orderID int) (models.TripStatus, error) {
	var tripStatus models.TripStatus
	if err := r.db.WithContext(ctx).
		Joins("JOIN trips ON orders.trip_id = trips.id").
		Where("orders.id = ?", orderID).
		Select("trips.status").
		First(&tripStatus).Error; err != nil {
		return 0, err
	}
	return tripStatus, nil
}
