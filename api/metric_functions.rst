<<<<<<< HEAD
Metric Functions
================
`Back To uAdmin Functions List`_

.. _Back To uAdmin Functions List: https://uadmin-docs.readthedocs.io/en/latest/api.html#api-reference

In this section, we will cover the following functions in-depth listed below:

* `uadmin.IncrementMetric`_
* `uadmin.NewMetric`_
* `uadmin.SetMetric`_
* `uadmin.SystemMetrics`_
* `uadmin.TimeMetric`_
* `uadmin.UserMetrics`_

uadmin.IncrementMetric
----------------------
`Back To Top`_

.. code-block:: go

    func IncrementMetric(name string)

IncrementMetric increments the value of a.

uadmin.NewMetric
----------------
`Back To Top`_

.. code-block:: go

    func NewMetric(name string, template string) error

NewMetric creates a new metric.

uadmin.SetMetric
----------------
`Back To Top`_

.. code-block:: go

    func SetMetric(name string, value float64)

SetMetric sets the value of a gauge metric.

uadmin.SystemMetrics
--------------------
`Back To Top`_

.. code-block:: go

    // Type: bool
    var SystemMetrics = false

SystemMetrics enables uAdmin system metrics to be recorded.

uadmin.TimeMetric
-----------------
`Back To Top`_

.. code-block:: go

    func TimeMetric(name string, div float64, f func())

TimeMetric runs a function and times it as a metric.

uadmin.UserMetrics
------------------
`Back To Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/api/metric_functions.html#metric-functions

.. code-block:: go

    // Type: bool
    var UserMetrics = false

UserMetrics enables the user metrics to be recorded.
=======
Metric Functions
================
`Back To uAdmin Functions List`_

.. _Back To uAdmin Functions List: https://uadmin-docs.readthedocs.io/en/latest/api.html#api-reference

In this section, we will cover the following functions in-depth listed below:

* `uadmin.IncrementMetric`_
* `uadmin.NewMetric`_
* `uadmin.SetMetric`_
* `uadmin.SystemMetrics`_
* `uadmin.TimeMetric`_
* `uadmin.UserMetrics`_

uadmin.IncrementMetric
----------------------
`Back To Top`_

.. code-block:: go

    func IncrementMetric(name string)

IncrementMetric increments the value of a.

uadmin.NewMetric
----------------
`Back To Top`_

.. code-block:: go

    func NewMetric(name string, template string) error

NewMetric creates a new metric.

uadmin.SetMetric
----------------
`Back To Top`_

.. code-block:: go

    func SetMetric(name string, value float64)

SetMetric sets the value of a gauge metric.

uadmin.SystemMetrics
--------------------
`Back To Top`_

.. code-block:: go

    // Type: bool
    var SystemMetrics = false

SystemMetrics enables uAdmin system metrics to be recorded.

uadmin.TimeMetric
-----------------
`Back To Top`_

.. code-block:: go

    func TimeMetric(name string, div float64, f func())

TimeMetric runs a function and times it as a metric.

uadmin.UserMetrics
------------------
`Back To Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/api/metric_functions.html#metric-functions

.. code-block:: go

    // Type: bool
    var UserMetrics = false

UserMetrics enables the user metrics to be recorded.
>>>>>>> de25cdd8a29ca2bb2c2df08be00b703b967aaed5
