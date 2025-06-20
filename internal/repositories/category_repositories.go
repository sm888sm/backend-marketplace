package repositories

import (
	"github.com/sm888sm/backend-marketplace/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(cat *models.Category) error
	List() []models.Category
	GetByID(id string) (*models.Category, error)
	GetByIDUint(id uint) (*models.Category, error)
	Update(cat *models.Category) error
	Delete(id string) error
	ExistsByName(name string) bool
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(cat *models.Category) error {
	return r.db.Create(cat).Error
}

func (r *categoryRepository) List() []models.Category {
	var cats []models.Category
	r.db.Find(&cats)
	return cats
}

func (r *categoryRepository) GetByID(id string) (*models.Category, error) {
	var cat models.Category
	if err := r.db.First(&cat, id).Error; err != nil {
		return nil, err
	}
	return &cat, nil
}

func (r *categoryRepository) GetByIDUint(id uint) (*models.Category, error) {
	var cat models.Category
	if err := r.db.First(&cat, id).Error; err != nil {
		return nil, err
	}
	return &cat, nil
}

func (r *categoryRepository) Update(cat *models.Category) error {
	return r.db.Save(cat).Error
}

func (r *categoryRepository) Delete(id string) error {
	result := r.db.Delete(&models.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *categoryRepository) ExistsByName(name string) bool {
	var count int64
	r.db.Model(&models.Category{}).Where("name = ?", name).Count(&count)
	return count > 0
}
