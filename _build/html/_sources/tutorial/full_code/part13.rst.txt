uAdmin Tutorial Part 13 - Accessing an HTML file (Current Progress)
===================================================================
`Back to Previous Page`_

.. _Back to Previous Page: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part13.html

Structure:

* `api`_
    * `add_friend.go`_
    * `api.go`_
    * `todo_list.go`_
    * `custom_list.go`_
* `models`_
    * `category.go`_
    * `friend.go`_
    * `item.go`_
    * `todo.go`_
* `templates`_
    * `todo.html`_
* `views`_
    * `todo_view.go`_
    * `view.go`_
* `main.go`_
* `uadmin.db`_
    * `Categories`_
    * `Friends`_
    * `Items`_
    * `Todos`_

api
---

**add_friend.go**
^^^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package api

    import (
        "net/http"

        // Specify the username that you used inside github.com folder
        "github.com/username/todo/models"
        "github.com/uadmin/uadmin"
    )

    // AddFriendAPIHandler !
    func AddFriendAPIHandler(w http.ResponseWriter, r *http.Request) {
        // Fetch data from Friend DB
        friend := models.Friend{}

        // Set the parameters of Name, Email, Password, and Nationality such that where nationality is
        // equivalent to the following:
        // 1 - Chinese
        // 2 - Filipino
        // 3 - Others
        friendName := r.FormValue("name")
        friendEmail := r.FormValue("email")
        friendPassword := r.FormValue("password")
        friendNationalityRaw := r.FormValue("nationality")

        // Convert the nationality to an integer.
        friendNationality, err := strconv.Atoi(friendNationalityRaw)

        // Validate if the friendName variable is empty.
        if friendName == "" {
            uadmin.ReturnJSON(w, r, map[string]interface{}{
                "status":  "error",
                "err_msg": "Name is required.",
            })
            return
        }

        // Store input into the Name, Email, and Password fields
        friend.Name = friendName
        friend.Email = friendEmail
        friend.Password = friendPassword
        friend.Nationality = models.Nationality(friendNationality)

        // Save input in the Friend model
        err = uadmin.Save(&friend)
        if err != nil {
            // Return an error message if the database did not save properly.
            uadmin.ReturnJSON(w, r, map[string]interface{}{
                "status":  "error",
                "err_msg": "Error saving the database : " + err.Error(),
            })
            return
        }

        // Return JSON to the client.
        uadmin.ReturnJSON(w, r, map[string]interface{}{
            "status": "ok",
        })
    }

**api.go**
^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package api

    import (
        "net/http"
        "strings"
    )

    // Handler !
    func Handler(w http.ResponseWriter, r *http.Request) {
        // r.URL.Path creates a new path called "/api/"
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")
        r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

        if strings.HasPrefix(r.URL.Path, "/todo_list") {
            TodoListAPIHandler(w, r)
            return
        }
        if strings.HasPrefix(r.URL.Path, "/custom_list") {
            CustomListAPIHandler(w, r)
            return
        }
        if strings.HasPrefix(r.URL.Path, "/add_friend") {
            AddFriendAPIHandler(w, r)
            return
        }
    }

**custom_list.go**
^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package api

    import (
        "net/http"

        // Specify the username that you used inside github.com folder
        "github.com/username/todo/models"
        "github.com/uadmin/uadmin"
    )

    // CustomListAPIHandler !
    func CustomListAPIHandler(w http.ResponseWriter, r *http.Request) {
        // Fetch Data from DB
        todo := []models.Todo{}

        // Assigns a map as a string of interface to store any types of values
        results := []map[string]interface{}{}

        // "id" - order the todo model by id
        // false - to sort in descending order
        // 0 - start at index 0
        // 5 - get five records
        // &todo - todo model to execute
        // "" - fetch the id of the model itself
        uadmin.AdminPage("id", false, 0, 5, &todo, "")

        // Loop to fetch the record of todo
        for i := range todo {
            // Accesses and fetches the record of the linking models in Todo
            uadmin.Preload(&todo[i])

            // Assigns the string of interface in each Todo fields
            results = append(results, map[string]interface{}{
                "ID":          todo[i].ID,
                "Name":        todo[i].Name,
                "Description": todo[i].Description,
                // This returns only the name of the Category model, not the
                // other fields
                "Category": todo[i].Category.Name,
                // This returns only the name of the Friend model, not the
                // other fields
                "Friend": todo[i].Friend.Name,
                // This returns only the name of the Item model, not the other
                // fields
                "Item":       todo[i].Item.Name,
                "TargetDate": todo[i].TargetDate,
                "Progress":   todo[i].Progress,
            })
        }

        // Prints the results in JSON format
        uadmin.ReturnJSON(w, r, results)
    }

**todo_list.go**
^^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package api

    import (
        "net/http"

        // Specify the username that you used inside github.com folder
        "github.com/username/todo/models"
        "github.com/uadmin/uadmin"
    )

    // TodoListAPIHandler !
    func TodoListAPIHandler(w http.ResponseWriter, r *http.Request) {
        // Fetch all records in the database
        todo := []models.Todo{}
        uadmin.All(&todo)

        // Accesses and fetches data from another model
        for t := range todo {
            uadmin.Preload(&todo[t])
        }

        // Return todo JSON object
        uadmin.ReturnJSON(w, r, todo)
    }

models
------

**category.go**
^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package models

    import (
        "github.com/uadmin/uadmin"
    )

    // Category Model !
    type Category struct {
        uadmin.Model
        Name string `uadmin:"required"`
        Icon string `uadmin:"image"`
    }

**friend.go**
^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package models

    import (
        "github.com/uadmin/uadmin"
    )

    // Nationality Field !
    type Nationality int

    // Chinese !
    func (Nationality) Chinese() Nationality {
        return 1
    }

    // Filipino !
    func (Nationality) Filipino() Nationality {
        return 2
    }

    // Others !
    func (Nationality) Others() Nationality {
        return 3
    }

    // Friend Model !
    type Friend struct {
        uadmin.Model
        Name        string `uadmin:"required"`
        Email       string `uadmin:"email"`
        Password    string `uadmin:"password;list_exclude"`
        Nationality Nationality
        Invite      string `uadmin:"link"`
    }

    // Save !
    func (f *Friend) Save() {
        f.Invite = "https://www.google.com/"
        uadmin.Save(f)
    }

**item.go**
^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package models

    import (
        "strings"

        "github.com/uadmin/uadmin"
    )

    // Item Model !
    type Item struct {
        uadmin.Model
        Name         string     `uadmin:"required;search;categorical_filter;filter;display_name:Product Name;default_value:Computer"`
        Description  string     `uadmin:"multilingual"`
        Category     []Category `uadmin:"list_exclude" gorm:"many2many:-"`
        CategoryList string     `uadmin:"read_only"`
        Cost         int        `uadmin:"money;pattern:^[0-9]*$;pattern_msg:Your input must be a number.;help:Input numeric characters only in this field."`
        Rating       int        `uadmin:"min:1;max:5"`
    }

    // Save !
    func (i *Item) Save() {
        // Add a new string array type variable called categoryList
        categoryList := []string{}

        // Append every element to the categoryList array
        for c := range i.Category {
            categoryList = append(categoryList, i.Category[c].Name)
        }

        // Concatenate the categoryList to a single string separated by comma
        joinList := strings.Join(categoryList, ", ")

        // Store the joined string to the CategoryList field
        i.CategoryList = joinList

        // Save it to the database
        uadmin.Save(i)
    }


**todo.go**
^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package models

    import (
        "time"

        "github.com/uadmin/uadmin"
    )

    // Todo Model !
    type Todo struct {
        uadmin.Model
        Name        string
        Description string `uadmin:"html"`
        Category    Category
        CategoryID  uint
        Friend      Friend `uadmin:"help:Who will be a part of your activity?"`
        FriendID    uint
        Item        Item `uadmin:"help:What are the requirements needed in order to accomplish your activity?"`
        ItemID      uint
        TargetDate  time.Time
        Progress    int `uadmin:"progress_bar"`
    }

templates
---------

**todo.html**
^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: html

    <!DOCTYPE html>
    <html lang="en">
      <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">

        <!-- Latest compiled and minified CSS -->
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-F3w7mX95PdgyTmZZMECAngseQB83DfGTowi0iMjiWaeVhAn4FJkqJByhZMI3AhiU" crossorigin="anonymous">

        <title>Todo List</title>
      </head>
      <body>
        <div class="container-fluid">
          <table class="table table-striped">
            <!-- Todo Fields -->
            <thead>
              <tr>
                <th>Name</th>
                <th>Description</th>
                <th>Category</th>
                <th>Friend</th>
                <th>Item</th>
                <th>Target Date</th>
                <th>Progress</th>
              </tr>
            </thead>
            <tbody>
        
            </tbody>
          </table>
        </div>
      </body>
    </html>

views
-----

**todo_view.go**
^^^^^^^^^^^^^^^^
`Back to Top`_

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

**view.go**
^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package views

    import (
        "net/http"
        "strings"
    )

    // PageHandler !
    func PageHandler(w http.ResponseWriter, r *http.Request) {
        // r.URL.Path creates a new path called "/page"
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/page")
        r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

        if strings.HasPrefix(r.URL.Path, "/todo") {
            TodoHandler(w, r)
            return
        }
    }

main.go
-------
`Back to Top`_

.. code-block:: go

    package main

    import (
        "net/http"

        // Specify the username that you used inside github.com folder
        "github.com/username/todo/api"
        "github.com/username/todo/models"
        "github.com/username/todo/views"

        "github.com/uadmin/uadmin"
    )

    func main() {
        // Register Models
        uadmin.Register(
            models.Todo{},
            models.Category{},
            models.Friend{},
            models.Item{},
        )

        // Register Inlines
        uadmin.RegisterInlines(models.Category{}, map[string]string{
            "Todo": "CategoryID",
        })
        uadmin.RegisterInlines(models.Friend{}, map[string]string{
            "Todo": "FriendID",
        })
        uadmin.RegisterInlines(models.Item{}, map[string]string{
            "Todo": "ItemID",
        })

        // Call InitializeRootURL function to change the RootURL value in the Settings model.
        InitializeRootURL()

        // API Handler
        http.HandleFunc("/api/", uadmin.Handler(api.Handler))

        // Page Handler
        http.HandleFunc("/page/", uadmin.Handler(views.HTTPHandler))

        uadmin.StartServer()
    }

    func InitializeRootURL() {
        // Initialize Setting model
        setting := uadmin.Setting{}

        // Get the code
        uadmin.Get(&setting, "code = ?", "uAdmin.RootURL")

        // Assign the value as "/admin/"
        setting.ParseFormValue([]string{"/admin/"})

        // Save changes
        setting.Save()
    }

uadmin.db
---------

**Categories**
^^^^^^^^^^^^^^
`Back to Top`_

.. image:: assets/categorymodelupdate2.png

**Friends**
^^^^^^^^^^^
`Back to Top`_

.. image:: assets/friendmodelupdate3.png

**Items**
^^^^^^^^^
`Back to Top`_

.. image:: assets/itemmodelupdate3.png

**Todos**
^^^^^^^^^
`Back to Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part13.html#uadmin-tutorial-part-13-accessing-an-html-file-current-progress

.. image:: assets/todomodelupdate4.png