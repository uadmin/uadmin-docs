uAdmin Tutorial Part 2 - Internal vs. External Models (Current Progress)
========================================================================
`Back to Previous Page`_

.. _Back to Previous Page: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part2.html

Structure:

* `models`_
    * `todo.go`_
* `main.go`_
* `uadmin.db`_
    * `Todos`_

.. _models: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part2.html#id1
.. _todo.go: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part2.html#id2
.. _main.go: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part2.html#id3
.. _uadmin.db: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part2.html#id4
.. _Todos: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part2.html#id5

models
------

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
        )

        uadmin.StartServer()
    }

uadmin.db
---------

**Todos**
^^^^^^^^^
`Back to Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part2.html#uadmin-tutorial-part-2-internal-vs-external-models-current-progress

.. image:: ../../assets/todomodeloutput.png
