package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

// LoginHandler verifies login data and creating sessions for users.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize the fields that we need in the custom struct.
	type Context struct {
		Err         string
		ErrExists   bool
		OTPRequired bool
		Username    string
		Password    string
	}
	// Call the Context struct.
	c := Context{}

	// If the request method is POST
	if r.Method == "POST" {
		// This is a login request from the user.
		username := r.PostFormValue("username")
		username = strings.TrimSpace(strings.ToLower(username))
		password := r.PostFormValue("password")
		otp := r.PostFormValue("otp")

		// Login2FA login using username, password and otp for users with OTPRequired = true.
		session := uadmin.Login2FA(r, username, password, otp)

		// Check whether the session returned is nil or the user is not active.
		if session == nil || !session.User.Active {
			/* Assign the login validation here that will be used for UI displaying. ErrExists and
			Err fields are coming from the Context struct. */
			c.ErrExists = true
			c.Err = "Invalid username/password or inactive user"

		} else {
			// If the user has OTPRequired enabled, it will print the username and OTP in the terminal.
			if session.PendingOTP {
				uadmin.Trail(uadmin.INFO, "User: %s OTP: %s", session.User.Username, session.User.GetOTP())
			}

			/* As long as the username and password is valid, it will create a session cookie in the
			browser. */
			cookie, _ := r.Cookie("session")
			if cookie == nil {
				cookie = &http.Cookie{}
			}
			cookie.Name = "session"
			cookie.Value = session.Key
			cookie.Path = "/"
			cookie.SameSite = http.SameSiteStrictMode
			http.SetCookie(w, cookie)

			// Check for OTP
			if session.PendingOTP {
				/* After the user enters a valid username and password in the first part of the form, these
				values will be used on the second part in the UI where the OTP input field will be
				displayed afterwards. */
				c.Username = username
				c.Password = password
				c.OTPRequired = true

			} else {
				// If the next value is empty, redirect the page that omits the logout keyword in the last part.
				if r.URL.Query().Get("next") == "" {
					http.Redirect(w, r, strings.TrimSuffix(r.RequestURI, "logout"), http.StatusSeeOther)
					return
				}

				// Redirect to the page depending on the value of the next.
				http.Redirect(w, r, r.URL.Query().Get("next"), http.StatusSeeOther)
				return
			}
		}
	}

	// Render the login filepath and pass the context data object to the HTML file.
	uadmin.RenderHTML(w, r, "templates/login.html", c)
}
