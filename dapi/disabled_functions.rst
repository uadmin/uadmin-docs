Disabled Functions
==================
`Back To uAdmin Functions List`_

.. _Back To uAdmin Functions List: https://uadmin-docs.readthedocs.io/en/latest/dapi.html#uadmin-functions

In this section, we will cover the following functions in-depth listed below:

* `uadmin.APIDisabledAdd`_
* `uadmin.APIDisabledAdder`_
* `uadmin.APIDisabledDelete`_
* `uadmin.APIDisabledDeleter`_
* `uadmin.APIDisabledEdit`_
* `uadmin.APIDisabledEditor`_
* `uadmin.APIDisabledRead`_
* `uadmin.APIDisabledReader`_
* `uadmin.APIDisabledSchema`_
* `uadmin.APIDisabledSchemer`_

uadmin.APIDisabledAdd
---------------------
`Back To Top`_

APIDisabledAdd controls the data API's disabled for add commands. By default, APIDisabledAdd returns **false**.

Type:

.. code-block:: go

    bool

uadmin.APIDisabledAdder
-----------------------
`Back To Top`_

APIDisabledAdder is an interface for models to disable access to add function in dAPI.

Type:

.. code-block:: go

    interface

uadmin.APIDisabledDelete
------------------------
`Back To Top`_

APIDisabledDelete controls the data API's disabled for delete commands. By default, APIDisabledDelete returns **false**.

Type:

.. code-block:: go

    bool

uadmin.APIDisabledDeleter
-------------------------
`Back To Top`_

APIDisabledDeleter is an interface for models to disable access to delete function in dAPI.

Type:

.. code-block:: go

    interface

uadmin.APIDisabledEdit
----------------------
`Back To Top`_

APIDisabledEdit controls the data API's disabled for edit commands. By default, APIDisabledEdit returns **false**.

Type:

.. code-block:: go

    bool

uadmin.APIDisabledEditor
------------------------
`Back To Top`_

APIDisabledEditor is an interface for models to disable access to edit function in dAPI.

Type:

.. code-block:: go

    interface

uadmin.APIDisabledRead
----------------------
`Back To Top`_

APIDisabledRead controls the data API's disabled for read commands. By default, APIDisabledRead returns **false**.

Type:

.. code-block:: go

    bool

uadmin.APIDisabledReader
------------------------
`Back To Top`_

APIDisabledReader is an interface for models to disable access to read function in dAPI.

Type:

.. code-block:: go

    interface

uadmin.APIDisabledSchema
------------------------
`Back To Top`_

APIDisabledSchema controls the data API's disabled for schema commands. By default, APIDisabledSchema returns **false**.

Type:

.. code-block:: go

    bool

uadmin.APIDisabledSchemer
-------------------------
`Back To Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/dapi/disabled_functions.html#disabled-functions

APIDisabledSchemer is an interface for models to disable access to schema function in dAPI.

Type:

.. code-block:: go

    interface
