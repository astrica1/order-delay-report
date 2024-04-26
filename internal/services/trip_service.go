package services

import (
	"context"

	"github.com/astrica1/order-delay-report/internal/models"
	"github.com/astrica1/order-delay-report/internal/repositories"
)

type TripService interface {
	GetTripByID(ctx context.Context, id int) (*models.Trip, error)
	GetAllTrips(ctx context.Context) ([]models.Trip, error)
	GetTripByOrderID(ctx context.Context, orderID int) (*models.Trip, error)
}

type tripService struct {
	tripRepository repositories.TripRepository
}

func NewTripService(tripRepository repositories.TripRepository) TripService {
	return &tripService{
		tripRepository: tripRepository,
	}
}

func (s *tripService) GetTripByID(ctx context.Context, id int) (*models.Trip, error) {
	return s.tripRepository.Get(ctx, id)
}

func (s *tripService) GetAllTrips(ctx context.Context) ([]models.Trip, error) {
	return s.tripRepository.GetAll(ctx)
}

func (s *tripService) GetTripByOrderID(ctx context.Context, orderID int) (*models.Trip, error) {
	return s.tripRepository.Get(ctx, orderID)
}
