package controller

import (
	"golang-chap47/config"
	"golang-chap47/database"
	"golang-chap47/service"

	"go.uber.org/zap"
)

type Controller struct {
	Product ProductController
	Order   OrderController
}

func NewController(service service.Service, log *zap.Logger, cacher database.Cacher, config config.Configuration) *Controller {
	return &Controller{
		Product: *NewProductController(service, log),
		Order:   *NewOrderController(service, log),
	}

}
