package wishlist

import (
	"context"
	"go-clean-arch-test/go-clean-arch-test/models"
)

type UseCase interface {
	CreateWish(ctx context.Context, title string, text string) error 
	GetAllWishes(ctx context.Context) ([]*models.Wish, error)
	DeleteWish(ctx context.Context, id int) error
}