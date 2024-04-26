package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/astrica1/order-delay-report/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DelayReportRepository interface {
	Get(ctx context.Context, id int) (*models.DelayReport, error)
	GetAll(ctx context.Context) ([]models.DelayReport, error)
	Create(ctx context.Context, delayReport *models.DelayReport) error
	Update(ctx context.Context, delayReport *models.DelayReport) error
	Delete(ctx context.Context, id int) error
	GetByOrderID(ctx context.Context, id int) (*models.DelayReport, error)
	GetWeeklyDelayReport(ctx context.Context) ([]models.DelayReport, error)
	PopReport(ctx context.Context, agentID int) (*models.DelayReport, error)
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

func (r *delayReportRepository) GetByOrderID(ctx context.Context, orderID int) (*models.DelayReport, error) {
	var delayReport models.DelayReport
	if err := r.db.WithContext(ctx).
		Where("order_id = ?", orderID).
		Preload("Order").
		Preload("Agent").
		Last(&delayReport).Error; err != nil {
		return nil, err
	}

	return &delayReport, nil
}

func (r *delayReportRepository) GetWeeklyDelayReport(ctx context.Context) ([]models.DelayReport, error) {
	var delayReports []models.DelayReport
	startOfLastWeek := time.Now().AddDate(0, 0, -7).Truncate(24 * time.Hour)
	if err := r.db.WithContext(ctx).
		Where("report_time > ?", startOfLastWeek).
		Preload("Vendor").
		Preload("Order").
		Preload("Agent").
		Find(&delayReports).Error; err != nil {
		return nil, err
	}

	return delayReports, nil
}

func (r *delayReportRepository) PopReport(ctx context.Context, agentID int) (*models.DelayReport, error) {
	var delayReports []models.DelayReport
	if err := r.db.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&delayReports).Error; err != nil {
		return nil, err
	}

	candidate := -1
	for i, r := range delayReports {
		if r.AgentID == agentID && r.Status < 3 {
			return &r, fmt.Errorf("you have already a report to check")
		}

		if candidate < 0 && r.Status == 1 && r.AgentID == 0 {
			candidate = i
		}
	}

	delayReport := delayReports[candidate]
	delayReport.AgentID = agentID
	if err := r.db.Clauses(clause.Locking{Strength: "UPDATE"}).Save(&delayReport).Error; err != nil {
		return nil, err
	}

	return &delayReport, nil
}
