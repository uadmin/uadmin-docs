Log Functions
=============
`Back To uAdmin Functions List`_

.. _Back To uAdmin Functions List: https://uadmin-docs.readthedocs.io/en/latest/api.html#api-reference

In this section, we will cover the following functions in-depth listed below:

* `uadmin.Action`_
* `uadmin.HTTPLogFormat`_
* `uadmin.Log`_
* `uadmin.LogAdd`_
* `uadmin.LogDelete`_
* `uadmin.LogEdit`_
* `uadmin.LogHTTPRequests`_
* `uadmin.LogRead`_
* `uadmin.LogTrail`_
* `uadmin.Syslogf`_
* `uadmin.TrailLoggingLevel`_

uadmin.Action
-------------
`Back To Top`_

Action is the process of doing something where you can check the status of your activities in the uAdmin project.

Type:

.. code-block:: go

    int

There are 12 types of actions:

* **Added** - Saved a new record
* **Custom** - For any other action that you would like to log
* **Deleted** - Deleted a record
* **GetSchema** - Opened the full schema of the model
* **LoginDenied** - User invalid login
* **LoginSuccessful** - User login
* **Logout** - User logout
* **Modified** - Save an existing record
* **PasswordResetDenied** - A password reset attempt was rejected
* **PasswordResetRequest** - A password reset was received
* **PasswordResetSuccessful** - A password was reset
* **Read** - Opened a record

Open "LOGS" in the uAdmin dashboard. You can see the Action field inside it as shown below.

.. image:: /assets/actionhighlighted.png

|

Now go to the main.go. Let's add each methods of actions in the log.

.. code-block:: go

    func main(){
        // Some codes
        for i := 0; i < 12; i++ {
            // Initialize the log model
            log := uadmin.Log{}

            // Call each methods of action based on the specific loop count
            switch i {
            case 0:
                log.Action = uadmin.Action(0).Added()
            case 1:
                log.Action = uadmin.Action(0).Custom()
            case 2:
                log.Action = uadmin.Action(0).Deleted()
            case 3:
                log.Action = uadmin.Action(0).GetSchema()
            case 4:
                log.Action = uadmin.Action(0).LoginDenied()
            case 5:
                log.Action = uadmin.Action(0).LoginSuccessful()
            case 6:
                log.Action = uadmin.Action(0).Logout()
            case 7:
                log.Action = uadmin.Action(0).Modified()
            case 8:
                log.Action = uadmin.Action(0).PasswordResetDenied()
            case 9:
                log.Action = uadmin.Action(0).PasswordResetRequest()
            case 10:
                log.Action = uadmin.Action(0).PasswordResetSuccessful()
            default:
                log.Action = uadmin.Action(0).Read()
            }

            // Add the method to the logs
            log.Save()
        }
    }

Once you are done, rebuild your application. Check your "LOGS" again to see the result.

.. image:: /assets/actionlist.png

|

As expected, all types of actions were added in the logs. Good job man!

More examples of this function can be found in `uadmin.Log`_.

Quiz:

* `Action and Log`_

uadmin.HTTPLogFormat
--------------------
`Back To Top`_

HTTPLogFormat is the format used to log HTTP access.

Type:

.. code-block:: go

    string

Format:

.. code-block:: bash

    %a: Client IP address
    %{remote}p: Client port
    %A: Server hostname/IP
    %{local}p: Server port
    %U: Path
    %c: All coockies
    %{NAME}c: Cookie named 'NAME'
    %{GET}f: GET request parameters
    %{POST}f: POST request parameters
    %B: Response length
    %>s: Response code
    %D: Time taken in microseconds
    %T: Time taken in seconds
    %I: Request length

uadmin.Log
----------
`Back To Top`_

Log is a system in uAdmin that is used to add, modify, and delete the status of the user activities.

Structure:

.. code-block:: go

    type Log struct {
        Model
        Username  string    `uadmin:"filter;read_only"`
        Action    Action    `uadmin:"filter;read_only"`
        TableName string    `uadmin:"filter;read_only"`
        TableID   int       `uadmin:"filter;read_only"`
        Activity  string    `uadmin:"code;read_only" gorm:"type:longtext"`
        RollBack  string    `uadmin:"link;"`
        CreatedAt time.Time `uadmin:"filter;read_only"`
    }

There are 12 types of actions:

* **Added** - Saved a new record
* **Custom** - For any other action that you would like to log
* **Deleted** - Deleted a record
* **GetSchema** - Opened the full schema of the model
* **LoginDenied** - User invalid login
* **LoginSuccessful** - User login
* **Logout** - User logout
* **Modified** - Save an existing record
* **PasswordResetDenied** - A password reset attempt was rejected
* **PasswordResetRequest** - A password reset was received
* **PasswordResetSuccessful** - A password was reset
* **Read** - Opened a record

There are 5 functions that you can use in Log:

**ParseRecord** - It means to analyze a record specifically. It uses this format as shown below:

.. code-block:: go

    func(a reflect.Value, modelName string, ID uint, user *User, action Action, r *http.Request) (err error)

Parameters:

* **a reflect.Value**: An interface initialized in NewModel function
* **modelName string**: The name of the model in lowercase letters
* **ID uint**: The ID of the model
* **user \*User**: What account is using in the session
* **action Action**: An activity status
* **r \*http.Request**: A data structure that represents the client HTTP request

Go to `Example #2: ParseRecord function`_ to see how ParseRecord works.

**PasswordReset** - It keeps track when the user resets his password. It uses this format as shown below:

.. code-block:: go

    func(user string, action Action, r *http.Request) (err error)

Parameters:

* **user string**: An account username
* **action Action**: An activity status
* **r \*http.Request**: A data structure that represents the client HTTP request

Go to `Example #3: PasswordReset function`_ to see how PasswordReset works.

**Save()** - Saves the object in the database

**SignIn** - It keeps track when the user signs in his account. It uses this format as shown below:

.. code-block:: go

    func(user string, action Action, r *http.Request) (err error)

Parameters:

* **user string**: An account username
* **action Action**: An activity status
* **r \*http.Request**: A data structure that represents the client HTTP request

Go to `Example #4: SignIn function`_ to see how SignIn works.

**String()** - Returns the Log ID

Examples:

* `Example #1: Assigning values in Log fields`_
* `Example #2: ParseRecord function`_
* `Example #3: PasswordReset function`_
* `Example #4: SignIn function`_

.. _Example #1\: Assigning values in Log fields: https://uadmin-docs.readthedocs.io/en/latest/api/log-functions/log.html#example-1-assigning-values-in-log-fields
.. _Example #2\: ParseRecord function: https://uadmin-docs.readthedocs.io/en/latest/api/log-functions/log.html#example-2-parserecord-function
.. _Example #3\: PasswordReset function: https://uadmin-docs.readthedocs.io/en/latest/api/log-functions/log.html#example-3-passwordreset-function
.. _Example #4\: SignIn function: https://uadmin-docs.readthedocs.io/en/latest/api/log-functions/log.html#example-4-signin-function

Page:

.. toctree::
   :maxdepth: 1

   log-functions/log

Quiz:

* `Action and Log`_

.. _Action and Log: https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/action-and-log.html

uadmin.LogAdd
-------------
`Back To Top`_

LogAdd adds a log when a record is added.

Type:

.. code-block:: go

    bool

To assign a value within an application, visit `Log Add`_ page for an example.

.. _Log Add: https://uadmin-docs.readthedocs.io/en/latest/system-reference/setting.html#log-add

To assign a value in the code, follow this approach:

Go to the main.go and apply this function to "true”. Put it above the uadmin.Register.

.. code-block:: go

    func main() {
        // NOTE: This code works only on first build.
        uadmin.LogAdd = true

        // ----- IF YOU RUN YOUR APPLICATION AGAIN, DO THIS BELOW -----

        // Assign the Log Add value to true
        setting := uadmin.Setting{}
        uadmin.Get(&setting, "code = ?", "uAdmin.LogAdd")
        setting.ParseFormValue([]string{"true"})
        setting.Save()

        uadmin.Register(
            // Some codes
        )

Run your application and go to "LOGS" model.

.. image:: assets/logshighlighted.png

|

Suppose that you have this record in your logs as shown below:

.. image:: assets/loginitialrecord.png

|

Go back to uAdmin dashboard then select "TODOS".

.. image:: /assets/todoshighlightedlog.png

|

Click "Add New Todo".

.. image:: /assets/addnewtodo.png

|

Input the name value in the text box (e.g. Read a book). Click Save button afterwards.

.. image:: /assets/readabook.png

|

Result

.. image:: /assets/readabookoutput.png

|

Now go back to the "LOGS" to see the result.

.. image:: /assets/logaddtrueresult.png

|

Exit your application for a while. Go to the main.go once again. This time, apply this function to "false".

.. code-block:: go

    func main() {
        // Assign the Log Add value to false
        setting := uadmin.Setting{}
        uadmin.Get(&setting, "code = ?", "uAdmin.LogAdd")
        setting.ParseFormValue([]string{"false"})
        setting.Save()

        uadmin.Register(
            // Some codes
        )

|

Rebuild and run your application. Go to "TODOS" model and add another data inside it.

.. image:: /assets/buildarobot.png

|

Result

.. image:: /assets/buildarobotoutput.png

|

Now go back to the "LOGS" to see the result.

.. image:: /assets/logaddfalseresult.png

|

As you can see, the log content remains the same. Well done!

See `uadmin.LogRead`_ for the continuation.

uadmin.LogDelete
----------------
`Back To Top`_

LogAdd adds a log when a record is deleted.

Type:

.. code-block:: go

    bool

To assign a value within an application, visit `Log Delete`_ page for an example.

.. _Log Delete: https://uadmin-docs.readthedocs.io/en/latest/system-reference/setting.html#log-delete

To assign a value in the code, follow this approach:

Before you proceed to this example, see `uadmin.LogEdit`_.

Go to the main.go and apply the LogDelete function to "true”. Put it above the uadmin.Register.

.. code-block:: go

    func main() {
        uadmin.LogAdd = false
        uadmin.LogRead = false
        uadmin.LogEdit = false

        // NOTE: This code works only on first build.
        uadmin.LogDelete = true

        // ----- IF YOU RUN YOUR APPLICATION AGAIN, DO THIS BELOW -----

        // Assign the Log Delete value to true
        setting := uadmin.Setting{}
        uadmin.Get(&setting, "code = ?", "uAdmin.LogDelete")
        setting.ParseFormValue([]string{"true"})
        setting.Save()

        uadmin.Register(
            // Some codes
        )

Run your application and go to "LOGS" model.

.. image:: assets/logshighlighted.png

|

Suppose that you have this record in your logs as shown below:

.. image:: /assets/logeditfalseresult.png

|

Go back to uAdmin dashboard then select "TODOS".

.. image:: /assets/todoshighlightedlog.png

|

Select any of your existing data that you wish to delete (e.g. Washing the dishes)

.. image:: /assets/washingthedishesdelete.png

|

Now go back to the "LOGS" to see the result.

.. image:: /assets/logdeletetrueresult.png

|

Exit your application for a while. Go to the main.go once again. This time, apply the LogDelete function to "false".

.. code-block:: go

    func main() {
        uadmin.LogAdd = false
        uadmin.LogRead = false
        uadmin.LogEdit = false

        // Assign the Log Delete value to false
        setting := uadmin.Setting{}
        uadmin.Get(&setting, "code = ?", "uAdmin.LogDelete")
        setting.ParseFormValue([]string{"false"})
        setting.Save()

        uadmin.Register(
            // Some codes
        )

Rebuild and run your application. Go to "TODOS" model and delete the remaining data (e.g. Read a book).

.. image:: /assets/readabookdelete.png

|

Now go back to the "LOGS" to see the result.

.. image:: /assets/logdeletefalseresult.png

|

As you can see, the log content remains the same. Well done!

Quiz:

* `Log Permissions`_

.. _Log Permissions: https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/log-permissions.html

uadmin.LogEdit
--------------
`Back To Top`_

LogAdd adds a log when a record is edited.

Type:

.. code-block:: go

    bool

To assign a value within an application, visit `Log Edit`_ page for an example.

.. _Log Edit: https://uadmin-docs.readthedocs.io/en/latest/system-reference/setting.html#log-edit

To assign a value in the code, follow this approach:

Before you proceed to this example, see `uadmin.LogRead`_.

Go to the main.go and apply the LogEdit function to "true”. Put it above the uadmin.Register.

.. code-block:: go

    func main() {
        uadmin.LogAdd = false
        uadmin.LogRead = false

        // NOTE: This code works only on first build.
        uadmin.LogEdit = true

        // ----- IF YOU RUN YOUR APPLICATION AGAIN, DO THIS BELOW -----

        // Assign the Log Edit value to true
        setting := uadmin.Setting{}
        uadmin.Get(&setting, "code = ?", "uAdmin.LogEdit")
        setting.ParseFormValue([]string{"true"})
        setting.Save()

        uadmin.Register(
            // Some codes
        )

Run your application and go to "LOGS" model.

.. image:: assets/logshighlighted.png

|

Suppose that you have this record in your logs as shown below:

.. image:: /assets/logreadfalseresult.png

|

Go back to uAdmin dashboard then select "TODOS".

.. image:: /assets/todoshighlightedlog.png

|

Select any of your existing data (e.g. Build a robot)

.. image:: /assets/todoexistingdata.png

|

Change it to "Assembling the CPU" for instance.

.. image:: /assets/assemblingthecpu.png

|

Result

.. image:: /assets/assemblingthecpuoutput.png

|

Now go back to the "LOGS" to see the result.

.. image:: /assets/logedittrueresult.png

|

Exit your application for a while. Go to the main.go once again. This time, apply the LogEdit function to "false".

.. code-block:: go

    func main() {
        uadmin.LogAdd = false
        uadmin.LogRead = false

        // Assign the Log Edit value to false
        setting := uadmin.Setting{}
        uadmin.Get(&setting, "code = ?", "uAdmin.LogEdit")
        setting.ParseFormValue([]string{"false"})
        setting.Save()

        uadmin.Register(
            // Some codes
        )

Rebuild and run your application. Go to "TODOS" model and modify any of your existing data (e.g. Assembling the CPU).

.. image:: /assets/buildarobot.png

|

Change it to "Washing the dishes" for instance.

.. image:: /assets/washingthedishes.png

|

Result

.. image:: /assets/washingthedishesresult.png

|

Now go back to the "LOGS" to see the result.

.. image:: /assets/logeditfalseresult.png

|

As you can see, the log content remains the same. Well done!

See `uadmin.LogDelete`_ for the continuation.

uadmin.LogHTTPRequests
----------------------
`Back To Top`_

LogHTTPRequests logs http requests to syslog.

Type:

.. code-block:: go

    bool

uadmin.LogRead
--------------
`Back To Top`_

LogRead adds a log when a record is read.

Type:

.. code-block:: go

    bool

To assign a value within an application, visit `Log Read`_ page for an example.

.. _Log Read: https://uadmin-docs.readthedocs.io/en/latest/system-reference/setting.html#log-read

To assign a value in the code, follow this approach:

Before you proceed to this example, see `uadmin.LogAdd`_.

Go to the main.go and apply the LogRead function to "true”. Put it above the uadmin.Register.

.. code-block:: go

    func main() {
        uadmin.LogAdd = false

        // NOTE: This code works only on first build.
        uadmin.LogRead = true

        // ----- IF YOU RUN YOUR APPLICATION AGAIN, DO THIS BELOW -----

        // Assign the Log Read value to true
        setting := uadmin.Setting{}
        uadmin.Get(&setting, "code = ?", "uAdmin.LogRead")
        setting.ParseFormValue([]string{"true"})
        setting.Save()

        uadmin.Register(
            // Some codes
        )

Run your application and go to "LOGS" model.

.. image:: assets/logshighlighted.png

|

Suppose that you have this record in your logs as shown below:

.. image:: /assets/logaddfalseresult.png

|

Go back to uAdmin dashboard then select "TODOS".

.. image:: /assets/todoshighlightedlog.png

|

Select any of your existing data.

.. image:: /assets/todoexistingdata.png

|

Result

.. image:: /assets/readabook.png

|

Now go back to the "LOGS" to see the result.

.. image:: /assets/logreadtrueresult.png

|

Exit your application for a while. Go to the main.go once again. This time, apply the LogRead function to "false".

.. code-block:: go

    func main() {
        uadmin.LogAdd = false

        // Assign the Log Read value to false
        setting := uadmin.Setting{}
        uadmin.Get(&setting, "code = ?", "uAdmin.LogRead")
        setting.ParseFormValue([]string{"false"})
        setting.Save()

        uadmin.Register(
            // Some codes
        )

Rebuild and run your application. Go to "TODOS" model and add select any of your existing data.

.. image:: /assets/todoexistingdata.png

|

Result

.. image:: /assets/readabook.png

|

Now go back to the "LOGS" to see the result.

.. image:: /assets/logreadfalseresult.png

|

As you can see, the log content remains the same. Well done!

See `uadmin.LogEdit`_ for the continuation.

uadmin.LogTrail
---------------
`Back To Top`_

LogTrail stores Trail logs to syslog.

Type:

.. code-block:: go

    bool

uadmin.Syslogf
--------------
`Back To Top`_

Function:

.. code-block:: go

    // For Mac / Linux
    func(level int, msg string, a ...interface{})

    // For Windows
    func(level int, msg string, a ...interface{}) error

uadmin.TrailLoggingLevel
------------------------
`Back To Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/api/log_functions.html#log-functions

TrailLoggingLevel is the minimum level to be logged into syslog.

Type:

.. code-block:: go

    int
