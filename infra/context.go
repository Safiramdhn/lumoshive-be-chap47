package infra

import (
	"golang-chap47/config"
	"golang-chap47/controller"
	"golang-chap47/database"
	"golang-chap47/helper"
	"golang-chap47/repository"
	"golang-chap47/service"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Cfg    config.Configuration
	DB     *gorm.DB
	Ctl    controller.Controller
	Log    *zap.Logger
	Cacher database.Cacher
	Cron   *cron.Cron
	// Middleware *middleware.AuthMiddleware
}

func NewServiceContext() (*ServiceContext, error) {

	handlerError := func(err error) (*ServiceContext, error) {
		return nil, err
	}

	// instance config
	config, err := config.ReadConfig()
	if err != nil {
		handlerError(err)
	}

	// instance looger
	log, err := helper.InitZapLogger()
	if err != nil {
		handlerError(err)
	}

	// instance database
	db, err := database.InitDB(config)
	if err != nil {
		handlerError(err)
	}

	rdb := database.NewCacher(config, 60*60)

	// middleware := middleware.NewMiddleware(log, rdb)

	// instance repository
	repository := repository.NewRepository(db, log)

	// instance service
	service := service.NewService(*repository)

	// instance controller
	Ctl := controller.NewController(*service, log, rdb, config)

	// Initialize middleware
	// authMiddleware := middleware.NewAuthMiddleware(log, rdb, config.JwtSecret)

	// initialize cron
	cronScheduler := cron.New()

	// Return service context
	return &ServiceContext{
		Cfg:    config,
		DB:     db,
		Ctl:    *Ctl,
		Log:    log,
		Cacher: rdb,
		Cron:   cronScheduler,
		// Middleware: authMiddleware,
	}, nil
}
