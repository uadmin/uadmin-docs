Login System Tutorial Part 5 - Session and Cookie (Current Progress)
====================================================================
`Back to Previous Page`_

.. _Back to Previous Page: https://uadmin-docs.readthedocs.io/en/latest/login_system/tutorial/part5.html

Structure:

* `templates`_
    * `login.html`_
* `views`_
    * `login.go`_
* `main.go`_

.. _templates: https://uadmin-docs.readthedocs.io/en/latest/login_system/tutorial/full_code/part5.html#id1
.. _login.html: https://uadmin-docs.readthedocs.io/en/latest/login_system/tutorial/full_code/part5.html#id2
.. _views: https://uadmin-docs.readthedocs.io/en/latest/login_system/tutorial/full_code/part5.html#id3
.. _login.go: https://uadmin-docs.readthedocs.io/en/latest/login_system/tutorial/full_code/part5.html#id4
.. _main.go: https://uadmin-docs.readthedocs.io/en/latest/login_system/tutorial/full_code/part5.html#id5

templates
---------
**login.html**
^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: html

    <!DOCTYPE html>
    <html lang="en">
    <head>
      <meta charset="UTF-8">
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <meta http-equiv="X-UA-Compatible" content="ie=edge">
      <title>Login Form</title>
    </head>
    <body>
      <form method="POST">
        <!-- The assigned name attribute is equivalent to r.FormValue in
        Golang while the assigned value attribute is the value of the
        r.FormValue. (e.g. r.FormValue("request") = "login") -->
        <input type="text" name="username" placeholder="Username"><br>
        <input type="password" name="password" placeholder="Password"><br>
        <input type="text" name="otp_pass" placeholder="OTP Password"><br><br>
        <button type="submit" name="request" value="login">Login</button><br>
      </form>
    </body>
    </html>

views
-----
**login.go**
^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

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
        // r.URL.Path creates a new path called "/login/"
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
                        }

                        // Check whether the OTP value from Login function is false
                        // and the OTP Password is not assigned
                        if otp == false && otpPass == "" {
                            // Set the "user_session" cookie to the IP Address
                            http.SetCookie(w, usersession)
                        }
                    }
                }
            }
        }

        // Pass the userContext data object to the HTML file
        uadmin.RenderHTML(w, r, "templates/login.html", userContext)
    }

main.go
-------
`Back to Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/login_system/tutorial/full_code/part5.html#login-system-tutorial-part-5-session-and-cookie-current-progress

.. code-block:: go

    package main

    import (
        "net/http"

        // Specify the username that you used inside github.com folder
        "github.com/username/login_system/views"
        "github.com/uadmin/uadmin"
    )

    func main() {
        // Assign RootURL value as "/admin/" and Site Name as "Login System"
        // NOTE: This code works only on first build.
        uadmin.RootURL = "/admin/"
        uadmin.SiteName = "Login System"

        // Login Handler
        http.HandleFunc("/login/", uadmin.Handler(views.LoginHandler))

        // Run the server
        uadmin.StartServer()
    }
