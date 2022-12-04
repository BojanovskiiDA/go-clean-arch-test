package transport

import (
	"go-clean-arch-test/models"
	"go-clean-arch-test/wishlist"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase wishlist.UseCase
}

func NewHandler(uc wishlist.UseCase) *Handler {
	return &Handler{
		useCase: uc,
	}
}

type CreateWishInput struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func (h *Handler) CreateWish(c *gin.Context) {
	input := new(CreateWishInput)
	if err := c.BindJSON(input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err:= h.useCase.CreateWish(c,input.Title,input.Text); err != nil{
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

type GetAllWishesOut struct{
	Wishes []*models.Wish `json:"wishes"`
}

func (h *Handler) GetAllWishes(c *gin.Context) {
	wishes, err := h.useCase.GetAllWishes(c)
	if err != nil{
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &GetAllWishesOut{Wishes: wishes})
}

type DeleteWishByIDInput struct{
	Id string `json:"id"`
}

func (h *Handler) DeleteWishByID(c *gin.Context) {
	input:= new(DeleteWishByIDInput)
	if err := c.BindJSON(input); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := h.useCase.DeleteWishByID(c, input.Id); err != nil{
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
