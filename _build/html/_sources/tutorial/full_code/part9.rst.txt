uAdmin Tutorial Part 9 - Introduction to API (Current Progress)
===============================================================
`Back to Previous Page`_

.. _Back to Previous Page: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part9.html

Structure:

* `api`_
    * `api.go`_
    * `todo_list.go`_
* `models`_
    * `category.go`_
    * `friend.go`_
    * `item.go`_
    * `todo.go`_
* `main.go`_
* `uadmin.db`_
    * `Categories`_
    * `Friends`_
    * `Items`_
    * `Todos`_

api
---

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

        // API Handler
        http.HandleFunc("/api/", uadmin.Handler(api.Handler))

        // Call InitializeRootURL function to change the RootURL value in the Settings model.
        InitializeRootURL()

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

.. image:: assets/friendmodelupdate.png

**Items**
^^^^^^^^^
`Back to Top`_

.. image:: assets/itemmodelupdate2.png

**Todos**
^^^^^^^^^
`Back to Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part9.html#uadmin-tutorial-part-9-introduction-to-api-current-progress

.. image:: assets/todomodelupdate3.png