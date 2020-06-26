uAdmin Tutorial Part 10 - Customizing your API Handler (Current Progress)
=========================================================================
`Back to Previous Page`_

.. _Back to Previous Page: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part10.html

Structure:

* `api`_
    * `api.go`_
    * `custom_list.go`_
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

.. _api: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part10.html#id1
.. _api.go: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part10.html#id2
.. _custom_list.go: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part10.html#id3
.. _todo_list.go: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part10.html#id4
.. _models: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part10.html#id5
.. _category.go: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part10.html#id6
.. _friend.go: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part10.html#id7
.. _item.go: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part10.html#id8
.. _todo.go: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part10.html#id9
.. _main.go: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part10.html#id10
.. _uadmin.db: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part10.html#id11
.. _Categories: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part10.html#id12
.. _Friends: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part10.html#id13
.. _Items: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part10.html#id14
.. _Todos: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part10.html#id15

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
        if strings.HasPrefix(r.URL.Path, "/custom_list") {
            CustomListAPIHandler(w, r)
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
        Category     []Category `uadmin:"list_exclude"`
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
        uadmin.Register(
            models.Todo{},
            models.Category{},
            models.Friend{},
            models.Item{},
        )

        uadmin.RegisterInlines(models.Category{}, map[string]string{
            "Todo": "CategoryID",
        })
        uadmin.RegisterInlines(models.Friend{}, map[string]string{
            "Todo": "FriendID",
        })
        uadmin.RegisterInlines(models.Item{}, map[string]string{
            "Todo": "ItemID",
        })

        // Initialize Setting model
        setting := uadmin.Setting{}

        // Get the code
        uadmin.Get(&setting, "code = ?", "uAdmin.RootURL")

        // Assign the value as "/admin/"
        setting.ParseFormValue([]string{"/admin/"})

        // Save changes
        setting.Save()

        // API Handler
        http.HandleFunc("/api/", uadmin.Handler(api.Handler))

        uadmin.StartServer()
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

.. image:: assets/friendmodelupdate2.png

**Items**
^^^^^^^^^
`Back to Top`_

.. image:: assets/itemmodelupdate3.png

**Todos**
^^^^^^^^^
`Back to Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part10.html#uadmin-tutorial-part-10-customizing-your-api-handler-current-progress

.. image:: assets/todomodelupdate4.png