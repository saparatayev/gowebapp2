package routes

import (
	"gowebapp2/models"
	"gowebapp2/sessions"
	"gowebapp2/utils"
	"net/http"
)

func registerGetHandler(w http.ResponseWriter, r *http.Request) {
	message := sessions.Flash(w, r)

	utils.ExecuteTemplate(w, "register.html", struct {
		Message string
	}{
		Message: message,
	})
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

	session, _ := sessions.Store.Get(r, "session")
	session.Values["MESSAGE"] = "Succesfully registered"
	session.Save(r, w)

	http.Redirect(w, r, "/register", 302)

	// utils.ToJson(w, user)
}
