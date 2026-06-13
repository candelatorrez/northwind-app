package api

import (
	"net/http"

	"github.com/candelatorrez/northwind-app/internal/service"
	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	dashboardService *service.DashboardService
}

func NewDashboardHandler(dashboardService *service.DashboardService) *DashboardHandler {
	return &DashboardHandler{
		dashboardService: dashboardService,
	}
}

func (h *DashboardHandler) GetMetrics(c *gin.Context) {
	metrics, err := h.dashboardService.GetMetrics()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, metrics)
}

func (h *DashboardHandler) GetClientsOverview(c *gin.Context) {

	clients, err := h.dashboardService.GetClientsOverview()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, clients)
}
