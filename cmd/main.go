package main

import (
	"log"
	"speakbuddy/internal/bootstrap"
)

func main() {
	app := bootstrap.InitializeApp()
	log.Println("Server running on http://localhost:8080")
	app.Listen(":8080")
}
