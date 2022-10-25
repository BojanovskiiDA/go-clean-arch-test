package localslice

import (
	"go-clean-arch-test/go-clean-arch-test/models"
	"sync"
	"context"
)

type WishLocalSlice struct {
	Wishlist []*models.Wish
	Mutex    *sync.Mutex
}

func NewWishLocalSlice() *WishLocalSlice {
	return &WishLocalSlice{Wishlist: make([]*models.Wish, 10),
		Mutex: new(sync.Mutex),
	}
}


func (w *WishLocalSlice) CreateWish(ctx context.Context, wish *models.Wish) error {
	w.Mutex.Lock()
	w.Wishlist = append(w.Wishlist, wish)
	w.Mutex.Unlock()
	return nil
}

func (w *WishLocalSlice) GetAllWishes(ctx context.Context) ([]*models.Wish, error) {
	return w.Wishlist, nil
}

func (w *WishLocalSlice) DeleteWish(ctx context.Context, wish *models.Wish) error {
	w.Mutex.Lock()
	for i, wishInRepo := range w.Wishlist {
		if wishInRepo.ID == wish.ID {
			w.Wishlist = append(w.Wishlist[:i], w.Wishlist[i+1:]...)
			break
		}
	}
	w.Mutex.Unlock()
	return nil
}