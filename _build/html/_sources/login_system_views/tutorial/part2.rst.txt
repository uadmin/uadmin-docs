Login System Tutorial Part 2 - Login Form
=========================================
In this, we will discuss about creating a login form in HTML.

Before you proceed, make sure you have at least the basic knowledge of HTML. If you are not familiar with HTML, we advise you to go over `W3Schools`_.

.. _W3Schools: https://www.w3schools.com/


First of all, create a new file in the templates folder named **login.html**. Inside login.html, create a login form containing the username and password input fields, and a submit button with the method of **POST**.

.. code-block:: html

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

          <br><button type="submit">Login</button><br>
        </form>
      </body>
    </html>

Now create a new file in the views folder named **main.go**. Inside main.go, create a MainHandler function that calls the /login_system/ URL path.

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
    }

Create another file in the views folder named **login.go**. Inside login.go, create a LoginHandler function that renders the login HTML file and prepares the fields that we need in context struct. We will use that in the later chapter of this tutorial.

.. code-block:: go

    package views

    import (
        "net/http"

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

        // Render the login filepath and pass the context data object to the HTML file.
        uadmin.RenderHTML(w, r, "templates/login.html", c)
    }

Call the LoginHandler function in **views/main.go**.

.. code-block:: go

    // MainHandler is the main handler for the login system.
    func MainHandler(w http.ResponseWriter, r *http.Request) {
        // Some codes

        // LoginHandler verifies login data and creating sessions for users.
        LoginHandler(w, r)
    }

Establish a connection in the main.go to the views by using http.HandleFunc. It should be placed before the StartServer.

.. code-block:: go

    import (
        "net/http"

        // Specify the username that you used inside github.com folder
        "github.com/username/project_name/views"

        "github.com/uadmin/uadmin"
    )

    func main() {
        // Some codes

        // Login System Main Handler
        http.HandleFunc("/login_system/", uadmin.Handler(views.MainHandler))
    }

Now run your application. Go to the login system path (e.g. http://localhost:8080/login_system/) and see what happens.

.. image:: assets/customloginform.png
   :align: center

|

Click `here`_ to view our progress so far.

In the `next part`_, we will talk about sending data from login form in HTML to the LoginHandler.

.. _here: https://uadmin-docs.readthedocs.io/en/latest/login_system_views/tutorial/full_code/part2.html

.. _next part: https://uadmin-docs.readthedocs.io/en/latest/login_system_views/tutorial/part3.html

.. toctree::
   :maxdepth: 1

   full_code/part2
