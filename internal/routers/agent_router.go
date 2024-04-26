package routers

import (
	"github.com/astrica1/order-delay-report/internal/handlers"
	"github.com/astrica1/order-delay-report/internal/repositories"
	"github.com/astrica1/order-delay-report/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAgentRouter(router *gin.Engine, db *gorm.DB) {
	delayReportRepository := repositories.NewDelayReportRepository(db)

	agentRepository := repositories.NewAgentRepository(db)
	agentService := services.NewAgentService(agentRepository)
	agentHandler := handlers.NewAgentHandler(agentService, delayReportRepository)

	agentGroup := router.Group("/agent")
	{
		agentGroup.GET("/list", agentHandler.GetAllAgents)
		agentGroup.GET("/:id", agentHandler.GetAgentByID)
		agentGroup.GET("/username/:username", agentHandler.GetAgentByUsername)

		agentGroup.POST("/new", agentHandler.CreateNewAgent)
		agentGroup.POST("/:id/update", agentHandler.UpdateExistingAgent)

		agentGroup.DELETE("/:id/delete", agentHandler.DeleteExistingAgent)

		agentGroup.POST("/:id/report/pop", agentHandler.CreateNewAgent)
	}
}
