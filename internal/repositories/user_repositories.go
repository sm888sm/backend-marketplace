package repositories

import (
	"github.com/sm888sm/backend-marketplace/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) ListUsers(page, perPage int, filter map[string]interface{}) ([]models.User, int64, error) {
	var users []models.User
	var total int64
	query := r.db.Model(&models.User{})
	for k, v := range filter {
		query = query.Where(k+" = ?", v)
	}
	query.Count(&total)
	if page > 0 && perPage > 0 {
		offset := (page - 1) * perPage
		query = query.Offset(offset).Limit(perPage)
	}
	if err := query.Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
