package main

import (
	"fmt"
	"gowebapp/models"
	"gowebapp/routes"
	"gowebapp/utils"
	"log"
	"net/http"
)

const PORT = ":8081"

func main() {

	models.TestConnection()

	fmt.Printf("Listening Port %s\n", PORT)

	utils.LoadTemplates("views/*.html")

	r := routes.NewRouter()

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(PORT, nil))
}
