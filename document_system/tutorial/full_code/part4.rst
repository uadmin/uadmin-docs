<<<<<<< HEAD
Document System Tutorial Part 4 - Creating Records in Folders (Current Progress)
================================================================================
`Back to Previous Page`_

.. _Back to Previous Page: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/part4.html

Structure:

* `models`_
    * `folder_group.go`_
    * `folder_user.go`_
    * `folder.go`_
* `main.go`_
* `uadmin.db`_
    * `Folder Groups`_
    * `Folder Users`_
    * `Folders`_
    * `Users`_

.. _models: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#id1
.. _folder_group.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#id2
.. _folder_user.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#id3
.. _folder.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#id4
.. _main.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#id5
.. _uadmin.db: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#id6
.. _Folder Groups: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#id7
.. _Folder Users: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#id8
.. _Folders: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#id9
.. _Users: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#id10

models
------
**folder_group.go**
^^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package models

    import (
        "github.com/uadmin/uadmin"
    )

    // FolderGroup !
    type FolderGroup struct {
        uadmin.Model
        Group    uadmin.UserGroup
        GroupID  uint
        Folder   Folder
        FolderID uint
        Read     bool
        Add      bool
        Edit     bool
        Delete   bool
    }

    // FolderGroup function that returns string value
    func (f *FolderGroup) String() string {

        // Gives access to the fields in another model
        uadmin.Preload(f)

        // Returns the GroupName from the Group model
        return f.Group.GroupName
    }

**folder_user.go**
^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package models

    import (
        "github.com/uadmin/uadmin"
    )

    // FolderUser !
    type FolderUser struct {
        uadmin.Model
        User     uadmin.User
        UserID   uint
        Folder   Folder
        FolderID uint
        Read     bool
        Add      bool
        Edit     bool
        Delete   bool
    }

    // FolderUser function that returns string value
    func (f *FolderUser) String() string {

        // Gives access to the fields in another model
        uadmin.Preload(f)

        // Returns the full name from the User model
        return f.User.String()
    }

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
            models.FolderGroup{},
            models.FolderUser{},
        )

        // Assign Site Name value as "Document System"
        // NOTE: This code works only on first build.
        uadmin.SiteName = "Document System"

        // Activates a uAdmin server
        uadmin.StartServer()
    }

uadmin.db
---------
**Folder Groups**
^^^^^^^^^^^^^^^^^
`Back to Top`_

.. image:: assets/foldergroupmodelupdate.png

**Folder Users**
^^^^^^^^^^^^^^^^
`Back to Top`_

.. image:: assets/folderusermodelupdate.png

**Folders**
^^^^^^^^^^^
`Back to Top`_

.. image:: assets/foldermodelupdate.png

**Users**
^^^^^^^^^
`Back to Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#document-system-tutorial-part-4-creating-records-in-folders-current-progress

.. image:: assets/usermodelupdate.png
=======
Document System Tutorial Part 4 - Creating Records in Folders (Current Progress)
================================================================================
`Back to Previous Page`_

.. _Back to Previous Page: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/part4.html

Structure:

* `models`_
    * `folder_group.go`_
    * `folder_user.go`_
    * `folder.go`_
* `main.go`_
* `uadmin.db`_
    * `Folder Groups`_
    * `Folder Users`_
    * `Folders`_
    * `Users`_

.. _models: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#id1
.. _folder_group.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#id2
.. _folder_user.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#id3
.. _folder.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#id4
.. _main.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#id5
.. _uadmin.db: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#id6
.. _Folder Groups: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#id7
.. _Folder Users: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#id8
.. _Folders: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#id9
.. _Users: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#id10

models
------
**folder_group.go**
^^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package models

    import (
        "github.com/uadmin/uadmin"
    )

    // FolderGroup !
    type FolderGroup struct {
        uadmin.Model
        Group    uadmin.UserGroup
        GroupID  uint
        Folder   Folder
        FolderID uint
        Read     bool
        Add      bool
        Edit     bool
        Delete   bool
    }

    // FolderGroup function that returns string value
    func (f *FolderGroup) String() string {

        // Gives access to the fields in another model
        uadmin.Preload(f)

        // Returns the GroupName from the Group model
        return f.Group.GroupName
    }

**folder_user.go**
^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package models

    import (
        "github.com/uadmin/uadmin"
    )

    // FolderUser !
    type FolderUser struct {
        uadmin.Model
        User     uadmin.User
        UserID   uint
        Folder   Folder
        FolderID uint
        Read     bool
        Add      bool
        Edit     bool
        Delete   bool
    }

    // FolderUser function that returns string value
    func (f *FolderUser) String() string {

        // Gives access to the fields in another model
        uadmin.Preload(f)

        // Returns the full name from the User model
        return f.User.String()
    }

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
            models.FolderGroup{},
            models.FolderUser{},
        )

        // Assign Site Name value as "Document System"
        // NOTE: This code works only on first build.
        uadmin.SiteName = "Document System"

        // Activates a uAdmin server
        uadmin.StartServer()
    }

uadmin.db
---------
**Folder Groups**
^^^^^^^^^^^^^^^^^
`Back to Top`_

.. image:: assets/foldergroupmodelupdate.png

**Folder Users**
^^^^^^^^^^^^^^^^
`Back to Top`_

.. image:: assets/folderusermodelupdate.png

**Folders**
^^^^^^^^^^^
`Back to Top`_

.. image:: assets/foldermodelupdate.png

**Users**
^^^^^^^^^
`Back to Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part4.html#document-system-tutorial-part-4-creating-records-in-folders-current-progress

.. image:: assets/usermodelupdate.png
>>>>>>> de25cdd8a29ca2bb2c2df08be00b703b967aaed5
