package repositories

import (
	"context"

	"github.com/astrica1/order-delay-report/internal/models"
	"gorm.io/gorm"
)

type DelayReportRepository struct {
	db *gorm.DB
}

func NewDelayReportRepository(db *gorm.DB) *DelayReportRepository {
	return &DelayReportRepository{
		db: db,
	}
}

func (r *DelayReportRepository) Get(ctx context.Context, id int) (*models.DelayReport, error) {
	var delayReport models.DelayReport
	if err := r.db.WithContext(ctx).First(&delayReport, id).Error; err != nil {
		return nil, err
	}

	return &delayReport, nil
}

func (r *DelayReportRepository) GetAll(ctx context.Context) ([]models.DelayReport, error) {
	var delayReport []models.DelayReport
	if err := r.db.WithContext(ctx).Find(&delayReport).Error; err != nil {
		return nil, err
	}

	return delayReport, nil
}

func (r *DelayReportRepository) Create(ctx context.Context, delayReport *models.DelayReport) error {
	if err := r.db.WithContext(ctx).Create(delayReport).Error; err != nil {
		return err
	}

	return nil
}

func (r *DelayReportRepository) Update(ctx context.Context, delayReport *models.DelayReport) error {
	if err := r.db.WithContext(ctx).Save(delayReport).Error; err != nil {
		return err
	}

	return nil
}

func (r *DelayReportRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&models.DelayReport{}, id).Error; err != nil {
		return err
	}

	return nil
}
