package repository

type Cart interface {
}

type cartRepository struct {
}

func NewCartRepository() Cart {
	return &cartRepository{}
}
