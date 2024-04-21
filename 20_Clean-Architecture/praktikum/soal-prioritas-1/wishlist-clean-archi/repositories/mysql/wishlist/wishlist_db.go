package wishlist

import (
	wishlistEntities "clean/entities/wishlist"
	"time"

	"gorm.io/gorm"
)

type Wishlist struct {
	ID         uint           `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Title      string
	IsAchieved bool
}

func FromWishlistEntitiesToWishlistDb(wishlist wishlistEntities.Wishlist) *Wishlist {
	return &Wishlist{
		ID: wishlist.ID,
		CreatedAt:  wishlist.CreatedAt,
		UpdatedAt:  wishlist.UpdatedAt,
		DeletedAt:  gorm.DeletedAt{},
		Title:      wishlist.Title,
		IsAchieved: wishlist.IsAchieved,
	}
}

func (wishlist *Wishlist) FromWishlistDbToWishlistEntities() *wishlistEntities.Wishlist {
	return &wishlistEntities.Wishlist{
		ID:         wishlist.ID,
		CreatedAt:  wishlist.CreatedAt,
		UpdatedAt:  wishlist.UpdatedAt,
		DeletedAt:  gorm.DeletedAt{},
		Title:      wishlist.Title,
		IsAchieved: wishlist.IsAchieved,
	}
}