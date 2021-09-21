Public Functions
================
`Back To uAdmin Functions List`_

.. _Back To uAdmin Functions List: https://uadmin-docs.readthedocs.io/en/latest/dapi.html#uadmin-functions

In this section, we will cover the following functions in-depth listed below:

* `uadmin.APIPublicAdd`_
* `uadmin.APIPublicAdder`_
* `uadmin.APIPublicDelete`_
* `uadmin.APIPublicDeleter`_
* `uadmin.APIPublicEdit`_
* `uadmin.APIPublicEditor`_
* `uadmin.APIPublicRead`_
* `uadmin.APIPublicReader`_
* `uadmin.APIPublicSchema`_
* `uadmin.APIPublicSchemer`_

uadmin.APIPublicAdd
-------------------
`Back To Top`_

APIPublicAdd controls the data API's public for add commands. By default, APIPublicAdd returns **false**.

Type:

.. code-block:: go

    bool

uadmin.APIPublicAdder
---------------------
`Back To Top`_

APIPublicAdder is an interface for models to control public access to add function in dAPI.

Type:

.. code-block:: go

    interface

uadmin.APIPublicDelete
----------------------
`Back To Top`_

APIPublicDelete controls the data API's public for delete commands. By default, APIPublicDelete returns **false**.

Type:

.. code-block:: go

    bool

uadmin.APIPublicDeleter
-----------------------
`Back To Top`_

APIPublicDeleter is an interface for models to control public access to delete function in dAPI.

Type:

.. code-block:: go

    interface

uadmin.APIPublicEdit
--------------------
`Back To Top`_

APIPublicEdit controls the data API's public for edit commands. By default, APIPublicEdit returns **false**.

Type:

.. code-block:: go

    bool

uadmin.APIPublicEditor
----------------------
`Back To Top`_

APIPublicEditor is an interface for models to control public access to edit function in dAPI.

Type:

.. code-block:: go

    interface

uadmin.APIPublicRead
--------------------
`Back To Top`_

APIPublicRead controls the data API's public for read commands. By default, APIPublicRead returns **false**.

Type:

.. code-block:: go

    bool

uadmin.APIPublicReader
----------------------
`Back To Top`_

APIPublicReader is an interface for models to control public access to read function in dAPI.

Type:

.. code-block:: go

    interface

uadmin.APIPublicSchema
----------------------
`Back To Top`_

APIPublicSchema controls the data API's public for schema commands. By default, APIPublicSchema returns **false**.

Type:

.. code-block:: go

    bool

uadmin.APIPublicSchemer
-----------------------
`Back To Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/dapi/public_functions.html#public-functions

APIPublicSchemer is an interface for models to control public access to schema function in dAPI.

Type:

.. code-block:: go

    interface
