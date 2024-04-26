package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/astrica1/order-delay-report/internal/models"
	"github.com/astrica1/order-delay-report/internal/repositories"
	"github.com/astrica1/order-delay-report/internal/services"
	"github.com/astrica1/order-delay-report/pkg/messages"
	"github.com/astrica1/order-delay-report/pkg/validator"
	"github.com/gin-gonic/gin"
)

type AgentHandler struct {
	agentService          services.AgentService
	delayReportRepository repositories.DelayReportRepository
}

func NewAgentHandler(agentService services.AgentService, delayReportRepository repositories.DelayReportRepository) *AgentHandler {
	return &AgentHandler{
		agentService:          agentService,
		delayReportRepository: delayReportRepository,
	}
}

func (h *AgentHandler) GetAgentByID(c *gin.Context) {
	agentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": messages.IS_INVALID.AsError("agent ID").Error(),
		})
		return
	}

	agent, err := h.agentService.GetAgentByID(c.Request.Context(), agentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": messages.DOES_NOT_EXISTS.AsError("agent").Error()})
		return
	}

	c.JSON(http.StatusOK, agent)
}

func (h *AgentHandler) GetAgentByUsername(c *gin.Context) {
	agentUsername := c.Param("username")

	agent, err := h.agentService.GetAgentByUsername(c.Request.Context(), agentUsername)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get agent"})
		return
	}

	c.JSON(http.StatusOK, agent)
}

func (h *AgentHandler) GetAllAgents(c *gin.Context) {
	agents, err := h.agentService.GetAllAgents(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get agents list"})
		return
	}

	type maskedResponse struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		IsActive bool   `json:"is_active"`
	}

	var agentResponses []maskedResponse
	for _, agent := range agents {
		agentResponse := maskedResponse{
			ID:       agent.ID,
			Username: agent.Username,
			IsActive: agent.IsActive,
		}

		agentResponses = append(agentResponses, agentResponse)
	}

	c.JSON(http.StatusOK, agentResponses)
}

func (h *AgentHandler) CreateNewAgent(c *gin.Context) {
	var agent models.Agent
	if err := c.ShouldBindJSON(&agent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	agents, err := h.agentService.GetAllAgents(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update agent right now"})
		return
	}

	for _, ag := range agents {
		if ag.Username == agent.Username {
			c.JSON(http.StatusInternalServerError, gin.H{"error": messages.ALREADY_EXISTS.AsError("Username").Error()})
			return
		}
	}

	if err := validator.ValidateUsername(agent.Username); err != nil {
		log.Printf("error creating agent: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.agentService.CreateNewAgent(c.Request.Context(), &agent); err != nil {
		log.Printf("error creating agent: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create agent"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Agent created successfully", "agent": agent})
}

func (h *AgentHandler) UpdateExistingAgent(c *gin.Context) {
	agentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid agent ID",
		})
		return
	}

	var newAgent models.Agent
	if err := c.ShouldBindJSON(&newAgent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	agents, err := h.agentService.GetAllAgents(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update agent right now"})
		return
	}

	var agent *models.Agent = nil
	for _, ag := range agents {
		if ag.ID == agentID {
			agent = &ag
			continue
		}

		if ag.Username == newAgent.Username {
			c.JSON(http.StatusInternalServerError, gin.H{"error": messages.ALREADY_EXISTS.AsError("Username").Error()})
			return
		}
	}

	if agent == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": messages.DOES_NOT_EXISTS.AsError("Agent").Error()})
		return
	}

	if err := validator.ValidateUsername(agent.Username); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.agentService.UpdateExistingAgent(c.Request.Context(), agentID, agent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Agent updated successfully", "agent": agent})
}

func (h *AgentHandler) DeleteExistingAgent(c *gin.Context) {
	agentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid agent ID",
		})
		return
	}

	_, err = h.agentService.GetAgentByID(c.Request.Context(), agentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": messages.DOES_NOT_EXISTS.AsError("Agent").Error()})
		return
	}

	if err := h.agentService.DeleteExistingAgent(c.Request.Context(), agentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete agent"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Agent deleted successfully"})
}

func (h *AgentHandler) GetNewReport(c *gin.Context) {
	agentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": messages.IS_INVALID.AsError("agent ID").Error(),
		})
		return
	}

	agent, err := h.delayReportRepository.PopReport(c.Request.Context(), agentID)
	if err != nil && agent == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err != nil && agent != nil {
		c.JSON(http.StatusBadRequest, agent)
		return
	}

	c.JSON(http.StatusOK, agent)
}
