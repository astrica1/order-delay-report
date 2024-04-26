package repositories

import (
	"context"

	"github.com/astrica1/order-delay-report/internal/models"
	"gorm.io/gorm"
)

type DelayReportRepository interface {
	Get(ctx context.Context, id int) (*models.DelayReport, error)
	GetAll(ctx context.Context) ([]models.DelayReport, error)
	Create(ctx context.Context, delayReport *models.DelayReport) error
	Update(ctx context.Context, delayReport *models.DelayReport) error
	Delete(ctx context.Context, id int) error
}

type delayReportRepository struct {
	db *gorm.DB
}

func NewDelayReportRepository(db *gorm.DB) DelayReportRepository {
	return &delayReportRepository{
		db: db,
	}
}

func (r *delayReportRepository) Get(ctx context.Context, id int) (*models.DelayReport, error) {
	var delayReport models.DelayReport
	if err := r.db.WithContext(ctx).First(&delayReport, id).Error; err != nil {
		return nil, err
	}

	return &delayReport, nil
}

func (r *delayReportRepository) GetAll(ctx context.Context) ([]models.DelayReport, error) {
	var delayReport []models.DelayReport
	if err := r.db.WithContext(ctx).Find(&delayReport).Error; err != nil {
		return nil, err
	}

	return delayReport, nil
}

func (r *delayReportRepository) Create(ctx context.Context, delayReport *models.DelayReport) error {
	if err := r.db.WithContext(ctx).Create(delayReport).Error; err != nil {
		return err
	}

	return nil
}

func (r *delayReportRepository) Update(ctx context.Context, delayReport *models.DelayReport) error {
	if err := r.db.WithContext(ctx).Save(delayReport).Error; err != nil {
		return err
	}

	return nil
}

func (r *delayReportRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&models.DelayReport{}, id).Error; err != nil {
		return err
	}

	return nil
}
