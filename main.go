package main

import (
	"communication-app/connector"
	"communication-app/models"
	"communication-app/router"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return
		}
	}
	models.DbURL = os.Getenv("URL")

	connector.Init()

	fmt.Println("Starting communication app....")
	r := router.Router()

	err := http.ListenAndServe(os.Getenv("port"), r)
	if err != nil {
		log.Fatal("error while starting the app")
	}
}
