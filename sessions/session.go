package sessions

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("S3CR3TK3Y"))

func Flash(w http.ResponseWriter, r *http.Request) (string, string) {
	message, alert := "", ""

	session, _ := Store.Get(r, "session")

	untypedMessage := session.Values["MESSAGE"]
	message, ok := untypedMessage.(string)
	if !ok {
		return "", ""
	}

	untypedAlert := session.Values["ALERT"]
	alert, ok = untypedAlert.(string)
	if !ok {
		return "", ""
	}

	delete(session.Values, "MESSAGE")
	delete(session.Values, "ALERT")

	session.Save(r, w)

	return message, alert
}
