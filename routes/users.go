package routes

import (
	"gowebapp2/models"
	"gowebapp2/utils"
	"net/http"
)

func registerGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}

func registerPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var user models.User

	user.Firstname = r.PostForm.Get("firstname")
	user.Lastname = r.PostForm.Get("lastname")
	user.Email = r.PostForm.Get("email")
	user.Password = r.PostForm.Get("password")

	_, err := models.NewUser(user)
	if err != nil {
		utils.InternalServerError(w)
		return
	}

	http.Redirect(w, r, "/register", 302)

	// utils.ToJson(w, user)
}
