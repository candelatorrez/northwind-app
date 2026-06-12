package api

import (
	"net/http"

	"github.com/candelatorrez/northwind-app/internal/service"
	"github.com/gin-gonic/gin"
)

type InvoiceHandler struct {
	invoiceService *service.InvoiceService
}

func NewInvoiceHandler(invoiceService *service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{
		invoiceService: invoiceService,
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

	c.JSON(http.StatusOK, invoice)
}
