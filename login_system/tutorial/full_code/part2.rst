Login System Tutorial Part 2 - Login Form (Full Source Code)
============================================================

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
        // r.URL.Path creates a new path called /login
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/login")

        // Initialize the UserContext struct that we have created
        userContext := UserContext{}

        // Pass the userContext data object to the HTML file
        uadmin.RenderHTML(w, userContext, "templates/login.html")
        return
    }
