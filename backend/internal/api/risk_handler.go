package api

import (
	"net/http"

	"github.com/candelatorrez/northwind-app/internal/service"
	"github.com/gin-gonic/gin"
)

type RiskHandler struct {
	riskService *service.RiskService
}

func NewRiskHandler(riskService *service.RiskService) *RiskHandler {
	return &RiskHandler{
		riskService: riskService,
	}
}

func (h *RiskHandler) GetLatestSnapshot(c *gin.Context) {
	clientID, err := parseUintParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid client id"})
		return
	}

	snapshot, err := h.riskService.GetLatestSnapshot(clientID)
	if err != nil {
		switch err {
		case service.ErrRiskSnapshotNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "risk snapshot not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, snapshot)
}

func (h *RiskHandler) CreateSnapshot(c *gin.Context) {
	clientID, err := parseUintParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid client id"})
		return
	}

	snapshot, err := h.riskService.SnapshotClient(clientID)
	if err != nil {
		switch err {
		case service.ErrClientNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "client not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, snapshot)
}
