package api

import (
	"net/http"
	"strconv"

	"github.com/candelatorrez/northwind-app/internal/service"
	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	clientService *service.ClientService
}

func NewClientHandler(clientService *service.ClientService) *ClientHandler {
	return &ClientHandler{
		clientService: clientService,
	}
}

func (h *ClientHandler) GetClients(c *gin.Context) {
	clients, err := h.clientService.GetAllClients()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, clients)
}

func (h *ClientHandler) GetClientByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid client id",
		})

		return
	}

	client, err := h.clientService.GetClientByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "client not found",
		})
		return
	}

	c.JSON(http.StatusOK, client)
}
