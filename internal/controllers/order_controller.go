package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sm888sm/backend-marketplace/internal/models"
	"github.com/sm888sm/backend-marketplace/internal/services"
	"github.com/sm888sm/backend-marketplace/pkg/utils"
)

type OrderController struct {
	orderService   services.OrderService
	productService services.ProductService
}

func NewOrderController(orderService services.OrderService, productService services.ProductService) *OrderController {
	return &OrderController{orderService: orderService, productService: productService}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(uint)
	roleVal, exists := ctx.Get("userRole")
	if !exists {
		utils.APIError(ctx, http.StatusUnauthorized, "Akses ditolak: role tidak ditemukan")
		return
	}
	role := roleVal.(string)
	if role != "customer" {
		utils.APIError(ctx, http.StatusForbidden, "Hanya customer yang bisa membuat order")
		return
	}
	var req struct {
		Items []models.OrderItemRequest `json:"items" binding:"required,dive"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.APIError(ctx, http.StatusBadRequest, "Format request tidak valid atau items kosong")
		return
	}
	if len(req.Items) == 0 {
		utils.APIError(ctx, http.StatusBadRequest, "Order harus memiliki minimal satu item")
		return
	}
	for _, item := range req.Items {
		if item.Quantity <= 0 {
			utils.APIError(ctx, http.StatusBadRequest, "Quantity harus lebih dari 0")
			return
		}
	}
	order, err := c.orderService.CreateOrder(userID, req.Items)
	if err != nil {
		utils.APIError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	utils.APIData(ctx, http.StatusCreated, order, "Order berhasil dibuat")
}

func (c *OrderController) GetOrdersByCustomer(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(uint)
	roleVal, exists := ctx.Get("userRole")
	if !exists {
		utils.APIError(ctx, http.StatusUnauthorized, "Akses ditolak: role tidak ditemukan")
		return
	}
	role := roleVal.(string)
	if role != "customer" {
		utils.APIError(ctx, http.StatusForbidden, "Hanya customer yang bisa melihat daftar order miliknya")
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
	orders, err := c.orderService.GetOrdersByCustomer(userID)
	if err != nil {
		utils.APIError(ctx, http.StatusInternalServerError, "Gagal mengambil data order: "+err.Error())
		return
	}
	total := len(orders)
	totalPages := (total + perPage - 1) / perPage
	start := (page - 1) * perPage
	end := start + perPage
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}
	pagedOrders := orders[start:end]
	meta := &utils.Meta{
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: totalPages,
	}
	utils.APIDataWithMeta(ctx, http.StatusOK, pagedOrders, "Daftar order customer", meta)
}

func (c *OrderController) GetBuyersByMerchant(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(uint)
	roleVal, exists := ctx.Get("userRole")
	if !exists {
		utils.APIError(ctx, http.StatusUnauthorized, "Akses ditolak: role tidak ditemukan")
		return
	}
	role := roleVal.(string)
	if role != "merchant" {
		utils.APIError(ctx, http.StatusForbidden, "Hanya merchant yang bisa melihat daftar pembeli")
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
	buyers, err := c.orderService.GetBuyersByMerchant(userID)
	if err != nil {
		utils.APIError(ctx, http.StatusInternalServerError, "Gagal mengambil data pembeli: "+err.Error())
		return
	}
	total := len(buyers)
	totalPages := (total + perPage - 1) / perPage
	start := (page - 1) * perPage
	end := start + perPage
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}
	pagedBuyers := buyers[start:end]
	meta := &utils.Meta{
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: totalPages,
	}
	utils.APIDataWithMeta(ctx, http.StatusOK, pagedBuyers, "Daftar customer yang pernah membeli produk merchant", meta)
}

func (c *OrderController) GetOrderByID(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(uint)
	roleVal, exists := ctx.Get("userRole")
	if !exists {
		utils.APIError(ctx, http.StatusUnauthorized, "Akses ditolak: role tidak ditemukan")
		return
	}
	role := roleVal.(string)
	if role != "customer" {
		utils.APIError(ctx, http.StatusForbidden, "Hanya customer yang bisa melihat detail order ini")
		return
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.APIError(ctx, http.StatusBadRequest, "ID order tidak valid")
		return
	}
	order, err := c.orderService.GetOrderByID(uint(id))
	if err != nil {
		utils.APIError(ctx, http.StatusNotFound, "Order tidak ditemukan")
		return
	}
	if order.CustomerID != userID {
		utils.APIError(ctx, http.StatusForbidden, "Anda tidak berhak melihat order ini")
		return
	}
	utils.APIData(ctx, http.StatusOK, order, "Detail order")
}

func (c *OrderController) GetOrderByIDMerchant(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(uint)
	roleVal, exists := ctx.Get("userRole")
	if !exists {
		utils.APIError(ctx, http.StatusUnauthorized, "Akses ditolak: role tidak ditemukan")
		return
	}
	role := roleVal.(string)
	if role != "merchant" {
		utils.APIError(ctx, http.StatusForbidden, "Hanya merchant yang bisa melihat detail order ini")
		return
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.APIError(ctx, http.StatusBadRequest, "ID order tidak valid")
		return
	}
	order, err := c.orderService.GetOrderByID(uint(id))
	if err != nil {
		utils.APIError(ctx, http.StatusNotFound, "Order tidak ditemukan")
		return
	}

	var merchantItems []models.OrderItem
	for _, item := range order.Items {
		product, _ := c.productService.GetProductByID(item.ProductID)
		if product != nil && product.MerchantID == userID {
			merchantItems = append(merchantItems, item)
		}
	}
	order.Items = merchantItems
	utils.APIData(ctx, http.StatusOK, order, "Detail order untuk merchant")
}
