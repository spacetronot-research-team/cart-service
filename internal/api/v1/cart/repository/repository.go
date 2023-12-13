package repository

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ICartRepository interface {
	Foo(ctx *gin.Context) (string, error)
	Bar(ctx *gin.Context, isShouldError string) (string, error)
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &cartRepository{db: db}
}

func (c *cartRepository) Foo(ctx *gin.Context) (string, error) {
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

func (c *cartRepository) Bar(ctx *gin.Context, isShouldError string) (string, error) {
	if isShouldError == "true" {
		return "", errors.New("dummy error from repository")
	} else {
		return "hidayat bar", nil
	}
}
