package main

import (
	"Basicgo/internal/controllers"
	"Basicgo/internal/seed"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var server = controllers.Server{}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		log.Println(" We are getting the env success")
	}

	// initialize server connect database
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	//... excute seeding data
	seed.Load(server.DB)

	// Run server api controller
	server.Run(":8080")
}