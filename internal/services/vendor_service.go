package services

import (
	"context"

	"github.com/astrica1/order-delay-report/internal/models"
	"github.com/astrica1/order-delay-report/internal/repositories"
)

type VendorService interface {
	GetVendorByID(ctx context.Context, id int) (*models.Vendor, error)
	GetAllVendors(ctx context.Context) ([]models.Vendor, error)
}

type vendorService struct {
	vendorRepository   repositories.VendorRepository
	delayReportService DelayReportService
}

func NewVendorService(vendorRepository repositories.VendorRepository, delayReportService DelayReportService) VendorService {
	return &vendorService{
		vendorRepository:   vendorRepository,
		delayReportService: delayReportService,
	}
}

func (s *vendorService) GetVendorByID(ctx context.Context, id int) (*models.Vendor, error) {
	return s.vendorRepository.Get(ctx, id)
}

func (s *vendorService) GetAllVendors(ctx context.Context) ([]models.Vendor, error) {
	return s.vendorRepository.GetAll(ctx)
}
