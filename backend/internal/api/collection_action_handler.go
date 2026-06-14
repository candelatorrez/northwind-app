package api

import (
	"net/http"

	"github.com/candelatorrez/northwind-app/internal/domain"
	"github.com/candelatorrez/northwind-app/internal/service"
	"github.com/gin-gonic/gin"
)

type CollectionActionHandler struct {
	actionService *service.CollectionActionService
}

func NewCollectionActionHandler(actionService *service.CollectionActionService) *CollectionActionHandler {
	return &CollectionActionHandler{
		actionService: actionService,
	}
}

type createCollectionActionRequest struct {
	Type        domain.ActionType `json:"type" binding:"required"`
	Notes       string            `json:"notes"`
	PerformedBy string            `json:"performed_by" binding:"required"`
}

func (h *CollectionActionHandler) CreateAction(c *gin.Context) {
	clientID, err := parseUintParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid client id"})
		return
	}

	var req createCollectionActionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	action, err := h.actionService.CreateAction(clientID, req.Type, req.Notes, req.PerformedBy)
	if err != nil {
		switch err {
		case service.ErrClientNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "client not found"})
		case service.ErrInvalidActionType:
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid action type"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, action)
}

func (h *CollectionActionHandler) GetActions(c *gin.Context) {
	clientID, err := parseUintParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid client id"})
		return
	}

	actions, err := h.actionService.GetActionsByClientID(clientID)
	if err != nil {
		switch err {
		case service.ErrClientNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "client not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, actions)
}
