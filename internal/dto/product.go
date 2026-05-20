// Package dto This file defines the data transfer objects (DTOs) for products and categories in the CartQL application.
package dto

// CreateCategoryRequest represents the request payload for creating a new category, containing the name and description of the category.
type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// UpdateCategoryRequest represents the request payload for updating an existing category, allowing changes to the name, description, and active status of the category.
type UpdateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	IsActive    *bool  `json:"is_active"`
}

// CategoryResponse represents the response structure for category-related API endpoints, including the category's ID, name, description, and active status.
type CategoryResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

// CreateProductRequest represents the request payload for creating a new product, containing the category ID, name, description, price, stock quantity, and SKU of the product.
type CreateProductRequest struct {
	CategoryID  uint    `json:"category_id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"min=0"`
	SKU         string  `json:"sku" binding:"required"`
}

// UpdateProductRequest represents the request payload for updating an existing product, allowing changes to the category ID, name, description, price, stock quantity, and active status of the product.
type UpdateProductRequest struct {
	CategoryID  uint    `json:"category_id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"min=0"`
	IsActive    *bool   `json:"is_active"`
}

// ProductResponse represents the response structure for product-related API endpoints, including the product's ID, category ID, name, description, price, stock quantity, SKU, active status, associated category details, and product images.
type ProductResponse struct {
	ID          uint                   `json:"id"`
	CategoryID  uint                   `json:"category_id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Price       float64                `json:"price"`
	Stock       int                    `json:"stock"`
	SKU         string                 `json:"sku"`
	IsActive    bool                   `json:"is_active"`
	Category    CategoryResponse       `json:"category"`
	Images      []ProductImageResponse `json:"images"`
}

// ProductImageResponse represents the response structure for product image-related API endpoints, including the image's ID, URL, alternative text, and whether it is the primary image for the product.
type ProductImageResponse struct {
	ID        uint   `json:"id"`
	URL       string `json:"url"`
	AltText   string `json:"alt_text"`
	IsPrimary bool   `json:"is_primary"`
}
