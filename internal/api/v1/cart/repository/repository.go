package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type ICartRepository interface {
	Foo(ctx context.Context) (string, error)
	Bar(ctx context.Context, isShouldError string) (string, error)
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &cartRepository{db: db}
}

func (c *cartRepository) Foo(ctx context.Context) (string, error) {
	type Foo struct {
		Name string `gorm:"name"`
	}
	foo := Foo{}
	q := c.db.Raw("SELECT 'hidayat foo' as name").Scan(&foo)
	if q.Error != nil {
		return "", q.Error
	}
	return foo.Name, nil
}

func (c *cartRepository) Bar(ctx context.Context, isShouldError string) (string, error) {
	if isShouldError == "true" {
		return "", errors.New("dummy error from repository")
	} else {
		return "hidayat bar", nil
	}
}
