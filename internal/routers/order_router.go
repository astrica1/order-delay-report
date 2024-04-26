package routers

import (
	"github.com/astrica1/order-delay-report/internal/handlers"
	"github.com/astrica1/order-delay-report/internal/repositories"
	"github.com/astrica1/order-delay-report/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupOrderRouter(router *gin.Engine, db *gorm.DB) {
	delayReportRepository := repositories.NewDelayReportRepository(db)
	delayReportService := services.NewDelayReportService(delayReportRepository)

	tripRepository := repositories.NewTripRepository(db)
	tripService := services.NewTripService(tripRepository)

	orderRepository := repositories.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepository, delayReportService, tripService)
	orderHandler := handlers.NewOrderHandler(orderService)

	orderGroup := router.Group("/order")
	{
		orderGroup.GET("/list", orderHandler.GetAllOrders)
		orderGroup.GET("/:id", orderHandler.GetOrderByID)

		orderGroup.POST("/:id/report-delay", orderHandler.ReportDelayForOrder) // Report delay for the order
	}
}
