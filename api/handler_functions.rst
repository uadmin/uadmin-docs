Handler Functions
=================
`Back To uAdmin Functions List`_

.. _Back To uAdmin Functions List: https://uadmin-docs.readthedocs.io/en/latest/api.html#api-reference

In this section, we will cover the following functions in-depth listed below:

* `uadmin.Handler`_
* `uadmin.StaticHandler`_
* `uadmin.UploadImageHandler`_

uadmin.Handler
--------------
`Back To Top`_

Handler is a function that allows access to another handler.

Function:

.. code-block:: go

    func(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request)

Used in the tutorial:

* `uAdmin Tutorial Part 9 - Introduction to API`_
* `uAdmin Tutorial Part 13 - Accessing an HTML file`_

.. _uAdmin Tutorial Part 9 - Introduction to API: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part9.html
.. _uAdmin Tutorial Part 13 - Accessing an HTML file: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part13.html

uadmin.StaticHandler
--------------------
`Back To Top`_

StaticHandler is a function that serves static files.

Function:

.. code-block:: go

    func(w http.ResponseWriter, r *http.Request)

Parameters:

    **w http.ResponseWriter:** Assembles the HTTP server's response; by writing to it, we send data to the HTTP client

    **r \*http.Request** Is a data structure that represents the client HTTP request

See `Part 2: Static Handler`_ for the example.

.. _Part 2\: Static Handler: https://uadmin-docs.readthedocs.io/en/latest/api/abtest.html#part-2-static-handler

uadmin.UploadImageHandler
-------------------------
`Back To Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/api/handler_functions.html#handler-functions

UploadImageHandler handles files sent from Tiny MCE's photo uploader.

Function:

.. code-block:: go

    func(w http.ResponseWriter, r *http.Request, session *Session)

Parameters:

    **w http.ResponseWriter:** Assembles the HTTP server's response; by writing to it, we send data to the HTTP client

    **r \*http.Request** Is a data structure that represents the client HTTP request

    **session \*Session** Is an activity that a user with a unique IP address spends on a Web site during a specified period of time
