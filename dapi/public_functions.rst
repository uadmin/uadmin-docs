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

APIPublicAdd controls the data API's public for add commands.

Type:

.. code-block:: go

    bool

uadmin.APIPublicAdder
---------------------
`Back To Top`_

APIPublicAdder is an interface that contains the method APIPublicAdd with the type of func(\*http.Request) bool. If the value is true, anyone will have permission access to add the record in the model using the dAPI add method even if the account is not logged in.

Type:

.. code-block:: go

    interface

uadmin.APIPublicDelete
----------------------
`Back To Top`_

APIPublicDelete controls the data API's public for delete commands.

Type:

.. code-block:: go

    bool

uadmin.APIPublicDeleter
-----------------------
`Back To Top`_

APIPublicDeleter is an interface that contains the method APIPublicDelete with the type of func(\*http.Request) bool. If the value is true, anyone will have permission access to delete the record in the model using the dAPI delete method even if the account is not logged in.

Type:

.. code-block:: go

    interface

uadmin.APIPublicEdit
--------------------
`Back To Top`_

APIPublicEdit controls the data API's public for edit commands.

Type:

.. code-block:: go

    bool

uadmin.APIPublicEditor
----------------------
`Back To Top`_

APIPublicEditor is an interface that contains the method APIPublicEdit with the type of func(\*http.Request) bool. If the value is true, anyone will have permission access to edit the record in the model using the dAPI edit method even if the account is not logged in.

Type:

.. code-block:: go

    interface

uadmin.APIPublicRead
--------------------
`Back To Top`_

APIPublicRead controls the data API's public for read commands.

Type:

.. code-block:: go

    bool

uadmin.APIPublicReader
----------------------
`Back To Top`_

APIPublicReader is an interface that contains the method APIPublicRead with the type of func(\*http.Request) bool. If the value is true, anyone will have permission access to read the record in the model using the dAPI read method even if the account is not logged in.

Type:

.. code-block:: go

    interface

uadmin.APIPublicSchema
----------------------
`Back To Top`_

APIPublicSchema controls the data API's public for schema commands.

Type:

.. code-block:: go

    bool

uadmin.APIPublicSchemer
-----------------------
`Back To Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/dapi/public_functions.html#public-functions

APIPublicSchemer is an interface that contains the method APIPublicSchema with the type of func(\*http.Request) bool. If the value is true, anyone will have permission access to read the full schema of the model using the dAPI schema method even if the account is not logged in.

Type:

.. code-block:: go

    interface
