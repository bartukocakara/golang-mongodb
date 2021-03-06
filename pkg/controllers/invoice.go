package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"restaurant-management/pkg/database"
	"restaurant-management/pkg/models"
	"time"

	"restaurant-management/pkg/httputil"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InvoiceViewFormat struct {
	Invoice_id string
	Payment_method string
	Order_id string
	Payment_status *string
	Payment_due interface{}
	Table_number interface{}
	Payment_due_date time.Time
	Order_details interface{}
}

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := invoiceCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while listing invoice items"})
		}

		var allInvoices []bson.M
		if err = result.All(ctx, &allInvoices); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allInvoices)
	}
}


// ShowInvoice godoc
// @Summary Show an Invoice
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path int true "Invoice ID"
// @Success 200 {object} models.Invoice
// @Router /invoice/{invoice_id} [get]
func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		invoiceId := c.Param("invoice_id")

		var invoice models.Invoice

		err := invoiceCollection.FindOne(ctx, bson.M{"invoice_id": invoiceId}).Decode(&invoice)

		defer cancel()

		if err != nil {
			httputil.NewError(c, http.StatusNotFound, err)
			return
		}

		var invoiceView InvoiceViewFormat

		allOrderItems, err := ItemsByOrder(invoice.Order_id)
		invoiceView.Order_id = invoice.Order_id
		invoice.Payment_due_date = invoice.Payment_due_date

		invoiceView.Payment_method = "null"
		if invoice.Payment_method != nil {
			invoiceView.Payment_method = *invoice.Payment_method
		}

		invoiceView.Invoice_id = invoice.Invoice_id
		invoiceView.Payment_status = *&invoice.Payment_status
		invoiceView.Payment_due = allOrderItems[0]["payment_due"]
		invoiceView.Table_number = allOrderItems[0]["table_number"]
		invoiceView.Order_details = allOrderItems[0]["order_items"]
		
		c.JSON(http.StatusOK, invoiceView)
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var invoice models.Invoice

		if err := c.BindJSON(&invoice); err != nil {
			httputil.NewError(c, http.StatusBadRequest, err)
			return
		}

		var order models.Order

		err := orderCollection.FindOne(ctx, bson.M{"order_id": invoice.Order_id}).Decode(&order)

		defer cancel()

		if err != nil {
			httputil.NewError(c, http.StatusBadRequest, err)
		}

		status := "PENDING"

		if invoice.Payment_status == nil {
			invoice.Payment_status = &status
		}

		invoice.Payment_due_date, _ = time.Parse(time.RFC3339, time.Now().AddDate(0, 0, 1).Format(time.RFC3339))
		invoice.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		invoice.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		invoice.ID = primitive.NewObjectID()
		invoice.Invoice_id = invoice.ID.Hex()

		validationErr := validate.Struct(invoice)
		if validationErr != nil {
			httputil.NewError(c, http.StatusBadRequest, err)
			return
		}

		result, insertErr := invoiceCollection.InsertOne(ctx, invoice)

		if insertErr != nil {
			msg := fmt.Sprintf("invoice item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		}

		defer cancel()

		c.JSON(http.StatusOK, result)

	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func DeleteInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}