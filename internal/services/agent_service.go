package services

import (
	"context"

	"github.com/astrica1/order-delay-report/internal/models"
	"github.com/astrica1/order-delay-report/internal/repositories"
	"github.com/astrica1/order-delay-report/pkg/messages"
)

type AgentService interface {
	GetAgentByID(ctx context.Context, id int) (*models.Agent, error)
	GetAgentByUsername(ctx context.Context, username string) (*models.Agent, error)
	GetAllAgents(ctx context.Context) ([]models.Agent, error)
	CreateNewAgent(ctx context.Context, agent *models.Agent) error
	UpdateExistingAgent(ctx context.Context, id int, agent *models.Agent) error
	DeleteExistingAgent(ctx context.Context, id int) error
}

type agentService struct {
	agentRepository repositories.AgentRepository
}

func NewAgentService(agentRepository repositories.AgentRepository) AgentService {
	return &agentService{
		agentRepository: agentRepository,
	}
}

func (s *agentService) GetAgentByID(ctx context.Context, id int) (*models.Agent, error) {
	return s.agentRepository.Get(ctx, id)
}

func (s *agentService) GetAgentByUsername(ctx context.Context, username string) (*models.Agent, error) {
	return s.agentRepository.GetByUsername(ctx, username)
}

func (s *agentService) GetAllAgents(ctx context.Context) ([]models.Agent, error) {
	return s.agentRepository.GetAll(ctx)
}

func (s *agentService) CreateNewAgent(ctx context.Context, agent *models.Agent) error {
	agents, err := s.GetAllAgents(ctx)
	if err != nil {
		return err
	}

	for _, ag := range agents {
		if ag.Username == agent.Username {
			return messages.ALREADY_EXISTS.AsError("Username")
		}
	}

	return s.agentRepository.Create(ctx, agent)
}

func (s *agentService) UpdateExistingAgent(ctx context.Context, id int, agent *models.Agent) error {
	agents, err := s.GetAllAgents(ctx)
	if err != nil {
		return err
	}

	for _, ag := range agents {
		if ag.Username == agent.Username && ag.ID != agent.ID {
			return messages.ALREADY_EXISTS.AsError("Username")
		}
	}

	return s.agentRepository.Update(ctx, agent)
}

func (s *agentService) DeleteExistingAgent(ctx context.Context, id int) error {
	return s.agentRepository.Delete(ctx, id)
}
