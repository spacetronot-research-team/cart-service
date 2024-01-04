package service

import "github.com/spacetronot-research-team/cart-service/internal/repository"

type Cart interface {
	AddCart()
	GetCart()
	DeleteCart()
}

type cartService struct {
	cartRepository repository.Cart
}

func NewCartService(cartRepository repository.Cart) Cart {
	return &cartService{
		cartRepository: cartRepository,
	}
}

// AddCart implements Cart.
func (*cartService) AddCart() {
	panic("unimplemented")
}

// DeleteCart implements Cart.
func (*cartService) DeleteCart() {
	panic("unimplemented")
}

// GetCart implements Cart.
func (*cartService) GetCart() {
	panic("unimplemented")
}
