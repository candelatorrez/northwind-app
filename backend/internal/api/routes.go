package api

import "github.com/gin-gonic/gin"

type Handlers struct {
	ClientHandler           *ClientHandler
	DashboardHandler        *DashboardHandler
	CollectionActionHandler *CollectionActionHandler
	InvoiceHandler          *InvoiceHandler
	RiskHandler             *RiskHandler
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

	router.GET("/clients/:id/actions", handlers.CollectionActionHandler.GetActions)

	router.POST("/clients/:id/actions", handlers.CollectionActionHandler.CreateAction)

	router.GET("/clients/:id/risk", handlers.RiskHandler.GetLatestSnapshot)

	router.POST("/clients/:id/risk-snapshots", handlers.RiskHandler.CreateSnapshot)

	router.POST("/invoices/:id/pay", handlers.InvoiceHandler.MarkAsPaid)

	router.GET("/dashboard/metrics", handlers.DashboardHandler.GetMetrics)

	router.GET("/dashboard/clients", handlers.DashboardHandler.GetClientsOverview)
}
