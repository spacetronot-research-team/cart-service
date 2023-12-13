package application

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spacetronot-research-team/cart-service/internal/config"
	"github.com/spacetronot-research-team/cart-service/internal/config/envvar"
	"github.com/spacetronot-research-team/cart-service/internal/infrastructure/database"
	"github.com/spacetronot-research-team/cart-service/internal/router"
)

func RunApp() {
	if err := config.LoadEnvFile(); err != nil {
		panic(err)
	}

	app := getGinEngineApp()

	db, err := database.GetDB()
	if err != nil {
		panic(err)
	}

	router.AddRouter(app, db)

	if err := app.Run(getTcpNetworkAddress()); err != nil {
		panic(err)
	}
}

func getGinEngineApp() *gin.Engine {
	if os.Getenv(envvar.MODE) == "local" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.Default()

	return app
}

func getTcpNetworkAddress() string {
	address := "0.0.0.0:" + os.Getenv(envvar.PORT)
	return address
}
