package routes

import (
	"restaurant-management/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controllers.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controllers.GetInvoice())
	incomingRoutes.POST("/invoices", controllers.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:Invoice_id", controllers.UpdateInvoice())
	incomingRoutes.DELETE("/invoices/:Invoice_id", controllers.DeleteInvoice())

}