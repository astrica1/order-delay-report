package services

import (
	"context"
	"errors"
	"sort"
	"time"

	"github.com/astrica1/order-delay-report/internal/models"
	"github.com/astrica1/order-delay-report/internal/repositories"
	"gorm.io/gorm"
)

type DelayReportService interface {
	GetDelayReportByOrderID(ctx context.Context, orderID int) (*models.DelayReport, error)
	CreateDelayReport(ctx context.Context, orderID int) error
	GetWeeklyDelayReport(ctx context.Context) ([]string, error)
}

type delayReportService struct {
	delayReportRepository repositories.DelayReportRepository
}

func NewDelayReportService(delayReportRepository repositories.DelayReportRepository) DelayReportService {
	return &delayReportService{
		delayReportRepository: delayReportRepository,
	}
}

func (s *delayReportService) GetDelayReportByOrderID(ctx context.Context, orderID int) (*models.DelayReport, error) {
	return s.delayReportRepository.GetByOrderID(ctx, orderID)
}

func (s *delayReportService) CreateDelayReport(ctx context.Context, orderID int) error {
	return s.delayReportRepository.Create(ctx, &models.DelayReport{
		OrderID:    orderID,
		ReportTime: time.Now(),
		Status:     models.ReportStatusUnknown,
	})
}

func (s *delayReportService) GetWeeklyDelayReport(ctx context.Context) ([]string, error) {
	reports, err := s.delayReportRepository.GetWeeklyDelayReport(ctx)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		return nil, nil
	}

	orderDelay := make(map[string]time.Duration)

	for _, r := range reports {
		if r.ResolvedTime.After(r.ReportTime) {
			_, exists := orderDelay[r.Order.Vendor.Name]
			if !exists {
				orderDelay[r.Order.Vendor.Name] = r.ResolvedTime.Sub(r.ReportTime)
			} else {
				orderDelay[r.Order.Vendor.Name] += r.ResolvedTime.Sub(r.ReportTime)
			}
		}
	}

	sortedVendors := make([]string, 0, len(orderDelay))
	for key := range orderDelay {
		sortedVendors = append(sortedVendors, key)
	}

	sort.Slice(sortedVendors, func(i, j int) bool {
		return orderDelay[sortedVendors[i]] > orderDelay[sortedVendors[j]]
	})

	return sortedVendors, nil
}
