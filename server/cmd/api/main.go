package main

import (
	"log"
	"github.com/code-guy21/PingUp/server/internal/app"
)

func main(){
	// Initialize app
	app, err := app.NewApp()

	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Start server
	if err:= app.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}