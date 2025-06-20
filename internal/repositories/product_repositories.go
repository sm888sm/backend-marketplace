package repositories

import (
	"errors"

	"github.com/sm888sm/backend-marketplace/internal/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) (*models.Product, error)
	Update(product *models.Product) (*models.Product, error)
	Delete(productID uint) error
	GetByID(productID uint) (*models.Product, error)
	GetByMerchantID(merchantID uint) ([]models.Product, error)
	GetAll() ([]models.Product, error)
	GetByIDTx(tx *gorm.DB, productID uint) (*models.Product, error)
	AddImage(image *models.ProductImage) error
	ListImages(productID uint) ([]models.ProductImage, error)
	SearchProducts(query string, minPrice, maxPrice float64, categoryID uint, sort string, page, perPage int) ([]models.Product, int64, error)
	GetCategoryByID(categoryID uint) (*models.Category, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *models.Product) (*models.Product, error) {
	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) Update(product *models.Product) (*models.Product, error) {
	result := r.db.Model(product).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (r *productRepository) Delete(productID uint) error {
	result := r.db.Delete(&models.Product{}, productID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}
	return nil
}

func (r *productRepository) GetByID(productID uint) (*models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, productID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) GetByMerchantID(merchantID uint) ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Where("merchant_id = ?", merchantID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) GetByIDTx(tx *gorm.DB, productID uint) (*models.Product, error) {
	var product models.Product
	if err := tx.First(&product, productID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) AddImage(image *models.ProductImage) error {
	return r.db.Create(image).Error
}

func (r *productRepository) ListImages(productID uint) ([]models.ProductImage, error) {
	var images []models.ProductImage
	if err := r.db.Where("product_id = ?", productID).Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}

func (r *productRepository) SearchProducts(query string, minPrice, maxPrice float64, categoryID uint, sort string, page, perPage int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64
	q := r.db.Model(&models.Product{})
	if query != "" {
		q = q.Where("name LIKE ?", "%"+query+"%")
	}
	if minPrice > 0 {
		q = q.Where("price >= ?", minPrice)
	}
	if maxPrice > 0 {
		q = q.Where("price <= ?", maxPrice)
	}
	if categoryID > 0 {
		q = q.Where("category_id = ?", categoryID)
	}
	q.Count(&total)
	if sort == "price_asc" {
		q = q.Order("price ASC")
	} else if sort == "price_desc" {
		q = q.Order("price DESC")
	}
	if page > 0 && perPage > 0 {
		offset := (page - 1) * perPage
		q = q.Offset(offset).Limit(perPage)
	}
	if err := q.Find(&products).Error; err != nil {
		return nil, 0, err
	}
	return products, total, nil
}

func (r *productRepository) GetCategoryByID(categoryID uint) (*models.Category, error) {
	var category models.Category
	if err := r.db.First(&category, categoryID).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
