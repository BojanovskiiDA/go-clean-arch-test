package main

import (
	"go-clean-arch-test/config"
	"go-clean-arch-test/server"
	"log"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatalf("%s", err)
	}

	app, err := server.NewApp()
	if err != nil {
		log.Fatal(err)
	}
	if err := app.Run(); err != nil{
		log.Fatalf("%s", err.Error())
	}

}