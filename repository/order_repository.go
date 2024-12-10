package repository

import (
	"golang-chap47/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderRepository interface {
	GetAll() ([]models.Order, error)
	Create(orderInput models.Order) error
}

type orderRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

// Create implements OrderRepository.
func (o *orderRepository) Create(orderInput models.Order) error {
	// Insert the order into the orders table
	return o.db.Transaction(func(tx *gorm.DB) error { // tx is *gorm.DB, not *gorm.Tx
		// Insert the order into the orders table
		if err := tx.Create(&orderInput).Error; err != nil {
			return err
		}

		// Ensure the OrderID is set before inserting order items
		// for i := range orderInput.OrderItems {
		// 	// Set the OrderID for each item (do not set ID, as it will be auto-incremented)
		// 	orderInput.OrderItems[i].OrderID = orderInput.ID

		// 	// Insert the order items associated with the order (if any)
		// 	if err := tx.Create(&orderInput.OrderItems[i]).Error; err != nil {
		// 		return err
		// 	}
		// }

		return nil
	})
}

// GetAll implements OrderRepository.
func (o *orderRepository) GetAll() ([]models.Order, error) {
	var orders []models.Order

	// Fetch all orders from the orders table
	err := o.db.Find(&orders).Error
	if err != nil {
		return nil, err
	}

	// Fetch order items for each order and manually associate them
	for i := range orders {
		var items []models.OrderItem
		// Fetch items for the current order
		err := o.db.Where("order_id = ?", orders[i].ID).Find(&items).Error
		if err != nil {
			return nil, err
		}
		// Assign the fetched items to the current order (this assumes you're using a slice of OrderItems in your Order model)
		orders[i].OrderItems = items
	}

	// Return the list of orders with their associated order items
	return orders, nil
}

func NewOrderRepository(db *gorm.DB, log *zap.Logger) OrderRepository {
	return &orderRepository{db: db, log: log}
}
