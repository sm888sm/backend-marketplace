package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sm888sm/backend-marketplace/internal/models"
	"github.com/sm888sm/backend-marketplace/internal/services"
	"github.com/sm888sm/backend-marketplace/pkg/utils"
)

type ProductController struct {
	productService  services.ProductService
	categoryService services.CategoryService
	authService     *services.AuthService
}

func NewProductController(productService services.ProductService, categoryService services.CategoryService, authService *services.AuthService) *ProductController {
	return &ProductController{
		productService:  productService,
		categoryService: categoryService,
		authService:     authService,
	}
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	merchantID := ctx.MustGet("userID").(uint)

	user, err := c.authService.GetUserByID(merchantID)
	if err != nil || user == nil || user.Role != "merchant" {
		utils.APIError(ctx, http.StatusBadRequest, "Merchant tidak valid")
		return
	}

	var req models.ProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.APIError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	categoryID := req.CategoryID

	if categoryID == 0 {
		utils.APIError(ctx, http.StatusBadRequest, "category_id wajib diisi")
		return
	}

	category, err := c.categoryService.GetByIDUint(categoryID)
	if err != nil || category == nil {
		utils.APIError(ctx, http.StatusBadRequest, "Kategori tidak ditemukan")
		return
	}
	product := &models.Product{
		MerchantID:  merchantID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  categoryID,
	}

	createdProduct, err := c.productService.CreateProduct(product)
	if err != nil {
		utils.APIError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.APIData(ctx, http.StatusCreated, createdProduct, "Produk berhasil dibuat")
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	merchantID := ctx.MustGet("userID").(uint)
	productID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.APIError(ctx, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var req models.ProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.APIError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	categoryID := req.CategoryID

	if categoryID == 0 {
		utils.APIError(ctx, http.StatusBadRequest, "category_id wajib diisi")
		return
	}

	category, err := c.categoryService.GetByIDUint(categoryID)
	if err != nil || category == nil {
		utils.APIError(ctx, http.StatusBadRequest, "Kategori tidak ditemukan")
		return
	}
	product := &models.Product{
		MerchantID:  merchantID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  categoryID,
	}
	product.ID = uint(productID)

	updatedProduct, err := c.productService.UpdateProduct(product)
	if err != nil {
		if err.Error() == "product not found" {
			utils.APIError(ctx, http.StatusNotFound, err.Error())
		} else {
			utils.APIError(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}
	utils.APIData(ctx, http.StatusOK, updatedProduct, "Produk berhasil diupdate")
}

func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	merchantID := ctx.MustGet("userID").(uint)
	productID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.APIError(ctx, http.StatusBadRequest, "Invalid product ID")
		return
	}

	err = c.productService.DeleteProduct(uint(productID), merchantID)
	if err != nil {
		if err.Error() == "product not found" {
			utils.APIError(ctx, http.StatusNotFound, err.Error())
		} else {
			utils.APIError(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.APIData(ctx, http.StatusOK, nil, "Produk berhasil dihapus")
}

func (c *ProductController) GetMerchantProducts(ctx *gin.Context) {
	merchantID := ctx.MustGet("userID").(uint)

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(ctx.DefaultQuery("per_page", "10"))
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}

	products, err := c.productService.GetProductsByMerchant(merchantID)
	if err != nil {
		utils.APIError(ctx, http.StatusInternalServerError, "Gagal mengambil produk merchant")
		return
	}

	total := len(products)
	totalPages := (total + perPage - 1) / perPage
	start := (page - 1) * perPage
	end := start + perPage
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}
	pagedProducts := products[start:end]

	meta := &utils.Meta{
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: totalPages,
	}
	utils.APIDataWithMeta(ctx, http.StatusOK, pagedProducts, "Daftar produk merchant", meta)
}

func (c *ProductController) GetAllProducts(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(ctx.DefaultQuery("per_page", "10"))
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}

	products, err := c.productService.GetAllProducts()
	if err != nil {
		utils.APIError(ctx, http.StatusInternalServerError, "Gagal mengambil daftar produk")
		return
	}
	total := len(products)
	totalPages := (total + perPage - 1) / perPage
	start := (page - 1) * perPage
	end := start + perPage
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}
	pagedProducts := products[start:end]

	meta := &utils.Meta{
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: totalPages,
	}
	utils.APIDataWithMeta(ctx, http.StatusOK, pagedProducts, "Daftar semua produk", meta)
}

func (c *ProductController) GetProductByID(ctx *gin.Context) {
	productID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.APIError(ctx, http.StatusBadRequest, "ID produk tidak valid")
		return
	}

	product, err := c.productService.GetProductByID(uint(productID))
	if err != nil {
		if err.Error() == "product not found" {
			utils.APIError(ctx, http.StatusNotFound, "Produk tidak ditemukan")
		} else {
			utils.APIError(ctx, http.StatusInternalServerError, "Gagal mengambil detail produk")
		}
		return
	}

	utils.APIData(ctx, http.StatusOK, product, "Detail produk")
}

func (c *ProductController) UploadImage(ctx *gin.Context) {
	merchantID := ctx.MustGet("userID").(uint)
	role := ""
	roleVal, exists := ctx.Get("userRole")
	if exists {
		role = roleVal.(string)
	}
	if role != "merchant" {
		utils.APIError(ctx, http.StatusForbidden, "Hanya merchant yang bisa upload gambar produk")
		return
	}
	productID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.APIError(ctx, http.StatusBadRequest, "ID produk tidak valid")
		return
	}
	file, err := ctx.FormFile("image")
	if err != nil {
		utils.APIError(ctx, http.StatusBadRequest, "File gambar tidak ditemukan")
		return
	}

	path := "uploads/" + file.Filename
	if err := ctx.SaveUploadedFile(file, path); err != nil {
		utils.APIError(ctx, http.StatusInternalServerError, "Gagal menyimpan file gambar")
		return
	}
	img, err := c.productService.UploadImage(uint(productID), merchantID, path)
	if err != nil {
		utils.APIError(ctx, http.StatusBadRequest, "Gagal upload gambar: "+err.Error())
		return
	}
	utils.APIData(ctx, http.StatusCreated, img, "Gambar produk berhasil diupload")
}

func (c *ProductController) ListImages(ctx *gin.Context) {
	productID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.APIError(ctx, http.StatusBadRequest, "ID produk tidak valid")
		return
	}
	images, err := c.productService.ListImages(uint(productID))
	if err != nil {
		utils.APIError(ctx, http.StatusInternalServerError, "Gagal mengambil gambar produk")
		return
	}
	utils.APIData(ctx, http.StatusOK, images, "Daftar gambar produk")
}

func (c *ProductController) ExploreProducts(ctx *gin.Context) {
	query := ctx.DefaultQuery("search", "")
	minPrice, _ := strconv.ParseFloat(ctx.DefaultQuery("min_price", "0"), 64)
	maxPrice, _ := strconv.ParseFloat(ctx.DefaultQuery("max_price", "0"), 64)
	categoryID, _ := strconv.ParseUint(ctx.DefaultQuery("category_id", "0"), 10, 64)
	sort := ctx.DefaultQuery("sort", "")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(ctx.DefaultQuery("per_page", "10"))
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}
	role := ""
	roleVal, exists := ctx.Get("userRole")
	if exists {
		role = roleVal.(string)
	}
	userID := uint(0)
	if role == "customer" {
		userID = ctx.MustGet("userID").(uint)
	}
	products, total, err := c.productService.SearchProducts(query, minPrice, maxPrice, uint(categoryID), sort, page, perPage)
	if err != nil {
		utils.APIError(ctx, http.StatusInternalServerError, "Gagal mengambil produk")
		return
	}

	if role == "customer" && userID > 0 {
		var filtered []models.Product
		for _, p := range products {
			if p.MerchantID != userID {
				filtered = append(filtered, p)
			}
		}
		products = filtered
	}
	totalPages := int((total + int64(perPage) - 1) / int64(perPage))
	meta := &utils.Meta{
		Page:       page,
		PerPage:    perPage,
		Total:      int(total),
		TotalPages: totalPages,
	}
	utils.APIDataWithMeta(ctx, http.StatusOK, products, "Hasil pencarian produk", meta)
}
