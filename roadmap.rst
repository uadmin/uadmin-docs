Roadmap
=======
* `Version 0.6.1 Beetle (2021-04-27)`_
* `Version 0.5.3 Atlas Moth (2021-04-08)`_
* `Version 0.5.2 Atlas Moth (2020-08-20)`_
* `Version 0.5.1 Atlas Moth (2020-08-07)`_
* `Version 0.5.0 Atlas Moth (2020-08-02)`_
* `Version 0.4.0 (2019-12-26)`_
* `Version 0.3.0 (2019-07-02)`_
* `Version 0.2.2 (2019-06-24)`_
* `Version 0.2.0 (2019-06-20)`_
* `Version 0.1.1 (2019-05-11)`_
* `Version 0.1.0 (2018-12-23)`_

Version 0.6.1 Beetle (2021-04-27)
---------------------------------
`Back to Top`_

**Added**
^^^^^^^^^
* Support for getting static files when running `uadmin prepare` from modules instead of source folder
* `uadmin prepare` now accepts a new parameter `--src` that allows to overide the default behavior and get the static files from source instead of module
* `uadmin prepare` takes go.mod into consideration to decide where to get the static files from. This allows developers to develop multiple applications that uses different version of uAdmin. If you use replace directive inside go.mod, the prepare tool will copy the files from the local folder as instructed in the replace directive.
* Two new settings "Logo" and "FavIcon" to customize your application even further.

**Changed**
^^^^^^^^^^^
* Added `.DS_Store` and `.vscode` to .gitignore

**Removed**
^^^^^^^^^^^
* `uadmin publish` is removed. It used to be an internal development tool and was cool to have it runnig publicly for a while

**Fixed**
^^^^^^^^^
* dAPI didn't have docs for `method` and fixed some typos
* Linting (So much linting)

**Security**
^^^^^^^^^^^^
* The system reads two envirnment variables `UADMIN_USER` and `UADMIN_PASS` for new deployments to create admin username and password. If these environment variables do not exist, uAdmin will user "admin" and "admin" for username and password.



Version 0.5.3 Atlas Moth (2021-04-08)
-------------------------------------
`Back to Top`_

**Changed**
^^^^^^^^^^^
* Drop support for GO version 1.11 and 1.12

**Fixed**
^^^^^^^^^
* Fixed adding more than 10 items in dAPI
* Fixed saving of language not removing the previous default language



Version 0.5.2 Atlas Moth (2020-08-20)
-------------------------------------
`Back to Top`_

**Fixed**
^^^^^^^^^
* Added CSRF token to inlines form for deleting
* User was overwritten with old user on logout when using cache sessions



Version 0.5.1 Atlas Moth (2020-08-07)
-------------------------------------
`Back to Top`_

**Changed**
^^^^^^^^^^^
* dAPI function `method` can return a value if the method called has a return. Note: if you have a return, you cannot use `$next` to redirect.

**Fixed**
^^^^^^^^^
* Fixed false positive SQL Injection in dAPI join.
* Fixed false detection in `customGet` for private fields of type `[]struct` as an M2M field.
* Typo in uadmin command line tool.

**Security**
^^^^^^^^^^^^
* CSRF protection in dAPI in functions: `add`, `edit`, `delete` and `method`.
* Template function `CSRF` implemented in `uadmin.RenderHTML` and `uadmin.RenderMultiHTML`. It returns anti CSRF token.
* `uadmin.IsAuthenticated` recognizes `nouser` sessions. These sessions are for users who are not authenticated in the system. To set a session cookie, user `SetSessionCookie`.
* `uadmin.SetSessionCookie` receives a pointer to a session and sets the session cookie in a secure way. If you pass a `nil` to the session, the session will be created as a `nouser` session which is still a session but gives the user to access as an authenticated user. These sessions can be used to protect against CSRF attacks in case you have a public API.



Version 0.5.0 Atlas Moth (2020-08-02)
-------------------------------------
`Back to Top`_

**Added**
^^^^^^^^^
* PreQuery and PostQuery handler for dAPI
* `method` command in dAPI to run model's methods
* Windows support for syslog
* `$preload` in dAPI
* RenderMultiHTML to render nested templates
* `$choices` in schema command in dAPI to preload FK and M2M choices
* Support for golang Modules
* Added search in dAPI using `$q`
* Reset button for ABTest
* Added delete functionality for file and image from UI and dAPI
* dAPI schema now translates the model based on your language cookie
* dAPI now supports M2M in add and delete functions
* `$distinct` in dAPI read function

**Changed**
^^^^^^^^^^^
* Dropping support for Golang 1.10
* Changed Excel export library to excelize
* GetDefaultLanguage and GetActiveLanguages are public now
* TranslateSchema is public now
* Model method `GetImageSize() (int, int)` to customize image size

**Fixed**
^^^^^^^^^
* Fixed bug with dAPI __ filters
* Fixed last insert ID in MySQL
* Fixed dAPI clearing file and image fields
* Fixed a bug with Aggregate column
* HideInDashboard works for existing models
* DashboardMenu changes icon size to 128X128 pixels

**Security**
^^^^^^^^^^^^
* PasswordAttempts and PasswordTimeout settings to protect limit invalid password attempts
* CheckRateLimit limits whole IP instead of IP and port combination
* CSRF protection in UI and dAPI and public function `CheckCSRF`
* SQL injection checking in dAPI, export and public function `SQLInjection`
* Added AllowedHosts setting to limit the domains/IPs for password reset
* Link fields get `x-scrf-token` added automatically in UI
* `session` cookie uses SameSite=SameSiteStrictMode
* Prevent navigation attacks in Theme setting and file upload



Version 0.4.0 (2019-12-26)
--------------------------
`Back to Top`_

**Added**
^^^^^^^^^
* Implemented request rate limits to protect from DDoS
* Implemented AB/Testing system
* Implemented dAPI which is an API to access model data
* Implemented CacheSessions and CachePermissions for direct in memory access to sessions and permissions
* Implemented a Metrics system using the following function: `SetMetric`, `IncrementMetric`, `TimeMetric` and `NewMetric`
* Trail can log to syslog
* HTTP requests can be logged to syslog
* Added `uadmin.Handler(func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request)` to enable syslog for HTTP requests
* Add OptimizeSQLQuery mode
* Handle ctrl+F in Home and List view to focus on the search field instead of the native search
* Added `stringer` meta tag
* Added two API end points `/api/get_models` and `/api/get_fields?m={MODEL_NAME}`

**Changed**
^^^^^^^^^^^
* Changed `Login(*http.Request, string, string) (*User, bool)` to `Login(*http.Request, string, string) (*Session, bool)`
* Changed `Login2FA(*http.Request, string, string, string) *User` to `Login2FA(*http.Request, string, string, string) *Session`
* Changed `HTMLContext` to `RenderHTML` that has support for templates functions
* URL filter place holders are capital letters `{USERNAME}`, `{USERID}` and `{NOW}`
* Search is allowed for `list_exclude` fields.
* Changed the search API path to `/api/search/`
* Make static handler a public function `uadmin.StaticHandler(http.ResponseWriter, *http.Request)`
* Added three new level to `Trail` for compatibility with syslog which are `Critical`, `Alert` and `Emergency`
* If database doesn't exist for mysql, uAdmin will try to create a new database.

**Deprecated**
^^^^^^^^^^^^^^
* `User.HasPermission` will be private starting `0.6.0`
* `UserGroup.HasPermission` will be private starting `0.6.0`

**Fixed**
^^^^^^^^^
* Fixed image crop modal conflict in list view with delete modal and add it to form and inlines
* Fixed FK in approvals
* Remove required from fields with pending approval
* Support filtering/searching by NULL value for `time.Time` pointer
* Fixed filtering by FK

**Security**
^^^^^^^^^^^^
* Restrict access to inlines based on user model permissions
* Search API escapes HTML results



Version 0.3.0 (2019-07-02)
--------------------------
`Back to Top`_

**Added**
^^^^^^^^^
* Approval System: By adding `approval` tag to any field, the field will required a special permission to edit. If the user does not have this permission, the edit will generate an Approval record that can be approved by a user with access to the new Approval model.
* Upload files using drag and drop into the field.
* Image and File fields now allow a new tag called `webcam` which adds web can access directly from the field.
* Model specific themes: You can use `ModelSchema.FormTheme` and `ModelSchema.ListTheme` to choose a theme for a model.
* Settings automatically adds all uAdmin global variables as settings.
* Added .gitignore which should have been there all along!!

**Changed**
^^^^^^^^^^^
* cropImageHandler now returns JSON with a status letting the for success and failure.
* Improved performance for saving settings page.

**Fixed**
^^^^^^^^^
* Fixed filters for list view. Now you can apply multiple filters. Also now you can filter by foreign key.
* Removed all `fmt.Println` and `log.Println` for printing errors and replaced it with `Trail`.
* Export to excel date time type now takes time zone into account.
* Fixed bug with method fields names rendering with missing letters.
* Settings URL takes into account `RootURL`
* Edit logs will only be saved if the form didn't have errors.
* Fixed bug with `required` file and image fields where you had to choose a file everytime you save even if there was a files saved in the record.
* Fixed bug with `required` foreign key and list type where it was not working before.

**Security**
^^^^^^^^^^^^
* revertLogHandler required the requires to be authenticated and the user to have access to edit the model and have read access to logs.



Version 0.2.2 (2019-06-24)
--------------------------
`Back to Top`_

**Added**
^^^^^^^^^
* GetValue() and ApplyValue() functions in uadmin.Setting

**Changed**
^^^^^^^^^^^
* Enhanced System Settings that contains uAdmin functions

**Fixed**
^^^^^^^^^
* Git submodule update --init --recursive



Version 0.2.0 (2019-06-20)
--------------------------
`Back to Top`_

**Added**
^^^^^^^^^
* Global Variable:
    * Settings model which can be used system wide to keep application settings.
    * GetSetting function to read settings from the settings model.
    * User.GetAccess function can evaluate the user's permission to a model using user level and group level and return a UserPermission instance.
    * ErrorHandleFunc: A function that can be passed for handling errors system wide that gets called when Trail is called.
    * AllowedIPs: is a list of allowed IPs to access uAdmin interface
    * BlockedIPs: is a list of blocked IPs from accessing uAdmin interface
    * RestrictSessionIP: restricts access to the system if the IP of the user changes after login
    * RetainMediaVersions: is to allow the system to keep files uploaded even after they are 

**Changed**
^^^^^^^^^^^
* Forgot password now sends HTML emails.
* GetString function can return the name of a static list item.

**Fixed**
^^^^^^^^^
* Bug fix for read only fields for new and edit.
    * Model names in dashboard are generated using proper plural function instead of just adding "s"



Version 0.1.1 (2019-05-11)
--------------------------
`Back to Top`_

**Added**
^^^^^^^^^
* Global Variable:
    * DefaultMediaPermission [0644]: The default permission for files uploaded through uAdmin



Version 0.1.0 (2018-12-23)
--------------------------
`Back to Top`_

.. _Back to Top: https://uadmin-docs.readthedocs.io/en/latest/roadmap.html#roadmap

- FK no tag
- DB error handling
- Self relation fk/many2many
- Auto/manual preload
- Global schema
- Preloaded templates/static
- DB Inspect
- Auto create dashboard menu on register
- Linear Gradient for progress_bar
- PageLength per model
- Cascade Delete/Update
- Translate help text for field
- Add salt
- Adjust active languages on change
- Customizable admin path
- Investigate Custom Widgets
- Move Static inside a folder named "admin"
- Check code for safety
- Global Email Config
- Process Display Name in Schema to be used in:
    * Form field label
    * Table columns names
- Rate Limit
    * Request / IP / Minute
    * Login requests / IP / Minute
- Add automated translation file sync
- BUG FIX: HTML Interface
- Clean up CLCI file
- Add session to 404 parameters
- Make translate function public
- Switch the welcome message
- Make new model public
- BUG FIX: Changing active languages does not reflect system wide1
- Public Access to Media
- Allow tags in HTML
- Remove CLI colors for windows
- File Field Type BUG FIX. closes #4
- Add salt, BUG FIX: M2M for mysql select
- BUG FIX: Pattern message does not clear out
- Add auto-translation for ErrMsg
- BUG FIX: Fix some uadmin tags to the new format
- BUG FIX: TinyMCE Upload API pathwq
- BUG FIX: Read log's table name is not lowered
- Add encrypt tag
- Remove debug lines
- Add Support for FormModifier and ListModifier
- BUG FIX: Fix add new link for fk
- SECURITY: Fix pass and salt order in Login
- Register Handlers
- Unit tests for 0.1.0-RC1