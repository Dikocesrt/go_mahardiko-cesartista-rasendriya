package wishlist

import (
	wishlistEntities "clean/entities/wishlist"

	"gorm.io/gorm"
)

type WishlistRepo struct {
	DB *gorm.DB
}

func NewWishlistRepo(db *gorm.DB) *WishlistRepo {
	return &WishlistRepo{
		DB: db,
	}
}

func (wishlistRepo *WishlistRepo) GetAll(wishlists []wishlistEntities.Wishlist) ([]wishlistEntities.Wishlist, error){
	err := wishlistRepo.DB.Find(&wishlists).Error

	if err != nil {
		return []wishlistEntities.Wishlist{}, err
	}

	return wishlists, nil
}

func (wishlistRepo *WishlistRepo) Create(wishlist wishlistEntities.Wishlist) (wishlistEntities.Wishlist, error) {
	wishlistDB := FromWishlistEntitiesToWishlistDb(wishlist)
	err := wishlistRepo.DB.Create(&wishlistDB).Error
	if err != nil {
		return wishlistEntities.Wishlist{}, err
	}
	wishlist = *wishlistDB.FromWishlistDbToWishlistEntities()
	return wishlist, nil
}