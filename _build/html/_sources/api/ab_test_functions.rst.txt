A/B Test Functions
==================
`Back To uAdmin Functions List`_

.. _Back To uAdmin Functions List: https://uadmin-docs.readthedocs.io/en/latest/api.html#api-reference

In this section, we will cover the following functions in-depth listed below:

* `uadmin.ABTest`_
    * `func (ABTest) Reset`_
    * `func (*ABTest) Save`_
* `uadmin.ABTestClick`_
* `uadmin.ABTestType`_
    * `func (ABTestType) Model`_
    * `func (ABTestType) Static`_
* `uadmin.ABTestValue`_
    * `func (*ABTestValue) ClickThroughRate`_
    * `func (ABTestValue) ClickThroughRate__Form__List`_
    * `func (ABTestValue) HideInDashboard`_
    * `func (ABTestValue) Preview__Form__List`_
    * `func (*ABTestValue) String`_
* `uadmin.FieldList`_
* `uadmin.GetABTest`_
* `uadmin.ModelList`_

uadmin.ABTest
-------------
`Back To Top`_

.. code-block:: go

    type ABTest struct {
        Model
        Name        string     `uadmin:"required"`
        Type        ABTestType `uadmin:"required"`
        StaticPath  string
        ModelName   ModelList
        Field       FieldList
        PrimaryKey  int
        Active      bool
        Group       string
        ResetABTest string `uadmin:"link"`
    }

ABTest is a model that stores an A/B test.

Here are the following fields and their definitions:

* **Name** - The name of the A/B Test
* **Type** - A list of test types from a dropdown menu
* **StaticPath** - The path assigned in the static
* **ModelName** - A list of registered models
* **Field** - A list of fields from schema for a registered model
* **PrimaryKey** - Used to uniquely identify each row in the table
* **Active** - Checks whether the A/B Test is Active
* **Group** - The name of the group
* **ResetABTest** - A button to reset an AB Test

**func (ABTest) Reset**
^^^^^^^^^^^^^^^^^^^^^^^
`Back To Top`_

.. code-block:: go

    func (a ABTest) Reset()

Reset resets the impressions and clicks to 0 based on a specified AB Test ID.

**func (\*ABTest) Save**
^^^^^^^^^^^^^^^^^^^^^^^^
`Back To Top`_

.. code-block:: go

    func (a *ABTest) Save()

Save saves the AB Test model to database.

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

.. code-block:: go

    func ABTestClick(r *http.Request, group string)

ABTestClick is a function to register a click for an ABTest group.

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

.. code-block:: go

    type ABTestType int

ABTestType is the type of the AB testing: model or static.

ABTestType has 2 functions:

**func (ABTestType) Model**
^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back To Top`_

.. code-block:: go

    func (ABTestType) Model() ABTestType

Model is used to do AB testing for model values coming from database.

**func (ABTestType) Static**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back To Top`_

.. code-block:: go

    func (ABTestType) Static() ABTestType

Static is used to do AB testing for static assets (images, js, css, ...).

See `Part 3: A/B Test Function for Static`_ and `Part 2: A/B Test Function for Model`_ for examples.

uadmin.ABTestValue
------------------
`Back To Top`_

.. code-block:: go

    type ABTestValue struct {
        Model
        ABTest      ABTest
        ABTestID    uint
        Value       string `uadmin:"list_exclude"`
        Active      bool
        Impressions int
        Clicks      int
    }

ABTestValue is a model to represent a possible value of an AB test.

Here are the following fields and their definitions:

* **ABTest** - A model that stores an A/B test
* **ABTestID** - An ID of the A/B Test model
* **Value** - The value that you want to assign in A/B Test
* **Active** - Checks whether the A/B Test is Active
* **Impressions** - The number of visits
* **Clicks** - The number of clicks

**func (\*ABTestValue) ClickThroughRate**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    func (a *ABTestValue) ClickThroughRate() float64

ClickThroughRate returns the rate of click through of this value.

**func (ABTestValue) ClickThroughRate__Form__List**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    func (a ABTestValue) ClickThroughRate__Form__List() string

ClickThroughRate__Form__List shows the click through rate in form and list views.

**func (ABTestValue) HideInDashboard**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    func (ABTestValue) HideInDashboard() bool

HideInDashboard to hide it from dashboard.

**func (ABTestValue) Preview__Form__List**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    func (a ABTestValue) Preview__Form__List() string

Preview__Form__List shows a preview of the AB test's value.

**func (\*ABTestValue) String**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back to Top`_

.. code-block:: go

    func (a ABTestValue) Preview__Form__List() string

String returns a value.

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

.. code-block:: go

    type FieldList int

FieldList is a list of fields from schema for a registered model.

See `Part 3: A/B Test Function for Static`_ for the example.

uadmin.GetABTest
----------------
`Back To Top`_

.. code-block:: go

    func GetABTest(r *http.Request, a interface{}, query interface{}, args ...interface{}) (err error)

GetABTest is like Get function but implements AB testing for the results.

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

.. code-block:: go

    type ModelList int

ModelList is a list of registered models.

See `Part 3: A/B Test Function for Static`_ for the example.
