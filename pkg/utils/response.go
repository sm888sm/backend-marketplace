package utils

import "github.com/gin-gonic/gin"

type Meta struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"per_page,omitempty"`
	Total      int `json:"total,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type DataResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"Message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"Message"`
}

func APIError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ErrorResponse{Success: false, Error: message})
}

func APISuccess(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, SuccessResponse{Success: true, Message: message})
}

func APIData(c *gin.Context, statusCode int, data interface{}, message string) {
	c.JSON(statusCode, DataResponse{Success: true, Message: message, Data: data})
}

func APIDataWithMeta(c *gin.Context, statusCode int, data interface{}, message string, meta *Meta) {
	c.JSON(statusCode, DataResponse{Success: true, Message: message, Data: data, Meta: meta})
}
