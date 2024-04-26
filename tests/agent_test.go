package services

import (
	"context"
	"errors"
	"testing"

	"github.com/astrica1/order-delay-report/internal/models"
	"github.com/astrica1/order-delay-report/internal/services"
	"github.com/stretchr/testify/assert"
)

type MockAgentRepository struct {
	GetFunc           func(ctx context.Context, id int) (*models.Agent, error)
	GetByUsernameFunc func(ctx context.Context, username string) (*models.Agent, error)
	GetAllFunc        func(ctx context.Context) ([]models.Agent, error)
	CreateFunc        func(ctx context.Context, agent *models.Agent) error
	UpdateFunc        func(ctx context.Context, agent *models.Agent) error
	DeleteFunc        func(ctx context.Context, id int) error
}

func (m *MockAgentRepository) Get(ctx context.Context, id int) (*models.Agent, error) {
	return m.GetFunc(ctx, id)
}

func (m *MockAgentRepository) GetByUsername(ctx context.Context, username string) (*models.Agent, error) {
	return m.GetByUsernameFunc(ctx, username)
}

func (m *MockAgentRepository) GetAll(ctx context.Context) ([]models.Agent, error) {
	return m.GetAllFunc(ctx)
}

func (m *MockAgentRepository) Create(ctx context.Context, agent *models.Agent) error {
	return m.CreateFunc(ctx, agent)
}

func (m *MockAgentRepository) Update(ctx context.Context, agent *models.Agent) error {
	return m.UpdateFunc(ctx, agent)
}

func (m *MockAgentRepository) Delete(ctx context.Context, id int) error {
	return m.DeleteFunc(ctx, id)
}

func TestAgentService_GetAgentByID(t *testing.T) {
	mockAgent := &models.Agent{ID: 1, Username: "testuser"}
	mockRepo := &MockAgentRepository{
		GetFunc: func(ctx context.Context, id int) (*models.Agent, error) {
			if id == 1 {
				return mockAgent, nil
			}
			return nil, errors.New("Agent not found")
		},
	}
	service := services.NewAgentService(mockRepo)

	agent, err := service.GetAgentByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, agent)
	assert.Equal(t, mockAgent.ID, agent.ID)
	assert.Equal(t, mockAgent.Username, agent.Username)

	agent, err = service.GetAgentByID(context.Background(), 2)
	assert.Error(t, err)
	assert.Nil(t, agent)
}

func TestAgentService_GetAgentByUsername(t *testing.T) {
	mockAgent := &models.Agent{ID: 1, Username: "testuser"}
	mockRepo := &MockAgentRepository{
		GetByUsernameFunc: func(ctx context.Context, username string) (*models.Agent, error) {
			if username == "testuser" {
				return mockAgent, nil
			}
			return nil, errors.New("Agent not found")
		},
	}
	service := services.NewAgentService(mockRepo)

	agent, err := service.GetAgentByUsername(context.Background(), "testuser")
	assert.NoError(t, err)
	assert.NotNil(t, agent)
	assert.Equal(t, mockAgent.ID, agent.ID)
	assert.Equal(t, mockAgent.Username, agent.Username)

	agent, err = service.GetAgentByUsername(context.Background(), "nonexistent")
	assert.Error(t, err)
	assert.Nil(t, agent)
}

func TestAgentService_GetAllAgents(t *testing.T) {
	mockAgents := []models.Agent{
		{ID: 1, Username: "user1"},
		{ID: 2, Username: "user2"},
	}
	mockRepo := &MockAgentRepository{
		GetAllFunc: func(ctx context.Context) ([]models.Agent, error) {
			return mockAgents, nil
		},
	}
	service := services.NewAgentService(mockRepo)

	agents, err := service.GetAllAgents(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, agents)
	assert.Equal(t, len(mockAgents), len(agents))
	for i, mockAgent := range mockAgents {
		assert.Equal(t, mockAgent.ID, agents[i].ID)
		assert.Equal(t, mockAgent.Username, agents[i].Username)
	}
}
