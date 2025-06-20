package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Role     string `gorm:"type:enum('merchant','customer');not null"`
}

type Product struct {
	gorm.Model
	MerchantID  uint   `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string
	Price       float64        `gorm:"type:decimal(10,2);not null"`
	Stock       int            `gorm:"not null"`
	CategoryID  uint           `json:"category_id"`
	Category    Category       `gorm:"foreignKey:CategoryID"`
	Images      []ProductImage `gorm:"foreignKey:ProductID"`
}

type Order struct {
	gorm.Model
	CustomerID     uint    `gorm:"not null"`
	TotalAmount    float64 `gorm:"type:decimal(10,2);not null"`
	DiscountAmount float64 `gorm:"type:decimal(10,2);default:0"`
	ShippingFee    float64 `gorm:"type:decimal(10,2);default:0"`
	FinalAmount    float64 `gorm:"type:decimal(10,2);not null"`
	Status         string  `gorm:"type:enum('pending','completed','cancelled');default:'pending'"`
	Items          []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"type:decimal(10,2);not null"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Role     string `json:"role" binding:"required,oneof=merchant customer"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	CategoryID  uint    `json:"category_id" binding:"required"`
}

type OrderItemRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique;not null" json:"name"`
	// Bisa ditambah field lain seperti description
	Products []Product `gorm:"foreignKey:CategoryID"`
}

type ProductImage struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Path      string  `gorm:"not null" json:"path"`
	Product   Product `gorm:"foreignKey:ProductID"`
}
