Register Functions
==================
`Back To uAdmin Functions List`_

.. _Back To uAdmin Functions List: https://uadmin-docs.readthedocs.io/en/latest/api.html#api-reference

In this section, we will cover the following functions in-depth listed below:

* `uadmin.Register`_
* `uadmin.RegisterInlines`_

uadmin.Register
---------------
`Back To Top`_

.. code-block:: go

    func Register(m ...interface{})

Register is used to register models to uAdmin.

Parameter:

    **m ...interface{}:** Is the model that you want to add in the dashboard

Used in the tutorial:

* `Document System Tutorial Part 2 - Creating and Registering a Model`_
* `Document System Tutorial Part 3 - Linking Models (Folders)`_
* `Document System Tutorial Part 5 - Linking Models (Documents)`_
* `uAdmin Tutorial Part 2 - Internal vs. External Models`_
* `uAdmin Tutorial Part 4 - Linking Models`_
* `uAdmin Tutorial Part 6 - Applying uAdmin Tags`_

.. _Document System Tutorial Part 2 - Creating and Registering a Model: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/part2.html
.. _Document System Tutorial Part 3 - Linking Models (Folders): https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/part3.html
.. _Document System Tutorial Part 5 - Linking Models (Documents): https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/part5.html
.. _uAdmin Tutorial Part 2 - Internal vs. External Models: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part2.
.. _uAdmin Tutorial Part 4 - Linking Models: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part4.html
.. _uAdmin Tutorial Part 6 - Applying uAdmin Tags: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part6.html

Create an internal Todo model inside the main.go. Afterwards, call the Todo{} inside the uadmin.Register so that the application will identify the Todo model to be added in the dashboard.

.. code-block:: go

    // Todo model ...
    type Todo struct {
        uadmin.Model
        Name        string
        Description string `uadmin:"html"`
        TargetDate  time.Time
        Progress    int `uadmin:"progress_bar"`
    }

    func main() {
        uadmin.Register(Todo{}) // <-- place it here
    }

Output

.. image:: ../tutorial/assets/uadmindashboard.png

If you click the Todos model, it will display this result as shown below.

.. image:: ../assets/todomodel.png

|

Quiz:

* `Register`_

.. _Register: https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/register.html

uadmin.RegisterInlines
----------------------
`Back To Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/api/register_functions.html#register-functions

.. code-block:: go

    func RegisterInlines(model interface{}, fk map[string]string)

RegisterInlines is a function to register a model as an inline for another model.

Parameters:

    **model (struct instance):** Is the model that you want to add inlines to.

    **fk (map[interface{}]string):** This is a map of the inlines to be added to the model. The map's key is the name of the model of the inline and the value of the map is the foreign key field's name.

Used in the tutorial:

* `Document System Tutorial Part 7 - Register Inlines`_
* `uAdmin Tutorial Part 5 - Register Inlines and Drop Down List`_
* `uAdmin Tutorial Part 6 - Applying uAdmin Tags`_

.. _Document System Tutorial Part 7 - Register Inlines: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/part7.html
.. _uAdmin Tutorial Part 5 - Register Inlines and Drop Down List: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part5.html
.. _uAdmin Tutorial Part 6 - Applying uAdmin Tags: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part6.html

Example:

.. code-block:: go

    type Person struct {
        uadmin.Model
        Name string
    }

    type Card struct {
        uadmin.Model
        PersonID uint
        Person   Person
    }

    func main() {
        // ...
        uadmin.RegisterInlines(Person{}, map[string]string{
            "Card": "PersonID",
        })
        // ...
    }

Quiz:

* `Foreign Key and Register Inlines`_

.. _Foreign Key and Register Inlines: https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/foreign-key-and-register-inline.html
