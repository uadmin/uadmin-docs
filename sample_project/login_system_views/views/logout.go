package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

// LogoutHandler handles the logout process for the user.
func LogoutHandler(w http.ResponseWriter, r *http.Request, session *uadmin.Session) {
	// Log out the user.
	uadmin.Logout(r)

	// Expire all cookies on logout by setting MaxAge to be less than 0.
	for _, cookie := range r.Cookies() {
		c := &http.Cookie{
			Name:   cookie.Name,
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}

		http.SetCookie(w, c)
	}
	http.Redirect(w, r, "/login_system/", http.StatusSeeOther)
}
