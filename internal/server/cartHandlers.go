// server package , this is cart handler
package server

import (
	"strconv"

	"github.com/UjjwalBaranwal/CartQL/internal/dto"
	"github.com/UjjwalBaranwal/CartQL/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) getCart(c *gin.Context) {
	userId := c.GetUint("user_id")
	cart, err := s.cartService.GetCart(userId)
	if err != nil {
		utils.NotFoundResponse(c, "Cart not found")
		return
	}
	utils.SuccessResponse(c, "Cart retrieved successfully", cart)
}

func (s *Server) addToCart(c *gin.Context) {
	userId := c.GetUint("user_id")
	var req dto.AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}
	cart, err := s.cartService.AddToCart(userId, &req)
	if err != nil {
		utils.BadRequestResponse(c, "Failed to add item to cart", err)
		return
	}

	utils.SuccessResponse(c, "Item added to cart successfully", cart)
}

func (s *Server) updateCartItem(c *gin.Context) {
	userID := c.GetUint("user_id")

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid cart item ID", err)
		return
	}

	var req dto.UpdateCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}

	cart, err := s.cartService.UpdateCartItem(userID, uint(id), &req)
	if err != nil {
		utils.BadRequestResponse(c, "Failed to update cart item", err)
		return
	}

	utils.SuccessResponse(c, "Cart item updated successfully", cart)
}

func (s *Server) removeFromCart(c *gin.Context) {
	userID := c.GetUint("user_id")

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid cart item ID", err)
		return
	}

	if err := s.cartService.RemoveFromCart(userID, uint(id)); err != nil {
		utils.InternalServerErrorResponse(c, "Failed to remove item from cart", err)
		return
	}

	utils.SuccessResponse(c, "Item removed from cart successfully", nil)
}
