package api

import (
	"net/http"

	"github.com/candelatorrez/northwind-app/internal/service"
	"github.com/gin-gonic/gin"
)

type InvoiceHandler struct {
	invoiceService *service.InvoiceService
	riskService    *service.RiskService
}

func NewInvoiceHandler(invoiceService *service.InvoiceService, riskService *service.RiskService) *InvoiceHandler {
	return &InvoiceHandler{
		invoiceService: invoiceService,
		riskService:    riskService,
	}
}

func (h *InvoiceHandler) MarkAsPaid(c *gin.Context) {
	invoiceID, err := parseUintParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid invoice id"})
		return
	}

	invoice, err := h.invoiceService.MarkAsPaid(invoiceID)
	if err != nil {
		switch err {
		case service.ErrInvoiceNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "invoice not found"})
		case service.ErrInvoiceAlreadyPaid:
			c.JSON(http.StatusConflict, gin.H{"error": "invoice already paid"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Recalculate risk snapshot for client after marking invoice as paid.
	go func() {
		if invoice != nil {
			_, _ = h.riskService.SnapshotClient(invoice.ClientID)
		}
	}()

	c.JSON(http.StatusOK, invoice)
}

func (h *InvoiceHandler) GetInvoicesByClientID(c *gin.Context) {
	clientID, err := parseUintParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid client id"})
		return
	}

	invoices, err := h.invoiceService.GetInvoicesByClientID(clientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, invoices)
}
