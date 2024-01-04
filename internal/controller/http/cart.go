package http

import "github.com/spacetronot-research-team/cart-service/internal/service"

type CartController struct {
	cartService service.Cart
}

func NewCartController(cartService service.Cart) *CartController {
	return &CartController{}
}
