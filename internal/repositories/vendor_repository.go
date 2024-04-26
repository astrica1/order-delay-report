package repositories

import (
	"context"

	"github.com/astrica1/order-delay-report/internal/models"
	"gorm.io/gorm"
)

type VendorRepository struct {
	db *gorm.DB
}

func NewVendorRepository(db *gorm.DB) *VendorRepository {
	return &VendorRepository{
		db: db,
	}
}

func (r *VendorRepository) Get(ctx context.Context, id int) (*models.Vendor, error) {
	var vendor models.Vendor
	if err := r.db.WithContext(ctx).First(&vendor, id).Error; err != nil {
		return nil, err
	}

	return &vendor, nil
}

func (r *VendorRepository) GetAll(ctx context.Context) ([]models.Vendor, error) {
	var vendors []models.Vendor
	if err := r.db.WithContext(ctx).Find(&vendors).Error; err != nil {
		return nil, err
	}

	return vendors, nil
}

func (r *VendorRepository) Create(ctx context.Context, vendor *models.Vendor) error {
	if err := r.db.WithContext(ctx).Create(vendor).Error; err != nil {
		return err
	}

	return nil
}

func (r *VendorRepository) Update(ctx context.Context, vendor *models.Vendor) error {
	if err := r.db.WithContext(ctx).Save(vendor).Error; err != nil {
		return err
	}

	return nil
}

func (r *VendorRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&models.Vendor{}, id).Error; err != nil {
		return err
	}

	return nil
}
