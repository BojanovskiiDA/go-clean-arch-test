package usecase

import (
	"context"
	"go-clean-arch-test/models"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetAllWishesOk(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockRepository(ctrl)
	usecase := NewUseCase(repo)
	
	expect := []*models.Wish{
		{
			ID:        "1",
			WishTitle: "my 1 wish title",
			WishText:  "my 1 wish text",
		},
		{
			ID:        "2",
			WishTitle: "my 2 wish title",
			WishText:  "my 2 wish text",
		},
	}
	
	repo.EXPECT().GetAllWishes(context.Background()).Return(expect, nil)
	//repo.EXPECT().GetAllWishes().Return(expect, nil)
	wishes, err := usecase.GetAllWishes(context.Background())

	if err != nil {
		t.Errorf("failed on GetAllWishes: %s", err)
		return
	}

	if 	wishes[0].ID != expect[0].ID || wishes[0].WishTitle != expect[0].WishTitle || wishes[0].WishText != expect[0].WishText ||
		wishes[1].ID != expect[1].ID || wishes[1].WishTitle != expect[1].WishTitle || wishes[1].WishText != expect[1].WishText {
			t.Errorf("result not match, want %v, get %v", wishes, expect)
			return
		}
	
}
