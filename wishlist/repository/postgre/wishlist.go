package postgre

import (
	"context"
	"database/sql"
	"go-clean-arch-test/models"

	_ "github.com/lib/pq"
)

type WishPG struct {
	DB *sql.DB
}

func NewWishPG(db *sql.DB) *WishPG {
	return &WishPG{
		DB: db,
	}

}

func (wpg *WishPG) CreateWish(ctx context.Context, wish *models.Wish) error {
	_, err := wpg.DB.Exec("INSERT INTO wishes (title, text) VALUES ($1, $2)", wish.WishText, wish.WishTitle)
	return err
}

func (wpg *WishPG) GetAllWishes(ctx context.Context) ([]*models.Wish, error) {
	rows, err := wpg.DB.Query("SELECT * FROM wishes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var outList []*models.Wish
	for rows.Next() {
		wish := &models.Wish{}
		err = rows.Scan(&wish.ID, &wish.WishTitle, &wish.WishText)
		if err != nil {
			return nil, err
		}
		outList = append(outList, wish)
	}
	return outList, nil
}

func (wpg *WishPG) DeleteWishByID(ctx context.Context, wish *models.Wish) error {
	_, err := wpg.DB.Exec("DELETE FROM wishes WHERE id = $1", wish.ID)
	return err
}
