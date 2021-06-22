uadmin.Schema
=============
`Back to Customizing Records Page`_

.. _Back to Customizing Records Page: https://uadmin-docs.readthedocs.io/en/latest/api/customizing_records.html

Schema is the global schema of the system.

Structure:

.. code-block:: go

    map[string]uadmin.ModelSchema

Used in the tutorial:

* `Document System Tutorial Part 15 - Schema Form Modifier`_

.. _Document System Tutorial Part 15 - Schema Form Modifier: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/part15.html

Examples:

* `Approval`_
* `CategoricalFilter`_
* `Choices`_
* `DefaultValue`_
* `DisplayName`_
* `Encrypt`_
* `ErrMsg`_
* `Filter`_
* `FormDisplay`_
* `Help`_
* `Hidden`_
* `ListDisplay`_
* `Max`_
* `Min`_
* `Pattern`_
* `PatternMsg`_
* `ProgressBar`_
* `ReadOnly`_
* `Required`_
* `Stringer`_
* `Type`_
* `UploadTo`_
* `WebCam`_

**Approval**
^^^^^^^^^^^^
`Back to Top`_

A feature used to set an approval permission in the field

Type:

.. code-block:: go

    bool

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Approval

Suppose you have "johndoe" account in the User model that is not an admin.

.. image:: ../assets/johndoerecord.png

|

User permissions are set for that user in Approvals and Friends models without an approval access.

.. image:: ../assets/johndoeuserpermissionapproval.png

|

And you have the Name field in the Friend model that has a primary key of 1 as shown below:

.. image:: ../assets/friendnamedefault.png
   :align: center

|

Let’s set a feature for an approval permission in the field. In order to do that, create a schema function of Friend model where the field name is Name then access Approval.

.. code-block:: go

    func main(){
        // Some codes

        // friend - Model Name
        // Name - Field Name
        uadmin.Schema["friend"].FieldByName("Name").Approval = true
    }

Run your application and login your account using "johndoe".

.. image:: ../assets/johndoelogin.png
   :align: center

|

As you can see, only the Approvals and Friends models are accessible in the dashboard. It is based on the user permission that was assigned on that user. Now click on "FRIENDS".

.. image:: ../assets/friendsapprovalhighlighted.png

|

Click "Add New Friend" to create a new record.

.. image:: ../assets/addnewfriend.png

|

Let's input the following field for this record.

.. image:: ../assets/johndoefriendrecord.png
   :align: center

|

Result

.. image:: ../assets/johndoenameempty.png

|

Based on the result, the name does not show up because we need an approval to someone who has approval access. Now logout johndoe account then login an admin account.

.. image:: ../assets/loginformadmin.png
   :align: center

|

From uAdmin dashboard, go to the Friends model, click the record that you have created, and in the input box of the Name field, there is a yellow warning sign on the left side that means it needs an approval to someone who has approval access. Now click the highlighted area below.

.. image:: ../assets/johndoeapprovalbutton.png
   :align: center

|

The admin will review the record that was created by a "johndoe" user. If you think his record is satisfactory, choose Approved in Approval Action then click Save and Continue on the bottom right corner of the screen.

.. image:: ../assets/johndoeapprovalreview.png

|

It is shown that the one who approved the record is an admin with an approved date. Now click View Record button to see the result.

.. image:: ../assets/johndoeviewrecord.png
   :align: center

|

The input Name field has a checkmark sign that means the record created by "johndoe" was approved.

.. image:: ../assets/johndoeapprovedrecord.png
   :align: center

**CategoricalFilter**
^^^^^^^^^^^^^^^^^^^^^
`Back to Top`_

A section of code that is designed to process user input and output request to produce a new data structure containing exactly those elements of the original data structure in the form of combo box

Type:

.. code-block:: go

    bool

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").CategoricalFilter

See `Filter`_ for the example.

**Choices**
^^^^^^^^^^^
`Back to Top`_

A struct for the list of choices

Type:

.. code-block:: go

    []uadmin.Choice

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Choices

Suppose you have the given source code in friend.go where Nationality is the type of the drop down list:

.. code-block:: go

    // Nationality ...
    type Nationality int

    // Chinese ...
    func (Nationality) Chinese() Nationality {
        return 1
    }

    // Filipino ...
    func (Nationality) Filipino() Nationality {
        return 2
    }

    // Others ...
    func (Nationality) Others() Nationality {
        return 3
    }

Let’s build a choice that includes Chinese and Filipino and excludes Others. In order to do that, create a schema function of Friend model where the field name is Nationality then access Choices.

.. code-block:: go

    func main(){
        // Some codes
        
        // friend - Model Name
        // Nationality - Field Name
        uadmin.Schema["friend"].FieldByName("Nationality").Choices = []uadmin.Choice{
            // K is the ID of the choice.
            // V is the value of the choice.
            {K: 0, V: " - "},
            {K: 1, V: "Chinese"},
            {K: 2, V: "Filipino"},
        }
    }

Run your application, go to the Friend model and click Add New Friend button on the top right corner of the screen. As expected, Chinese and Filipino choices are included in the list.

.. image:: ../assets/friendnationalitychoices.png

**DefaultValue**
^^^^^^^^^^^^^^^^
`Back to Top`_

A value assigned automatically if you want to add a new record

Type:

.. code-block:: go

    string

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").DefaultValue

Let's set a feature that assigns a value automatically when creating a new record. In order to do that, create a schema function of Friend model where the field name is Nationality then access DefaultValue.

.. code-block:: go

    func main(){
        // Some codes
        
        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").DefaultValue = "Type here"
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, "Type here” value has assigned automatically in the Name field.

.. image:: ../assets/categorydefaultvalue.png
   :align: center

**DisplayName**
^^^^^^^^^^^^^^^
`Back to Top`_

The name that you want to display in the model. It is an alias name.

Type:

.. code-block:: go

    string

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").DisplayName

Let’s replace the actual field name. In order to do that, create a schema function of Category model where the field name is Name then access DisplayName.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").DisplayName = "Display Name"
    }

Run your application and go to Category model. As expected, the name has changed to "CATEGORY NAME”.

.. image:: ../assets/categorydisplayname.png

**Encrypt**
^^^^^^^^^^^
`Back to Top`_

A feature used to encrypt the value in the database

Type:

.. code-block:: go

    bool

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Encrypt

Suppose you have two records in the Category model as shown below:

.. image:: ../assets/categorynametworecords.png

|

Let's encrypt the value of the Name field in the Category Model. In order to do that, create a schema function of Category model where the field name is Name then access Encrypt.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").Encrypt = true
    }

Run your application. From your project folder, open uadmin.db with DB Browser for SQLite.

.. image:: ../assets/uadmindbsqlite.png
   :align: center

|

Click on Execute SQL.

.. image:: ../assets/executesqlhighlighted.png
   :align: center

|

Get all records by typing this command: **SELECT \* FROM categories** then click the right arrow icon to execute your SQL command.

.. image:: ../assets/selectfromcategories.png
   :align: center

|

As expected, the Name value is encrypted in the database.

.. image:: ../assets/categorynameencrypt.png
   :align: center


**ErrMsg**
^^^^^^^^^^
`Back to Top`_

An error message displayed beneath the input field

Type:

.. code-block:: go

    string

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").ErrMsg

Let's set a feature where an error message will be displayed beneath the input Name field. In order to do that, create a schema function of Category model where the field name is Name then access ErrMsg.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").ErrMsg = "This field cannot be modified."
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, the error message was displayed beneath the input Name field.

.. image:: ../assets/categorynameerrmsg.png
   :align: center

**Filter**
^^^^^^^^^^
`Back to Top`_

A section of code that is designed to process user input and output request to produce a new data structure containing exactly those elements of the original data structure in the form of fill-up text

Type:

.. code-block:: go

    bool

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Filter

Let's set a feature where the user can filter the name in the Category model. In order to do that, create a schema function of Category model where the field name is Name then access Filter for input and CategoricalFilter for display.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").Filter = true
        uadmin.Schema["category"].FieldByName("Name").CategoricalFilter = true
    }

Run your application and go to the Category model. As expected, the combo box form highlighted on the right side is the CategoricalFilter to notify the user that the Category Name is the field that will be filtered. Now click the Filter button. Suppose you have two records as shown below:

.. image:: ../assets/categoryfilter.png

|

Assign "Work" in the Category Name. Click Filter button on the bottom right corner of the modal and see what happens.

.. image:: ../assets/categoryfilterwork.png

|

As expected, the Category record has filtered out where the name contains "Work".

.. image:: ../assets/categorynamefilterresult.png

**FormDisplay**
^^^^^^^^^^^^^^^
`Back to Top`_

A feature that will hide the field in the editing section of the model if the value returns false

Type:

.. code-block:: go

    bool

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").FormDisplay

Let's set a feature that will hide the field in the editing section of the Category model. In order to do that, create a schema function of Category model where the field name is Name then access FormDisplay.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").FormDisplay = false
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, the Name Field is now invisible in the Category model.

.. image:: ../assets/categorynameformdisplay.png
   :align: center

**Help**
^^^^^^^^
`Back to Top`_

A feature that gives solution(s) to solve advanced tasks

Type:

.. code-block:: go

    string

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Help

Let’s assign a help note in the Name field to instruct the user what to do on that field. In order to do that, create a schema function of Category model where the field name is Name then access Help.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").Help = "Input a category name for your Todo List."
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, the help note was displayed below the input Name field.

.. image:: ../assets/categorynamehelp.png
   :align: center

**Hidden**
^^^^^^^^^^
`Back to Top`_

A feature to hide the component in the editing section of the form

Type:

.. code-block:: go

    bool

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Hidden

Unlike in FormDisplay, the field will hide if the value is true. In order to hide the Name field in the Category model, create a schema function of Category model where the field name is Name then access Hidden.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").Hidden = true
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, the Name Field is now invisible in the Category model.

.. image:: ../assets/categorynameformdisplay.png

**ListDisplay**
^^^^^^^^^^^^^^^
`Back to Top`_

A feature that will hide the field in the viewing section of the model if the value returns false

Type:

.. code-block:: go

    bool

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").ListDisplay

Let's set a feature that will hide the field or column name in the viewing section of the Category model. In order to hide the Name field in the Category model, create a schema function of Category model where the field name is Name then access ListDisplay.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").ListDisplay = false
    }

Run your application and go to the Category model. As expected, the Name Field in Category Model is now invisible in the list.

.. image:: ../assets/categorynamelistdisplay.png
   :align: center

**Max**
^^^^^^^
`Back to Top`_

The maximum value the user can assign. It is applicable for numeric characters.

Type:

.. code-block:: go

    interface{}

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Max

Let's set a limitation where the user can assign a value up to 100. In order to do that, create a schema function of Todo model where the field name is Progress then access Max.

.. code-block:: go

    func main(){
        // Some codes

        // todo - Model Name
        // Progress - Field Name
        uadmin.Schema["todo"].FieldByName("Progress").Max = "100"
    }

Run your application and go to the Todo model. Let's put a numeric value beyond the maximum limit in the Progress field and see what happens.

.. image:: ../assets/todoprogressmax.png

**Min**
^^^^^^^
`Back to Top`_

The minimum value the user can assign. It is applicable for numeric characters.

Type:

.. code-block:: go

    interface{}

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Min

Let's set a limitation where the user can assign a value at least 0. In order to do that, create a schema function of Todo model where the field name is Progress then access Min.

.. code-block:: go

    func main(){
        // Some codes

        // todo - Model Name
        // Progress - Field Name
        uadmin.Schema["todo"].FieldByName("Progress").Min = "0"
    }

Run your application and go to the Todo model. Let's put a numeric value beyond the minimum limit in the Progress field and see what happens.

.. image:: ../assets/todoprogressmin.png

**Pattern**
^^^^^^^^^^^
`Back to Top`_

A regular expression

Type:

.. code-block:: go

    string

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Pattern

Let's set a feature where the user can assign letters only in the Name field. In order to do that, create a schema function of Category model where the field name is Name then access Pattern for regular expression and PatternMsg for an error message if the user did not match the requested format.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").Pattern = "^[a-zA-Z _]*$"
        uadmin.Schema["category"].FieldByName("Name").PatternMsg = "Your input must be a letter."
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. Let's assign a numeric value in the Name field. If you click Save, the system will prompt the user the the value of the Name field must assign letters only.

.. image:: ../assets/categorynamepattern.png
   :align: center

**PatternMsg**
^^^^^^^^^^^^^^
`Back to Top`_

An error message if the user assigns a value that did not match the requested format

Type:

.. code-block:: go

    string

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").PatternMsg

See `Pattern`_ for an example.

**ProgressBar**
^^^^^^^^^^^^^^^
`Back to Top`_

A feature used to measure the progress of the activity

Type:

.. code-block:: go

    map[float64]string

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").ProgressBar

Let's assign the value and the color of the progress bar. In order to do that, create a schema function of Todo model where the field name is Progress then access ProgressBar.

.. code-block:: go

    func main(){
        // Some codes

        // todo - Model Name
        // Progress - Field Name
        // 100.0 - maximum value
        // #07c - blue color
        uadmin.Schema["todo"].FieldByName("Progress").ProgressBar = map[float64]string{100.0: "#07c"}
    }

Run your application and go to the Todo model. As expected, the assigned values were applied to the progress bar.

.. image:: ../assets/todoprogressbar.png

**ReadOnly**
^^^^^^^^^^^^
`Back to Top`_

A field that cannot be modified

Type:

.. code-block:: go

    string

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").ReadOnly

Let's set a feature where the user cannot modify a Name field in the Category model. In order to do that, create a schema function of Category model where the field name is Name then access ReadOnly.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").ReadOnly = "true"
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. As expected, the Name field is now Read Only that means the value cannot be modified.

.. image:: ../assets/categorynamereadonly.png
   :align: center

**Required**
^^^^^^^^^^^^
`Back to Top`_

A field that user must perform the given task(s). It cannot be skipped or left empty.

Type:

.. code-block:: go

    bool

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Required

Let's set a feature where the user needs to fill up the Name field. If the value is empty, the user will prompt the user that the value of the Name field should be assigned. In order to do that, create a schema function of Category model where the field name is Name then access Required.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // Name - Field Name
        uadmin.Schema["category"].FieldByName("Name").Required = true
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. If you notice, there is an asterisk (\*) symbol located on the top right after the "Name:". Let's leave the Name field value as it is. If you click Save, the system will prompt the user that the Name must be filled out.

.. image:: ../assets/categorynamerequired.png
   :align: center

**Stringer**
^^^^^^^^^^^^
`Back to Top`_

A feature that assigns a field as a unique type

Type:

.. code-block:: go

    bool

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Stringer

Let's set a feature that assigns a field as a unique type. In order to do that, create a schema function of Friend model where the field name is Name then access Stringer.


.. code-block:: go

    func main(){
        // Some codes

        // friend - Model Name
        // Name - Field Name
        uadmin.Schema["friend"].FieldByName("Name").Stringer = true
    }

Go to `uadmin.Stringer`_ in the API Reference for the continuation.

.. _uadmin.Stringer: https://uadmin-docs.readthedocs.io/en/latest/api.html#uadmin-getstringer

**Type**
^^^^^^^^
`Back to Top`_

The field type (e.g. file, list, progress_bar)

Type:

.. code-block:: go

    string

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Type

Suppose you have this field in the Todo model as shown below:

.. image:: ../assets/todoprogressdefault.png

|

Let's convert the input type to the progress bar. In order to do that, create a schema function of Todo model where the field name is Progress then access Type.

.. code-block:: go

    func main(){
        // Some codes

        // todo - Model Name
        // Progress - Field Name
        uadmin.Schema["todo"].FieldByName("Progress").Type = "progress_bar"
    }

Run your application and go to the Todo model. As expected, the field type has changed from regular to a progress bar. However, the appearance does not look good because we have not assigned the value and color of the progress bar yet.

.. image:: ../assets/todoprogresstype.png

|

Let's improvise the appearance by assigning the value and the color of the progress bar. In order to do that, create a schema function of Todo model where the field name is Progress then access ProgressBar.

.. code-block:: go

    func main(){
        // Some codes

        // todo - Model Name
        // Progress - Field Name
        // 100.0 - maximum value
        // #07c - blue color
        uadmin.Schema["todo"].FieldByName("Progress").ProgressBar = map[float64]string{100.0: "#07c"}
    }

Run your application and go to the Todo model. As expected, the appearance of the progress bar is now good enough.

.. image:: ../assets/todoprogressbar.png

**UploadTo**
^^^^^^^^^^^^
`Back to Top`_

A path where to save the uploaded files

Type:

.. code-block:: go

    string

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").UploadTo

Let's set a feature where the uploaded file will save in the specified path on your project folder. In order to do that, create a schema function of Category model where the field name is File then access UploadTo.

.. code-block:: go

    func main(){
        // Some codes

        // category - Model Name
        // File - Field Name
        uadmin.Schema["category"].FieldByName("File").UploadTo = "/media/files/"
    }

Run your application, go to the Category model and click Add New Category button on the top right corner of the screen. Let's add a new record that includes the uploaded file from your computer (e.g. Windows Installation.pdf).

.. image:: ../assets/categoryinstallationrecord.png
   :align: center

|

Result:

.. image:: ../assets/categoryinstallationrecordresult.png

|

From your project folder, go to /media/files/(generated_folder_name)/. As expected, the "Windows Installation.pdf" file was saved on that path.

.. image:: ../assets/categoryfileuploadto.png
   :align: center

**WebCam**
^^^^^^^^^^
`Back to Top`_

.. _Back to Top: https://uadmin-docs.readthedocs.io/en/latest/api/customizing-records/schema.html#uadmin-schema

A feature which adds web can access directly from the image and file fields

Type:

.. code-block:: go

    string

Structure:

.. code-block:: go

    uadmin.Schema[ModelName].FieldByName("FieldName").Webcam

Let's set a feature that accesses a webcam directly into the image field. In order to do that, create a schema function of Friend model where the field name is ProfilePic then access Webcam.


.. code-block:: go

    func main(){
        // Some codes

        // friend - Model Name
        // ProfilePic - Field Name
        uadmin.Schema["friend"].FieldByName("ProfilePic").Type = "image"
        uadmin.Schema["friend"].FieldByName("ProfilePic").WebCam = true
    }

Run your application, go to the Friend model and click Add New Friend button on the top right corner of the screen. As expected, there is a camera tag on the right side of the ProfilePic input field. If you have a webcam installed on your computer, click that icon and see it for yourself.

.. image:: ../assets/webcamiconhighlighted.png
