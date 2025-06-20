package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sm888sm/backend-marketplace/internal/services"
	"github.com/sm888sm/backend-marketplace/pkg/utils"
)

type CategoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	roleVal, exists := ctx.Get("userRole")
	fmt.Println("roleVal:", roleVal)
	if !exists {
		utils.APIError(ctx, http.StatusUnauthorized, "Akses ditolak: role tidak ditemukan")
		return
	}
	role := roleVal.(string)
	if role != "admin" && role != "merchant" {
		utils.APIError(ctx, http.StatusForbidden, "Hanya admin atau merchant yang bisa membuat kategori")
		return
	}
	var req struct {
		Name string `json:"name" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.APIError(ctx, http.StatusBadRequest, "Format request tidak valid")
		return
	}
	category, err := c.categoryService.CreateCategory(req.Name)
	if err != nil {
		utils.APIError(ctx, http.StatusBadRequest, "Gagal membuat kategori: "+err.Error())
		return
	}
	utils.APIData(ctx, http.StatusCreated, category, "Kategori berhasil dibuat")
}

func (c *CategoryController) ListCategories(ctx *gin.Context) {
	categories, err := c.categoryService.ListCategories()
	if err != nil {
		utils.APIError(ctx, http.StatusInternalServerError, "Gagal mengambil daftar kategori")
		return
	}
	utils.APIData(ctx, http.StatusOK, categories, "Daftar kategori produk")
}

func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	roleVal, exists := ctx.Get("userRole")
	if !exists {
		utils.APIError(ctx, http.StatusUnauthorized, "Akses ditolak: role tidak ditemukan")
		return
	}
	role := roleVal.(string)
	if role != "admin" && role != "merchant" {
		utils.APIError(ctx, http.StatusForbidden, "Hanya admin atau merchant yang bisa mengubah kategori")
		return
	}
	id := ctx.Param("id")
	var req struct {
		Name string `json:"name" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.APIError(ctx, http.StatusBadRequest, "Format request tidak valid")
		return
	}
	category, err := c.categoryService.UpdateCategory(id, req.Name)
	if err != nil {
		utils.APIError(ctx, http.StatusBadRequest, "Gagal mengubah kategori: "+err.Error())
		return
	}
	utils.APIData(ctx, http.StatusOK, category, "Kategori berhasil diubah")
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	roleVal, exists := ctx.Get("userRole")
	if !exists {
		utils.APIError(ctx, http.StatusUnauthorized, "Akses ditolak: role tidak ditemukan")
		return
	}
	role := roleVal.(string)
	if role != "admin" && role != "merchant" {
		utils.APIError(ctx, http.StatusForbidden, "Hanya admin atau merchant yang bisa menghapus kategori")
		return
	}
	id := ctx.Param("id")
	if err := c.categoryService.DeleteCategory(id); err != nil {
		if err.Error() == "record not found" {
			utils.APIError(ctx, http.StatusNotFound, "Kategori tidak ditemukan")
		} else {
			utils.APIError(ctx, http.StatusBadRequest, "Gagal menghapus kategori: "+err.Error())
		}
		return
	}
	utils.APISuccess(ctx, http.StatusOK, "Kategori berhasil dihapus")
}
