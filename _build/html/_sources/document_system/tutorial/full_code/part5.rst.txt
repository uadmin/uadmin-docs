Document System Tutorial Part 5 - Linking Models (Current Progress)
===================================================================
`Back to Previous Page`_

.. _Back to Previous Page: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/part5.html

Structure:

* `models`_
    * `channel.go`_
    * `document_group.go`_
    * `document_user.go`_
    * `document_version.go`_
    * `document.go`_
    * `folder_group.go`_
    * `folder_user.go`_
    * `folder.go`_
* `main.go`_
* `uadmin.db`_
    * `Folder Groups`_
    * `Folder Users`_
    * `Folders`_
    * `Users`_

.. _models: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part5.html#id1
.. _channel.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part5.html#id2
.. _document_group.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part5.html#id3
.. _document_user.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part5.html#id4
.. _document_version.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part5.html#id5
.. _document.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part5.html#id6
.. _folder_group.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part5.html#id7
.. _folder_user.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part5.html#id8
.. _folder.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part5.html#id9
.. _main.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part5.html#id10
.. _uadmin.db: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part5.html#id11
.. _Folder Groups: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part5.html#id12
.. _Folder Users: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part5.html#id13
.. _Folders: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part5.html#id14
.. _Users: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part5.html#id15

models
------
**channel.go**
^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package models

    import (
        "github.com/uadmin/uadmin"
    )

    // Channel !
    type Channel struct {
        uadmin.Model
        Name string `uadmin:"required"`
    }

**document_group.go**
^^^^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package models

    import (
        "github.com/uadmin/uadmin"
    )

    // DocumentGroup !
    type DocumentGroup struct {
        uadmin.Model
        Group      uadmin.UserGroup
        GroupID    uint
        Document   Document
        DocumentID uint
        Read       bool
        Add        bool
        Edit       bool
        Delete     bool
    }

    // DocumentGroup function that returns string value
    func (d *DocumentGroup) String() string {

        // Gives access to the fields in another model
        uadmin.Preload(d)

        // Returns the GroupName from the Group model
        return d.Group.GroupName
    }

**document_user.go**
^^^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package models

    import (
        "github.com/uadmin/uadmin"
    )

    // DocumentUser !
    type DocumentUser struct {
        uadmin.Model
        User       uadmin.User
        UserID     uint
        Document   Document
        DocumentID uint
        Read       bool
        Add        bool
        Edit       bool
        Delete     bool
    }

    // DocumentUser function that returns string value
    func (d *DocumentUser) String() string {

        // Gives access to the fields in another model
        uadmin.Preload(d)

        // Returns the full name from the User model
        return d.User.String()
    }

**document_version.go**
^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package models

    import (
        "fmt"
        "time"

        "github.com/uadmin/uadmin"
    )

    // DocumentVersion !
    type DocumentVersion struct {
        uadmin.Model
        Document   Document
        DocumentID uint
        File       string `uadmin:"file"`
        Number     int    `uadmin:"help:version number"`
        Date       time.Time
    }

    // Returns the version number
    func (d DocumentVersion) String() string {
        return fmt.Sprint(d.Number)
    }

**document.go**
^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package models

    import (
        "time"

        "github.com/uadmin/uadmin"
    )

    // Document !
    type Document struct {
        uadmin.Model
        Name        string
        File        string `uadmin:"file"`
        Description string `uadmin:"html"`
        RawText     string `uadmin:"list_exclude"`
        Folder      Folder `uadmin:"filter"`
        FolderID    uint
        CreatedDate time.Time
        Channel     Channel `uadmin:"list_exclude"`
        ChannelID   uint
        CreatedBy   string
    }

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
            models.Channel{},
            models.Document{},
            models.DocumentGroup{},
            models.DocumentUser{},
            models.DocumentVersion{},
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

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part5.html#document-system-tutorial-part-5-linking-models-current-progress

.. image:: assets/usermodelupdate.png
