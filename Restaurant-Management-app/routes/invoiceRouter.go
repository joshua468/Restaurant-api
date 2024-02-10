package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/joshua468/restaurant-management-app/controllers"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controller.GetInvoice())
	incomingRoutes.POST("/invoices", controller.CreateInvoice())
	incomingRoutes.PATCH("/invoices/invoice_id", controller.UpdateInvoice())
}
