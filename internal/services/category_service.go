package services

import (
	"errors"

	"github.com/sm888sm/backend-marketplace/internal/models"
	"github.com/sm888sm/backend-marketplace/internal/repositories"
)

type CategoryService interface {
	CreateCategory(name string) (*models.Category, error)
	ListCategories() ([]models.Category, error)
	UpdateCategory(id string, name string) (*models.Category, error)
	DeleteCategory(id string) error
	GetByIDUint(id uint) (*models.Category, error)
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) CreateCategory(name string) (*models.Category, error) {
	if s.repo.ExistsByName(name) {
		return nil, errors.New("nama kategori sudah digunakan")
	}
	cat := &models.Category{Name: name}
	if err := s.repo.Create(cat); err != nil {
		return nil, err
	}
	return cat, nil
}

func (s *categoryService) ListCategories() ([]models.Category, error) {
	return s.repo.List(), nil
}

func (s *categoryService) UpdateCategory(id string, name string) (*models.Category, error) {
	cat, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if s.repo.ExistsByName(name) && cat.Name != name {
		return nil, errors.New("nama kategori sudah digunakan")
	}
	cat.Name = name
	if err := s.repo.Update(cat); err != nil {
		return nil, err
	}
	return cat, nil
}

func (s *categoryService) DeleteCategory(id string) error {
	return s.repo.Delete(id)
}

func (s *categoryService) GetByIDUint(id uint) (*models.Category, error) {
	return s.repo.GetByIDUint(id)
}
