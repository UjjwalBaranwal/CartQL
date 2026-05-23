package services

import (
	"github.com/UjjwalBaranwal/CartQL/internal/dto"
	"github.com/UjjwalBaranwal/CartQL/internal/models"
	"github.com/UjjwalBaranwal/CartQL/internal/utils"
	"gorm.io/gorm"
)

// ProductService is the structure of the product service
type ProductService struct {
	db *gorm.DB
}

// NewProductService creates a new instance of ProductService
func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{
		db: db,
	}
}

// CreateCategory creates a new category in the database and returns the created category as a response.
func (s *ProductService) CreateCategory(req *dto.CreateCategoryRequest) (*dto.CategoryResponse, error) {

	category := models.Category{
		Name:        req.Name,
		Description: req.Description,
	}
	if err := s.db.Create(&category).Error; err != nil {
		return nil, err
	}
	return &dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive,
	}, nil
}

// GetCategories retrieves all active categories from the database and returns them as a list of category responses.
func (s *ProductService) GetCategories() ([]dto.CategoryResponse, error) {
	var categories []models.Category
	if err := s.db.Where("is_active = ?", true).Find(&categories).Error; err != nil {
		return nil, err
	}

	response := make([]dto.CategoryResponse, len(categories))
	for i := range categories {
		response[i] = dto.CategoryResponse{
			ID:          categories[i].ID,
			Name:        categories[i].Name,
			Description: categories[i].Description,
			IsActive:    categories[i].IsActive,
		}
	}

	return response, nil
}

// UpdateCategory updates an existing category in the database based on the provided category ID and update request, and returns the updated category as a response.
func (s *ProductService) UpdateCategory(id uint, req *dto.UpdateCategoryRequest) (*dto.CategoryResponse, error) {

	var category models.Category
	if err := s.db.First(&category, id).Error; err != nil {
		return nil, err
	}

	category.Name = req.Name
	category.Description = req.Description
	if req.IsActive != nil {
		category.IsActive = *req.IsActive
	}

	if err := s.db.Save(&category).Error; err != nil {
		return nil, err
	}

	return &dto.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		IsActive:    category.IsActive,
	}, nil
}

// DeleteCategory deletes a category from the database based on the provided category ID.
func (s *ProductService) DeleteCategory(id uint) error {
	return s.db.Delete(&models.Category{}, id).Error
}

func (s *ProductService) AddProductImage(productID uint, url, altText string) error {
	var count int64
	s.db.Model(&models.ProductImage{}).Where("product_id = ?", productID).Count(&count)
	image := models.ProductImage{
		ProductID: productID,
		URL:       url,
		AltText:   altText,
		IsPrimary: count == 0, // Set as primary if it's the first image
	}
	return s.db.Create(&image).Error
}

// CreateProduct creates a new product in the database based on the provided product creation request, and returns the created product as a response.
func (s *ProductService) CreateProduct(req *dto.CreateProductRequest) (*dto.ProductResponse, error) {
	product := models.Product{
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		SKU:         req.SKU,
	}
	if err := s.db.Create(&product).Error; err != nil {
		return nil, err
	}
	return s.GetProduct(product.ID)
}

// GetProducts retrieves a paginated list of active products from the database, along with pagination metadata, and returns them as a list of product responses.
func (s *ProductService) GetProducts(page, limit int) ([]dto.ProductResponse, *utils.PaginationMeta, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit
	var products []models.Product
	var total int64

	s.db.Model(&models.Product{}).Where("is_active = ?", true).Count(&total)

	if err := s.db.Preload("Category").Preload("Images").
		Where("is_active = ?", true).
		Offset(offset).Limit(limit).
		Find(&products).Error; err != nil {
		return nil, nil, err
	}

	response := make([]dto.ProductResponse, len(products))
	for i := range products {
		response[i] = s.convertToProductResponse(&products[i])
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))
	meta := &utils.PaginationMeta{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}

	return response, meta, nil
}

// GetProduct retrieves a product from the database based on the provided product ID and returns it as a product response.
func (s *ProductService) GetProduct(id uint) (*dto.ProductResponse, error) {
	var product models.Product
	if err := s.db.Preload("Category").Preload("Images").First(&product, id).Error; err != nil {
		return nil, err
	}

	response := s.convertToProductResponse(&product)
	return &response, nil
}

// UpdateProduct updates an existing product in the database based on the provided product ID and update request, and returns the updated product as a response.
func (s *ProductService) UpdateProduct(id uint, req *dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	var product models.Product
	if err := s.db.First(&product, id).Error; err != nil {
		return nil, err
	}

	product.CategoryID = req.CategoryID
	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock
	if req.IsActive != nil {
		product.IsActive = *req.IsActive
	}

	if err := s.db.Save(&product).Error; err != nil {
		return nil, err
	}

	return s.GetProduct(id)
}

// DeleteProduct deletes a product from the database based on the provided product ID.
func (s *ProductService) DeleteProduct(id uint) error {
	return s.db.Delete(&models.Product{}, id).Error
}

// convertToProductResponse converts a product model to a product response DTO, including its associated category and images.
func (s *ProductService) convertToProductResponse(product *models.Product) dto.ProductResponse {
	images := make([]dto.ProductImageResponse, len(product.Images))
	for i := range product.Images {
		images[i] = dto.ProductImageResponse{
			ID:        product.Images[i].ID,
			URL:       product.Images[i].URL,
			AltText:   product.Images[i].AltText,
			IsPrimary: product.Images[i].IsPrimary,
		}
	}

	return dto.ProductResponse{
		ID:          product.ID,
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		SKU:         product.SKU,
		IsActive:    product.IsActive,
		Category: dto.CategoryResponse{
			ID:          product.Category.ID,
			Name:        product.Category.Name,
			Description: product.Category.Description,
			IsActive:    product.Category.IsActive,
		},
		Images: images,
	}
}
