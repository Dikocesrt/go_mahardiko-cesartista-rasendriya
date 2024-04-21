package wishlist

import (
	"time"

	"gorm.io/gorm"
)

type Wishlist struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
	Title      string
	IsAchieved bool
}

type RepositoryInterface interface {
	GetAll(wishlists []Wishlist) ([]Wishlist, error)
	Create(wishlist Wishlist) (Wishlist, error)
}

type UseCaseInterface interface {
	GetAll(wishlists []Wishlist) ([]Wishlist, error)
	Create(wishlist Wishlist) (Wishlist, error)
}