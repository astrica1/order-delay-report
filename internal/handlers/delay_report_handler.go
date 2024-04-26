package handlers

import (
	"net/http"

	"github.com/astrica1/order-delay-report/internal/services"
	"github.com/astrica1/order-delay-report/pkg/messages"
	"github.com/gin-gonic/gin"
)

type DelayReportHandler struct {
	delayReportService services.DelayReportService
}

func NewDelayReportHandler(delayReportService services.DelayReportService) *DelayReportHandler {
	return &DelayReportHandler{
		delayReportService: delayReportService,
	}
}

func (h *DelayReportHandler) GetWeeklyDelays(c *gin.Context) {
	vendors, err := h.delayReportService.GetWeeklyDelayReport(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": messages.DOES_NOT_EXISTS.AsError("Vendors").Error()})
		return
	}

	c.JSON(http.StatusOK, vendors)
}
