package repositories

import (
	"context"

	"github.com/astrica1/order-delay-report/internal/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Get(ctx context.Context, id int) (*models.Order, error)
	GetAll(ctx context.Context) ([]models.Order, error)
	Create(ctx context.Context, order *models.Order) error
	Update(ctx context.Context, order *models.Order) error
	Delete(ctx context.Context, id int) error
	GetTripStatusByOrderID(ctx context.Context, orderID int) (models.TripStatus, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) Get(ctx context.Context, id int) (*models.Order, error) {
	var order models.Order
	if err := r.db.WithContext(ctx).First(&order, id).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *orderRepository) GetAll(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.WithContext(ctx).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *orderRepository) Create(ctx context.Context, order *models.Order) error {
	if err := r.db.WithContext(ctx).Create(order).Error; err != nil {
		return err
	}

	return nil
}

func (r *orderRepository) Update(ctx context.Context, order *models.Order) error {
	if err := r.db.WithContext(ctx).Save(order).Error; err != nil {
		return err
	}

	return nil
}

func (r *orderRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&models.Order{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *orderRepository) GetTripStatusByOrderID(ctx context.Context, orderID int) (models.TripStatus, error) {
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
