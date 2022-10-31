package main

import (
	"context"
	"errors"
	"fmt"
	"go-clean-arch-test/models"
	"go-clean-arch-test/wishlist/repository/localslice"
)

func main() {
	wls := localslice.NewWishLocalSlice()
	wish := &models.Wish{
		ID: "0",
		WishTitle: "title 0",
		WishText: "text 0",
	}
	ctx := context.Background()
	err := wls.CreateWish(ctx, wish)
	fmt.Println(err)
	//assert.NoError(t, err)
	gwl, err := wls.GetAllWishes(ctx)
	fmt.Println(err)
	//assert.NoError(t, err)
	fmt.Println(gwl[0])
	if gwl[0].ID != wish.ID && gwl[0].WishText != wish.WishText && gwl[0].WishTitle != wish.WishTitle{
		err = errors.New("returned wish is not equal to input")
		fmt.Println(err)
	} 
	//assert.NoError(t, err)
}