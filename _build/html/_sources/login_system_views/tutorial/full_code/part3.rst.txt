Login System Tutorial Part 3 - Sending Request (Current Progress)
=================================================================
`Back to Previous Page`_

.. _Back to Previous Page: https://uadmin-docs.readthedocs.io/en/latest/login_system_views/tutorial/part3.html

Structure:

* `templates`_
    * `login.html`_
* `views`_
    * `login.go`_
    * `views_main.go`_
* `main.go`_

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
          <!-- The assigned name attribute is equivalent to r.PostFormValue in
          Golang (e.g. r.PostFormValue("username") = "admin") -->
          <input type="text" name="username" placeholder="Username"><br>
          <input type="password" name="password" placeholder="Password"><br>
          <input type="text" name="otp" placeholder="OTP"><br>

          <br><button type="submit">Login</button><br>
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

            // Display the results here.
            uadmin.Trail(uadmin.DEBUG, "Session: %s", session)
        }

        // Render the login filepath and pass the context data object to the HTML file.
        uadmin.RenderHTML(w, r, "templates/login.html", c)
    }

**views_main.go**
^^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package views

    import (
        "net/http"
        "strings"
    )

    // MainHandler is the main handler for the login system.
    func MainHandler(w http.ResponseWriter, r *http.Request) {
        // r.URL.Path creates a new path called "/login_system/"
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/login_system")
        r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

        // LoginHandler verifies login data and creating sessions for users.
        LoginHandler(w, r)
    }


main.go
-------
`Back to Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/login_system_views/tutorial/full_code/part3.html#login-system-tutorial-part-3-sending-request-current-progress

.. code-block:: go

    package main

    import (
        "net/http"

        // Specify the username that you used inside github.com folder
        "github.com/username/project_name/views"

        "github.com/uadmin/uadmin"
    )

    func main() {
        // Assign RootURL value as "/admin/" and Site Name as "Login System"
        // NOTE: This code works only on first build.
        uadmin.RootURL = "/admin/"
        uadmin.SiteName = "Login System"

        // Login System Main Handler
        http.HandleFunc("/login_system/", uadmin.Handler(views.MainHandler))

        // Run the server
        uadmin.StartServer()
    }
