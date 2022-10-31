package localslice

import (
	"context"
	"go-clean-arch-test/models"
	"strconv"
	"sync"
)

type Wish struct{
	WishTitle string
	WishText string
}

type WishLocalSlice struct {
	Wishlist []*Wish
	Mutex    *sync.Mutex
}

func NewWishLocalSlice() *WishLocalSlice {
	return &WishLocalSlice{Wishlist: make([]*Wish, 10),
		Mutex: new(sync.Mutex),
	}
}


func (w *WishLocalSlice) CreateWish(ctx context.Context, wish *models.Wish) error {
	w.Mutex.Lock()
	localWish, _, _ := toLocalWish(wish)
	w.Wishlist = append(w.Wishlist, localWish)
	w.Mutex.Unlock()
	return nil
}

func (w *WishLocalSlice) GetAllWishes(ctx context.Context) (gwl []*models.Wish, err error) {
	for i, v := range w.Wishlist {
		lw := v
		gw := toGlobalWish(lw, i)
		gwl = append(gwl, gw)
	}
	return gwl, err
}

func (w *WishLocalSlice) DeleteWish(ctx context.Context, gw *models.Wish) error {
	_, id, _ := toLocalWish(gw)
	w.Mutex.Lock()
	for i := range w.Wishlist {
		if i == id {
			w.Wishlist = append(w.Wishlist[:i], w.Wishlist[i+1:]...)
			break
		}
	}
	w.Mutex.Unlock()
	return nil
}

func toGlobalWish(lw* Wish, id int) (gb *models.Wish){
	gb.ID = strconv.Itoa(id)
	gb.WishText = lw.WishText
	gb.WishTitle = lw.WishTitle
	return gb
}

func toLocalWish(gb *models.Wish) (lw *Wish, id int, err error){
	lw.WishText = gb.WishText
	lw.WishTitle = gb.WishTitle
	id, err = strconv.Atoi(gb.ID)
	if err != nil {
		return lw, 0, err
	}
	return lw, id, nil
}
