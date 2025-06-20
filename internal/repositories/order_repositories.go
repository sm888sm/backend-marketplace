package repositories

import (
	"github.com/sm888sm/backend-marketplace/internal/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) error
	GetOrdersByCustomer(customerID uint) ([]models.Order, error)
	GetBuyersByMerchant(merchantID uint) ([]models.User, error)
	GetOrderByID(id uint) (*models.Order, error)
	GetDB() *gorm.DB
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) GetOrdersByCustomer(customerID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Items").Where("customer_id = ?", customerID).Find(&orders).Error
	return orders, err
}

func (r *orderRepository) GetBuyersByMerchant(merchantID uint) ([]models.User, error) {
	var users []models.User
	r.db.Raw(`SELECT DISTINCT u.* FROM users u JOIN orders o ON u.id = o.customer_id JOIN order_items oi ON o.id = oi.order_id JOIN products p ON oi.product_id = p.id WHERE p.merchant_id = ?`, merchantID).Scan(&users)
	return users, nil
}

func (r *orderRepository) GetOrderByID(id uint) (*models.Order, error) {
	var order models.Order
	if err := r.db.Preload("Items").First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) GetDB() *gorm.DB {
	return r.db
}
