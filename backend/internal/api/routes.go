package api

import "github.com/gin-gonic/gin"

type Handlers struct {
	ClientHandler    *ClientHandler
	DashboardHandler *DashboardHandler
}

func RegisterRoutes(
	router *gin.Engine,
	handlers Handlers,
) {

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.GET("/clients", handlers.ClientHandler.GetClients)

	router.GET("/clients/:id", handlers.ClientHandler.GetClientByID)

	router.GET("/dashboard/metrics", handlers.DashboardHandler.GetMetrics)
}
