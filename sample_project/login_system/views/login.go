<<<<<<< HEAD
package views

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

// UserContext !
type UserContext struct {
	User    *uadmin.User
	OTP     bool
	Message string
}

// LoginHandler !
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called /login/
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/login")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

	// Initialize the UserContext struct that we have created
	userContext := UserContext{}

	// Initialize the User model from uAdmin
	user := uadmin.User{}

	// Check if the user submits request in HTML form
	if r.Method == "POST" {
		// Check if the value of the request is login
		if r.FormValue("request") == "login" {
			// Create the parameter of "username"
			username := r.FormValue("username")

			// Get the user record where username is the assigned value
			uadmin.Get(&user, "username=?", username)

			// Check if the fetched record is existing in the User model
			if user.ID > 0 {
				// Create the parameters of "password" and "otp_pass"
				password := r.FormValue("password")
				otpPass := r.FormValue("otp_pass")

				// Pass the requested username and password in Login function to
				// return the session and the boolean value for IsOTPRequired
				login, otp := uadmin.Login(r, username, password)

				// Initialize Login2FA that returns the Session
				login2fa := &uadmin.Session{}

				// Check whether the OTP value from Login function is true
				// and the OTP Password is valid
				if otp == true && user.VerifyOTP(otpPass) {
					// Pass the requested username, password, and OTP Password in
					// Login2FA function to return the session
					login2fa = uadmin.Login2FA(r, username, password, otpPass)

					// Print the result
					uadmin.Trail(uadmin.DEBUG, "Login with 2FA as: %s", login2fa.User)
				}

				// Check if the session is fetched from either login or login2fa
				if login != nil {
					// Create a cookie named "user_session" with the value of
					// UserID
					usersession := &http.Cookie{
						Name:  "user_session",
						Value: fmt.Sprint(user.ID),
					}

					// Check whether the OTP value from Login function is true
					// and the OTP Password is valid
					if otp == true && user.VerifyOTP(otpPass) {
						// Set the "user_session" cookie to the IP Address
						http.SetCookie(w, usersession)

						// Assign the full name of the user and OTP boolean value to the
						// userContext
						userContext = UserContext{
							User: &login2fa.User,
							OTP:  otp,
						}
					}

					// Check whether the OTP value from Login function is false
					// and the OTP Password is not assigned
					if otp == false && otpPass == "" {
						// Set the "user_session" cookie to the IP Address
						http.SetCookie(w, usersession)

						// Assign the full name of the user and OTP boolean value to the
						// userContext
						userContext = UserContext{
							User: &login.User,
							OTP:  otp,
						}
					}

					// Pass the userContext data object to the HTML file
					uadmin.RenderHTML(w, r, "templates/home.html", userContext)
					return
				}
			}
		}

		// Check if the request submitted is logout
		if r.FormValue("request") == "logout" {
			// Assign the message to the Message field of userContext
			userContext.Message = "User has logged out."

			// Logout the user in uAdmin
			uadmin.Logout(r)

			// Deletes the cookie
			usersession := &http.Cookie{
				Name:   "user_session",
				Value:  "",
				MaxAge: -1,
			}
			http.SetCookie(w, usersession)

			// Pass the userContext data object to the HTML file
			uadmin.RenderHTML(w, r, "templates/login.html", userContext)
			return
		}
	}

	// Read the cookie of "user_session"
	cookie, _ := r.Cookie("user_session")

	// Check if the fetched cookie is existing
	if cookie != nil {
		// Get the user record based on the value of the cookie
		uadmin.Get(&user, "id = ?", cookie.Value)

		// Assign the full name of the user and OTP boolean value to the
		// userContext
		userContext = UserContext{
			User: &user,
			OTP:  user.OTPRequired,
		}

		// Pass the userContext data object to the HTML file
		uadmin.RenderHTML(w, r, "templates/home.html", userContext)
		return
	}

	// Pass the userContext data object to the HTML file
	uadmin.RenderHTML(w, r, "templates/login.html", userContext)
}
=======
package views

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

// UserContext !
type UserContext struct {
	User    *uadmin.User
	OTP     bool
	Message string
}

// LoginHandler !
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called /login/
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/login")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

	// Initialize the UserContext struct that we have created
	userContext := UserContext{}

	// Initialize the User model from uAdmin
	user := uadmin.User{}

	// Check if the user submits request in HTML form
	if r.Method == "POST" {
		// Check if the value of the request is login
		if r.FormValue("request") == "login" {
			// Create the parameter of "username"
			username := r.FormValue("username")

			// Get the user record where username is the assigned value
			uadmin.Get(&user, "username=?", username)

			// Check if the fetched record is existing in the User model
			if user.ID > 0 {
				// Create the parameters of "password" and "otp_pass"
				password := r.FormValue("password")
				otpPass := r.FormValue("otp_pass")

				// Pass the requested username and password in Login function to
				// return the session and the boolean value for IsOTPRequired
				login, otp := uadmin.Login(r, username, password)

				// Initialize Login2FA that returns the Session
				login2fa := &uadmin.Session{}

				// Check whether the OTP value from Login function is true
				// and the OTP Password is valid
				if otp == true && user.VerifyOTP(otpPass) {
					// Pass the requested username, password, and OTP Password in
					// Login2FA function to return the session
					login2fa = uadmin.Login2FA(r, username, password, otpPass)

					// Print the result
					uadmin.Trail(uadmin.DEBUG, "Login with 2FA as: %s", login2fa.User)
				}

				// Check if the session is fetched from either login or login2fa
				if login != nil {
					// Create a cookie named "user_session" with the value of
					// UserID
					usersession := &http.Cookie{
						Name:  "user_session",
						Value: fmt.Sprint(user.ID),
					}

					// Check whether the OTP value from Login function is true
					// and the OTP Password is valid
					if otp == true && user.VerifyOTP(otpPass) {
						// Set the "user_session" cookie to the IP Address
						http.SetCookie(w, usersession)

						// Assign the full name of the user and OTP boolean value to the
						// userContext
						userContext = UserContext{
							User: &login2fa.User,
							OTP:  otp,
						}
					}

					// Check whether the OTP value from Login function is false
					// and the OTP Password is not assigned
					if otp == false && otpPass == "" {
						// Set the "user_session" cookie to the IP Address
						http.SetCookie(w, usersession)

						// Assign the full name of the user and OTP boolean value to the
						// userContext
						userContext = UserContext{
							User: &login.User,
							OTP:  otp,
						}
					}

					// Pass the userContext data object to the HTML file
					uadmin.RenderHTML(w, r, "templates/home.html", userContext)
					return
				}
			}
		}

		// Check if the request submitted is logout
		if r.FormValue("request") == "logout" {
			// Assign the message to the Message field of userContext
			userContext.Message = "User has logged out."

			// Logout the user in uAdmin
			uadmin.Logout(r)

			// Deletes the cookie
			usersession := &http.Cookie{
				Name:   "user_session",
				Value:  "",
				MaxAge: -1,
			}
			http.SetCookie(w, usersession)

			// Pass the userContext data object to the HTML file
			uadmin.RenderHTML(w, r, "templates/login.html", userContext)
			return
		}
	}

	// Read the cookie of "user_session"
	cookie, _ := r.Cookie("user_session")

	// Check if the fetched cookie is existing
	if cookie != nil {
		// Get the user record based on the value of the cookie
		uadmin.Get(&user, "id = ?", cookie.Value)

		// Assign the full name of the user and OTP boolean value to the
		// userContext
		userContext = UserContext{
			User: &user,
			OTP:  user.OTPRequired,
		}

		// Pass the userContext data object to the HTML file
		uadmin.RenderHTML(w, r, "templates/home.html", userContext)
		return
	}

	// Pass the userContext data object to the HTML file
	uadmin.RenderHTML(w, r, "templates/login.html", userContext)
}
>>>>>>> de25cdd8a29ca2bb2c2df08be00b703b967aaed5
