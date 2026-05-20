// Package dto defines the data transfer objects (DTOs) for orders and cart items in the CartQL application.
package dto

// AddToCartRequest represents the request payload for adding a product to the cart, containing the product ID and quantity to be added.
type AddToCartRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

// UpdateCartItemRequest represents the request payload for updating the quantity of a cart item, containing the new quantity for the specified cart item.
type UpdateCartItemRequest struct {
	Quantity int `json:"quantity" binding:"required,min=1"`
}

// CartResponse represents the response structure for cart-related API endpoints, including the cart's ID, user ID, list of cart items, and total amount for the cart.
type CartResponse struct {
	ID        uint               `json:"id"`
	UserID    uint               `json:"user_id"`
	CartItems []CartItemResponse `json:"cart_items"`
	Total     float64            `json:"total"`
}

// CartItemResponse represents the response structure for cart item-related API endpoints, including the cart item's ID, associated product details, quantity added to the cart, and subtotal for the cart item.
type CartItemResponse struct {
	ID       uint            `json:"id"`
	Product  ProductResponse `json:"product"`
	Quantity int             `json:"quantity"`
	Subtotal float64         `json:"subtotal"`
}

// OrderResponse represents the response structure for order-related API endpoints, including the order's ID, user ID, status, total amount, list of order items, and creation timestamp.
type OrderResponse struct {
	ID          uint                `json:"id"`
	UserID      uint                `json:"user_id"`
	Status      string              `json:"status"`
	TotalAmount float64             `json:"total_amount"`
	OrderItems  []OrderItemResponse `json:"order_items"`
	CreatedAt   string              `json:"created_at"`
}

// OrderItemResponse represents the response structure for order item-related API endpoints, including the order item's ID, associated product details, quantity ordered, and price for the order item.
type OrderItemResponse struct {
	ID       uint            `json:"id"`
	Product  ProductResponse `json:"product"`
	Quantity int             `json:"quantity"`
	Price    float64         `json:"price"`
}
