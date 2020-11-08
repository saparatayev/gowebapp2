package routes

import (
	"gowebapp2/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", homeGetHandler).Methods("GET")
	r.HandleFunc("/", homePostHandler).Methods("POST")
	r.HandleFunc("/register", registerGetHandler).Methods("GET")
	r.HandleFunc("/register", registerPostHandler).Methods("POST")
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
	r.HandleFunc("/admin", middleware.AuthRequired(adminGetHandler)).Methods("GET")
	r.HandleFunc("/logout", middleware.AuthRequired(logoutGetHandler)).Methods("GET")
	r.HandleFunc("/products", middleware.AuthRequired(productsGetHandler)).Methods("GET")
	r.HandleFunc("/products/create", middleware.AuthRequired(productsCreateGetHandler)).Methods("GET")
	r.HandleFunc("/products/create", middleware.AuthRequired(productsCreatePostHandler)).Methods("POST")
	r.HandleFunc("/product/edit", middleware.AuthRequired(productEditGetHandler)).Methods("GET")

	fileServer := http.FileServer(http.Dir("./assets/"))

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
