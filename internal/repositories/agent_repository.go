package repositories

import (
	"context"

	"github.com/astrica1/order-delay-report/internal/models"
	"gorm.io/gorm"
)

type AgentRepository interface {
	Get(ctx context.Context, id int) (*models.Agent, error)
	GetByUsername(ctx context.Context, username string) (*models.Agent, error)
	GetAll(ctx context.Context) ([]models.Agent, error)
	Create(ctx context.Context, agent *models.Agent) error
	Update(ctx context.Context, agent *models.Agent) error
	Delete(ctx context.Context, id int) error
}

type agentRepository struct {
	db *gorm.DB
}

func NewAgentRepository(db *gorm.DB) AgentRepository {
	return &agentRepository{
		db: db,
	}
}

func (r *agentRepository) Get(ctx context.Context, id int) (*models.Agent, error) {
	var agent models.Agent
	if err := r.db.WithContext(ctx).First(&agent, id).Error; err != nil {
		return nil, err
	}

	return &agent, nil
}

func (r *agentRepository) GetByUsername(ctx context.Context, username string) (*models.Agent, error) {
	var agent models.Agent
	if err := r.db.WithContext(ctx).First(&agent, username).Error; err != nil {
		return nil, err
	}

	return &agent, nil
}

func (r *agentRepository) GetAll(ctx context.Context) ([]models.Agent, error) {
	var agents []models.Agent
	if err := r.db.WithContext(ctx).Find(&agents).Error; err != nil {
		return nil, err
	}

	return agents, nil
}

func (r *agentRepository) Create(ctx context.Context, agent *models.Agent) error {
	if err := r.db.WithContext(ctx).Create(agent).Error; err != nil {
		return err
	}

	return nil
}

func (r *agentRepository) Update(ctx context.Context, agent *models.Agent) error {
	if err := r.db.WithContext(ctx).Save(agent).Error; err != nil {
		return err
	}

	return nil
}

func (r *agentRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&models.Agent{}, id).Error; err != nil {
		return err
	}

	return nil
}
