package routes

import (
	"golang-chap47/infra"
	// "golang-chap47/middleware"

	"github.com/gin-gonic/gin"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {
	router := gin.Default()

	controller := ctx.Ctl

	router.GET("/products", controller.Product.GetAllProductsController)

	router.POST("/orders", controller.Order.CreateOrder)

	return router
}
