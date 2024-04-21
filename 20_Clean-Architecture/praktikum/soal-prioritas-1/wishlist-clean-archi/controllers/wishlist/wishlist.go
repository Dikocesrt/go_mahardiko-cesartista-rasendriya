package wishlist

import (
	wishlistReqRes "clean/controllers/wishlist/reqres"
	wishlistEntities "clean/entities/wishlist"
	"clean/utilities/base"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WishlistController struct {
	wishlistUseCase wishlistEntities.UseCaseInterface
}

func NewWishlistController(wishlistUseCase wishlistEntities.UseCaseInterface) *WishlistController {
	return &WishlistController{
		wishlistUseCase: wishlistUseCase,
	}
}

func (wishlistController *WishlistController) GetAll(c echo.Context) error {
	var wishlists []wishlistEntities.Wishlist

	wishlists, err := wishlistController.wishlistUseCase.GetAll(wishlists)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	wishlistsRes := make([]wishlistReqRes.Wishlist, len(wishlists))
	for i, wishlist := range wishlists {
		wishlistsRes[i] = wishlistReqRes.Wishlist{
			ID:         wishlist.ID,
			CreatedAt:  wishlist.CreatedAt,
			UpdatedAt:  wishlist.UpdatedAt,
			DeletedAt:  wishlist.DeletedAt,
			Title:      wishlist.Title,
			IsAchieved: wishlist.IsAchieved,
		}
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Get All Wishlists", wishlistsRes))
}

func (wishlistController *WishlistController) Create(c echo.Context) error {
	var wishlistReq wishlistReqRes.Wishlist
	c.Bind(&wishlistReq)

	wishlist := wishlistEntities.Wishlist{
		ID:         wishlistReq.ID,
		CreatedAt:  wishlistReq.CreatedAt,
		UpdatedAt:  wishlistReq.UpdatedAt,
		DeletedAt:  wishlistReq.DeletedAt,
		Title:      wishlistReq.Title,
		IsAchieved: wishlistReq.IsAchieved,
	}

	wishlist, err := wishlistController.wishlistUseCase.Create(wishlist)
	if err != nil {
		return c.JSON(base.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	wishlistRes := wishlistReqRes.Wishlist{
		ID:         wishlist.ID,
		CreatedAt:  wishlist.CreatedAt,
		UpdatedAt:  wishlist.UpdatedAt,
		DeletedAt:  wishlist.DeletedAt,
		Title:      wishlist.Title,
		IsAchieved: wishlist.IsAchieved,
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Success Create Wishlists", wishlistRes))
}
