package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spacetronot-research-team/cart-service/internal/api/v1/cart/service"
	"github.com/spacetronot-research-team/cart-service/internal/pkg/carterr"
	"github.com/spacetronot-research-team/cart-service/pkg/gin/response"
)

type ICartController interface {
	Baz(*gin.Context)
}

type cartController struct {
	cartService service.ICartService
}

func NewCartController(cartService service.ICartService) ICartController {
	return &cartController{cartService: cartService}
}

func (c *cartController) Baz(ctx *gin.Context) {
	isShouldError := ctx.Param("is_should_error")
	name, err := c.cartService.Baz(ctx, isShouldError)
	if errors.Is(err, carterr.ErrBar) {
		response.Write(ctx, http.StatusBadRequest, nil, "CART5001001")
		return
	} else if err != nil {
		response.Write(ctx, http.StatusBadRequest, nil, "CART5001000")
		return
	}
	response.Write(ctx, http.StatusBadRequest, name, "CART2001000")
}
