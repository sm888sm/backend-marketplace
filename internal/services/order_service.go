package services

import (
	"errors"
	"fmt"

	"github.com/sm888sm/backend-marketplace/internal/models"
	"github.com/sm888sm/backend-marketplace/internal/repositories"
	"gorm.io/gorm"
)

// Masih di hard code
const (
	DefaultShippingFee    = 5000.0
	FreeShippingThreshold = 15000.0
	DiscountThreshold     = 50000.0
	DiscountRate          = 0.10
)

type OrderService interface {
	CreateOrder(customerID uint, items []models.OrderItemRequest) (*models.Order, error)
	GetOrdersByCustomer(customerID uint) ([]models.Order, error)
	GetBuyersByMerchant(merchantID uint) ([]models.User, error)
	GetOrderByID(id uint) (*models.Order, error)
}

type orderService struct {
	repo     repositories.OrderRepository
	prodRepo repositories.ProductRepository
}

func NewOrderService(repo repositories.OrderRepository, prodRepo repositories.ProductRepository) OrderService {
	return &orderService{repo: repo, prodRepo: prodRepo}
}

func (s *orderService) CreateOrder(customerID uint, items []models.OrderItemRequest) (*models.Order, error) {
	if len(items) == 0 {
		return nil, errors.New("order harus memiliki minimal satu item")
	}

	productMap := make(map[uint]bool)
	for _, item := range items {
		if productMap[item.ProductID] {
			return nil, fmt.Errorf("produk dengan ID %d duplikat dalam order", item.ProductID)
		}
		productMap[item.ProductID] = true
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("quantity untuk produk ID %d harus lebih dari 0", item.ProductID)
		}
	}

	tx := s.repo.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var total float64
	var orderItems []models.OrderItem

	for _, req := range items {
		product, err := s.prodRepo.GetByIDTx(tx, req.ProductID)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("produk dengan ID %d tidak ditemukan", req.ProductID)
		}
		if product.Stock < req.Quantity {
			tx.Rollback()
			return nil, fmt.Errorf("stok produk %s kurang (tersisa %d)", product.Name, product.Stock)
		}
		if product.MerchantID == customerID {
			tx.Rollback()
			return nil, fmt.Errorf("tidak bisa membeli produk milik sendiri: %s", product.Name)
		}
		orderItems = append(orderItems, models.OrderItem{
			ProductID: req.ProductID,
			Quantity:  req.Quantity,
			Price:     product.Price,
		})
		total += product.Price * float64(req.Quantity)
	}

	shipping := DefaultShippingFee
	if total > FreeShippingThreshold {
		shipping = 0
	}
	discount := 0.0
	if total > DiscountThreshold {
		discount = total * DiscountRate
	}
	final := total - discount + shipping

	order := &models.Order{
		CustomerID:     customerID,
		TotalAmount:    total,
		DiscountAmount: discount,
		ShippingFee:    shipping,
		FinalAmount:    final,
		Status:         "pending",
		Items:          orderItems,
	}

	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, req := range items {
		if err := tx.Model(&models.Product{}).
			Where("id = ? AND stock >= ?", req.ProductID, req.Quantity).
			UpdateColumn("stock", gorm.Expr("stock - ?", req.Quantity)).Error; err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("gagal mengurangi stok produk ID %d", req.ProductID)
		}
	}

	tx.Commit()
	return order, nil
}

func (s *orderService) GetOrdersByCustomer(customerID uint) ([]models.Order, error) {
	return s.repo.GetOrdersByCustomer(customerID)
}

func (s *orderService) GetBuyersByMerchant(merchantID uint) ([]models.User, error) {
	return s.repo.GetBuyersByMerchant(merchantID)
}

func (s *orderService) GetOrderByID(id uint) (*models.Order, error) {
	return s.repo.GetOrderByID(id)
}
