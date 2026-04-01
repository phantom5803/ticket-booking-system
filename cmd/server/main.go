package main

import (
	"context"
	"log"

	"airline-booking/internal/config"
	"airline-booking/internal/platform/db"
	"airline-booking/internal/repository/postgres"
	"airline-booking/internal/transport/http/router"
)

func main() {
	cfg := config.Load()

	pool, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	if err := db.MigrateAndSeed(context.Background(), pool); err != nil {
		log.Fatal(err)
	}

	flightRepo := postgres.NewFlightRepository(pool)
	bookingRepo := postgres.NewBookingRepository(pool)
	server := router.New(flightRepo, bookingRepo)

	log.Printf("server running on :%s", cfg.Port)
	if err := server.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
