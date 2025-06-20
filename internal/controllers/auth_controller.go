package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sm888sm/backend-marketplace/internal/models"
	"github.com/sm888sm/backend-marketplace/internal/services"
	"github.com/sm888sm/backend-marketplace/pkg/utils"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req models.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.APIError(ctx, http.StatusBadRequest, "Payload permintaan tidak valid")
		return
	}

	user := &models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Role:     req.Role,
	}

	createdUser, err := c.authService.Register(user)
	if err != nil {
		msg := err.Error()
		if msg == "username already exists" {
			msg = "Username sudah terdaftar"
		} else if msg == "email already registered" {
			msg = "Email sudah terdaftar"
		} else if msg == "gagal hash password" {
			msg = "Gagal mengenkripsi password"
		} else {
			msg = "Registrasi gagal: " + msg
		}
		utils.APIError(ctx, http.StatusBadRequest, msg)
		return
	}

	resp := gin.H{
		"id":         createdUser.ID,
		"username":   createdUser.Username,
		"email":      createdUser.Email,
		"role":       createdUser.Role,
		"created_at": createdUser.CreatedAt,
	}
	utils.APIData(ctx, http.StatusCreated, resp, "Registrasi berhasil")
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.APIError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, token, err := c.authService.Login(req.Username, req.Password)
	if err != nil {
		utils.APIError(ctx, http.StatusUnauthorized, "Login gagal: "+err.Error())
		return
	}
	resp := gin.H{
		"token": token,
		"user": gin.H{
			"id":         user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"role":       user.Role,
			"created_at": user.CreatedAt,
		},
	}
	utils.APIData(ctx, http.StatusOK, resp, "Login berhasil")
}

func (c *AuthController) ListUsers(ctx *gin.Context) {
	roleVal, exists := ctx.Get("userRole")
	if !exists {
		utils.APIError(ctx, http.StatusUnauthorized, "Akses ditolak: role tidak ditemukan")
		return
	}
	role := roleVal.(string)
	if role != "admin" {
		utils.APIError(ctx, http.StatusForbidden, "Hanya admin yang bisa melihat daftar user")
		return
	}
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(ctx.DefaultQuery("per_page", "10"))
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}
	filter := make(map[string]interface{})
	if roleParam := ctx.Query("role"); roleParam != "" {
		filter["role"] = roleParam
	}
	users, total, err := c.authService.ListUsers(page, perPage, filter)
	if err != nil {
		utils.APIError(ctx, http.StatusInternalServerError, "Gagal mengambil daftar user")
		return
	}
	totalPages := int((total + int64(perPage) - 1) / int64(perPage))
	meta := &utils.Meta{
		Page:       page,
		PerPage:    perPage,
		Total:      int(total),
		TotalPages: totalPages,
	}
	utils.APIDataWithMeta(ctx, http.StatusOK, users, "Daftar user", meta)
}

func (c *AuthController) GetUserByID(ctx *gin.Context) {
	roleVal, exists := ctx.Get("userRole")
	if !exists {
		utils.APIError(ctx, http.StatusUnauthorized, "Akses ditolak: role tidak ditemukan")
		return
	}
	role := roleVal.(string)
	if role != "admin" {
		utils.APIError(ctx, http.StatusForbidden, "Hanya admin yang bisa melihat detail user")
		return
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.APIError(ctx, http.StatusBadRequest, "ID user tidak valid")
		return
	}
	user, err := c.authService.GetUserByID(uint(id))
	if err != nil {
		utils.APIError(ctx, http.StatusNotFound, "User tidak ditemukan")
		return
	}
	utils.APIData(ctx, http.StatusOK, user, "Detail user")
}
