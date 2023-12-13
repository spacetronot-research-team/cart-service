package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spacetronot-research-team/cart-service/internal/api/v1/cart/controller"
	"github.com/spacetronot-research-team/cart-service/internal/api/v1/cart/repository"
	"github.com/spacetronot-research-team/cart-service/internal/api/v1/cart/service"
	"github.com/spacetronot-research-team/cart-service/internal/middleware"
	"gorm.io/gorm"
)

func AddRouter(app *gin.Engine, db *gorm.DB) {
	app.Use(middleware.GetCorsConfig())
	app.Use(middleware.TraceIDMiddleware())

	apiV1(app, db)
}

func apiV1(app *gin.Engine, db *gorm.DB) {
	cartController := initCart(db)

	apiV1 := app.Group("/api/v1")
	apiV1.GET("/baz/:is_should_error", cartController.Baz)
}

func initCart(db *gorm.DB) controller.ICartController {
	cartRepository := repository.NewCartRepository(db)
	cartService := service.NewCartService(cartRepository)
	return controller.NewCartController(cartService)
}
