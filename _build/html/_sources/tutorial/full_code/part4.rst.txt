uAdmin Tutorial Part 4 - Linking Models (Current Progress)
==========================================================
`Back to Previous Page`_

.. _Back to Previous Page: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part4.html

Structure:

* `models`_
    * `category.go`_
    * `friend.go`_
    * `todo.go`_
* `main.go`_
* `uadmin.db`_
    * `Categories`_
    * `Friends`_
    * `Todos`_

.. _models: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part4.html#id1
.. _category.go: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part4.html#id2
.. _friend.go: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part4.html#id3
.. _todo.go: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part4.html#id4
.. _main.go: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part4.html#id5
.. _uadmin.db: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part4.html#id6
.. _Categories: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part4.html#id7
.. _Friends: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part4.html#id8
.. _Todos: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part4.html#id9

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

    // Friend Model !
    type Friend struct {
        uadmin.Model
        Name     string `uadmin:"required"`
        Email    string `uadmin:"email"`
        Password string `uadmin:"password;list_exclude"`
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
        TargetDate  time.Time
        Progress    int `uadmin:"progress_bar"`
    }

main.go
-------
`Back to Top`_

.. code-block:: go

    package main

    import (
        // Specify the username that you used inside github.com folder
        "github.com/username/todo/models"

        "github.com/uadmin/uadmin"
    )

    func main() {
        uadmin.Register(
            models.Todo{},
            models.Category{},
            models.Friend{},
        )

        uadmin.StartServer()
    }

uadmin.db
---------

**Categories**
^^^^^^^^^^^^^^
`Back to Top`_

.. image:: assets/categorymodelupdate.png

**Friends**
^^^^^^^^^^^
`Back to Top`_

.. image:: ../assets/friendsdataoutput.png

**Todos**
^^^^^^^^^
`Back to Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part4.html#uadmin-tutorial-part-4-linking-models-current-progress

.. image:: assets/todomodelupdate.png
