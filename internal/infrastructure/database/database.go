package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/spacetronot-research-team/cart-service/internal/config/envvar"
	"github.com/spacetronot-research-team/cart-service/pkg/eraerr"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv(envvar.DB_HOST), os.Getenv(envvar.DB_USER),
		os.Getenv(envvar.DB_PASSWORD), os.Getenv(envvar.DB_NAME),
		os.Getenv(envvar.DB_PORT), os.Getenv(envvar.SSL_MODE),
		os.Getenv(envvar.DB_TZ),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, eraerr.Wrap(err, errors.New("error gorm open db"))
	}
	return db, nil
}
