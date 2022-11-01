package usecase

import (
	"context"
	"go-clean-arch-test/models"
	"go-clean-arch-test/wishlist"
)
type WishListUseCase struct {
	WishlistRepository wishlist.Repository
	
}

func NewUseCase (wlr wishlist.Repository) *WishListUseCase{
	return &WishListUseCase{
		WishlistRepository: wlr,
	}
}



func (wluc *WishListUseCase) CreateWish(ctx context.Context, title string, text string) error {
	wish := &models.Wish{
		WishTitle: title,
		WishText: text,
	}
	return wluc.WishlistRepository.CreateWish(ctx, wish)
}

func (wluc *WishListUseCase) GetAllWishes(ctx context.Context) ([]*models.Wish, error){
	return wluc.WishlistRepository.GetAllWishes(ctx)
}

func (wluc *WishListUseCase) DeleteWishByID(ctx context.Context, id string) error {
	wish := &models.Wish{
		ID: id,
	}
	return wluc.WishlistRepository.DeleteWishByID(ctx, wish)
}