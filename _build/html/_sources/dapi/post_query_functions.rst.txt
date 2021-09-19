Post Query Functions
====================
`Back To uAdmin Functions List`_

.. _Back To uAdmin Functions List: https://uadmin-docs.readthedocs.io/en/latest/dapi.html#uadmin-functions

In this section, we will cover the following functions in-depth listed below:

* `uadmin.APIPostQueryAdd`_
* `uadmin.APIPostQueryAdder`_
* `uadmin.APIPostQueryDelete`_
* `uadmin.APIPostQueryDeleter`_
* `uadmin.APIPostQueryEdit`_
* `uadmin.APIPostQueryEditor`_
* `uadmin.APIPostQueryRead`_
* `uadmin.APIPostQueryReader`_
* `uadmin.APIPostQuerySchema`_
* `uadmin.APIPostQuerySchemer`_

uadmin.APIPostQueryAdd
----------------------
`Back To Top`_

APIPostQueryAdd controls the data API's post query for add commands.

Type:

.. code-block:: go

    bool

uadmin.APIPostQueryAdder
------------------------
`Back To Top`_

APIPostQueryAdder is an interface for models to run after processing add function in dAPI and before returning the results. Returning false stops the rest of the process from happening.

Type:

.. code-block:: go

    interface

uadmin.APIPostQueryDelete
-------------------------
`Back To Top`_

APIPostQueryDelete controls the data API's post query for delete commands.

Type:

.. code-block:: go

    bool

uadmin.APIPostQueryDeleter
--------------------------
`Back To Top`_

APIPostQueryDeleter is an interface for models to run after processing delete function in dAPI and before returning the results. Returning false stops the rest of the process from happening.

Type:

.. code-block:: go

    interface

uadmin.APIPostQueryEdit
-----------------------
`Back To Top`_

APIPostQueryEdit controls the data API's post query for edit commands.

Type:

.. code-block:: go

    bool

uadmin.APIPostQueryEditor
-------------------------
`Back To Top`_

APIPostQueryEditor is an interface for models to run after processing edit function in dAPI and before returning the results. Returning false stops the rest of the process from happening.

Type:

.. code-block:: go

    interface

uadmin.APIPostQueryRead
-----------------------
`Back To Top`_

APIPostQueryRead controls the data API's post query for read commands.

Type:

.. code-block:: go

    bool

uadmin.APIPostQueryReader
-------------------------
`Back To Top`_

APIPostQueryReader is an interface for models to run after processing read function in dAPI and before returning the results. Returning false stops the rest of the process from happening.

Type:

.. code-block:: go

    interface

uadmin.APIPostQuerySchema
-------------------------
`Back To Top`_

APIPostQuerySchema controls the data API's post query for schema commands.

Type:

.. code-block:: go

    bool

uadmin.APIPostQuerySchemer
--------------------------
`Back To Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/dapi/post_query_functions.html#post-query-functions

APIPostQuerySchemer is an interface for models to run after processing schema function in dAPI and before returning the results. Returning false stops the rest of the process from happening.

Type:

.. code-block:: go

    interface
