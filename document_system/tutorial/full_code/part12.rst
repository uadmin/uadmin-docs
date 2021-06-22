<<<<<<< HEAD
Document System Tutorial Part 12 - Custom Count function (Current Progress)
===========================================================================
`Back to Previous Page`_

.. _Back to Previous Page: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/part12.html

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
    * `format.go`_
* `main.go`_
* `uadmin.db`_
    * `Channels`_
    * `Documents`_
    * `Document Groups`_
    * `Document Users`_
    * `Document Versions`_
    * `Folder Groups`_
    * `Folder Users`_
    * `Folders`_
    * `Group Permissions`_
    * `Users`_

.. _models: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id1
.. _channel.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id2
.. _document_group.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id3
.. _document_user.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id4
.. _document_version.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id5
.. _document.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id6
.. _folder_group.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id7
.. _folder_user.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id8
.. _folder.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id9
.. _format.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id10
.. _main.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id11
.. _uadmin.db: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id12
.. _Channels: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id13
.. _Documents: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id14
.. _Document Groups: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id15
.. _Document Users: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id16
.. _Document Versions: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id17
.. _Folder Groups: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id18
.. _Folder Users: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id19
.. _Folders: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id20
.. _Group Permissions: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id21
.. _Users: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id22

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
        Format     Format
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
        Format      Format `uadmin:"list_exclude"`
        Folder      Folder `uadmin:"filter"`
        FolderID    uint
        CreatedDate time.Time
        Channel     Channel `uadmin:"list_exclude"`
        ChannelID   uint
        CreatedBy   string
    }

    // Save !
    func (d *Document) Save() {
        // Initialized variables
        docChange := false
        newDoc := false
        // Checks whether the document record is new or existing
        if d.ID != 0 {
            // Initializes the Document model
            oldDoc := Document{}

            // Gets the ID of the old Document
            uadmin.Get(&oldDoc, "id = ?", d.ID)

            // Checks if the file is changed or updated
            if d.File != oldDoc.File {
                docChange = true
            }
        } else {
            // New document record
            docChange = true
            newDoc = true
        }

        // Save the document
        uadmin.Save(d)

        // Checks whether the document record has changed
        if docChange {
            // Prints the result
            uadmin.Trail(uadmin.DEBUG, "The document has changed.")

            // Sets the document value to the DocumentVersion
            ver := DocumentVersion{}
            ver.Date = time.Now()
            ver.DocumentID = d.ID
            ver.File = d.File
            ver.Format = d.Format

            // Counts the version number by DocumentID and increment it by 1
            ver.Number = uadmin.Count([]DocumentVersion{}, "document_id = ?", d.ID) + 1

            // Save the document version
            uadmin.Save(&ver)

            // Checks whether the document is a new record
            if newDoc {
                // Initializes the User model
                user := uadmin.User{}

                // Gets the username of the user to display in CreatedBy
                uadmin.Get(&user, "username = ?", d.CreatedBy)

                // Sets values to the DocumentUser model fields
                creator := DocumentUser{
                    UserID:     user.ID,
                    DocumentID: d.ID,
                    Read:       true,
                    Edit:       true,
                    Add:        true,
                    Delete:     true,
                }

                // Save the document user
                uadmin.Save(&creator)
            }
        }
    }

    // GetPermissions !
    func (d Document) GetPermissions(user uadmin.User) (Read bool, Add bool, Edit bool, Delete bool) {
        // Check whether the user is an admin
        if user.Admin {
            // Set all permissions to true
            Read = true
            Add = true
            Edit = true
            Delete = true
        }

        // Since Folder is a foreign key to the Document model, we need to check
        // whether there is a Folder specified in the Document model.
        // We will check for folder permissions first
        // Then we will check for document permissions after that
        if d.FolderID != 0 {
            // Initialize the FolderGroup model
            folderGroup := FolderGroup{}

            // Get data by GroupID and FolderID
            uadmin.Get(&folderGroup, "group_id = ? AND folder_id = ?", user.UserGroupID, d.FolderID)

            // Check whether there is a FolderGroup recird
            if folderGroup.ID != 0 {
                // Assign FolderGroup permission values to the variables
                Read = folderGroup.Read
                Add = folderGroup.Add
                Edit = folderGroup.Edit
                Delete = folderGroup.Delete
            }

            // Initialize the FolderUser model
            folderUser := FolderUser{}

            // Get data by UserID and FolderID
            uadmin.Get(&folderUser, "user_id = ? AND folder_id = ?", user.ID, d.FolderID)

            // Check whether there is a FolderUser record
            if folderUser.ID != 0 {
                // Assign FolderUser permission values to the variables
                Read = folderUser.Read
                Add = folderUser.Add
                Edit = folderUser.Edit
                Delete = folderUser.Delete
            }
        }

        // Document Permissions
        // Initialize the DocumentGroup model
        documentGroup := DocumentGroup{}

        // Get data by GroupID and DocumentID
        uadmin.Get(&documentGroup, "group_id = ? AND document_id = ?", user.UserGroupID, d.ID)

        // Check whether there is a DocumentGroup record
        if documentGroup.ID != 0 {
            // Assign DocumentGroup permission values to the variables
            Read = documentGroup.Read
            Add = documentGroup.Add
            Edit = documentGroup.Edit
            Delete = documentGroup.Delete
        }

        // Initialize the DocumentUser model
        documentUser := DocumentUser{}

        // // Get data by UserID and DocumentID
        uadmin.Get(&documentUser, "user_id = ? AND document_id = ?", user.ID, d.ID)

        // Check whether there is a DocumentUser record
        if documentUser.ID != 0 {
            // Assign DocumentUser permission values to the variables
            Read = documentUser.Read
            Add = documentUser.Add
            Edit = documentUser.Edit
            Delete = documentUser.Delete
        }

        // Return Read, Add, Edit, and Delete values
        return
    }

    // Count !
    func (d Document) Count(a interface{}, query interface{}, args ...interface{}) int {
        // Converts the query into a string
        Q := fmt.Sprint(query)

        // Checks whether the string contains a query and a UserID
        if strings.Contains(Q, "user_id = ?") {
            // Split the query part by part
            qParts := strings.Split(Q, " AND ")

            // Initialize tempArgs as an interface and tempQuery as a
            // string
            tempArgs := []interface{}{}
            tempQuery := []string{}

            // Loop the query every part
            for i := range qParts {
                // Checks whether the specific query part is not
                // equal to the UserID value
                if qParts[i] != "user_id = ?" {
                    // Append the arguments into the tempArgs
                    // variable
                    tempArgs = append(tempArgs, args[i])

                    // Append the specific query part into the
                    // tempQuery variable
                    tempQuery = append(tempQuery, qParts[i])
                }
            }
            // Concatenate the query to create a single string
            query = strings.Join(tempQuery, " AND ")

            // Assign tempArgs object into the args variable
            args = tempArgs
        }

        // Return the a, query, and args... inside the Count function
        // parameters
        return uadmin.Count(a, query, args...)
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

**folder.go**
^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package models

    // Format is the name of the drop down list ...
    type Format int

    // PDF is the name of the drop down list value ...
    func (Format) PDF() Format {
        return 1
    }

    // TXT is the name of the drop down list value ...
    func (Format) TXT() Format {
        return 2
    }

    // Others is the name of the drop down list value ...
    func (Format) Others() Format {
        return 3
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

        // Register FolderGroup and FolderUser to Folder model
        uadmin.RegisterInlines(
            models.Folder{},
            map[string]string{
                "foldergroup": "FolderID",
                "folderuser":  "FolderID",
            },
        )

        // Register DocumentVersion, DocumentGroup, and DocumentUser to Document
        // model
        uadmin.RegisterInlines(
            models.Document{},
            map[string]string{
                "documentgroup":   "DocumentID",
                "documentuser":    "DocumentID",
                "documentversion": "DocumentID",
            },
        )

        // Assign Site Name value as "Document System"
        // NOTE: This code works only on first build.
        uadmin.SiteName = "Document System"

        // Activates a uAdmin server
        uadmin.StartServer()
    }


uadmin.db
---------
**Channels**
^^^^^^^^^^^^
`Back to Top`_

.. image:: assets/channelmodelupdate.png

**Documents**
^^^^^^^^^^^^^
`Back to Top`_

.. image:: assets/documentmodelupdate3.png

**Document Groups**
^^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. image:: assets/documentgroupmodelupdate.png

**Document Users**
^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. image:: assets/documentusermodelupdate.png

**Document Versions**
^^^^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. image:: assets/documentversionmodelupdate3.png

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

**Group Permissions**
^^^^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. image:: assets/grouppermissionmodelupdate.png

**Users**
^^^^^^^^^
`Back to Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#document-system-tutorial-part-12-custom-count-function-current-progress

.. image:: assets/usermodelupdate.png
=======
Document System Tutorial Part 12 - Custom Count function (Current Progress)
===========================================================================
`Back to Previous Page`_

.. _Back to Previous Page: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/part12.html

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
    * `format.go`_
* `main.go`_
* `uadmin.db`_
    * `Channels`_
    * `Documents`_
    * `Document Groups`_
    * `Document Users`_
    * `Document Versions`_
    * `Folder Groups`_
    * `Folder Users`_
    * `Folders`_
    * `Group Permissions`_
    * `Users`_

.. _models: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id1
.. _channel.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id2
.. _document_group.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id3
.. _document_user.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id4
.. _document_version.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id5
.. _document.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id6
.. _folder_group.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id7
.. _folder_user.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id8
.. _folder.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id9
.. _format.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id10
.. _main.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id11
.. _uadmin.db: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id12
.. _Channels: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id13
.. _Documents: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id14
.. _Document Groups: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id15
.. _Document Users: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id16
.. _Document Versions: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id17
.. _Folder Groups: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id18
.. _Folder Users: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id19
.. _Folders: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id20
.. _Group Permissions: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id21
.. _Users: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#id22

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
        Format     Format
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
        Format      Format `uadmin:"list_exclude"`
        Folder      Folder `uadmin:"filter"`
        FolderID    uint
        CreatedDate time.Time
        Channel     Channel `uadmin:"list_exclude"`
        ChannelID   uint
        CreatedBy   string
    }

    // Save !
    func (d *Document) Save() {
        // Initialized variables
        docChange := false
        newDoc := false
        // Checks whether the document record is new or existing
        if d.ID != 0 {
            // Initializes the Document model
            oldDoc := Document{}

            // Gets the ID of the old Document
            uadmin.Get(&oldDoc, "id = ?", d.ID)

            // Checks if the file is changed or updated
            if d.File != oldDoc.File {
                docChange = true
            }
        } else {
            // New document record
            docChange = true
            newDoc = true
        }

        // Save the document
        uadmin.Save(d)

        // Checks whether the document record has changed
        if docChange {
            // Prints the result
            uadmin.Trail(uadmin.DEBUG, "The document has changed.")

            // Sets the document value to the DocumentVersion
            ver := DocumentVersion{}
            ver.Date = time.Now()
            ver.DocumentID = d.ID
            ver.File = d.File
            ver.Format = d.Format

            // Counts the version number by DocumentID and increment it by 1
            ver.Number = uadmin.Count([]DocumentVersion{}, "document_id = ?", d.ID) + 1

            // Save the document version
            uadmin.Save(&ver)

            // Checks whether the document is a new record
            if newDoc {
                // Initializes the User model
                user := uadmin.User{}

                // Gets the username of the user to display in CreatedBy
                uadmin.Get(&user, "username = ?", d.CreatedBy)

                // Sets values to the DocumentUser model fields
                creator := DocumentUser{
                    UserID:     user.ID,
                    DocumentID: d.ID,
                    Read:       true,
                    Edit:       true,
                    Add:        true,
                    Delete:     true,
                }

                // Save the document user
                uadmin.Save(&creator)
            }
        }
    }

    // GetPermissions !
    func (d Document) GetPermissions(user uadmin.User) (Read bool, Add bool, Edit bool, Delete bool) {
        // Check whether the user is an admin
        if user.Admin {
            // Set all permissions to true
            Read = true
            Add = true
            Edit = true
            Delete = true
        }

        // Since Folder is a foreign key to the Document model, we need to check
        // whether there is a Folder specified in the Document model.
        // We will check for folder permissions first
        // Then we will check for document permissions after that
        if d.FolderID != 0 {
            // Initialize the FolderGroup model
            folderGroup := FolderGroup{}

            // Get data by GroupID and FolderID
            uadmin.Get(&folderGroup, "group_id = ? AND folder_id = ?", user.UserGroupID, d.FolderID)

            // Check whether there is a FolderGroup recird
            if folderGroup.ID != 0 {
                // Assign FolderGroup permission values to the variables
                Read = folderGroup.Read
                Add = folderGroup.Add
                Edit = folderGroup.Edit
                Delete = folderGroup.Delete
            }

            // Initialize the FolderUser model
            folderUser := FolderUser{}

            // Get data by UserID and FolderID
            uadmin.Get(&folderUser, "user_id = ? AND folder_id = ?", user.ID, d.FolderID)

            // Check whether there is a FolderUser record
            if folderUser.ID != 0 {
                // Assign FolderUser permission values to the variables
                Read = folderUser.Read
                Add = folderUser.Add
                Edit = folderUser.Edit
                Delete = folderUser.Delete
            }
        }

        // Document Permissions
        // Initialize the DocumentGroup model
        documentGroup := DocumentGroup{}

        // Get data by GroupID and DocumentID
        uadmin.Get(&documentGroup, "group_id = ? AND document_id = ?", user.UserGroupID, d.ID)

        // Check whether there is a DocumentGroup record
        if documentGroup.ID != 0 {
            // Assign DocumentGroup permission values to the variables
            Read = documentGroup.Read
            Add = documentGroup.Add
            Edit = documentGroup.Edit
            Delete = documentGroup.Delete
        }

        // Initialize the DocumentUser model
        documentUser := DocumentUser{}

        // // Get data by UserID and DocumentID
        uadmin.Get(&documentUser, "user_id = ? AND document_id = ?", user.ID, d.ID)

        // Check whether there is a DocumentUser record
        if documentUser.ID != 0 {
            // Assign DocumentUser permission values to the variables
            Read = documentUser.Read
            Add = documentUser.Add
            Edit = documentUser.Edit
            Delete = documentUser.Delete
        }

        // Return Read, Add, Edit, and Delete values
        return
    }

    // Count !
    func (d Document) Count(a interface{}, query interface{}, args ...interface{}) int {
        // Converts the query into a string
        Q := fmt.Sprint(query)

        // Checks whether the string contains a query and a UserID
        if strings.Contains(Q, "user_id = ?") {
            // Split the query part by part
            qParts := strings.Split(Q, " AND ")

            // Initialize tempArgs as an interface and tempQuery as a
            // string
            tempArgs := []interface{}{}
            tempQuery := []string{}

            // Loop the query every part
            for i := range qParts {
                // Checks whether the specific query part is not
                // equal to the UserID value
                if qParts[i] != "user_id = ?" {
                    // Append the arguments into the tempArgs
                    // variable
                    tempArgs = append(tempArgs, args[i])

                    // Append the specific query part into the
                    // tempQuery variable
                    tempQuery = append(tempQuery, qParts[i])
                }
            }
            // Concatenate the query to create a single string
            query = strings.Join(tempQuery, " AND ")

            // Assign tempArgs object into the args variable
            args = tempArgs
        }

        // Return the a, query, and args... inside the Count function
        // parameters
        return uadmin.Count(a, query, args...)
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

**folder.go**
^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    package models

    // Format is the name of the drop down list ...
    type Format int

    // PDF is the name of the drop down list value ...
    func (Format) PDF() Format {
        return 1
    }

    // TXT is the name of the drop down list value ...
    func (Format) TXT() Format {
        return 2
    }

    // Others is the name of the drop down list value ...
    func (Format) Others() Format {
        return 3
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

        // Register FolderGroup and FolderUser to Folder model
        uadmin.RegisterInlines(
            models.Folder{},
            map[string]string{
                "foldergroup": "FolderID",
                "folderuser":  "FolderID",
            },
        )

        // Register DocumentVersion, DocumentGroup, and DocumentUser to Document
        // model
        uadmin.RegisterInlines(
            models.Document{},
            map[string]string{
                "documentgroup":   "DocumentID",
                "documentuser":    "DocumentID",
                "documentversion": "DocumentID",
            },
        )

        // Assign Site Name value as "Document System"
        // NOTE: This code works only on first build.
        uadmin.SiteName = "Document System"

        // Activates a uAdmin server
        uadmin.StartServer()
    }


uadmin.db
---------
**Channels**
^^^^^^^^^^^^
`Back to Top`_

.. image:: assets/channelmodelupdate.png

**Documents**
^^^^^^^^^^^^^
`Back to Top`_

.. image:: assets/documentmodelupdate3.png

**Document Groups**
^^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. image:: assets/documentgroupmodelupdate.png

**Document Users**
^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. image:: assets/documentusermodelupdate.png

**Document Versions**
^^^^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. image:: assets/documentversionmodelupdate3.png

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

**Group Permissions**
^^^^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. image:: assets/grouppermissionmodelupdate.png

**Users**
^^^^^^^^^
`Back to Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part12.html#document-system-tutorial-part-12-custom-count-function-current-progress

.. image:: assets/usermodelupdate.png
>>>>>>> de25cdd8a29ca2bb2c2df08be00b703b967aaed5
