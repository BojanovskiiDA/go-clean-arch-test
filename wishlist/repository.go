package wishlist

import (
	"context"
	"go-clean-arch-test/go-clean-arch-test/models"
)

type Repository interface {
	CreateWish(ctx context.Context, wish *models.Wish) error
	GetAllWishes(ctx context.Context) ([]*models.Wish, error)
	DeleteWish(ctx context.Context, wish *models.Wish) error
}
