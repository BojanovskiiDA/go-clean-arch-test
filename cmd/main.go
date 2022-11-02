package main

import (
	"go-clean-arch-test/server"
	"log"
)

func main() {
	
	app := server.NewApp()

	if err := app.Run(); err != nil{
		log.Fatalf("%s", err.Error())
	}

}