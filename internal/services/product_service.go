package services

import (
	"errors"

	"github.com/sm888sm/backend-marketplace/internal/models"
	"github.com/sm888sm/backend-marketplace/internal/repositories"
)

type ProductService interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	UpdateProduct(product *models.Product) (*models.Product, error)
	DeleteProduct(productID, merchantID uint) error
	GetProductsByMerchant(merchantID uint) ([]models.Product, error)
	GetAllProducts() ([]models.Product, error)
	GetProductByID(productID uint) (*models.Product, error)
	UploadImage(productID, merchantID uint, path string) (*models.ProductImage, error)
	ListImages(productID uint) ([]models.ProductImage, error)
	SearchProducts(query string, minPrice, maxPrice float64, categoryID uint, sort string, page, perPage int) ([]models.Product, int64, error)
}

type productService struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (s *productService) CreateProduct(product *models.Product) (*models.Product, error) {
	if product.Name == "" {
		return nil, errors.New("product name is required")
	}
	if product.Price <= 0 {
		return nil, errors.New("product price must be positive")
	}
	if product.Stock < 0 {
		return nil, errors.New("product stock cannot be negative")
	}

	return s.productRepo.Create(product)
}

func (s *productService) UpdateProduct(product *models.Product) (*models.Product, error) {
	existingProduct, err := s.productRepo.GetByID(product.ID)
	if err != nil {
		return nil, err
	}

	if existingProduct.MerchantID != product.MerchantID {
		return nil, errors.New("unauthorized to update this product")
	}

	if product.Name == "" {
		return nil, errors.New("product name is required")
	}
	if product.Price <= 0 {
		return nil, errors.New("product price must be positive")
	}
	if product.Stock < 0 {
		return nil, errors.New("product stock cannot be negative")
	}

	return s.productRepo.Update(product)
}

func (s *productService) DeleteProduct(productID, merchantID uint) error {
	product, err := s.productRepo.GetByID(productID)
	if err != nil {
		return err
	}

	if product.MerchantID != merchantID {
		return errors.New("unauthorized to delete this product")
	}

	return s.productRepo.Delete(productID)
}

func (s *productService) GetProductsByMerchant(merchantID uint) ([]models.Product, error) {
	return s.productRepo.GetByMerchantID(merchantID)
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.productRepo.GetAll()
}

func (s *productService) GetProductByID(productID uint) (*models.Product, error) {
	return s.productRepo.GetByID(productID)
}

func (s *productService) UploadImage(productID, merchantID uint, path string) (*models.ProductImage, error) {
	product, err := s.productRepo.GetByID(productID)
	if err != nil {
		return nil, err
	}
	if product.MerchantID != merchantID {
		return nil, errors.New("hanya merchant pemilik produk yang bisa upload gambar")
	}
	img := &models.ProductImage{ProductID: productID, Path: path}
	if err := s.productRepo.AddImage(img); err != nil {
		return nil, err
	}
	return img, nil
}

func (s *productService) ListImages(productID uint) ([]models.ProductImage, error) {
	return s.productRepo.ListImages(productID)
}

func (s *productService) SearchProducts(query string, minPrice, maxPrice float64, categoryID uint, sort string, page, perPage int) ([]models.Product, int64, error) {
	return s.productRepo.SearchProducts(query, minPrice, maxPrice, categoryID, sort, page, perPage)
}
