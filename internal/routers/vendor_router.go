package routers

import (
	"github.com/astrica1/order-delay-report/internal/handlers"
	"github.com/astrica1/order-delay-report/internal/repositories"
	"github.com/astrica1/order-delay-report/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupVendorRouter(router *gin.Engine, db *gorm.DB) {
	delayReportRepository := repositories.NewDelayReportRepository(db)
	delayReportService := services.NewDelayReportService(delayReportRepository)

	vendorRepository := repositories.NewVendorRepository(db)
	vendorService := services.NewVendorService(vendorRepository, delayReportService)
	vendorHandler := handlers.NewVendorHandler(vendorService)

	orderGroup := router.Group("/order")
	{
		orderGroup.GET("/list", vendorHandler.GetAllVendors)
		orderGroup.GET("/:id", vendorHandler.GetVendorByID)

		orderGroup.GET("/:id/report-delay")  // Get report status
		orderGroup.POST("/:id/report-delay") // Report delay for the order
	}
}
