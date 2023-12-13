package service

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spacetronot-research-team/cart-service/internal/api/v1/cart/repository"
	"github.com/spacetronot-research-team/cart-service/internal/pkg/carterr"
	"github.com/spacetronot-research-team/cart-service/pkg/eraerr"
)

type ICartService interface {
	Baz(ctx *gin.Context, isShouldError string) (string, error)
}

type cartService struct {
	cartRepository repository.ICartRepository
}

func NewCartService(cartRepository repository.ICartRepository) ICartService {
	return &cartService{cartRepository: cartRepository}
}

func (c *cartService) Baz(ctx *gin.Context, isShouldError string) (string, error) {
	foo, err := c.cartRepository.Foo(ctx)
	if err != nil {
		return "", eraerr.Wrap(err, errors.New("error foo"))
	}
	bar, err := c.cartRepository.Bar(ctx, isShouldError)
	if err != nil {
		return "", eraerr.Wrap(err, carterr.ErrBar)
	}
	return fmt.Sprintf("%s||%s", foo, bar), nil
}