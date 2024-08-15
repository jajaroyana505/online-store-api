package main

import (
	"log"
	"net/http"
	"online-store/models"
	"online-store/routes"
)

func main() {
	models.ConnectDatabase()
	r := routes.InitRoutes()

	log.Fatal(http.ListenAndServe(":5000", r))
}
