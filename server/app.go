package server

import (
	"context"
	"database/sql"
	"fmt"
	"go-clean-arch-test/wishlist"

	//"go-clean-arch-test/wishlist/repository/localslice"
	"go-clean-arch-test/wishlist/repository/memcache"
	//"go-clean-arch-test/wishlist/repository/postgre"
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

func NewApp() (*App, error) {
	// db, err := initPGDB()
	// if err != nil {
	// 	return nil, err
	// }
	//wishRepo := postgre.NewWishPG(db)
	
	//need to remake for config file
	wishRepo := memcache.NewWishMemCache("127.0.0.1:49153")
	err := wishRepo.MCClient.Ping()
	if err != nil {
		return nil, err
	}
	wishUC := usecase.NewUseCase(wishRepo)

	return &App{
		wishListUC: wishUC,
	}, err
}

func (a *App) Run() error {

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

func initPGDB() (*sql.DB, error){
	host:= viper.GetString("postgre.host")
	port:= viper.GetString("postgre.port")
	user:= viper.GetString("postgre.user")
	pswd:= viper.GetString("postgre.pswd")
	dbname:= viper.GetString("postgre.bdname")
	conStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",host, port, user, pswd, dbname) 

	db, err := sql.Open("postgres", conStr)
	if err!=nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)
	
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}