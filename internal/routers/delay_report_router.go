package routers

import (
	"github.com/astrica1/order-delay-report/internal/handlers"
	"github.com/astrica1/order-delay-report/internal/repositories"
	"github.com/astrica1/order-delay-report/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupDelayReportRouter(router *gin.Engine, db *gorm.DB) {
	delayReportRepository := repositories.NewDelayReportRepository(db)
	delayReportService := services.NewDelayReportService(delayReportRepository)
	delayReportHandler := handlers.NewDelayReportHandler(delayReportService)

	delayReportGroup := router.Group("/reports")
	{
		delayReportGroup.GET("/weekly", delayReportHandler.GetWeeklyDelays)
	}
}
