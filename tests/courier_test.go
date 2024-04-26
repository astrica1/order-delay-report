package services

import (
	"context"
	"errors"
	"testing"

	"github.com/astrica1/order-delay-report/internal/models"
	"github.com/astrica1/order-delay-report/internal/services"
	"github.com/stretchr/testify/assert"
)

type MockCourierRepository struct {
	GetFunc    func(ctx context.Context, id int) (*models.Courier, error)
	GetAllFunc func(ctx context.Context) ([]models.Courier, error)
	CreateFunc func(ctx context.Context, courier *models.Courier) error
	UpdateFunc func(ctx context.Context, courier *models.Courier) error
	DeleteFunc func(ctx context.Context, id int) error
}

func (m *MockCourierRepository) Get(ctx context.Context, id int) (*models.Courier, error) {
	return m.GetFunc(ctx, id)
}

func (m *MockCourierRepository) GetAll(ctx context.Context) ([]models.Courier, error) {
	return m.GetAllFunc(ctx)
}

func (m *MockCourierRepository) Create(ctx context.Context, courier *models.Courier) error {
	return m.CreateFunc(ctx, courier)
}

func (m *MockCourierRepository) Update(ctx context.Context, courier *models.Courier) error {
	return m.UpdateFunc(ctx, courier)
}

func (m *MockCourierRepository) Delete(ctx context.Context, id int) error {
	return m.DeleteFunc(ctx, id)
}

func TestCourierService_GetCourierByID(t *testing.T) {
	mockCourier := &models.Courier{
		ID:          1,
		FullName:    "Full Name",
		PhoneNumber: "9981221256",
		NationalID:  "0017548952",
		IsActive:    true,
	}
	mockRepo := &MockCourierRepository{
		GetFunc: func(ctx context.Context, id int) (*models.Courier, error) {
			if id == 1 {
				return mockCourier, nil
			}
			return nil, errors.New("Courier not found")
		},
	}
	service := services.NewCourierService(mockRepo)

	courier, err := service.GetCourierByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, courier)
	assert.Equal(t, mockCourier.ID, courier.ID)
	assert.Equal(t, mockCourier.FullName, courier.FullName)
	assert.Equal(t, mockCourier.PhoneNumber, courier.PhoneNumber)

	courier, err = service.GetCourierByID(context.Background(), 2)
	assert.Error(t, err)
	assert.Nil(t, courier)
}

func TestCourierService_GetAllCouriers(t *testing.T) {
	mockCouriers := []models.Courier{
		{
			ID:          1,
			FullName:    "Name1",
			PhoneNumber: "9981221256",
			NationalID:  "0017548952",
			IsActive:    true,
		},
		{
			ID:          2,
			FullName:    "Name2",
			PhoneNumber: "9997253256",
			NationalID:  "2117559282",
			IsActive:    true,
		},
	}
	mockRepo := &MockCourierRepository{
		GetAllFunc: func(ctx context.Context) ([]models.Courier, error) {
			return mockCouriers, nil
		},
	}
	service := services.NewCourierService(mockRepo)

	couriers, err := service.GetAllCouriers(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, couriers)
	assert.Equal(t, len(mockCouriers), len(couriers))
	for i, mockCourier := range mockCouriers {
		assert.Equal(t, mockCourier.ID, couriers[i].ID)
		assert.Equal(t, mockCourier.PhoneNumber, couriers[i].PhoneNumber)
		assert.Equal(t, mockCourier.NationalID, couriers[i].NationalID)
	}
}
