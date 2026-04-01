package main

import (
	"log"
	"os"

	"airline-booking/internal/api"
	"airline-booking/internal/store"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	flightStore := store.NewFlightStore()
	bookingStore := store.NewBookingStore()
	server := api.NewServer(flightStore, bookingStore)

	log.Printf("server running on :%s", port)
	if err := server.Routes().Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
