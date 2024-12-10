package controller

import (
	"golang-chap47/models"
	"golang-chap47/service"
	"golang-chap47/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OrderController struct {
	service service.Service
	log     *zap.Logger
}

func NewOrderController(service service.Service, log *zap.Logger) *OrderController {
	return &OrderController{service, log}
}

func (c *OrderController) ExportOrderReports(filePath string) error {
	orders, err := c.service.Order.GetAllOrders()
	if err != nil {
		c.log.Error("Error getting orders", zap.Error(err))
		return err
	}
	return utils.ExportOrdersToExcel(orders, filePath)
}

func (ctrl *OrderController) CreateOrder(c *gin.Context) {
	var orderInput models.Order

	// Bind the request body to the Order model
	if err := c.ShouldBindJSON(&orderInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Call the service to create the order
	if err := ctrl.service.Order.CreateOrder(orderInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully"})
}
