package localslice

import (
	"context"
	"errors"
	"fmt"
	"go-clean-arch-test/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateWish(t *testing.T) {
	wls := NewWishLocalSlice()
	wish := &models.Wish{
		ID:        "0",
		WishTitle: "title 0",
		WishText:  "text 0",
	}
	ctx := context.Background()
	err := wls.CreateWish(ctx, wish)
	assert.NoError(t, err)
	gwl, err := wls.GetAllWishes(ctx)
	assert.NoError(t, err)

	if gwl[0].ID != wish.ID && gwl[0].WishText != wish.WishText && gwl[0].WishTitle != wish.WishTitle {
		err = errors.New("returned wish is not equal to input")
	}
	assert.NoError(t, err)

}

func TestGetAllWishes(t *testing.T) {
	ctx := context.Background()
	wls := NewWishLocalSlice()
	for i := 0; i < 10; i++ {
		gw := &models.Wish{
			ID:        fmt.Sprintf("%d", i),
			WishTitle: fmt.Sprintf("title - %d", i),
			WishText:  fmt.Sprintf("text - %d", i),
		}
		err := wls.CreateWish(ctx, gw)
		assert.NoError(t, err)
	}
	gwl, err := wls.GetAllWishes(ctx)
	assert.NoError(t, err)
	assert.Equal(t, len(gwl), 10)

	for i, v := range gwl {
		assert.Equal(t, v.WishText, wls.Wishlist[i].WishText)
		assert.Equal(t, v.WishTitle, wls.Wishlist[i].WishTitle)
	}
}

func TestDeleteWishByID(t *testing.T){
	ctx := context.Background()
	wls := NewWishLocalSlice()
	for i := 0; i < 10; i++ {
		gw := &models.Wish{
			ID:        fmt.Sprintf("%d", i),
			WishTitle: fmt.Sprintf("title - %d", i),
			WishText:  fmt.Sprintf("text - %d", i),
		}
		err := wls.CreateWish(ctx, gw)
		assert.NoError(t, err)
	}
	gw := &models.Wish{
		ID: "9",
	}
	wls.DeleteWishByID(ctx, gw)
	gw.ID = "0"
	wls.DeleteWishByID(ctx, gw)
	gw.ID = "4"
	wls.DeleteWishByID(ctx, gw)
	assert.Equal(t, 7, len(wls.Wishlist))
}