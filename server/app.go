package server

import (
	"context"
	"fmt"
	"go-clean-arch-test/config"
	"go-clean-arch-test/wishlist"
	"go-clean-arch-test/wishlist/repository/localslice"
	"go-clean-arch-test/wishlist/transport"
	"go-clean-arch-test/wishlist/usecase"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type App struct {
	httpServer *http.Server
	wishListUC wishlist.UseCase
}

func NewApp() *App {
	wishRepo := localslice.NewWishLocalSlice()
	wishUC := usecase.NewUseCase(wishRepo)

	return &App{
		wishListUC: wishUC,
	}
}

func (a *App) Run() error {
	err := config.Init()
	if err != nil {
		log.Fatalf("%s", err)
	}
	settings := viper.AllSettings()
	port, ok := settings["port"].(int)
	if !ok {
		log.Fatal("Port is not found")
	}
	router := gin.Default()
	transport.RegisterWishListEndpoints(router, a.wishListUC)
	a.httpServer = &http.Server{
		Addr:         ":"+ strconv.Itoa(port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
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
