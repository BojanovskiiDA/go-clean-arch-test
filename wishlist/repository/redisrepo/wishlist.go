package redisrepo

import (
	"context"
	"encoding/json"
	//"fmt"
	"go-clean-arch-test/models"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)


type redModelWish struct {
	WishTitle string `json:"Title"`
	WishText  string `json:"Text"`
}

type WishRedis struct {
	DB redis.Conn
}

func NewWishRedis(db redis.Conn) *WishRedis {
	return &WishRedis{
		DB: db,
	}

}

func (wr *WishRedis) CreateWish(ctx context.Context, wish *models.Wish) error {
	key := uuid.NewString()
	mcWish := modelToMC(wish)
	val, err := json.Marshal(mcWish)
	if err != nil {
		log.Println(err)
		return err
	}
	ok, err := redis.String(wr.DB.Do("SET", key, val))
	if ok != "OK" {return err}
	return nil
}

func (wr *WishRedis) GetAllWishes(ctx context.Context) ([]*models.Wish, error) {
	keysSlice, err := redis.ByteSlices(wr.DB.Do("KEYS", "*"))
	if err!= nil {
		return nil, err
	}

	var outList []*models.Wish
	var redMW redModelWish
	for _, v := range keysSlice {
		val, err := redis.Bytes(wr.DB.Do("GET", v))
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(val, &redMW)
		wish := mcToModel(&redMW)
		wish.ID = string(v)
		outList = append(outList, &wish)
	}

	return outList, nil
}

func (wr *WishRedis) DeleteWishByID(ctx context.Context, wish *models.Wish) error {
	_, err := wr.DB.Do("DEL", wish.ID)
	return err
}


func modelToMC(wish *models.Wish) redModelWish {
	return redModelWish{
		WishTitle: wish.WishTitle,
		WishText:  wish.WishText,
	}
}

func mcToModel (mcWish *redModelWish) models.Wish{
	return models.Wish{
		WishTitle: mcWish.WishTitle,
		WishText: mcWish.WishText,
	}
}