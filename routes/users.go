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
	// session, _ := sessions.Store.Get(r, "session")

	message := "Succesfully registered"

	if err != nil {
		switch err {
		case models.ErrRequiredFirstname,
			models.ErrRequiredLastname,
			models.ErrRequiredEmail,
			models.ErrRequiredPassword,
			models.ErrInvalidEmail,
			models.ErrMaxLimit,
			models.ErrDuplicateKeyEmail:
			message = fmt.Sprintf("%s", err)
		default:
			utils.InternalServerError(w)
			fmt.Println(err)
			return
		}

		sessions.Message(message, "danger", w, r)

		http.Redirect(w, r, "/register", 302)

		return
	}

	sessions.Message(message, "success", w, r)

	http.Redirect(w, r, "/login", 302)
}

func usersGetHandler(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetUsers()
	if err != nil {
		utils.InternalServerError(w)
		return
	}

	total := int64(len(users))

	utils.ExecuteTemplate(w, "users.html", struct {
		Users []models.User
		Total int64
		// Alert utils.Alert
	}{
		Users: users,
		Total: total,
		// Alert: utils.NewAlert(message, alert),
	})
}
