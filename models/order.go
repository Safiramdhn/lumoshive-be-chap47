package models

import "time"

// Order represents the order table
// Order represents the order table
type Order struct {
	ID            uint        `json:"id" gorm:"primaryKey;autoIncrement"`
	TotalQuantity int         `json:"quantity" gorm:"not null"`
	Total         float64     `json:"total" gorm:"not null"`
	Status        string      `json:"status" gorm:"size:20;check:status IN ('pending','shipped','completed','canceled');default:'created'" binding:"required" example:"pending"`
	CreatedAt     time.Time   `json:"created_at" gorm:"autoCreateTime"`
	OrderItems    []OrderItem `json:"items"` // A slice of related order items
}

// OrderItem represents the order_item table
type OrderItem struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID     uint      `json:"order_id" gorm:"not null"`
	ProductID   uint      `json:"product_id" gorm:"not null"`
	ProductName string    `json:"product_name" gorm:"not null"`
	Quantity    int       `json:"quantity" gorm:"not null"`
	Price       float64   `json:"price" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func OrderSeeder() []Order {
	return []Order{
		{TotalQuantity: 2, Total: 59.98, Status: "pending", CreatedAt: time.Now()},
		{TotalQuantity: 1, Total: 99.99, Status: "shipped", CreatedAt: time.Now()},
	}
}

func OrderItemSeeder() []OrderItem {
	return []OrderItem{
		{OrderID: 1, ProductID: 1, ProductName: "Product A", Quantity: 1, Price: 29.99, CreatedAt: time.Now()},
		{OrderID: 1, ProductID: 1, ProductName: "Product A", Quantity: 1, Price: 29.99, CreatedAt: time.Now()},
		{OrderID: 2, ProductID: 3, ProductName: "Product C", Quantity: 1, Price: 99.99, CreatedAt: time.Now()},
	}
}
