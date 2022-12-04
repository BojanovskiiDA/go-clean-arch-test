package memcache

import (
	"context"
	"encoding/json"
	"fmt"
	"go-clean-arch-test/models"
	"log"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/google/uuid"
)

type mcModelWish struct {
	WishTitle string `json:"Title"`
	WishText  string `json:"Text"`
}

type WishMemCache struct {
	MCClient *memcache.Client
}

func NewWishMemCache(srv string) *WishMemCache {
	return &WishMemCache{
		MCClient: memcache.New([]string{srv}...),
	}
}

func (wmc *WishMemCache) CreateWish(ctx context.Context, wish *models.Wish) error {
	key := uuid.NewString()
	mcWish := modelToMC(wish)
	val, err := json.Marshal(mcWish)
	if err != nil {
		log.Println(err)
		return err
	}
	err = wmc.MCClient.Set(&memcache.Item{
		Key: key,
		Value: val,
		Expiration: 0,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("WishMemCache.CreateWish - ", key, string(val))
	key += ","
	err = wmc.MCClient.Append(&memcache.Item{
		Key: "AllKeys",
		Value: []byte(key),
		Expiration: 0,
	})
	if err != nil {
		log.Println(err)
		log.Println(wmc.MCClient.Add(&memcache.Item{
			Key: "AllKeys",
			Value: []byte(key),
			Expiration: 0,
		}))
	}
	return nil
}

func (wmc *WishMemCache) GetAllWishes(ctx context.Context) ([]*models.Wish, error) {
	items, err := wmc.MCClient.Get("AllKeys")
	if err != nil {
		log.Println(err)
	}
	keys := strings.Split(string(items.Value),",")
	keys = append(keys[:len(keys)-1],)
	kv, err := wmc.MCClient.GetMulti(keys)
	if err != nil {
		fmt.Println(err)
	}
	var outList []*models.Wish
	var mcMW mcModelWish

	for _, v := range kv {
		err = json.Unmarshal(v.Value, &mcMW)
		if err != nil {
			return nil, err
		}
		wish := mcToModel(&mcMW)
		wish.ID = v.Key
		outList = append(outList, &wish)
	}
	return outList, nil
}

func (wmc *WishMemCache) DeleteWishByID(ctx context.Context, wish *models.Wish) error {
	//need delete key from AllKeys
	return wmc.MCClient.Delete(wish.ID)	
	
}

func modelToMC(wish *models.Wish) mcModelWish {
	return mcModelWish{
		WishTitle: wish.WishTitle,
		WishText:  wish.WishText,
	}
}

func mcToModel (mcWish *mcModelWish) models.Wish{
	return models.Wish{
		WishTitle: mcWish.WishTitle,
		WishText: mcWish.WishText,
	}
}