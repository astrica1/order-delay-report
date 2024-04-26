package services

import (
	"context"

	"github.com/astrica1/order-delay-report/internal/models"
	"github.com/astrica1/order-delay-report/internal/repositories"
)

type CourierService interface {
	GetCourierByID(ctx context.Context, id int) (*models.Courier, error)
	GetAllCouriers(ctx context.Context) ([]models.Courier, error)
	CreateNewCourier(ctx context.Context, courier *models.Courier) error
	UpdateExistingCourier(ctx context.Context, courier *models.Courier) error
	DeleteExistingCourier(ctx context.Context, id int) error
}

type courierService struct {
	courierRepository repositories.CourierRepository
}

func NewCourierService(courierRepository repositories.CourierRepository) CourierService {
	return &courierService{
		courierRepository: courierRepository,
	}
}

func (s *courierService) GetCourierByID(ctx context.Context, id int) (*models.Courier, error) {
	return s.courierRepository.Get(ctx, id)
}

func (s *courierService) GetAllCouriers(ctx context.Context) ([]models.Courier, error) {
	return s.courierRepository.GetAll(ctx)
}

func (s *courierService) CreateNewCourier(ctx context.Context, courier *models.Courier) error {
	return s.courierRepository.Create(ctx, courier)
}

func (s *courierService) UpdateExistingCourier(ctx context.Context, courier *models.Courier) error {
	return s.courierRepository.Update(ctx, courier)
}

func (s *courierService) DeleteExistingCourier(ctx context.Context, id int) error {
	return s.courierRepository.Delete(ctx, id)
}
