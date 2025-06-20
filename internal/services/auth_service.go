package services

import (
	"errors"

	"github.com/sm888sm/backend-marketplace/internal/models"
	"github.com/sm888sm/backend-marketplace/internal/repositories"
	"github.com/sm888sm/backend-marketplace/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo  *repositories.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo *repositories.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) Register(user *models.User) (*models.User, error) {
	existingUser, err := s.userRepo.GetByUsername(user.Username)
	if err == nil && existingUser != nil {
		return nil, errors.New("username already exists")
	}

	existingEmailUser, err := s.userRepo.GetByEmail(user.Email)
	if err == nil && existingEmailUser != nil {
		return nil, errors.New("email already registered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("gagal hash password")
	}
	user.Password = string(hashed)

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(username, password string) (*models.User, string, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil || user == nil {
		return nil, "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	token, err := auth.GenerateToken(user.ID, user.Role, s.jwtSecret)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	return user, token, nil
}

func (s *AuthService) ListUsers(page, perPage int, filter map[string]interface{}) ([]models.User, int64, error) {
	return s.userRepo.ListUsers(page, perPage, filter)
}

func (s *AuthService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}
