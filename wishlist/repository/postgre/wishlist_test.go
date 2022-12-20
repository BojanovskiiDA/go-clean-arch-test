package postgre

import (
	"context"
	"go-clean-arch-test/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAllWishesOk(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal("Mock was not created", err)
	}
	defer db.Close()

	wishPG := NewWishPG(db)

	rows := sqlmock.NewRows([]string{"id", "title", "text"})
	expect := []models.Wish{
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

	rows.AddRow(expect[0].ID, expect[0].WishTitle, expect[0].WishText)
	rows.AddRow(expect[1].ID, expect[1].WishTitle, expect[1].WishText)

	mock.
		ExpectQuery("SELECT * FROM wishes").
		WillReturnRows(rows)

	wishes, err := wishPG.GetAllWishes(context.Background())
	if err != nil {
		t.Error("Error in GetAllWishesMethod", err)
		return
	}

	if err = mock.ExpectationsWereMet(); err != nil{
		t.Errorf("Expectations were not meer: #{err}")
		return
	}
	 if wishes[0].ID != expect[0].ID || wishes[0].WishTitle != expect[0].WishTitle || wishes[0].WishText != expect[0].WishText ||
	 	wishes[1].ID != expect[1].ID || wishes[1].WishTitle != expect[1].WishTitle || wishes[1].WishText != expect[1].WishText {
			t.Errorf("result not match, want %v, get %v", wishes, expect)
			return
		}

	
	
}
