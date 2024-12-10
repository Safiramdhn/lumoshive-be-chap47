package service

import (
	"errors"
	"golang-chap47/models"
	"golang-chap47/repository"
)

type OrderService interface {
	GetAllOrders() ([]models.Order, error)
	CreateOrder(orderInput models.Order) error
}

type orderService struct {
	repo repository.Repository
}

func (s *orderService) CreateOrder(orderInput models.Order) error {
	// Validate order input (you can add more validation logic as needed)
	if orderInput.TotalQuantity <= 0 || orderInput.Total <= 0 {
		return errors.New("invalid order input")
	}

	// Call the repository to create the order and order items
	if err := s.repo.Order.Create(orderInput); err != nil {
		return err
	}

	return nil
}

// GetAllOrders implements OrderService.
func (o *orderService) GetAllOrders() ([]models.Order, error) {
	return o.repo.Order.GetAll()
}

func NewOrderService(repo repository.Repository) OrderService {
	return &orderService{repo}
}
