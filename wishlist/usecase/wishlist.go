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
	return nil
}

func (wluc *WishListUseCase) GetAllWishes(ctx context.Context) ([]*models.Wish, error){
	return nil, nil
}

func (wluc *WishListUseCase) DeleteWish(ctx context.Context, id int) error {
	return nil
}