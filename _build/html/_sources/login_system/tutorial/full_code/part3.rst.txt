Login System Tutorial Part 3 - Sending Request (Current Progress)
=================================================================
`Back to Previous Page`_

.. _Back to Previous Page: https://uadmin-docs.readthedocs.io/en/latest/login_system/tutorial/part3.html

Structure:

* `templates`_
    * `login.html`_
* `views`_
    * `login.go`_
* `main.go`_

.. _templates: https://uadmin-docs.readthedocs.io/en/latest/login_system/tutorial/full_code/part3.html#id1
.. _login.html: https://uadmin-docs.readthedocs.io/en/latest/login_system/tutorial/full_code/part3.html#id2
.. _views: https://uadmin-docs.readthedocs.io/en/latest/login_system/tutorial/full_code/part3.html#id3
.. _login.go: https://uadmin-docs.readthedocs.io/en/latest/login_system/tutorial/full_code/part3.html#id4
.. _main.go: https://uadmin-docs.readthedocs.io/en/latest/login_system/tutorial/full_code/part3.html#id5

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

                // Print the result
                uadmin.Trail(uadmin.DEBUG, "Username: %s", username)

                // Check if the fetched record is existing in the User model
                if user.ID > 0 {
                    // Create the parameters of "password" and "otp_pass"
                    password := r.FormValue("password")
                    otpPass := r.FormValue("otp_pass")

                    // Print results
                    uadmin.Trail(uadmin.DEBUG, "Password: %s", password)
                    uadmin.Trail(uadmin.DEBUG, "OTP Password: %s", otpPass)
                }
            }
        }

        // Pass the userContext data object to the HTML file
        uadmin.RenderHTML(w, r, "templates/login.html", userContext)
    }

main.go
-------
`Back to Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/login_system/tutorial/full_code/part3.html#login-system-tutorial-part-3-sending-request-current-progress

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
