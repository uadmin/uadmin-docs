uAdmin Settings
===============
`Back to Settings in System Reference`_

.. _Back to Settings in System Reference: https://uadmin-docs.readthedocs.io/en/latest/system_reference.html#setting

Here are all build-in uAdmin Settings, their format, and how to use them in the project.

* `Allowed Hosts`_
* `Allowed IPs`_
* `API Disabled Add`_
* `API Disabled Delete`_
* `API Disabled Edit`_
* `API Disabled Read`_
* `API Disabled Schema`_
* `API Log Add`_
* `API Log Delete`_
* `API Log Edit`_
* `API Log Read`_
* `API Log Schema`_
* `Blocked IPs`_
* `Cache Permissions`_
* `Cache Sessions`_
* `Cache Translation`_
* `DebugDB`_
* `Email From`_
* `Email Password`_
* `Email SMTP Server`_
* `Email SMTP Server Port`_
* `Email Username`_
* `HTTP Log Format`_
* `Log Add`_
* `Log Delete`_
* `Log Edit`_
* `Log HTTP Requests`_
* `Log Read`_
* `Log Trail`_
* `Max Image Width`_
* `Max Image Height`_
* `Max Upload File Size`_
* `Optimize SQL Query`_
* `OTP Algorithm`_
* `OTP Digits`_
* `OTP Period`_
* `OTP Skew`_
* `Page Length`_
* `Password Attempts`_
* `Password Timeout`_
* `Port`_
* `Public Media`_
* `Rate Limit`_
* `Rate Limit Burst`_
* `Report Time Stamp`_
* `Reporting Level`_
* `Restrict Session IP`_
* `Retain Media Versions`_
* `Root URL`_
* `Site Name`_
* `System Metrics`_
* `Theme`_
* `Trail Logging Level`_
* `User Metrics`_

Allowed Hosts
-------------
`Back To Top`_

Allowed Hosts is a comma separated list of allowed hosts for the server to work. The default value if only for development and production domain should be added before deployment.

* **Default Value**: 0.0.0.0,127.0.0.1,localhost,::1
* **Data Type**: String
* **Code**: `uAdmin.AllowedHosts`_

.. _uAdmin.AllowedHosts: https://uadmin-docs.readthedocs.io/en/latest/api/network_functions.html#uadmin-allowedhosts

Allowed IPs
-----------
`Back To Top`_

Allowed IPs is a list of allowed IPs to access uAdmin interface in one of the following formats:

- "*" = Allow all
- "" = Allow none
- "192.168.1.1" = Only allow this IP
- "192.168.1.0/24" = Allow all IPs from 192.168.1.1 to 192.168.1.254

You can also create a list of the above formats using comma to separate them.

For example: "192.168.1.1, 192.168.1.2, 192.168.0.0/24"

* **Default Value**: \*
* **Data Type**: String
* **Code**: `uAdmin.AllowedIPs`_

.. _uAdmin.AllowedIPs: https://uadmin-docs.readthedocs.io/en/latest/api/network_functions.html#uadmin-allowedips

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, assign your IP address connected on your PC in the AllowedIPs function.

.. image:: assets/allowedipssetting.png

|

Now link the host to your IP and see what happens.

.. image:: ../assets/allowedipsresult.png

API Disabled Add
----------------
`Back To Top`_

API Disabled Add controls the data API’s disabled for add commands.

* **Default Value**: 0
* **Data Type**: Boolean
* **Code**: `uAdmin.APIDisabledAdd`_

.. _uAdmin.APIDisabledAdd: https://uadmin-docs.readthedocs.io/en/latest/dapi/disabled_functions.html#uadmin-apidisabledadd

API Disabled Delete
-------------------
`Back To Top`_

API Disabled Delete controls the data API’s disabled for delete commands.

* **Default Value**: 0
* **Data Type**: Boolean
* **Code**: `uAdmin.APIDisabledDelete`_

.. _uAdmin.APIDisabledDelete: https://uadmin-docs.readthedocs.io/en/latest/dapi/disabled_functions.html#uadmin-apidisableddelete

API Disabled Edit
-----------------
`Back To Top`_

API Disabled Edit controls the data API’s disabled for edit commands.

* **Default Value**: 0
* **Data Type**: Boolean
* **Code**: `uAdmin.APIDisabledEdit`_

.. _uAdmin.APIDisabledEdit: https://uadmin-docs.readthedocs.io/en/latest/dapi/disabled_functions.html#uadmin-apidisablededit

API Disabled Read
-----------------
`Back To Top`_

API Disabled Read controls the data API’s disabled for read commands.

* **Default Value**: 0
* **Data Type**: Boolean
* **Code**: `uAdmin.APIDisabledRead`_

.. _uAdmin.APIDisabledRead: https://uadmin-docs.readthedocs.io/en/latest/dapi/disabled_functions.html#uadmin-apidisabledread

API Disabled Schema
-------------------
`Back To Top`_

API Disabled Schema controls the data API’s disabled for schema commands.

* **Default Value**: 0
* **Data Type**: Boolean
* **Code**: `uAdmin.APIDisabledSchema`_

.. _uAdmin.APIDisabledSchema: https://uadmin-docs.readthedocs.io/en/latest/dapi/disabled_functions.html#uadmin-apidisabledschema

API Log Add
-----------
`Back To Top`_

API Log Add controls the data API’s logging for add commands.

* **Default Value**: 1
* **Data Type**: Boolean
* **Code**: `uAdmin.APILogAdd`_

.. _uAdmin.APILogAdd: https://uadmin-docs.readthedocs.io/en/latest/dapi/log_functions.html#uadmin-apilogadd

Prerequisites:

* `Add Multiple`_
* `Add One`_

.. _Add Multiple: https://uadmin-docs.readthedocs.io/en/latest/dapi.html#add-multiple
.. _Add One: https://uadmin-docs.readthedocs.io/en/latest/dapi.html#add-one

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, enable the APILogAdd then click Save button on the bottom right corner.

.. image:: assets/apilogaddsettingenabled.png

|

Now let's call this URL in the address bar to add multiple records in the Document model with the following information below:

**First Record**

* Name: Golang
* Author: John

**Second Record**

* Name: uAdmin
* Author: Adam

.. code-block:: bash

    # document is a model name
    # name and author are field names
    # __0 is the first index
    # __1 is the second index
    http://api.example.com/api/d/document/add/?_name__0=Golang&_author__0=John&_name__1=uAdmin&_author__1=Adam

Result:

.. code-block:: JSON

    {
        "id": [
            1,
            2
        ],
        "rows_count": 2,
        "status": "ok"
    }

It returns an array with a list of IDs for the newly created records.

Now go back to the uAdmin dashboard.

.. image:: assets/dashboardfromsettings.png

|

From here, click on "LOGS".

.. image:: ../assets/logshighlighted.png

|

As expected, the user's action in adding records through HTTP API was recorded in the Log model.

.. image:: assets/apilogaddenabled.png

|

Now let's try disabling the API Log Add in the Settings.

.. image:: assets/apilogaddsettingdisabled.png

|

Let's call this URL to add a new record in the Document model with the following information below:

* Name: Programming
* Author: Admin

.. code-block:: bash

    # document is a model name
    # name and author are field names
    http://api.example.com/api/d/document/add/?_name=Programming&_author=Admin

Result:

.. code-block:: JSON

    {
        "id": 3,
        "rows_count": 1,
        "status": "ok"
    }

It returns the ID of the newly created record.

Check the "LOGS" to see the result.

.. image:: assets/apilogadddisabled.png

|

As expected, the user's action in adding a record through HTTP API was not recorded in the Log model.

API Log Delete
--------------
`Back To Top`_

API Log Delete controls the data API's logging for delete commands.

* **Default Value**: 1
* **Data Type**: Boolean
* **Code**: `uAdmin.APILogDelete`_

.. _uAdmin.APILogDelete: https://uadmin-docs.readthedocs.io/en/latest/dapi/log_functions.html#uadmin-apilogdelete

Prerequisites:

* `Delete Multiple`_
* `Delete One`_

.. _Delete Multiple: https://uadmin-docs.readthedocs.io/en/latest/dapi.html#delete-multiple
.. _Delete One: https://uadmin-docs.readthedocs.io/en/latest/dapi.html#delete-one

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, enable the APILogDelete then click Save button on the bottom right corner.

.. image:: assets/apilogdeletesettingenabled.png

|

Suppose you have five records in the Item model.

.. image:: ../api/assets/itemfiverecords.png

|

Call this URL in the address bar to delete records where the name of an item contains "iPad".

.. code-block:: bash

    # item is a model name
    # name is a field name
    # __contains is an operator that will search for string values that contract
    http://api.example.com/api/d/item/delete/?name__contains=iPad

Result:

.. code-block:: JSON

    {
        "rows_count": 2,
        "status": "ok",
    }

It returns the status and the rows affected by your query.

Now go back to the uAdmin dashboard.

.. image:: assets/dashboardfromsettings.png

|

From here, click on "LOGS".

.. image:: ../assets/logshighlighted.png

|

As expected, the user's action in deleting records through HTTP API that contains "iPad" in the item name was recorded in the Log model.

.. image:: assets/apilogdeleteenabled.png

|

Now let's try disabling the API Log Delete in the Settings.

.. image:: assets/apilogdeletesettingdisabled.png

|

Let's call this URL in the address bar to delete the fourth record in the database.

.. code-block:: bash

    # item is a model name
    # 4 is an ID number
    http://api.example.com/api/d/item/delete/4/

Result:

.. code-block:: JSON

    {
        "rows_count": 1,
        "status": "ok"
    }

It returns the status and the rows affected by your query.

Check the "LOGS" to see the result.

.. image:: assets/apilogdeletedisabled.png

|

As expected, the user's action in deleting the fourth record through HTTP API was not recorded in the Log model.

API Log Edit
------------
`Back To Top`_

API Log Edit controls the data API's logging for edit commands.

* **Default Value**: 1
* **Data Type**: Boolean
* **Code**: `uAdmin.APILogEdit`_

.. _uAdmin.APILogEdit: https://uadmin-docs.readthedocs.io/en/latest/dapi/log_functions.html#uadmin-apilogedit

Prerequisites:

* `Edit Multiple`_
* `Edit One`_

.. _Edit Multiple: https://uadmin-docs.readthedocs.io/en/latest/dapi.html#edit-multiple
.. _Edit One: https://uadmin-docs.readthedocs.io/en/latest/dapi.html#edit-one

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, enable the APILogEdit then click Save button on the bottom right corner.

.. image:: assets/apilogeditsettingenabled.png

|

Suppose you have five records in the Item model where all iPad items have a rating of 4.

.. image:: ../api/assets/itemipadoldrating.png

|

Call this URL to edit the rating of all iPad items to a value of 5.

.. code-block:: bash

    # item is a model name
    # name is a field name
    # __contains is an operator that will search for string values that contract
    # rating=4&_rating=5 means that where rating is equal to 4, change the
    # rating value to 5
    http://api.example.com/api/d/item/edit/?rating=4&_rating=5

Result:

.. code-block:: JSON

    {
        "rows_count": 2,
        "status": "ok"
    }

It returns the status and the rows affected by your query.

Now go back to the uAdmin dashboard.

.. image:: assets/dashboardfromsettings.png

|

From here, click on "LOGS".

.. image:: ../assets/logshighlighted.png

|

As expected, the user's action in editing records through HTTP API was recorded in the Log model.

.. image:: assets/apilogeditenabled.png

|

Now let's try disabling the API Log Edit in the Settings.

.. image:: assets/apilogeditsettingdisabled.png

|

Suppose the first record in the Item model is named as "Robot".

.. image:: ../api/assets/itemfirstrecordrobot.png

|

Call this URL to edit the name of the first record in the database from "Robot" to "Supercomputer".

.. code-block:: bash

    # item is a model name
    # 1 is an ID number
    # name is a field name
    http://api.example.com/api/d/item/edit/1/?_name=Supercomputer

Result:

.. code-block:: JSON

    {
        "rows_count": 1,
        "status": "ok"
    }

It returns the status and the rows affected by your query.

Check the "LOGS" to see the result.

.. image:: assets/apilogeditdisabled.png

|

As expected, the user's action in editing the first record through HTTP API was not recorded in the Log model.

API Log Read
------------
`Back To Top`_

API Log Read controls the data API's logging for read commands.

* **Default Value**: 0
* **Data Type**: Boolean
* **Code**: `uAdmin.APILogRead`_

.. _uAdmin.APILogRead: https://uadmin-docs.readthedocs.io/en/latest/dapi/log_functions.html#uadmin-apilogread

Prerequisites:

* `Read Multiple`_
* `Read One`_

.. _Read Multiple: https://uadmin-docs.readthedocs.io/en/latest/dapi.html#read-multiple
.. _Read One: https://uadmin-docs.readthedocs.io/en/latest/dapi.html#read-one

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, enable the APILogRead then click Save button on the bottom right corner.

.. image:: assets/apilogreadsettingenabled.png

|

Suppose you have five records in the Item model.

.. image:: ../api/assets/itemfiverecords.png

|

Call this URL to read record(s) where rating is equal to 3.

.. code-block:: bash

    # item is a model name
    # rating is a field name
    http://api.example.com/api/d/item/read/?rating=3

Result:

.. image:: ../dapi/assets/readmultipleresult.png
   :align: center

|

It returns a list of records where rating is equal to 3.

Now go back to the uAdmin dashboard.

.. image:: assets/dashboardfromsettings.png

|

From here, click on "LOGS".

.. image:: ../assets/logshighlighted.png

|

As expected, the user's action in reading records through HTTP API was recorded in the Log model.

.. image:: assets/apilogreadenabled.png

|

Now let's try disabling the API Log Read in the Settings.

.. image:: assets/apilogreadsettingdisabled.png

|

Call this URL to read the second record in the Item model.

.. code-block:: bash

    # item is a model name
    # 2 is an ID number
    http://api.example.com/api/d/item/read/2/

Result:

.. image:: ../dapi/assets/readoneresult.png
   :align: center

|

It returns a JSON object representing an item where ID=2.

Check the "LOGS" to see the result.

.. image:: assets/apilogreaddisabled.png

|

As expected, the user's action in reading the second record through HTTP API was not recorded in the Log model.

API Log Schema
--------------
`Back To Top`_

API Log Schema controls the data API's logging for schema commands.

* **Default Value**: 1
* **Data Type**: Boolean
* **Code**: `uAdmin.APILogSchema`_

.. _uAdmin.APILogSchema: https://uadmin-docs.readthedocs.io/en/latest/dapi/log_functions.html#uadmin-apilogschema

Prerequisite:

* `Schema`_

.. _Schema: https://uadmin-docs.readthedocs.io/en/latest/dapi.html#schema

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, enable the APILogSchema then click Save button on the bottom right corner.

.. image:: assets/apilogschemasettingenabled.png

|

Suppose you have five records in the Item model.

.. image:: ../api/assets/itemfiverecords.png

|

Call this URL to read the full schema of the Item model.

.. code-block:: bash

    # item is a model name
    http://api.example.com/api/d/item/schema/

Result:

.. image:: ../dapi/assets/schemaresult.png
   :align: center

|

It returns a JSON object representing uAdmin's ModelSchema of the Item model.

Now go back to the uAdmin dashboard.

.. image:: assets/dashboardfromsettings.png

|

From here, click on "LOGS".

.. image:: ../assets/logshighlighted.png

|

As expected, the user's action in getting the schema of the Item model through HTTP API was recorded in the Log model.

.. image:: assets/apilogschemaenabled.png

|

Now let's try disabling the API Log Schema in the Settings.

.. image:: assets/apilogschemasettingdisabled.png

|

Recall this URL to read the full schema of the Item model.

.. code-block:: bash

    # item is a model name
    http://api.example.com/api/d/item/schema/

Check the "LOGS" to see the result.

.. image:: assets/apilogschemadisabled.png

|

As expected, the user's action in getting the schema of the Item model through HTTP API was not recorded in the Log model.

Blocked IPs
-----------
`Back To Top`_

BlockedIPs is a list of blocked IPs from accessing uAdmin interface in one of the following formats:

- "*" = Block all
- "" = Block none
- "192.168.1.1" = Only block this IP
- "192.168.1.0/24" = Block all IPs from 192.168.1.1 to 192.168.1.254

You can also create a list of the above formats using comma to separate them.

For example: "192.168.1.1, 192.168.1.2, 192.168.0.0/24"

* **Default Value**: ""
* **Data Type**: String
* **Code**: `uAdmin.BlockedIPs`_

.. _uAdmin.BlockedIPs: https://uadmin-docs.readthedocs.io/en/latest/api/network_functions.html#uadmin-blockedips

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, assign your IP address connected on your PC in the BlockedIPs function.

.. image:: assets/blockedipssetting.png

|

Now link the host to your IP and see what happens.

.. image:: ../assets/blockedipsresult.png

Quiz:

* `Miscellaneous Functions (3)`_

Cache Permissions
-----------------
`Back To Top`_

Cache Permissions allows uAdmin to store permissions data in memory.

* **Default Value**: 1
* **Data Type**: Boolean
* **Code**: `uAdmin.CachePermissions`_

.. _uAdmin.CachePermissions: https://uadmin-docs.readthedocs.io/en/latest/api/user_functions.html#uadmin-cachepermissions

Cache Sessions
--------------
`Back To Top`_

Cache Sessions allows uAdmin to store sessions data in memory.

* **Default Value**: 1
* **Data Type**: Boolean
* **Code**: `uAdmin.CacheSessions`_

.. _uAdmin.CacheSessions: https://uadmin-docs.readthedocs.io/en/latest/api/user_functions.html#uadmin-cachesessions

Cache Translation
-----------------
`Back To Top`_

Cache Translation allows a translation to store data in a cache memory.

* **Default Value**: 0
* **Data Type**: Boolean
* **Code**: `uAdmin.CacheTranslation`_

.. _uAdmin.CacheTranslation: https://uadmin-docs.readthedocs.io/en/latest/api/language_functions.html#uadmin-cachetranslation

Enable

.. image:: assets/cachetranslationsettingenabled.png

|

Disable

.. image:: assets/cachetranslationsettingdisabled.png

DebugDB
-------
`Back To Top`_

Debug DB prints all SQL statements going to DB.

* **Default Value**: 0
* **Data Type**: Boolean
* **Code**: `uAdmin.DebugDB`_

.. _uAdmin.DebugDB: https://uadmin-docs.readthedocs.io/en/latest/api/database_functions.html#uadmin-debugdb

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, enable the Debug DB then click Save button on the bottom right corner.

.. image:: assets/debugdbsettingenabled.png

|

Check your terminal to see the result.

.. code-block:: bash

    (/home/dev1/go/src/github.com/uadmin/uadmin/db.go:246) 
    [2019-11-05 15:44:45]  [0.51ms]  SELECT * FROM "languages"  WHERE "languages"."deleted_at" IS NULL AND ((code='en')) ORDER BY "languages"."id" ASC LIMIT 1  
    [1 rows affected or returned ] 

    (/home/dev1/go/src/github.com/uadmin/uadmin/db.go:158) 
    [2019-11-05 15:44:45]  [0.20ms]  SELECT * FROM "setting_categories"  WHERE "setting_categories"."deleted_at" IS NULL  
    [1 rows affected or returned ] 

    (/home/dev1/go/src/github.com/uadmin/uadmin/db.go:436) 
    [2019-11-05 15:44:45]  [1.56ms]  SELECT * FROM "settings"  WHERE "settings"."deleted_at" IS NULL AND ((category_id = 1))  
    [38 rows affected or returned ] 

    (/home/dev1/go/src/github.com/uadmin/uadmin/db.go:436) 
    [2019-11-05 15:44:55]  [0.35ms]  SELECT * FROM "ab_tests"  WHERE "ab_tests"."deleted_at" IS NULL AND ((active = true))  
    [0 rows affected or returned ] 

Quiz:

* `Miscellaneous Functions`_

.. _Miscellaneous Functions: https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/miscellaneous-functions.html

Email From
----------
`Back To Top`_

Email From identifies where the email is coming from.

* **Default Value**: ""
* **Data Type**: String
* **Code**: `uAdmin.EmailFrom`_

.. _uAdmin.EmailFrom: https://uadmin-docs.readthedocs.io/en/latest/api/email_functions.html#uadmin-emailfrom

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, assign the following email configurations.

.. image:: assets/emailconfigurationsetting.png

|

Let's go back to the uAdmin dashboard, go to Users model, create your own user account and set the email address based on your assigned EmailFrom in the code above.

.. image:: ../tutorial/assets/useremailhighlighted.png

|

Log out your account. At the moment, you suddenly forgot your password. How can we retrieve our account? Click Forgot Password at the bottom of the login form.

.. image:: ../tutorial/assets/forgotpasswordhighlighted.png

|

Input your email address based on the user account you wish to retrieve it back.

.. image:: ../tutorial/assets/forgotpasswordinputemail.png

|

Once you are done, open your email account. You will receive a password reset notification from the Todo List support. To reset your password, click the link highlighted below.

.. image:: ../tutorial/assets/passwordresetnotification.png

|

You will be greeted by the reset password form. Input the following information in order to create a new password for you.

.. image:: ../tutorial/assets/resetpasswordform.png

Once you are done, you can now access your account using your new password.

Quiz:

* `Email Functions`_

.. _Email Functions: https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/email-functions.html

Email Password
--------------
`Back To Top`_

Email Password sets the password of an email.

* **Default Value**: ""
* **Data Type**: String
* **Code**: `uAdmin.EmailPassword`_

.. _uAdmin.EmailPassword: https://uadmin-docs.readthedocs.io/en/latest/api/email_functions.html#uadmin-emailpassword

See `Email From`_ for the example.

Email SMTP Server
-----------------
`Back To Top`_

Email SMTP Server sets the name of the SMTP Server in an email.

* **Default Value**: ""
* **Data Type**: String
* **Code**: `uAdmin.EmailSMTPServer`_

.. _uAdmin.EmailSMTPServer: https://uadmin-docs.readthedocs.io/en/latest/api/email_functions.html#uadmin-emailsmtpserver

See `Email From`_ for the example.

Email SMTP Server Port
----------------------
`Back To Top`_

Email SMTP Server Port sets the port number of an SMTP Server in an email.

* **Default Value**: 0
* **Data Type**: Integer
* **Code**: `uAdmin.EmailSMTPServerPort`_

.. _uAdmin.EmailSMTPServerPort: https://uadmin-docs.readthedocs.io/en/latest/api/email_functions.html#uadmin-emailsmtpserverport

See `Email From`_ for the example.

Email Username
--------------
`Back To Top`_

Email Username sets the username of an email.

* **Default Value**: ""
* **Data Type**: String
* **Code**: `uAdmin.EmailUsername`_

.. _uAdmin.EmailUsername: https://uadmin-docs.readthedocs.io/en/latest/api/email_functions.html#uadmin-emailusername

See `Email From`_ for the example.

HTTP Log Format
---------------
`Back To Top`_

HTTP Log Format is the format used to log HTTP access.

* **Default Value**: %a %>s %B %U %D
* **Data Type**: String
* **Code**: `uAdmin.HTTPLogFormat`_

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

.. _uAdmin.HTTPLogFormat: https://uadmin-docs.readthedocs.io/en/latest/api/log_functions.html#uadmin-httplogformat

Log Add
-------
`Back To Top`_

Log Add adds a log when a record is added.

* **Default Value**: 1
* **Data Type**: Boolean
* **Code**: `uAdmin.LogAdd`_

.. _uAdmin.LogAdd: https://uadmin-docs.readthedocs.io/en/latest/api/log_functions.html#uadmin-logadd

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, enable the Log Add then click Save button on the bottom right corner.

.. image:: assets/logaddsettingenabled.png

|

Now go back to the uAdmin dashboard.

.. image:: assets/dashboardfromsettings.png

|

From here, click on "LOGS".

.. image:: ../assets/logshighlighted.png

|

Suppose that you have this record in your logs as shown below:

.. image:: ../api/assets/loginitialrecord.png

|

Go back to uAdmin dashboard then select "TODOS".

.. image:: ../assets/todoshighlightedlog.png

|

Click "Add New Todo".

.. image:: ../assets/addnewtodo.png

|

Input the name value in the text box (e.g. Read a book). Click Save button afterwards.

.. image:: ../assets/readabook.png

|

Result

.. image:: ../assets/readabookoutput.png

|

Now go back to the "LOGS" to see the result.

.. image:: ../assets/logaddtrueresult.png

|

Now let's try disabling the Log Add in the Settings.

.. image:: assets/logaddsettingdisabled.png

|

Go back to the uAdmin dashboard. Click on "TODOS" model and add another data inside it.

.. image:: ../assets/buildarobot.png

|

Result

.. image:: ../assets/buildarobotoutput.png

|

Now go back to the "LOGS" to see the result.

.. image:: ../assets/logaddfalseresult.png

|

As you can see, the log content remains the same. Well done!

See `Log Read`_ for the continuation.

Log Delete
----------
`Back To Top`_

Log Delete adds a log when a record is deleted.

* **Default Value**: 1
* **Data Type**: Boolean
* **Code**: `uAdmin.LogDelete`_

.. _uAdmin.LogDelete: https://uadmin-docs.readthedocs.io/en/latest/api/log_functions.html#uadmin-logdelete

Before you proceed to this example, see `Log Edit`_.

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, enable the Log Delete then click Save button on the bottom right corner.

.. image:: assets/logdeletesettingenabled.png

|

Now go back to the uAdmin dashboard.

.. image:: assets/dashboardfromsettings.png

|

From here, click on "LOGS".

.. image:: ../assets/logshighlighted.png

|

Suppose that you have this record in your logs as shown below:

.. image:: ../assets/logeditfalseresult.png

|

Go back to uAdmin dashboard then select "TODOS".

.. image:: ../assets/todoshighlightedlog.png

|

Select any of your existing data that you wish to delete (e.g. Washing the dishes)

.. image:: ../assets/washingthedishesdelete.png

|

Now go back to the "LOGS" to see the result.

.. image:: ../assets/logdeletetrueresult.png

|

Now let's try disabling the Log Delete in the Settings.

.. image:: assets/logdeletesettingdisabled.png

|

Go back to the uAdmin dashboard. Click on "TODOS" model and delete the remaining data (e.g. Read a book).

.. image:: ../assets/readabookdelete.png

|

Now go back to the "LOGS" to see the result.

.. image:: ../assets/logdeletefalseresult.png

|

As you can see, the log content remains the same. Well done!

Quiz:

* `Log Permissions`_

.. _Log Permissions: https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/log-permissions.html

Log Edit
--------
`Back To Top`_

Log Edit adds a log when a record is edited.

* **Default Value**: 1
* **Data Type**: Boolean
* **Code**: `uAdmin.LogEdit`_

.. _uAdmin.LogEdit: https://uadmin-docs.readthedocs.io/en/latest/api/log_functions.html#uadmin-logedit

Before you proceed to this example, see `Log Read`_.

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, enable the Log Edit then click Save button on the bottom right corner.

.. image:: assets/logeditsettingenabled.png


Now go back to the uAdmin dashboard.

.. image:: assets/dashboardfromsettings.png

|

From here, click on "LOGS".

.. image:: ../assets/logshighlighted.png

|

Suppose that you have this record in your logs as shown below:

.. image:: ../assets/logreadfalseresult.png

|

Go back to uAdmin dashboard then select "TODOS".

.. image:: ../assets/todoshighlightedlog.png

|

Select any of your existing data (e.g. Build a robot)

.. image:: ../assets/todoexistingdata.png

|

Change it to "Assembling the CPU" for instance.

.. image:: ../assets/assemblingthecpu.png

|

Result

.. image:: ../assets/assemblingthecpuoutput.png

|

Now go back to the "LOGS" to see the result.

.. image:: ../assets/logedittrueresult.png

|

Now let's try disabling the Log Edit in the Settings.

.. image:: assets/logeditsettingdisabled.png

|

Go back to the uAdmin dashboard. Click on "TODOS" model and modify any of your existing data (e.g. Assembling the CPU).

.. image:: ../assets/buildarobot.png

|

Change it to "Washing the dishes" for instance.

.. image:: ../assets/washingthedishes.png

|

Result

.. image:: ../assets/washingthedishesresult.png

|

Now go back to the "LOGS" to see the result.

.. image:: ../assets/logeditfalseresult.png

|

As you can see, the log content remains the same. Well done!

See `Log Delete`_ for the continuation.

Log HTTP Requests
-----------------
`Back To Top`_

Logs http requests to syslog

* **Default Value**: 1
* **Data Type**: Boolean
* **Code**: `uAdmin.LogHTTPRequests`_

.. _uAdmin.LogHTTPRequests: https://uadmin-docs.readthedocs.io/en/latest/api/log_functions.html#uadmin-loghttprequests

Log Read
--------
`Back To Top`_

Log Read adds a log when a record is read.

* **Default Value**: 0
* **Data Type**: Boolean
* **Code**: `uAdmin.LogRead`_

.. _uAdmin.LogRead: https://uadmin-docs.readthedocs.io/en/latest/api/log_functions.html#uadmin-logread

Before you proceed to this example, see `Log Add`_.

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, enable the Log Read then click Save button on the bottom right corner.

.. image:: assets/logreadsettingenabled.png

|

Now go back to the uAdmin dashboard.

.. image:: assets/dashboardfromsettings.png

|

From here, click on "LOGS".

.. image:: ../assets/logshighlighted.png

|

Suppose that you have this record in your logs as shown below:

.. image:: ../assets/logaddfalseresult.png

|

Go back to uAdmin dashboard then select "TODOS".

.. image:: ../assets/todoshighlightedlog.png

|

Select any of your existing data.

.. image:: ../assets/todoexistingdata.png

|

Result

.. image:: ../assets/readabook.png

|

Now go back to the "LOGS" to see the result.

.. image:: ../assets/logreadtrueresult.png

|

Now let's try disabling the Log Read in the Settings.

.. image:: assets/logreadsettingdisabled.png

|

Go back to the uAdmin dashboard. Click on "TODOS" model and add select any of your existing data.

.. image:: ../assets/todoexistingdata.png

|

Result

.. image:: ../assets/readabook.png

|

Now go back to the "LOGS" to see the result.

.. image:: ../assets/logreadfalseresult.png

|

As you can see, the log content remains the same. Well done!

See `Log Edit`_ for the continuation.

Log Trail
---------
`Back To Top`_

Log Trail stores Trail logs to syslog.

* **Default Value**: 0
* **Data Type**: Boolean
* **Code**: `uAdmin.LogTrail`_

.. _uAdmin.LogTrail: https://uadmin-docs.readthedocs.io/en/latest/api/log_functions.html#uadmin-logtrail

Max Image Width
---------------
`Back To Top`_

Max Image Width sets the maximum width of an image.

* **Default Value**: 800
* **Data Type**: Integer
* **Code**: `uAdmin.MaxImageWidth`_

.. _uAdmin.MaxImageWidth: https://uadmin-docs.readthedocs.io/en/latest/api/basic_functions.html#uadmin-maximagewidth

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, set the Max Image Width to 360 pixels and the Max Image Height to 240 pixels. Click Save on the bottom right corner afterwards.

.. image:: assets/maximagewidthheightsetting.png

|

uAdmin has a feature that allows you to customize your own profile. In order to do that, click the profile icon on the top right corner then select admin as highlighted below.

.. image:: ../tutorial/assets/adminhighlighted.png

|

By default, there is no profile photo inserted on the top left corner. If you want to add it in your profile, click the Choose File button to browse the image on your computer.

.. image:: ../tutorial/assets/choosefilephotohighlighted.png

|

Let's pick a photo that surpasses the MaxImageWidth and MaxImageHeight values.

.. image:: ../tutorial/assets/widthheightbackground.png
   :align: center

|

Once you are done, click Save Changes on the left corner and refresh the webpage to see the output.

.. image:: ../tutorial/assets/profilepicadded.png

As expected, the profile pic will be uploaded to the user profile that automatically resizes to 360x240 pixels.

Quiz:

* `Max Functions`_

Max Image Height
----------------
`Back To Top`_

Max Image Height sets the maximum height of an image.

* **Default Value**: 600
* **Data Type**: Integer
* **Code**: `uAdmin.MaxImageHeight`_

.. _uAdmin.MaxImageHeight: https://uadmin-docs.readthedocs.io/en/latest/api/basic_functions.html#uadmin-maximageheight

See `Max Image Width`_ for the example.

Max Upload File Size
--------------------
`Back To Top`_

Max Upload File Size is the maximum upload file size in bytes.

1MB = 1024 * 1024

* **Default Value**: 26214400
* **Data Type**: Integer
* **Code**: `uAdmin.MaxUploadFileSize`_

.. _uAdmin.MaxUploadFileSize: https://uadmin-docs.readthedocs.io/en/latest/api/basic_functions.html#uadmin-maxuploadfilesize

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, set the Max Upload File Size value to 1 MB. It is 1 multiplied by 1024 (Kilobytes) multiplied by 1024 (Bytes) = 1048576 Bytes.

.. image:: assets/maxuploadfilesizesetting.png

|

Now go to your profile and upload an image that exceeds the Max Upload File Size limit. If you click Save changes...

.. image:: ../tutorial/assets/noprofilepic.png

|

The profile picture has failed to upload in the user profile because the file size is larger than the limit.

Quiz:

* `Max Functions`_

.. _Max Functions: https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/max-functions.html

Optimize SQL Query
------------------
`Back To Top`_

Optimize SQL Query selects columns during rendering a form a list to visible fields.

* **Default Value**: 1
* **Data Type**: Boolean
* **Code**: `uAdmin.OptimizeSQLQuery`_

.. _uAdmin.OptimizeSQLQuery: https://uadmin-docs.readthedocs.io/en/latest/api/database_functions.html#uadmin-optimizesqlquery

Enable

.. image:: assets/optimizesqlquerysettingenabled.png

|

Disable

.. image:: assets/optimizesqlquerysettingdisabled.png

OTP Algorithm
-------------
`Back To Top`_

OTP Algorithm is the hashing algorithm of OTP. Other options are sha256 and sha512.

* **Default Value**: sha1
* **Data Type**: String
* **Code**: `uAdmin.OTPAlgorithm`_

.. _uAdmin.OTPAlgorithm: https://uadmin-docs.readthedocs.io/en/latest/api/security_functions.html#uadmin-otpalgorithm

You can apply any of these in Settings.

.. image:: assets/otpalgorithmsetting.png
   :align: center

OTP Digits
----------
`Back To Top`_

OTP Digits is the number of digits for the OTP.

* **Default Value**: 6
* **Data Type**: Integer
* **Code**: `uAdmin.OTPDigits`_

.. _uAdmin.OTPDigits: https://uadmin-docs.readthedocs.io/en/latest/api/security_functions.html#uadmin-otpdigits

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, set the OTP Digits value to 8. Click Save button on the bottom right corner afterwards.

.. image:: assets/otpdigitssetting.png

|

Make sure that OTP Required on the account you are using is enabled in the User model.

.. image:: assets/otprequiredenabled.png

|

Logout your account, relogin your account, and check your terminal afterwards to see the OTP verification code assigned by your system.

.. code-block:: bash

    [  INFO  ]   User: admin OTP: 90401068

As shown above, it has 8 OTP digits.

Quiz:

* `OTP Functions`_

OTP Period
----------
`Back To Top`_

OTP Period is the number of seconds for the OTP to change.

* **Default Value**: 30
* **Data Type**: Integer
* **Code**: `uAdmin.OTPPeriod`_

.. _uAdmin.OTPPeriod: https://uadmin-docs.readthedocs.io/en/latest/api/security_functions.html#uadmin-otpperiod

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, set the OTP Period to 10 seconds. Click Save button on the bottom right corner afterwards.

.. image:: assets/otpperiodsetting.png

|

Make sure that OTP Required on the account you are using is enabled in the User model.

.. image:: assets/otprequiredenabled.png

|

Logout your account, relogin your account, and check your terminal afterwards to see how the OTP code changes every 10 seconds by refreshing your browser.

.. code-block:: bash

    // Before refreshing your browser
    [  INFO  ]   User: admin OTP: 433452

    // After refreshing your browser in more than 10 seconds
    [  INFO  ]   User: admin OTP: 185157

Quiz:

* `OTP Functions`_

OTP Skew
--------
`Back To Top`_

OTP Skew is the number of minutes to search around the OTP.

* **Default Value**: 5
* **Data Type**: Integer
* **Code**: `uAdmin.OTPSkew`_

.. _uAdmin.OTPSkew: https://uadmin-docs.readthedocs.io/en/latest/api/security_functions.html#uadmin-otpskew

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, set the OTP Skew value to 2 minutes. Click Save button on the bottom right corner afterwards.

.. image:: assets/otpskewsetting.png

|

Make sure that OTP Required on the account you are using is enabled in the User model.

.. image:: assets/otprequiredenabled.png

|

Logout your account, relogin your account, and check your terminal afterwards to see the OTP verification code assigned by your system. Wait for more than two minutes and check if the OTP code is still valid.

After waiting for more than two minutes,

.. image:: ../assets/loginformwithotp.png

It redirects to the same webpage which means your OTP code is no longer valid.

Quiz:

* `OTP Functions`_

.. _OTP Functions: https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/otp.html

Page Length
-----------
`Back To Top`_

Page Length is the list view max number of records.

* **Default Value**: 100
* **Data Type**: Integer
* **Code**: `uAdmin.PageLength`_

.. _uAdmin.PageLength: https://uadmin-docs.readthedocs.io/en/latest/api/basic_functions.html#uadmin-pagelength

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, assign the Page Length value to 4.

.. image:: assets/pagelengthsettingenabled.png

|

Go to the Item model. Inside it you have 6 total elements. The elements in the item model will display 4 elements per page.

.. image:: ../tutorial/assets/pagelength.png

|

Quiz:

* `Miscellaneous Functions`_

.. _Miscellaneous Functions: https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/miscellaneous-functions.html

Password Attempts
-----------------
`Back To Top`_

Password Attempts is the maximum number of invalid password attempts before the IP address is blocked for some time from using the system.

* **Default Value**: 5
* **Data Type**: Integer
* **Code**: `uAdmin.PasswordAttempts`_

.. _uadmin.PasswordAttempts: https://uadmin-docs.readthedocs.io/en/latest/api/security_functions.html#uadmin-passwordattempts

Password Timeout
----------------
`Back To Top`_

Password Attempts is the maximum number of invalid password attempts before the IP address is blocked for some time from using the system.

* **Default Value**: 5
* **Data Type**: Integer
* **Code**: `uAdmin.PasswordTimeout`_

.. _uadmin.PasswordTimeout: https://uadmin-docs.readthedocs.io/en/latest/api/security_functions.html#uadmin-passwordtimeout

Port
----
`Back To Top`_

Port is the port used for http or https server.

* **Default Value**: 8080
* **Data Type**: Integer
* **Code**: `uAdmin.Port`_

.. _uadmin.Port: https://uadmin-docs.readthedocs.io/en/latest/api/network_functions.html#uadmin-port

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, apply **8000** as a port number. Click Save button on the bottom right corner afterwards.

.. image:: assets/portsetting.png

|

Rebuild your application. Check your terminal to see the result.

.. code-block:: bash

    [   OK   ]   Initializing DB: [13/13]
    [   OK   ]   Synching System Settings: [46/46]
    [   OK   ]   Server Started: http://0.0.0.0:8000
             ___       __          _
      __  __/   | ____/ /___ ___  (_)___
     / / / / /| |/ __  / __  __ \/ / __ \
    / /_/ / ___ / /_/ / / / / / / / / / /
    \__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/

In the Server Started, it will redirect you to port number **8000**.

Quiz:

* `IP Configuration`_

.. _IP Configuration: https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/ip-configuration.html

Public Media
------------
`Back To Top`_

Public Media allows public access to media handler without authentication.

* **Default Value**: 0
* **Data Type**: Boolean
* **Code**: `uAdmin.PublicMedia`_

.. _uadmin.PublicMedia: https://uadmin-docs.readthedocs.io/en/latest/api/security_functions.html#uadmin-publicmedia

For instance, my account was not signed in.

.. image:: ../tutorial/assets/loginform.png

|

And you want to access **travel.png** inside your media folder.

.. image:: ../assets/mediapath.png

|

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, enable the Public Media then click Save button on the bottom right corner.

.. image:: assets/publicmediasettingenabled.png

|

Logout your account. Access the image path in the URL to see the result.

.. image:: ../assets/publicmediaimage.png

|

Quiz:

* `Miscellaneous Functions`_

.. _Miscellaneous Functions: https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/miscellaneous-functions.html

Rate Limit
----------
`Back To Top`_

Rate Limit is the maximum number of requests/second for any unique IP.

* **Default Value**: 3
* **Data Type**: Integer
* **Code**: `uAdmin.RateLimit`_

.. _uadmin.RateLimit: https://uadmin-docs.readthedocs.io/en/latest/api/network_functions.html#uadmin-ratelimit

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, assign the rate limit to 1. Click Save button on the bottom right corner afterwards.

.. image:: assets/ratelimitsetting.png

|

Now go back to the uAdmin dashboard.

.. image:: assets/dashboardfromsettings.png

|

From here, hold the Ctrl Key on your keyboard then click any dashboard menu in the form really fast to add in a new tab and see what happens.

.. image:: ../assets/ratelimithighlighttab.png

|

The title bar name looks different in the last two tabs. Click any of them to see the result.

.. image:: ../assets/ratelimitresult.png

|

The website is crashed as expected. In fact that our rate limit is 1, it might take a long time to bring the website back to normal. To increase the recovery rate, adjust the rate limit to a higher value (e.g. 100) in the Settings.

.. image:: assets/ratelimit100setting.png

|

Do the same process as shown above. Afterwards, click any button in the form and you will see that the website is back to normal much faster.

.. image:: ../assets/websitebacktonormal.png

|

Quiz:

* `Rate Limit Functions`_

Rate Limit Burst
----------------
`Back To Top`_

Rate Limit Burst is the maximum number of requests for an idle user.

* **Default Value**: 3
* **Data Type**: Integer
* **Code**: `uAdmin.RateLimitBurst`_

.. _uadmin.RateLimitBurst: https://uadmin-docs.readthedocs.io/en/latest/api/network_functions.html#uadmin-ratelimitburst

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, assign the rate limit burst to 3. Click Save button on the bottom right corner afterwards.

.. image:: assets/ratelimitburstsetting.png

|

Now go back to the uAdmin dashboard.

.. image:: assets/dashboardfromsettings.png

|

From here, hold the Ctrl Key on your keyboard then click any dashboard menu in the form really fast to add in a new tab and see what happens.

.. image:: ../assets/ratelimithighlighttab.png

|

The title bar name looks different in the last two tabs. Click any of them to see the result.

.. image:: ../assets/ratelimitresult.png

|

The website is crashed because our request exceeds the limit that we have assigned.

Quiz:

* `Rate Limit Functions`_

.. _Rate Limit Functions: https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/rate-limit-functions.html

Report Time Stamp
-----------------
`Back To Top`_

Report Time Stamp set this to true to have a time stamp in your logs.

* **Default Value**: 0
* **Data Type**: Boolean
* **Code**: `uAdmin.ReportTimeStamp`_

.. _uadmin.ReportTimeStamp: https://uadmin-docs.readthedocs.io/en/latest/api/print_functions.html#uadmin-reporttimestamp

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, enable the Report Time Stamp then click Save button on the bottom right corner.

.. image:: assets/reporttimestampsettingenabled.png

|

Rebuild your application. Check your terminal to see the result.

.. code-block:: bash

    [   OK   ]   Initializing DB: [13/13]
    [   OK   ]   Synching System Settings: [46/46]
    2018/11/07 08:52:14 [   OK   ]   Server Started: http://0.0.0.0:8080
             ___       __          _
      __  __/   | ____/ /___ ___  (_)___
     / / / / /| |/ __  / __  __ \/ / __ \
    / /_/ / ___ / /_/ / / / / / / / / / /
    \__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/

Quiz:

* `Miscellaneous Functions`_

Reporting Level
---------------
`Back To Top`_

ReportingLevel is the standard reporting level.

There are 9 different levels:

* DEBUG = 0
* WORKING = 1
* INFO = 2
* OK = 3
* WARNING = 4
* ERROR = 5
* CRITICAL = 6
* ALERT = 7
* EMERGENCY = 8

--------------------------------------

* **Default Value**: 0
* **Data Type**: Integer
* **Code**: `uAdmin.ReportingLevel`_

.. _uadmin.ReportingLevel: https://uadmin-docs.readthedocs.io/en/latest/api/print_functions.html#uadmin-reportinglevel

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, set the Reporting Level to 1 to show that the debugging process is working. Click Save button on the bottom right corner afterwards.

.. image:: assets/reportinglevel1settingenabled.png

|

Rebuild your application. Check your terminal to see the result.

.. code-block:: bash

    [   OK   ]   Initializing DB: [13/13]
    [   OK   ]   Synching System Settings: [46/46]
    [   OK   ]   Server Started: http://0.0.0.0:8080
             ___       __          _
      __  __/   | ____/ /___ ___  (_)___
     / / / / /| |/ __  / __  __ \/ / __ \
    / /_/ / ___ / /_/ / / / / / / / / / /
    \__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/

What if I set the value to 5?

.. image:: assets/reportinglevel5settingenabled.png

|

Result

.. code-block:: bash

    [   OK   ]   Initializing DB: [13/13]
    [   OK   ]   Synching System Settings: [46/46]
             ___       __          _
      __  __/   | ____/ /___ ___  (_)___
     / / / / /| |/ __  / __  __ \/ / __ \
    / /_/ / ___ / /_/ / / / / / / / / / /
    \__,_/_/  |_\__,_/_/ /_/ /_/_/_/ /_/

The database was initialized. The server has started. However the error message did not show up because the reporting level is assigned to 5 which is ERROR.

Quiz:

* `Miscellaneous Functions`_

Restrict Session IP
-------------------
`Back To Top`_

Restrict Session IP is to block access of a user if their IP changes from their original IP during login.

* **Default Value**: 0
* **Data Type**: Boolean
* **Code**: `uAdmin.RestrictSessionIP`_

.. _uadmin.RestrictSessionIP: https://uadmin-docs.readthedocs.io/en/latest/api/network_functions.html#uadmin-restrictsessionip

Enable

.. image:: assets/restrictsessionipsettingenabled.png

|

Disable

.. image:: assets/restrictsessionipsettingdisabled.png

Retain Media Versions
---------------------
`Back To Top`_

Retain Media Versions is to allow the system to keep files uploaded even after they are changed. This allows the system to "Roll Back" to an older version of the file.

* **Default Value**: 1
* **Data Type**: Boolean
* **Code**: `uAdmin.RetainMediaVersions`_

.. _uadmin.RetainMediaVersions: https://uadmin-docs.readthedocs.io/en/latest/api/basic_functions.html#uadmin-retainmediaversions

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, disable the Retain Media Versions then click Save button on the bottom right corner.

.. image:: assets/retainmediaversionssettingdisabled.png

|

Now go back to the uAdmin dashboard.

.. image:: assets/dashboardfromsettings.png

|

From here, go to the Category model then click Add New Category button on the top right corner of the screen. Let's add a new record that includes the uploaded file from your computer (e.g. Windows Installation.pdf).

.. image:: ../api/assets/categoryinstallationrecord.png
   :align: center

|

Result:

.. image:: ../api/assets/categoryinstallationrecordresult.png

|

From your project folder, go to /media/files/(generated_folder_name)/. As expected, the "Windows Installation.pdf" file was saved on that path.

.. image:: ../assets/categoryinstallationsaved.png
   :align: center

|

Go back to your application and click the existing record that you have (e.g. Installation).

.. image:: ../api/assets/categoryinstallationrecordresult.png

|

Now update the file on that record (e.g. PDF file to ODT file).

.. image:: ../assets/categoryinstallationupdateodt.png
   :align: center

|

Result:

.. image:: ../assets/categoryinstallationresultodt.png

|

From your project folder, go to /media/files/(generated_folder_name)/. As expected, the "Windows Installation.pdf" file was updated from "Windows Installation.pdf" to "Windows Installation.odt" on the same folder.

.. image:: ../assets/categoryinstallationsavedodt.png
   :align: center

|

Now let's try enabling the Retain Media Versions in the Settings.

.. image:: assets/retainmediaversionssettingenabled.png

|

Go back to the uAdmin dashboard then go to the Category model. Update the file of the Installation record back to PDF.

.. image:: ../api/assets/categoryinstallationrecord.png
   :align: center

|

Result:

.. image:: ../api/assets/categoryinstallationrecordresult.png

|

From your project folder, go to /media/files/ path. Inside it, there are two generated folders that means the old version of the file is kept and the new version was saved in the different folder.

.. image:: ../assets/categoryinstallationtwofolders.png
   :align: center

|

Quiz:

* `Miscellaneous Functions (3)`_

.. _Miscellaneous Functions (3): https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/miscellaneous-functions-3.html

Root URL
--------
`Back To Top`_

Root URL is where the listener is mapped to.

* **Default Value**: /
* **Data Type**: String
* **Code**: `uAdmin.RootURL`_

.. _uadmin.RootURL: https://uadmin-docs.readthedocs.io/en/latest/api/basic_functions.html#uadmin-rooturl

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, assign the RootURL value as **/admin/**. Click Save button on the bottom right corner afterwards.

.. image:: assets/rooturlsetting.png

|

Rebuild your application and go to the home page with the RootURL in the address bar to see the result.

.. image:: ../assets/rooturladmin.png

|

Quiz:

* `Root URL and Site Name`_

Site Name
---------
`Back To Top`_

Site Name is the name of the website that shows on title and dashboard.

* **Default Value**: uAdmin
* **Data Type**: String
* **Code**: `uAdmin.SiteName`_

.. _uadmin.SiteName: https://uadmin-docs.readthedocs.io/en/latest/api/basic_functions.html#uadmin-sitename

Run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, assign the SiteName value as **Todo List**. Click Save button on the bottom right corner afterwards.

.. image:: assets/sitenamesetting.png

|

Now go back to the uAdmin dashboard.

.. image:: assets/dashboardfromsettings.png

|

Result

.. image:: ../tutorial/assets/todolisttitle.png

|

Quiz:

* `Root URL and Site Name`_

.. _Root URL and Site Name: https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/root-url-and-site-name.html

System Metrics
--------------
`Back To Top`_

System Metrics enables uAdmin system metrics to be recorded.

* **Default Value**: 0
* **Data Type**: Boolean
* **Code**: `uAdmin.SystemMetrics`_

.. _uadmin.SystemMetrics: https://uadmin-docs.readthedocs.io/en/latest/api/metric_functions.html#uadmin-systemmetrics

Theme
-----
`Back To Top`_

Theme is the name of the theme used in uAdmin.

* **Default Value**: default
* **Data Type**: String
* **Code**: `uAdmin.Theme`_

.. _uadmin.Theme: https://uadmin-docs.readthedocs.io/en/latest/api/basic_functions.html#uadmin-theme

From your project folder, click on "templates".

.. image:: ../assets/templatesfolderhighlighted.png

|

Inside templates, click on "uadmin".

.. image:: ../assets/uadminfolder.png

|

Create a new folder named "custom".

.. image:: ../assets/customfolderhighlighted.png

|

Inside custom folder, create a new file named "home.html".

.. image:: ../assets/homehtml.png

|

Inside home.html file, apply the following codes below to display a header that shows "Welcome to Home Page".

.. code-block:: html

    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">
        <title>Home Page</title>
    </head>
    <body>
        <h1>Welcome to Home Page</h1>
    </body>
    </html>

Now run your application and login your account. From uAdmin dashboard, click the wrench icon located on the top right.

.. image:: assets/wrenchiconfromdashboard.png

|

In the Settings, assigns the theme name as "custom". "custom" is the name of the folder inside the templates/uadmin path that uAdmin will run when the user starts the server. Click Save button on the bottom right corner afterwards.

.. image:: assets/themesetting.png

|

Rebuild your application to see the result.

.. image:: ../assets/welcometohomepage.png
   :align: center

|

Quiz:

* `Miscellaneous Functions (2)`_

.. _Miscellaneous Functions (2): https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/miscellaneous-functions-2.html

Trail Logging Level
-------------------
`Back To Top`_

Trail Logging Level is the minimum level to be logged into syslog.

* **Default Value**: 2
* **Data Type**: Integer
* **Code**: `uAdmin.TrailLoggingLevel`_

.. _uadmin.TrailLoggingLevel: https://uadmin-docs.readthedocs.io/en/latest/api/log_functions.html#uadmin-traillogginglevel

User Metrics
------------
`Back To Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/system-reference/setting.html#uadmin-settings

User Metrics enables the user metrics to be recorded.

* **Default Value**: 0
* **Data Type**: Boolean
* **Code**: `uAdmin.UserMetrics`_

.. _uadmin.UserMetrics: https://uadmin-docs.readthedocs.io/en/latest/api/metric_functions.html#uadmin-usermetrics
