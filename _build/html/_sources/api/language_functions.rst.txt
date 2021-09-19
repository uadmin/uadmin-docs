Language Functions
==================
`Back To uAdmin Functions List`_

.. _Back To uAdmin Functions List: https://uadmin-docs.readthedocs.io/en/latest/api.html#api-reference

In this section, we will cover the following functions in-depth listed below:

* `uadmin.CacheTranslation`_
* `uadmin.CustomTranslation`_
* `uadmin.GetActiveLanguages`_
* `uadmin.GetDefaultLanguage`_
* `uadmin.Language`_
    * `func (*Language) Save`_
    * `func (Language) String`_
* `uadmin.Tf`_
* `uadmin.Translate`_
* `uadmin.TranslateSchema`_

uadmin.CacheTranslation
-----------------------
`Back To Top`_

.. code-block:: go

    var CacheTranslation = false

CacheTranslation allows a translation to store data in a cache memory.

To assign a value within an application, visit `Cache Translation`_ page for an example.

.. _Cache Translation: https://uadmin-docs.readthedocs.io/en/latest/system-reference/setting.html#cache-translation

To assign a value in the code, follow this approach:

Example:

.. code-block:: go

    package main

    import (
        "github.com/uadmin/uadmin"
    )

    func main() {
        // NOTE: This code works only on first build.
        // Allows a translation to store data in a cache memory
        uadmin.CacheTranslation = true

        // Prohibits a translation to store data in a cache memory
        uadmin.CacheTranslation = false

        // ----- IF YOU RUN YOUR APPLICATION AGAIN, DO THIS BELOW -----

        // Allows a translation to store data in a cache memory
        setting := uadmin.Setting{}
        uadmin.Get(&setting, "code = ?", "uAdmin.CacheTranslation")
        setting.ParseFormValue([]string{"true"})
        setting.Save()

        // Prohibits a translation to store data in a cache memory
        setting := uadmin.Setting{}
        uadmin.Get(&setting, "code = ?", "uAdmin.CacheTranslation")
        setting.ParseFormValue([]string{"false"})
        setting.Save()
    }

Quiz:

* `Miscellaneous Functions (2)`_

uadmin.CustomTranslation
------------------------
`Back To Top`_

.. code-block:: go

    var CustomTranslation = []string{
        "uadmin/system",
    }

CustomTranslation is where you can register custom translation files. To register a custom translation file, always assign it with its key in the this format "{CATEGORY_NAME}/{FILENAME}". For example:

.. code-block:: go

    uadmin.CustomTranslation = append(uadmin.CustomTranslation, "ui/billing")

This will register the file and you will be able to use it if `uadmin.Tf`. By default there is only one registed custom translation wich is "uadmin/system".

Let's create an HTML file named **index.html** inside the templates folder. It constructs two link texts for the languages which are English and Chinese, and a sample sentence in big heading text.

.. code-block:: html

    <!DOCTYPE html>
    <html lang="en">
      <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Document</title>
      </head>
      <body>
        <a href="?language=en">English</a>
        <a href="?language=zh">Chinese</a>

        <!-- Tf is function for translating strings into any given language. -->
        <!-- page/index is where to get the translation from. -->
        <!-- .Language is the language code coming from the PageContext struct in views/page.go. -->
        <!-- "This is a sentence." is the value that we are going to translate based on the given language code. -->
        <h1>{{Tf "page/index" .Language "This is a sentence."}}</h1>
      </body>
    </html>

Once you are done with this step, create a file called **page.go** to implement the PageHandler in the views folder.

.. code-block:: go

    package views

    import (
        "net/http"

        "github.com/uadmin/uadmin"
    )

    // Build a custom struct to store the language code here so that it can communicate to the Tf that we assigned in index.html.
    type PageContext struct {
        Language string
    }

    // PageHandler allows the system to create a connection to the HTML pages.
    func PageHandler(w http.ResponseWriter, r *http.Request) {
        // Get the value of the language key in the URL requested by client.
        // e.g. http://api.example.com/?language=en
        lang := r.URL.Query().Get("language")

        // Check whether the client passes the URL query where the given language key contains a value.
        if lang != "" {
            // Assign the name, value, and path to build the language cookie.
            langC := http.Cookie{
                Name:  "language",
                Value: lang,
                Path:  "/",
            }
            // Set the language cookie to the browser.
            http.SetCookie(w, &langC)

        } else {
            // If no language cookie has been found
            if langC, err := r.Cookie("language"); err != nil || (langC != nil && langC.Value == "") {
                // Get the default language code from the database instead and assign to this variable.
                lang = uadmin.GetDefaultLanguage().Code
            } else {
                // Assign the value from the language cookie in this variable.
                lang = langC.Value
            }
        }

        // Assign the language code in the custom struct to be passed in the front-end part
        context := PageContext{
            Language: lang,
        }

        // Pass the context to the specified HTML path.
        uadmin.RenderHTML(w, r, "templates/index.html", context)
    }

Establish a connection in the **main.go** to the views by using http.HandleFunc. It should be placed after the uadmin.Register and before the StartServer.

.. code-block:: go

    package main

    import (
        "net/http"

        // Specify the username that you used inside github.com folder
        "github.com/username/project_name/views"

        "github.com/uadmin/uadmin"
    )

    func main() {
        // To avoid having issues when running an application, register function should be called.
        uadmin.Register()

        // Assign the Root URL value to /admin/. This code allows to resolve multiple registration issue when running an application.
        uadmin.RootURL = "/admin/" // Applies on first build only
        setting := uadmin.Setting{}
        uadmin.Get(&setting, "code = ?", "uAdmin.RootURL")
        setting.ParseFormValue([]string{"/admin/"})
        setting.Save()

        // Assign the path where to register the custom translation files in static/i18n.
        // Format: "{CATEGORY_NAME}/{FILENAME}"
        uadmin.CustomTranslation = append(uadmin.CustomTranslation, "page/index")

        // Register the Page Handler
        http.HandleFunc("/", uadmin.Handler(views.PageHandler))

        uadmin.StartServer()
    }

Now run your application, go to http://localhost:8080/ and see what happens.

.. image:: assets/language_page_ui.png

|

At the moment, if you click the Chinese link text, you may notice that it doesn't translate automatically in the UI. It's because we need to activate the Chinese language in the Language model on uAdmin dashboard. In order to do that, go to http://localhost:8080/admin and open **LANGUAGES** model.

.. image:: ../assets/languageshighlighted.png

|

Activate the Chinese language in the form and save it. This should mark as checked in the table.

.. image:: assets/chineseactive.png

|

Go to **/static/i18n/page** path. This is where the translation files are registered and that's the usage of **uadmin.CustomTranslation**. Page is the category name and index.(code).json is the filename.

.. image:: assets/i18nindex.png

|

You may notice that even if the Chinese language is activated in the Language model, it still does not generate the translation file for Chinese in **/static/i18n/page** path. In order to take effect, you need to rebuild an application.

After you rebuild an application, revisit the **/static/i18n/page** path. You must be able to see the **index.zh.json** now. This is the translation file for Chinese. At the moment, both of them are empty object.

.. image:: assets/i18nindexzh.png

|

Go to http://localhost:8080/ and let's select the English and Chinese link texts back and forth.

.. image:: assets/englishchineselinktexthighlighted.png

|

Go back to **/static/i18n/page** path and you may notice that both English and Chinese translation files generated the key "This is a sentence." inside an object.

**index.en.json**

.. code-block:: JSON

    {
      "This is a sentence.": "This is a sentence."
    }

**index.zh.json**

.. code-block:: JSON

    {
      "This is a sentence.": "Translate me ---> This is a sentence."
    }

The keyword **Translate me** is something that you have to assign the translation manually. In this example, let's mock up the Chinese translated text for "This is a sentence." and apply it in the **index.zh.json** translation file.

.. code-block:: JSON

    {
      "This is a sentence.": "这是一个句子。"
    }

Now go back to http://localhost:8080/, select the Chinese link text and see what happens.

.. image:: assets/chineselinktexthighlighted.png

|

Result:

.. image:: assets/chineselanguageapplied.png

|

Congratulations! You have created a multilingual application covering the English and Chinese languages.

Quiz:

* `Miscellaneous Functions (2)`_

.. _Miscellaneous Functions (2): https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/miscellaneous-functions-2.html

uadmin.GetActiveLanguages
-------------------------
`Back To Top`_

.. code-block:: go

    func GetActiveLanguages() []Language

GetActiveLanguages returns a list of active languages from the Language model.

Suppose Chinese with the code "zh" and English with the code "en" are active languages in the Language model.

.. image:: assets/chineseenglishactive.png
   :align: center

|

Go to the main.go and apply the following codes below:

.. code-block:: go

    func main() {
        // Returns the list of active languages from the Language model
        activeLanguages := uadmin.GetActiveLanguages()

        // Print the result
        fmt.Println(activeLanguages)

        uadmin.StartServer()
    }

Run your application. Check your console for the result.

.. code-block:: bash

    [zh en]

uadmin.GetDefaultLanguage
-------------------------
`Back To Top`_

.. code-block:: go

    func GetDefaultLanguage() Language

GetDefaultLanguage returns the default language by code from the Language model.

Suppose Vietnamese is the default language with the code "vi" in the Language model.

.. image:: assets/vietnamesedefault.png
   :align: center

|

Go to the main.go and apply the following codes below:

.. code-block:: go

    func main() {
        // Returns the default language from the Language model
        defaultLanguage := uadmin.GetDefaultLanguage()

        // Print the result
        fmt.Println(defaultLanguage)

        uadmin.StartServer()
    }

Run your application. Check your console for the result.

.. code-block:: bash

    vi

uadmin.Language
---------------
`Back To Top`_

.. code-block:: go

    type Language struct {
        Model
        EnglishName    string `uadmin:"required;read_only;filter;search"`
        Name           string `uadmin:"required;read_only;filter;search"`
        Flag           string `uadmin:"image;list_exclude"`
        Code           string `uadmin:"filter;read_only;list_exclude"`
        RTL            bool   `uadmin:"list_exclude"`
        Default        bool   `uadmin:"help:Set as the default language;list_exclude"`
        Active         bool   `uadmin:"help:To show this in available languages;filter"`
        AvailableInGui bool   `uadmin:"help:The App is available in this language;read_only"`
    }

Language is a system in uAdmin that is used to add, modify and delete the elements of a language.

**func (\*Language) Save**
^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    func (l *Language) Save()

Save saves the object in the database.

**func (Language) String**
^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    func (l Language) String() string

String returns the code of the language.

There are 2 ways you can do for initialization process using this function: one-by-one and by group.

One-by-one initialization:

.. code-block:: go

    func main(){
        // Some codes
        language := uadmin.Language{}
        language.EnglishName = "English Name"
        language.Name = "Name"
    }

By group initialization:

.. code-block:: go

    func main(){
        // Some codes
        language := uadmin.Language{
            EnglishName: "English Name",
            Name: "Name",
        }
    }

In the following examples, we will use "by group” initialization process.

* `Example #1: Active`_
* `Example #2: Default`_
* `Example #3: RTL`_

.. _Example #1\: Active: https://uadmin-docs.readthedocs.io/en/latest/api/language-functions/language.html#example-1-active
.. _Example #2\: Default: https://uadmin-docs.readthedocs.io/en/latest/api/language-functions/language.html#example-2-default
.. _Example #3\: RTL: https://uadmin-docs.readthedocs.io/en/latest/api/language-functions/language.html#example-3-rtl

Page:

.. toctree::
   :maxdepth: 1

   language-functions/language

Quiz:

* `Language`_

.. _Language: https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/language.html

uadmin.Tf
---------
`Back To Top`_

.. code-block:: go

    func Tf(path string, lang string, term string, args ...interface{}) string

Tf is a function for translating strings into any given language.

Parameters:

    **path (string):** This is where to get the translation from. It is in the
    format of "GROUPNAME/FILENAME" for example: "uadmin/system"

    **lang (string):** Is the language code. If empty string is passed we will use
    the default language.

    **term (string):** The term to translate

    **args (...interface{}):** Is a list of arguments to fill the term with place holders

First of all, create a back-end validation function inside the todo.go.

.. code-block:: go

    // Validate !
    func (t Todo) Validate() (errMsg map[string]string) {
        // Initialize the error messages
        errMsg = map[string]string{}

        // Get any records from the database that matches the name of
        // this record and make sure the record is not the record we are
        // editing right now
        todo := Todo{}
        system := "system"
        if uadmin.Count(&todo, "name = ? AND id <> ?", t.Name, t.ID) != 0 {
            errMsg["Name"] = uadmin.Tf("models/Todo/Name/errMsg", "", fmt.Sprintf("This todo name is already in the %s", system))
        }
        return
    }

Run your application and login using "admin” as username and password.

.. image:: assets/loginformadmin.png
   :align: center

|

Open "LANGUAGES" model.

.. image:: /assets/languageshighlighted.png

|

Search whatever languages you want to be available in your application. For this example, let's choose Tagalog and set it to Active.

.. image:: assets/tagalogactive.png
   :align: center

|

Open "TODOS" model and create at least one record inside it.

.. image:: ../tutorial/assets/todomodeloutput.png

|

Logout your account and login again. Set your language to **Wikang Tagalog (Tagalog)**.

.. image:: /assets/loginformtagalog.png

|

Open "TODOS" model, create a duplicate record, save it and let's see what happens.

.. image:: /assets/duplicaterecord.png
   :align: center

|

The error message appears. Now rebuild your application and see what happens.

.. code-block:: go

    [   OK   ]   Initializing DB: [13/13]
    [   OK   ]   Synching System Settings: [51/51]
    [ WARNING]   Translation of tl at 0% [0/134]

It says tl is 0% which means we have not translated yet. 

From your project folder, go to static/i18n/models/todo.tl.json. Inside it, you will see a bunch of data in JSON format that says Translate Me. This is where you put your translated text. For this example, let's translate the err_msg value in Tagalog language then save it.

.. image:: /assets/errmsgtagalog.png

|

Once you are done, go back to your application, refresh your browser and see what happens.

.. image:: /assets/todotagalogtranslatedtf.png
   :align: center

|

And if you rebuild your application, you will notice that uAdmin has found 1 word we have translated and is telling us we are at 1% translation for the Tagalog language.

.. code-block:: bash

    [   OK   ]   Initializing DB: [13/13]
    [   OK   ]   Synching System Settings: [51/51]
    [ WARNING]   Translation of tl at 1% [1/134]

Congrats, now you know how to translate your sentence using uadmin.Tf.

Quiz:

* `Tf and Translate`_

uadmin.Translate
----------------
`Back To Top`_

.. code-block:: go

    func Translate(raw string, lang string, args ...bool) string

Translate is used to get a translation from a multilingual fields.

Parameters:

    **raw string:** Is the field of the model that you want to access to

    **lang string:** Is the code of the language

    **args ...bool:** Series of arguments that returns a boolean value

Before we proceed to the example, read `Tutorial Part 9 - Introduction to API`_ to familiarize how API works in uAdmin.

.. _Tutorial Part 9 - Introduction to API: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part9.html

Suppose I have two multilingual fields in my Item record.

.. image:: /assets/itementl.png

Create a file named custom_todo.go inside the api folder with the following codes below:

.. code-block:: go

    // CustomTodoHandler !
    func CustomTodoHandler(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimPrefix(r.URL.Path, "/custom_todo")
        r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

        res := map[string]interface{}{}

        item := models.Item{}

        results := []map[string]interface{}{}

        uadmin.Get(&item, "id = 1")

        results = append(results, map[string]interface{}{
            "Description (en)": uadmin.Translate(item.Description, "en"),
            "Description (tl)": uadmin.Translate(item.Description, "tl"),
        })

        res["status"] = "ok"
        res["item"] = results
        uadmin.ReturnJSON(w, r, res)
    }

Establish a connection in the main.go to the API by using http.HandleFunc. It should be placed after the uadmin.Register and before the StartServer.

.. code-block:: go

    func main() {
        // Some codes

        // CustomTodoHandler
        http.HandleFunc("/custom_todo/", uadmin.Handler(api.CustomTodoHandler)) // <-- place it here
    }

api is the folder name while CustomTodoHandler is the name of the function inside custom_todo.go.

Run your application and see what happens.

.. image:: /assets/translatejson.png

|

Quiz:

* `Tf and Translate`_

.. _Tf and Translate: https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/tf-and-translate.html

uadmin.TranslateSchema
----------------------
`Back To Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/api/language_functions.html#language-functions

.. code-block:: go

    func Translate(raw string, lang string, args ...bool) string

TranslateSchema translate a model schema.
