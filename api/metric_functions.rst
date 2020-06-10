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

Function:

.. code-block:: go

    func(name string)

uadmin.NewMetric
----------------
`Back To Top`_

Function:

.. code-block:: go

    func(name string, template string) error

uadmin.SetMetric
----------------
`Back To Top`_

Function:

.. code-block:: go

    func(name string, value float64)

uadmin.SystemMetrics
--------------------
`Back To Top`_

SystemMetrics enables uAdmin system metrics to be recorded.

Type:

.. code-block:: go

    bool

uadmin.TimeMetric
-----------------
`Back To Top`_

Function:

.. code-block:: go

    func(name string, div float64, f func())

uadmin.UserMetrics
------------------
`Back To Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/api/metric_functions.html#metric-functions

UserMetrics enables the user metrics to be recorded.

Type:

.. code-block:: go

    bool

