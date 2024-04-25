package repositories

import (
	"context"

	"github.com/astrica1/order-delay-report/internal/models"
	"gorm.io/gorm"
)

type AgentRepository struct {
	db *gorm.DB
}

func NewAgentRepository(db *gorm.DB) *AgentRepository {
	return &AgentRepository{
		db: db,
	}
}

func (r *AgentRepository) Get(ctx context.Context, id int) (*models.Agent, error) {
	var agent models.Agent
	if err := r.db.First(&agent, id).Error; err != nil {
		return nil, err
	}

	return &agent, nil
}

func (r *AgentRepository) GetAll(ctx context.Context) ([]models.Agent, error) {
	var agents []models.Agent
	if err := r.db.Find(&agents).Error; err != nil {
		return nil, err
	}

	return agents, nil
}

func (r *AgentRepository) Create(ctx context.Context, agent *models.Agent) error {
	if err := r.db.Create(agent).Error; err != nil {
		return err
	}

	return nil
}

func (r *AgentRepository) Update(ctx context.Context, agent *models.Agent) error {
	if err := r.db.Save(agent).Error; err != nil {
		return err
	}

	return nil
}

func (r *AgentRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.Delete(&models.Agent{}, id).Error; err != nil {
		return err
	}

	return nil
}
