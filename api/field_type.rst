uadmin.FieldType
================
`Back To uAdmin Functions List`_

.. _Back To uAdmin Functions List: https://uadmin-docs.readthedocs.io/en/latest/api.html#api-reference

FieldType is a list of field types from dropdown menu.

Type:

.. code-block:: go

    int

FieldType has 19 functions:

* **Boolean()** - A field type that has one of two possible values (usually denoted true and false), intended to represent the two truth values of logic and Boolean algebra
* **Code()** - A set of instructions that will be executed by a computer
* **DateTime()** - Provides functionality for measuring and displaying time. It is equivalent to time.Time.
* **DateTimePtr()** - Provides functionality for measuring and displaying time. It is equivalent to \*time.Time.
* **Email()** - Identifies an email box to which email messages are delivered
* **File()** - Enables the user to upload files/attachments in the model
* **Float()** - Used in various programming languages to define a variable with a fractional value
* **ForeignKey()** - The key used to link two models together
* **HTML()** - Allows the user to modify text in HTML format
* **Image()** - Mark a field as an image
* **Int()** - Used to represent a whole number that ranges from -2147483647 to 2147483647 for 9 or 10 digits of precision
* **Link()** - Displays a button in the model containing the assigned hyperlink
* **M2M()** - Allows you to select more than one element inside an input box field
* **Money()** - Displays a numeric value with two decimal places in the list
* **Multilingual()** - Allows the user to use more than two languages for input
* **Password()** - A string of characters that hides the input data for security
* **ProgressBar()** - A feature used to measure the progress of the activity
* **StaticList()** - Retains the exact record IDs that matched your criteria at the time the list was saved [#f1]_
* **String()** - Used to represent text rather than numbers

Reference
---------
.. [#f1] eTrigue Support. (n.d.). What is the difference between a dynamic and static list? Retrieved from https://support.etrigue.com/hc/en-us/articles/231552027-What-is-the-difference-between-a-dynamic-and-static-list-
