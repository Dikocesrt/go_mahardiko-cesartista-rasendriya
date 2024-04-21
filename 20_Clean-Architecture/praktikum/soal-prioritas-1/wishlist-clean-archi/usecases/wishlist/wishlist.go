package wishlist

import (
	"clean/constants"
	wishlistEntities "clean/entities/wishlist"
)

type WishlistUseCase struct {
	repository wishlistEntities.RepositoryInterface
}

func NewWishlistUseCase(repository wishlistEntities.RepositoryInterface) *WishlistUseCase {
	return &WishlistUseCase{
		repository: repository,
	}
}

func (wishlistUseCase *WishlistUseCase) GetAll(wishlists []wishlistEntities.Wishlist) ([]wishlistEntities.Wishlist, error) {
	wishlists, err := wishlistUseCase.repository.GetAll(wishlists)
	if err != nil {
		return []wishlistEntities.Wishlist{}, constants.ErrGetAllDatabase
	}
	return wishlists, nil
}

func (wishlistUseCase *WishlistUseCase) Create(wishlist wishlistEntities.Wishlist) (wishlistEntities.Wishlist, error) {
	if wishlist.Title == "" {
		return wishlistEntities.Wishlist{}, constants.ErrEmptyInput
	}
	wishlist, err := wishlistUseCase.repository.Create(wishlist)
	if err != nil {
		return wishlistEntities.Wishlist{}, constants.ErrInsertDatabase
	}
	return wishlist, nil
}