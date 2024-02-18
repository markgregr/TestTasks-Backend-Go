package main

import (
	"context"
	"log"

	"crypto-rates-client/internal/app"
)

func main() {
	log.Println("Application start!")
	// создает контекст
	ctx := context.Background()

	// создает Aplication
	application, err := app.New(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// запуск сервера
	application.Run()
	log.Println("Application terminated!")
}
