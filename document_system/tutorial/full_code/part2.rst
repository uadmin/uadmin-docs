Document System Tutorial Part 2 - Creating and Registering a Model (Current Progress)
=====================================================================================
`Back to Previous Page`_

.. _Back to Previous Page: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/part2.html

Structure:

* `models`_
    * `folder.go`_
* `main.go`_
* `uadmin.db`_
    * `Folders`_

models
------
**folder.go**
^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package models

    import (
        "github.com/uadmin/uadmin"
    )

    // Folder !
    type Folder struct {
        uadmin.Model
        Name     string
        Parent   *Folder
        ParentID uint
    }

main.go
-------
`Back to Top`_

.. code-block:: go

    package main

    import (
        // Specify the username that you used inside github.com folder
        "github.com/username/document_system/models"
        "github.com/uadmin/uadmin"
    )

    func main() {
        // Register models to uAdmin
        uadmin.Register(
            models.Folder{},
        )

        // Assign Site Name value as "Document System"
        // NOTE: This code works only if database does not exist yet.
        uadmin.SiteName = "Document System"

        // Activates a uAdmin server
        uadmin.StartServer()
    }

uadmin.db
---------
**Folders**
^^^^^^^^^^^
`Back to Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part2.html#document-system-tutorial-part-2-creating-and-registering-a-model-current-progress

.. image:: assets/foldermodelupdate.png
