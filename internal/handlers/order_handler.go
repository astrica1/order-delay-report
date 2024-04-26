package handlers

import (
	"net/http"
	"strconv"

	"github.com/astrica1/order-delay-report/internal/models"
	"github.com/astrica1/order-delay-report/internal/services"
	"github.com/astrica1/order-delay-report/pkg/messages"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService services.OrderService
}

func NewOrderHandler(orderService services.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid order ID",
		})
		return
	}

	order, err := h.orderService.GetOrderByID(c.Request.Context(), orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": messages.DOES_NOT_EXISTS.AsError("Order").Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	var orders []models.Order
	orders, err := h.orderService.GetAllOrders(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": messages.DOES_NOT_EXISTS.AsError("Order").Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) ReportDelayForOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid order ID",
		})
		return
	}

	err = h.orderService.CreateDelayReportForOrder(c.Request.Context(), orderID)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delay reported successfully",
	})
}
