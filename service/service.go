package service

import "golang-chap47/repository"

type Service struct {
	Product ProductService
	Order   OrderService
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		Product: NewProductService(repo),
		Order:   NewOrderService(repo),
	}
}
