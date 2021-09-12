Login System Tutorial Part 6 - Logout (Current Progress)
========================================================
`Back to Previous Page`_

.. _Back to Previous Page: https://uadmin-docs.readthedocs.io/en/latest/login_system_views/tutorial/part6.html

Structure:

* `templates`_
    * `home.html`_
    * `login.html`_
* `views`_
    * `home.go`_
    * `login.go`_
    * `logout.go`_
    * `main.go (views)`_
* `main.go`_

templates
---------
**home.html**
^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: html

    <!DOCTYPE html>
    <html lang="en">
      <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">
        <title>Home Page</title>
      </head>
      <body>
        <!-- .User is a field that came from the HomeHandler in Golang -->
        <h1>Login as {{.User}}

        <!-- Validate if the OTP is enabled in the user. Use "with" if the
        user has OTPRequired enabled in the database. Otherwise, use "without". -->
        {{if eq .OTPRequired true}} with {{else}} without {{end}}
        2FA Authentication</h1>

        <a href="/login_system/logout" type="button">Logout</a>
      </body>
    </html>

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
**home.go**
^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package views

    import (
        "net/http"

        "github.com/uadmin/uadmin"
    )

    // HomeHandler handles the home page.
    func HomeHandler(w http.ResponseWriter, r *http.Request, session *uadmin.Session) {
        // Initialize the fields that we need in the custom struct.
        type Context struct {
            User        string
            OTPRequired bool
        }

        // Call the custom struct and assign the full name in the User field under the context object.
        c := Context{}
        c.User = session.User.FirstName + " " + session.User.LastName

        // Check if the user has OTPRequired enabled in the database.
        if session.User.OTPRequired {
            /* Assign a boolean value to OTPRequired field. We will use this to manipulate the grammar
            in the UI. */
            c.OTPRequired = true
        }

        // Render the home filepath and pass the context data object to the HTML file.
        uadmin.RenderHTML(w, r, "templates/home.html", c)
        return
    }

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

**logout.go**
^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

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
                MaxAge: -1,
            }

            http.SetCookie(w, c)
        }
        http.Redirect(w, r, "/login_system/", http.StatusSeeOther)
    }

**main.go (views)**
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

        // Authentication : This session is preloaded with a user.
        session := uadmin.IsAuthenticated(r)
        if session == nil {
            // LoginHandler verifies login data and creating sessions for users.
            LoginHandler(w, r)
            return
        }

        if r.URL.Path == "" {
            // HomeHandler handles the home page.
            HomeHandler(w, r, session)
            return

        } else if r.URL.Path == "/logout" {
            /* If the request URL Path is /logout after the /login_system/, it will proceed to this part.
            e.g. localhost:8080/login_system/logout */

            // LogoutHandler handles the logout process for the user.
            LogoutHandler(w, r, session)
            return
        }
    }

main.go
-------
`Back to Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/login_system_views/tutorial/full_code/part6.html#login-system-tutorial-part-6-logout-current-progress

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
