package envvar

import (
	"errors"
	"os"
)

const (
	PORT        = "PORT"
	MODE        = "MODE"
	DB_USER     = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
	DB_HOST     = "DB_HOST"
	DB_NAME     = "DB_NAME"
	DB_PORT     = "DB_PORT"
	DB_TZ       = "DB_TZ"
	SSL_MODE    = "SSL_MODE"
)

func ValidateEnvVar() error {
	if err := osEnvMustHave(PORT); err != nil {
		return err
	}

	if err := osEnvMustHave(MODE); err != nil {
		return err
	}

	if err := osEnvMustHave(DB_USER); err != nil {
		return err
	}

	if err := osEnvMustHave(DB_PASSWORD); err != nil {
		return err
	}

	if err := osEnvMustHave(DB_HOST); err != nil {
		return err
	}

	if err := osEnvMustHave(DB_NAME); err != nil {
		return err
	}

	if err := osEnvMustHave(DB_PORT); err != nil {
		return err
	}

	if err := osEnvMustHave(DB_TZ); err != nil {
		return err
	}

	if err := osEnvMustHave(SSL_MODE); err != nil {
		return err
	}

	return nil
}

func osEnvMustHave(key string) error {
	if os.Getenv(key) == "" {
		return errors.New(key + " is not exist in env var")
	}
	return nil
}
