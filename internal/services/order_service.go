package services

import (
	"context"
	"errors"
	"time"

	"github.com/astrica1/order-delay-report/internal/models"
	"github.com/astrica1/order-delay-report/internal/repositories"
	"github.com/astrica1/order-delay-report/pkg/messages"
	"gorm.io/gorm"
)

type OrderService interface {
	GetOrderByID(ctx context.Context, id int) (*models.Order, error)
	GetAllOrders(ctx context.Context) ([]models.Order, error)
	CreateNewOrder(ctx context.Context, order *models.Order) error
	UpdateExistingOrder(ctx context.Context, order *models.Order) error
	DeleteExistingOrder(ctx context.Context, id int) error
	CreateDelayReportForOrder(ctx context.Context, orderID int) error
}

type orderService struct {
	orderRepository    repositories.OrderRepository
	delayReportService DelayReportService
	tripService        TripService
}

func NewOrderService(orderRepository repositories.OrderRepository, delayReportService DelayReportService, tripService TripService) OrderService {
	return &orderService{
		orderRepository:    orderRepository,
		delayReportService: delayReportService,
		tripService:        tripService,
	}
}

func (s *orderService) GetOrderByID(ctx context.Context, id int) (*models.Order, error) {
	return s.orderRepository.GetOrderByIDWithRelations(ctx, id)
}

func (s *orderService) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	return s.orderRepository.GetAllOrdersWithRelations(ctx)
}

func (s *orderService) CreateNewOrder(ctx context.Context, order *models.Order) error {
	return s.orderRepository.Create(ctx, order)
}

func (s *orderService) UpdateExistingOrder(ctx context.Context, order *models.Order) error {
	return s.orderRepository.Update(ctx, order)
}

func (s *orderService) DeleteExistingOrder(ctx context.Context, id int) error {
	return s.orderRepository.Delete(ctx, id)
}

func (s *orderService) CreateDelayReportForOrder(ctx context.Context, orderID int) error {
	report, err := s.delayReportService.GetDelayReportByOrderID(ctx, orderID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		report = nil
	}

	if report != nil && report.Status < 3 {
		return messages.ALREADY_EXISTS.AsError("request")
	}

	trip, err := s.tripService.GetTripByOrderID(ctx, orderID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		trip = nil
	}

	if trip != nil {
		if trip.Status == models.TripStatusDelivered {
			return messages.IS_INVALID.AsError("request")
		}

		if trip.Order.DeliveryTime.After(time.Now()) {
			return messages.IS_INVALID.AsError("request")
		}
	}

	order, err := s.orderRepository.Get(ctx, orderID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return messages.DOES_NOT_EXISTS.AsError("Order")
		}

		return err
	}

	if order.DeliveryTime.After(time.Now()) {
		return messages.IS_INVALID.AsError("request")
	}

	return s.delayReportService.CreateDelayReport(ctx, orderID)
}
