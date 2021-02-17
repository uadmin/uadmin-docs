A/B Test Functions
==================
`Back To uAdmin Functions List`_

.. _Back To uAdmin Functions List: https://uadmin-docs.readthedocs.io/en/latest/api.html#api-reference

In this section, we will cover the following functions in-depth listed below:

* `uadmin.ABTest`_
* `uadmin.ABTestClick`_
* `uadmin.ABTestType`_
* `uadmin.ABTestValue`_
* `uadmin.FieldList`_
* `uadmin.GetABTest`_
* `uadmin.ModelList`_

**uadmin.ABTest**
^^^^^^^^^^^^^^^^^
`Back To Top`_

ABTest is a model that stores an A/B test.

Structure:

.. code-block:: go

    type ABTest struct {
        Model
        Name       string   `uadmin:"required"`
        Type       ABTestType `uadmin:"required"`
        StaticPath string
        ModelName  ModelList
        Field      FieldList
        PrimaryKey int
        Active     bool
        Group      string
    }

Here are the following fields and their definitions:

* **Name** - The name of the A/B Test
* **Type** - A list of test types from a dropdown menu
* **StaticPath** - The path assigned in the static
* **ModelName** - A list of registered models
* **Field** - A list of fields from schema for a registered model
* **PrimaryKey** - Used to uniquely identify each row in the table
* **Active** - Checks whether the A/B Test is Active
* **Group** - The name of the group

There are 2 ways you can do for initialization process using this function: one-by-one and by group.

One-by-one initialization:

.. code-block:: go

    func main(){
        // Some codes
        abtest := uadmin.ABTest{}
        abtest.Name = "Name"
        abtest.Type = uadmin.ABTestType(0).Static()
        abtest.StaticPath = "Static Path"
    }

By group initialization:

.. code-block:: go

    func main(){
        // Some codes
        abtest := uadmin.ABTest{
            Name:       "Name",
            Type:       uadmin.ABTestType(0).Static(),
            StaticPath: "Static Path",
        }
    }

In the following examples, we will use "by group” initialization process.

* `Example #1: Static`_
    * `Part 1: HTML Template`_
    * `Part 2: Static Handler`_
    * `Part 3: A/B Test Function for Static`_
    * `Part 4: A/B Test Value Function for Static`_
    * `Part 5: Image Testing`_

* `Example #2: Models`_
    * `Part 1: Campaign Info Model`_
    * `Part 2: A/B Test Function for Model`_
    * `Part 3: A/B Test Value Function for Model`_
    * `Part 4: API Click Handler`_
    * `Part 5: Get A/B Test`_
    * `Part 6: Button Testing`_

.. _Example #1\: Static: https://uadmin-docs.readthedocs.io/en/latest/api/ab-test-functions/abtest.html#example-1-static
.. _Part 1\: HTML Template: https://uadmin-docs.readthedocs.io/en/latest/api/ab-test-functions/abtest.html#part-1-html-template
.. _Part 2\: Static Handler: https://uadmin-docs.readthedocs.io/en/latest/api/ab-test-functions/abtest.html#part-2-static-handler
.. _Part 3\: A/B Test Function for Static: https://uadmin-docs.readthedocs.io/en/latest/api/ab-test-functions/abtest.html#part-3-a-b-test-function-for-static
.. _Part 4\: A/B Test Value Function for Static: https://uadmin-docs.readthedocs.io/en/latest/api/ab-test-functions/abtest.html#part-4-a-b-test-value-function-for-static
.. _Part 5\: Image Testing: https://uadmin-docs.readthedocs.io/en/latest/api/ab-test-functions/abtest.html#part-5-image-testing

.. _Example #2\: Models: https://uadmin-docs.readthedocs.io/en/latest/api/ab-test-functions/abtest.html#example-2-models
.. _Part 1\: Campaign Info Model: https://uadmin-docs.readthedocs.io/en/latest/api/ab-test-functions/abtest.html#part-1-campaign-info-model
.. _Part 2\: A/B Test Function for Model: https://uadmin-docs.readthedocs.io/en/latest/api/ab-test-functions/abtest.html#part-2-a-b-test-function-for-model
.. _Part 3\: A/B Test Value Function for Model: https://uadmin-docs.readthedocs.io/en/latest/api/ab-test-functions/abtest.html#part-3-a-b-test-value-function-for-model
.. _Part 4\: API Click Handler: https://uadmin-docs.readthedocs.io/en/latest/api/ab-test-functions/abtest.html#part-4-api-click-handler
.. _Part 5\: Get A/B Test: https://uadmin-docs.readthedocs.io/en/latest/api/ab-test-functions/abtest.html#part-5-get-a-b-test
.. _Part 6\: Button Testing: https://uadmin-docs.readthedocs.io/en/latest/api/ab-test-functions/abtest.html#part-6-button-testing

Page:

.. toctree::
   :maxdepth: 1

   ab-test-functions/abtest

Quiz:

* `A/B Test Functions`_

uadmin.ABTestClick
------------------
`Back To Top`_

ABTestClick is a function to register a click for an ABTest group.

Function:

.. code-block:: go

    func(r *http.Request, group string)

Parameters:

    **r \*http.Request**: Is a data structure that represents the client HTTP request

    **group string** : Is the name of the group

See `Part 4: API Click Handler`_ for the example.

.. _Part 4\: API Click Handler: https://uadmin-docs.readthedocs.io/en/latest/api/ab-test-functions/abtest.html#part-4-api-click-handler

Quiz:

* `A/B Test Functions`_

uadmin.ABTestType
-----------------
`Back To Top`_

ABTestType is a list of test types from a dropdown menu.

Type:

.. code-block:: go

    int

ABTestType has 2 functions:

* **Static()** - Test static files
* **Model()** - Test registered models

See `Part 3: A/B Test Function for Static`_ and `Part 2: A/B Test Function for Model`_ for examples.

uadmin.ABTestValue
------------------
`Back To Top`_

ABTest is a model that stores a value for an A/B test.

Structure:

.. code-block:: go

    type ABTestValue struct {
        Model
        ABTest      ABTest
        ABTestID    uint
        Value       string
        Active      bool
        Impressions int
        Clicks      int
    }

Here are the following fields and their definitions:

* **ABTest** - A model that stores an A/B test
* **ABTestID** - An ID of the A/B Test model
* **Value** - The value that you want to assign in A/B Test
* **Active** - Checks whether the A/B Test is Active
* **Impressions** - The number of visits
* **Clicks** - The number of clicks

There are 4 functions that you can use in ABTestValue:

* **String()** - Returns the value
* **ClickThroughRate()** - Returns the rate/percentage of the user clicks
* **ClickThroughRate__Form__List()** - Displays the rate/percentage of the user clicks in the form and the list
* **HideInDashboard()** - Return true and auto hide this from dashboard

There are 2 ways you can do for initialization process using this function: one-by-one and by group.

One-by-one initialization:

.. code-block:: go

    func main(){
        // Some codes
        abtestvalue := uadmin.ABTestValue{}
        abtestvalue.ABTest = ABTest
        abtestvalue.ABTestID = 1
        abtestvalue.Value = "Value"
    }

By group initialization:

.. code-block:: go

    func main(){
        // Some codes
        abtestvalue := uadmin.ABTestValue{
            ABTest:   ABTest,
            ABTestID: 1,
            Value:    "Value",
        }
    }

In the following examples, we will use "by group” initialization process.

See `Part 4: A/B Test Value Function for Static`_ and `Part 3: A/B Test Value Function for Model`_ for examples.

.. _Part 4\: A/B Test Value Function for Static: https://uadmin-docs.readthedocs.io/en/latest/api/ab-test-functions/abtest.html#part-4-a-b-test-value-function-for-static
.. _Part 3\: A/B Test Value Function for Model: https://uadmin-docs.readthedocs.io/en/latest/api/ab-test-functions/abtest.html#part-3-a-b-test-value-function-for-model

Quiz:

* `A/B Test Functions`_

.. _A/B Test Functions: https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/a-b-test-functions.html

uadmin.FieldList
----------------
`Back To Top`_

FieldList is a list of fields from schema for a registered model.

Type:

.. code-block:: go

    int

See `Part 3: A/B Test Function for Static`_ for the example.

uadmin.GetABTest
----------------
`Back To Top`_

GetABTest is a function that gets an active A/B tests for any field in AB Test model.

Function:

.. code-block:: go

    func(r *http.Request, a interface{}, query interface{}, args ...interface{}) (err error)

Parameters:

    **r \*http.Request:** Is a data structure that represents the client HTTP request

    **a interface{}:** Is the variable where the model was initialized

    **query interface{}:** Is an action that you want to perform in your database

    **args ...interface{}:** Is the series of arguments for query input

See `Part 5: Get A/B Test`_ for the example.

.. _Part 5\: Get A/B Test: https://uadmin-docs.readthedocs.io/en/latest/api/ab-test-functions/abtest.html#part-5-get-a-b-test

uadmin.ModelList
----------------
`Back To Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/api/ab_test_functions.html#a-b-test-functions

ModelList a list of registered models.

Type:

.. code-block:: go

    int

See `Part 3: A/B Test Function for Static`_ for the example.
