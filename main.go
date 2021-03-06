package main

import (
	"fmt"
	"gowebapp2/models"
	"gowebapp2/routes"
	"gowebapp2/sessions"
	"gowebapp2/utils"
	"log"
	"net/http"
)

const PORT = ":8081"

func main() {

	models.TestConnection()

	fmt.Printf("Listening Port %s\n", PORT)

	utils.LoadTemplates("views/*.html")

	sessions.SessionOptions("localhost", "/", 3600, true)

	r := routes.NewRouter()

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(PORT, nil))
}
