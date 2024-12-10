package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Product ProductRepository
	Order   OrderRepository
}

func NewRepository(db *gorm.DB, log *zap.Logger) *Repository {
	return &Repository{
		Product: NewProductRepository(db, log),
		Order:   NewOrderRepository(db, log),
	}
}
