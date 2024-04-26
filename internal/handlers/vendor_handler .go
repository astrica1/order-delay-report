package handlers

import (
	"net/http"
	"strconv"

	"github.com/astrica1/order-delay-report/internal/models"
	"github.com/astrica1/order-delay-report/internal/services"
	"github.com/astrica1/order-delay-report/pkg/messages"
	"github.com/gin-gonic/gin"
)

type VendorHandler struct {
	vendorService services.VendorService
}

func NewVendorHandler(vendorService services.VendorService) *VendorHandler {
	return &VendorHandler{
		vendorService: vendorService,
	}
}

func (h *VendorHandler) GetVendorByID(c *gin.Context) {
	vendorID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid order ID",
		})
		return
	}

	vendor, err := h.vendorService.GetVendorByID(c.Request.Context(), vendorID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": messages.DOES_NOT_EXISTS.AsError("Vendor").Error()})
		return
	}

	c.JSON(http.StatusOK, vendor)
}

func (h *VendorHandler) GetAllVendors(c *gin.Context) {
	var vendors []models.Vendor
	vendors, err := h.vendorService.GetAllVendors(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": messages.DOES_NOT_EXISTS.AsError("Vendors").Error()})
		return
	}

	c.JSON(http.StatusOK, vendors)
}
