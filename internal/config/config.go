package config

import (
	"errors"

	"github.com/joho/godotenv"
	"github.com/spacetronot-research-team/cart-service/internal/config/envvar"
	"github.com/spacetronot-research-team/cart-service/pkg/eraerr"
)

func LoadEnvFile() error {
	if err := godotenv.Load(); err != nil {
		return eraerr.Wrap(err, errors.New("error read env file"))
	}

	if err := envvar.ValidateEnvVar(); err != nil {
		return eraerr.Wrap(err, errors.New("error validate env var"))
	}

	return nil
}
