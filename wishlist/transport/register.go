package transport

import (
	"go-clean-arch-test/wishlist"

	"github.com/gin-gonic/gin"
)

func RegisterWishListEndpoints(router *gin.Engine, uc wishlist.UseCase) {
	h := NewHandler(uc)

	wishListGroup := router.Group("/wishlist")
	{
		wishListGroup.POST("", h.CreateWish)
		wishListGroup.GET("", h.GetAllWishes)
		wishListGroup.DELETE("", h.DeleteWishByID)
	}

}