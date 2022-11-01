package wishlist

import (
	"context"
	"go-clean-arch-test/models"
)

type UseCase interface {
	CreateWish(ctx context.Context, title string, text string) error 
	GetAllWishes(ctx context.Context) ([]*models.Wish, error)
	DeleteWishByID(ctx context.Context, id string) error
}
