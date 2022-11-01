package server

import (
	"context"
	"go-clean-arch-test/wishlist"
	"go-clean-arch-test/wishlist/repository/localslice"
	"go-clean-arch-test/wishlist/transport"
	"go-clean-arch-test/wishlist/usecase"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

type App struct {
	httpServer *http.Server
	wishListUC wishlist.UseCase
}

func NewApp() *App{
	wishRepo := localslice.NewWishLocalSlice()
	wishUC := usecase.NewUseCase(wishRepo)

	return &App{
		wishListUC: wishUC,
	}
}

func (a *App) Run(port string) error{
	router := gin.Default()
	transport.RegisterWishListEndpoints(router, a.wishListUC)
	a.httpServer = &http.Server{
		Addr: ":" + port,
		Handler: router,
		ReadTimeout: 10* time.Second,
		WriteTimeout: 10* time.Second,
	}
	
	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}