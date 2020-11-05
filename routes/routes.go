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

	fileServer := http.FileServer(http.Dir("./assets/"))

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
