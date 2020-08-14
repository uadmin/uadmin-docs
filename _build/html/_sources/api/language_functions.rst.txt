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

CustomTranslation is where you can register custom translation files. To register a custom translation file, always assign it with it's key in the this format "category/name". For example:

.. code-block:: go

    uadmin.CustomTranslation = append(uadmin.CustomTranslation, "ui/billing")

This will register the file and you will be able to use it if `uadmin.Tf`. By default there is only one registed custom translation wich is "uadmin/system".

Another example: Suppose that English is the only active language in your application. Go to the main.go and apply the following codes below. It should be placed before uadmin.Register.

.. code-block:: go

    func main(){
        // Place it here
        uadmin.CustomTranslation = []string{"models/custom", "models/todo_custom"}

        uadmin.Register(
            // Some codes
        )
    }

From your project folder, go to static/i18n/models. You will notice that two JSON files are created in the models folder.

.. image:: ../assets/customtranslationcreate.png

Every JSON file is per language. In other words, if you have 2 languages available in your application, there will be a total of 4 created JSON files.

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

.. image:: ../assets/todomodeloutput.png

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
    [   OK   ]   Synching System Settings: [46/46]
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
    [   OK   ]   Synching System Settings: [46/46]
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
