package routes

import (
	"fmt"
	"gowebapp2/models"
	"gowebapp2/sessions"
	"gowebapp2/utils"
	"net/http"
)

func registerGetHandler(w http.ResponseWriter, r *http.Request) {
	message, alert := sessions.Flash(w, r)

	utils.ExecuteTemplate(w, "register.html", struct {
		Alert utils.Alert
	}{
		Alert: utils.NewAlert(message, alert),
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

	checkErrRegister(err, w, r)
}

func checkErrRegister(err error, w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")

	message := "Succesfully registered"

	if err != nil {
		switch err {
		case models.ErrRequiredFirstname,
			models.ErrRequiredLastname,
			models.ErrRequiredEmail,
			models.ErrRequiredPassword,
			models.ErrInvalidEmail:
			message = fmt.Sprintf("%s", err)
		default:
			utils.InternalServerError(w)
			return
		}

		session.Values["MESSAGE"] = message
		session.Values["ALERT"] = "danger"
		session.Save(r, w)

		http.Redirect(w, r, "/register", 302)

		return
	}

	session.Values["MESSAGE"] = message
	session.Values["ALERT"] = "success"
	session.Save(r, w)

	http.Redirect(w, r, "/login", 302)
}
