Pre Query Functions
===================
`Back To uAdmin Functions List`_

.. _Back To uAdmin Functions List: https://uadmin-docs.readthedocs.io/en/latest/dapi.html#uadmin-functions

In this section, we will cover the following functions in-depth listed below:

* `uadmin.APIPreQueryAdd`_
* `uadmin.APIPreQueryAdder`_
* `uadmin.APIPreQueryDelete`_
* `uadmin.APIPreQueryDeleter`_
* `uadmin.APIPreQueryEdit`_
* `uadmin.APIPreQueryEditor`_
* `uadmin.APIPreQueryRead`_
* `uadmin.APIPreQueryReader`_
* `uadmin.APIPreQuerySchema`_
* `uadmin.APIPreQuerySchemer`_

uadmin.APIPreQueryAdd
---------------------
`Back To Top`_

APIPreQueryAdd controls the data API's pre query for add commands.

Type:

.. code-block:: go

    bool

uadmin.APIPreQueryAdder
-----------------------
`Back To Top`_

APIPreQueryAdder is an interface for models to run before processing add function in dAPI. Returning false stops the rest of the process from happening.

Type:

.. code-block:: go

    interface

uadmin.APIPreQueryDelete
------------------------
`Back To Top`_

APIPreQueryDelete controls the data API's pre query for delete commands.

Type:

.. code-block:: go

    bool

uadmin.APIPreQueryDeleter
-------------------------
`Back To Top`_

APIPreQueryDeleter is an interface for models to run before processing delete function in dAPI. Returning false stops the rest of the process from happening.

Type:

.. code-block:: go

    interface

uadmin.APIPreQueryEdit
----------------------
`Back To Top`_

APIPreQueryEdit controls the data API's pre query for edit commands.

Type:

.. code-block:: go

    bool

uadmin.APIPreQueryEditor
------------------------
`Back To Top`_

APIPreQueryEditor is an interface for models to run before processing edit function in dAPI. Returning false stops the rest of the process from happening.

Type:

.. code-block:: go

    interface

uadmin.APIPreQueryRead
----------------------
`Back To Top`_

APIPreQueryRead controls the data API's pre query for read commands.

Type:

.. code-block:: go

    bool

uadmin.APIPreQueryReader
------------------------
`Back To Top`_

APIPreQueryReader is an interface for models to run before processing read function in dAPI. Returning false stops the rest of the process from happening.

Type:

.. code-block:: go

    interface

uadmin.APIPreQuerySchema
------------------------
`Back To Top`_

APIPreQuerySchema controls the data API's pre query for schema commands.

Type:

.. code-block:: go

    bool

uadmin.APIPreQuerySchemer
-------------------------
`Back To Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/dapi/pre_query_functions.html#pre-query-functions

APIPreQuerySchemer is an interface for models to run before processing schema function in dAPI. Returning false stops the rest of the process from happening.

Type:

.. code-block:: go

    interface
