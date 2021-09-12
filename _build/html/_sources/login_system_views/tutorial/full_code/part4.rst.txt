Login System Tutorial Part 4 - Login Access Debugging (Current Progress)
========================================================================
`Back to Previous Page`_

.. _Back to Previous Page: https://uadmin-docs.readthedocs.io/en/latest/login_system_views/tutorial/part4.html

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
          <input type="text" name="username" placeholder="Username" value="{{.Username}}"><br>
          <input type="password" name="password" placeholder="Password" value="{{.Password}}"><br>

          <!-- The OTP field will be displayed if the user has OTPRequired
            enabled in the database and after the user enters a valid username and
            password.-->
          {{if .OTPRequired}}
          <input type="text" name="otp" placeholder="OTP"><br>
          {{end}}

          <br><button type="submit">Login</button><br>
        </form>

        {{if .ErrExists}}
        <p>
          <b style="color:red">{{.Err}}</b>
        </p>
        {{end}}
      </body>
    </html>


views
-----
**login.go**
^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

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
                    uadmin.Trail(uadmin.DEBUG, "Your login credentials are valid.")
                }
            }
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

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/login_system_views/tutorial/full_code/part4.html#login-system-tutorial-part-4-login-access-debugging-current-progress

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
