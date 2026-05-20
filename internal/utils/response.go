// Package utils provides utility functions for handling API responses in a standardized format, including success and error responses, as well as paginated responses.
package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents the standard structure for API responses, including success status, message, data, and error information.
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

// PaginatedResponse extends the standard Response structure to include pagination metadata for paginated API responses.
type PaginatedResponse struct {
	Response
	Meta PaginationMeta `json:"meta"`
}

// PaginationMeta contains metadata about the pagination state, including the current page, limit per page, total items, and total pages.
type PaginationMeta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// SuccessResponse sends a successful JSON response with the provided message and data.
func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// CreatedResponse sends a JSON response with a 201 Created status, including the provided message and data.
func CreatedResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse sends a JSON response with the specified status code, message, and error details if provided.
func ErrorResponse(c *gin.Context, statusCode int, message string, err error) {
	response := Response{
		Success: false,
		Message: message,
	}

	if err != nil {
		response.Error = err.Error()
	}

	c.JSON(statusCode, response)
}

// BadRequestResponse sends a JSON response with a 400 Bad Request status, including the provided message and error details.
func BadRequestResponse(c *gin.Context, message string, err error) {
	ErrorResponse(c, http.StatusBadRequest, message, err)
}

// UnauthorizedResponse sends a JSON response with a 401 Unauthorized status, including the provided message.
func UnauthorizedResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, message, nil)
}

// ForbiddenResponse sends a JSON response with a 403 Forbidden status, including the provided message.
func ForbiddenResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusForbidden, message, nil)
}

// NotFoundResponse sends a JSON response with a 404 Not Found status, including the provided message.
func NotFoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message, nil)
}

// InternalServerErrorResponse sends a JSON response with a 500 Internal Server Error status, including the provided message and error details.
func InternalServerErrorResponse(c *gin.Context, message string, err error) {
	ErrorResponse(c, http.StatusInternalServerError, message, err)
}

// PaginatedSuccessResponse sends a successful JSON response with pagination metadata, including the provided message, data, and pagination details.
func PaginatedSuccessResponse(c *gin.Context, message string, data interface{}, meta PaginationMeta) {
	c.JSON(http.StatusOK, PaginatedResponse{
		Response: Response{
			Success: true,
			Message: message,
			Data:    data,
		},
		Meta: meta,
	})
}
