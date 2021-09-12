uAdmin Tutorial Part 13 - Accessing an HTML file
================================================
In this part, we will talk about establishing a connection to the Page Handler, setting the URL path name, and executing an HTML file.

Go to **view.go** inside the views folder with the following codes below:

.. code-block:: go

    package views

    import (
        "net/http"
        "strings"
    )

    // PageHandler !
    func PageHandler(w http.ResponseWriter, r *http.Request) {
        // r.URL.Path creates a new path called "/page/"
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/page")
        r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
    }

Establish a connection in the main.go to the views by using http.HandleFunc. It should be placed after the uadmin.Register and before the StartServer.

.. code-block:: go

    import (
        "net/http"

        // Specify the username that you used inside github.com folder
        "github.com/username/todo/api"
        "github.com/username/todo/models"

        // Import this library
        "github.com/username/todo/views"

        "github.com/uadmin/uadmin"
    )

    func main() {
        // Some codes

        // HTTP UI Handler
        http.HandleFunc("/page/", uadmin.Handler(views.PageHandler))
    }

Create a file named **todo_view.go** inside the views folder with the following codes below:

.. code-block:: go

    package views

    import (
        "net/http"

        "github.com/uadmin/uadmin"
    )

    // TodoList field inside the Context that will be used in Golang HTML template
    type Context struct {
        TodoList []map[string]interface{}
    }

    // TodoHandler !
    func TodoHandler(w http.ResponseWriter, r *http.Request) {
        // Assigns Context struct to the c variable
        c := Context{}

        // Pass TodoList data object to the specified HTML path
        uadmin.RenderHTML(w, r, "templates/todo.html", c)
    }

Finally, add this piece of code in the **view.go** shown below. This will establish a communication between the PageHandler and the TodoHandler.

.. code-block:: go

    // PageHandler !
    func PageHandler(w http.ResponseWriter, r *http.Request) {
        // r.URL.Path creates a new path called "/page/"
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/page")
        r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

        if strings.HasPrefix(r.URL.Path, "/todo") {
            TodoHandler(w, r)
            return
        }
    }

Now run your application, go to http://localhost:8080/page/todo path and see what happens.

.. image:: assets/todohtmlaccess.png

|

Click `here`_ to view our progress so far.

In the `next part`_, we will discuss about fetching the records in the API and migrating the data from API to HTML that will display the records using Go template.

.. _here: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part13.html
.. _next part: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part14.html

.. toctree::
   :maxdepth: 1

   full_code/part13
