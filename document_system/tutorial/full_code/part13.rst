Document System Tutorial Part 13 - Custom AdminPage function (Current Progress)
===============================================================================
`Back to Previous Page`_

.. _Back to Previous Page: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/part13.html

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

    // AdminPage !
    func (d Document) AdminPage(order string, asc bool, offset int, limit int, a interface{}, query interface{}, args ...interface{}) (err error) {
        // Checks whether the starting point is less than 0
        if offset < 0 {
            offset = 0
        }

        // Converts the userID into uint because SQL Database reads the model ID
        // as uint
        userID := uint(0)

        // Converts the query into a string
        Q := fmt.Sprint(query)

        // Checks whether the string contains a query and a UserID
        if strings.Contains(Q, "user_id = ?") {
            // Prints the result for debugging
            uadmin.Trail(uadmin.DEBUG, "1")

            // Split the query part by part
            qParts := strings.Split(Q, " AND ")

            // Initialize tempArgs as an interface and tempQuery as a string
            tempArgs := []interface{}{}
            tempQuery := []string{}

            // Loop the query every part
            for i := range qParts {
                // Checks whether the specific query part is not equal to the
                // UserID value
                if qParts[i] != "user_id = ?" {
                    // Append the arguments into the tempArgs variable
                    tempArgs = append(tempArgs, args[i])

                    // Append the specific query part into the tempQuery variable
                    tempQuery = append(tempQuery, qParts[i])
                } else {
                    // Prints the result for debugging
                    uadmin.Trail(uadmin.DEBUG, "UserID: %d", args[i])
                    
                    // A type assertion that provides access to an interface
                    // value's (args[i]) underlying concrete value (uint).
                    userID, _ = (args[i]).(uint)
                }
            }
            // Concatenate the query to create a single string
            query = strings.Join(tempQuery, " AND ")

            // Assign tempArgs object into the args variable
            args = tempArgs
        }

        // Checks whether the userID is equal to 0
        if userID == 0 {
            // Prints the result for debugging
            uadmin.Trail(uadmin.DEBUG, "2")

            // Fetch the error by using AdminPage function
            err = uadmin.AdminPage(order, asc, offset, limit, a, query, args...)

            // Returns an error
            return err
        }

        // Initialize the user variable that calls the User model
        user := uadmin.User{}

        // Gets the ID of the user
        uadmin.Get(&user, "id = ?", userID)

        // Initialize docList and tempList that calls the Document model
        docList := []Document{}
        tempList := []Document{}

        // Loop execution
        for {
            // Fetch the error by using AdminPage function
            err = uadmin.AdminPage(order, asc, offset, limit, &tempList, query, args)
            uadmin.Trail(uadmin.DEBUG, "8: offset:%d, limit:%d", offset, limit)

            // Checks whether an error is not equal to nil
            if err != nil {
                // Prints the result for debugging
                uadmin.Trail(uadmin.DEBUG, "3")

                // Cast a model of interface as an array of Document then assigns
                // the docList object
                *a.(*[]Document) = docList

                // Return an error
                return err
            }

            // Checks whether the length of tempList is equal to 0
            if len(tempList) == 0 {
                // Prints the result for debugging
                uadmin.Trail(uadmin.DEBUG, "4")

                // Cast a model of interface as an array of Document then assigns
                // the docList object
                *a.(*[]Document) = docList

                // Prints the result for debugging
                uadmin.Trail(uadmin.DEBUG, "a: %#v", a)

                // Returns nothing
                return nil
            }

            // Loop the tempList values
            for i := range tempList {
                // Initialize p variable as Read permission
                p, _, _, _ := tempList[i].GetPermissions(user)

                // Checks whether the Document has read permission access
                if p {
                    // Prints the result for debugging
                    uadmin.Trail(uadmin.DEBUG, "5")

                    // Append the tempList (Document) object to the docList
                    // variable
                    docList = append(docList, tempList[i])
                }

                // Checks whether the length of docList is equal to the limit
                if len(docList) == limit {
                    // Prints the result for debugging
                    uadmin.Trail(uadmin.DEBUG, "6")

                    // Cast a model of interface as an array of Document then
                    // assigns the docList object
                    *a.(*[]Document) = docList

                    // Returns nothing
                    return nil
                }
            }

            // Add limit values to the offset variable
            offset += limit
        }
        // Cast a model of interface as an array of Document then assigns the
        // docList object
        *a.(*[]Document) = docList

        // Prints the result for debugging
        uadmin.Trail(uadmin.DEBUG, "7")

        // Returns nothing
        return nil
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

**format.go**
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

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part13.html#document-system-tutorial-part-13-custom-adminpage-function-current-progress

.. image:: assets/usermodelupdate.png
