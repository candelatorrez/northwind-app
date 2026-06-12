package api

import "github.com/gin-gonic/gin"

type Handlers struct {
	ClientHandler           *ClientHandler
	DashboardHandler        *DashboardHandler
	CollectionActionHandler *CollectionActionHandler
	InvoiceHandler          *InvoiceHandler
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

	router.PATCH("/clients/:id/status", handlers.ClientHandler.UpdateClientStatus)

	router.POST("/clients/:id/actions", handlers.CollectionActionHandler.CreateAction)

	router.POST("/invoices/:id/pay", handlers.InvoiceHandler.MarkAsPaid)

	router.GET("/dashboard/metrics", handlers.DashboardHandler.GetMetrics)
}
