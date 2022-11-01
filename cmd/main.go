package main

import (
	"go-clean-arch-test/server"
)

func main() {
	app := server.NewApp()
	app.Run("8080")
}