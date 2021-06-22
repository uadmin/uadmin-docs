<<<<<<< HEAD
uAdmin Tutorial Part 11 - Inserting and Saving the Record
=========================================================
Create a file named **add_friend.go** inside the api folder with the following codes below:

.. code-block:: go

    package api

    import (
        "net/http"

        // Specify the username that you used inside github.com folder
        "github.com/username/todo/models"
        "github.com/uadmin/uadmin"
    )

    // AddFriendAPIHandler !
    func AddFriendAPIHandler(w http.ResponseWriter, r *http.Request) {
        res := map[string]interface{}{}

        // Fetch data from Friend DB
        friend := models.Friend{}

        // Set the parameters of Name, Email, and Password
        friendName := r.FormValue("name")
        friendEmail := r.FormValue("email")
        friendPassword := r.FormValue("password")

        // Validate if the friendName variable is empty.
        if friendName == "" {
            res["status"] = "ERROR"
            res["err_msg"] = "Name is required."
            uadmin.ReturnJSON(w, r, res)
            return
        }

        // Store input into the Name, Email, and Password fields
        friend.Name = friendName
        friend.Email = friendEmail
        friend.Password = friendPassword

        // Store input in the Friend model
        uadmin.Save(&friend)

        res["status"] = "ok"
        uadmin.ReturnJSON(w, r, res)
    }

Finally, add the following pieces of code in the api.go shown below. This will establish a communication between the AddFriendHandler and the APIHandler.

.. code-block:: go

    func APIHandler(w http.ResponseWriter, r *http.Request) {

        // Some codes contained in this part

        // --------------------- ADD THIS CODE ---------------------
        if strings.HasPrefix(r.URL.Path, "/add_friend") {
            AddFriendAPIHandler(w, r)
            return
        }
        // ---------------------------------------------------------
    }

Now run your application. In order to insert the information in the Friend model, put the **?** symbol after /api/add_friend path which means **WHERE** in SQL, followed by the parameter name. Set the value of each parameter to store your input and save into the Friend model. **&** symbol is equivalent to **AND** in SQL.

* name = Allen
* email = allen@gmail.com
* password = 123456

.. code-block:: bash

    http://localhost:8080/api/add_friend?name=Allen&email=allen@gmail.com&password=123456

.. code-block:: json

    {
      "status": "ok"
    }

|

Go back to the Friend model. You will notice that Allen was added inside it.

.. image:: assets/todomodeladdfriend.png

|

Congrats, now you know how to insert and save a record to the model in the API using multiple parameters.

See `API Reference`_ for more examples.

Click `here`_ to view our progress so far.

In the `next part`_, we will discuss about designing a table in HTML and setting up a template file.

.. _API Reference: https://uadmin-docs.readthedocs.io/en/latest/api.html
.. _here: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part11.html
.. _next part: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part12.html

.. toctree::
   :maxdepth: 1

   full_code/part11
=======
uAdmin Tutorial Part 11 - Inserting and Saving the Record
=========================================================
Create a file named **add_friend.go** inside the api folder with the following codes below:

.. code-block:: go

    package api

    import (
        "net/http"

        // Specify the username that you used inside github.com folder
        "github.com/username/todo/models"
        "github.com/uadmin/uadmin"
    )

    // AddFriendAPIHandler !
    func AddFriendAPIHandler(w http.ResponseWriter, r *http.Request) {
        res := map[string]interface{}{}

        // Fetch data from Friend DB
        friend := models.Friend{}

        // Set the parameters of Name, Email, and Password
        friendName := r.FormValue("name")
        friendEmail := r.FormValue("email")
        friendPassword := r.FormValue("password")

        // Validate if the friendName variable is empty.
        if friendName == "" {
            res["status"] = "ERROR"
            res["err_msg"] = "Name is required."
            uadmin.ReturnJSON(w, r, res)
            return
        }

        // Store input into the Name, Email, and Password fields
        friend.Name = friendName
        friend.Email = friendEmail
        friend.Password = friendPassword

        // Store input in the Friend model
        uadmin.Save(&friend)

        res["status"] = "ok"
        uadmin.ReturnJSON(w, r, res)
    }

Finally, add the following pieces of code in the api.go shown below. This will establish a communication between the AddFriendHandler and the APIHandler.

.. code-block:: go

    func APIHandler(w http.ResponseWriter, r *http.Request) {

        // Some codes contained in this part

        // --------------------- ADD THIS CODE ---------------------
        if strings.HasPrefix(r.URL.Path, "/add_friend") {
            AddFriendAPIHandler(w, r)
            return
        }
        // ---------------------------------------------------------
    }

Now run your application. In order to insert the information in the Friend model, put the **?** symbol after /api/add_friend path which means **WHERE** in SQL, followed by the parameter name. Set the value of each parameter to store your input and save into the Friend model. **&** symbol is equivalent to **AND** in SQL.

* name = Allen
* email = allen@gmail.com
* password = 123456

.. code-block:: bash

    http://localhost:8080/api/add_friend?name=Allen&email=allen@gmail.com&password=123456

.. code-block:: json

    {
      "status": "ok"
    }

|

Go back to the Friend model. You will notice that Allen was added inside it.

.. image:: assets/todomodeladdfriend.png

|

Congrats, now you know how to insert and save a record to the model in the API using multiple parameters.

See `API Reference`_ for more examples.

Click `here`_ to view our progress so far.

In the `next part`_, we will discuss about designing a table in HTML and setting up a template file.

.. _API Reference: https://uadmin-docs.readthedocs.io/en/latest/api.html
.. _here: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part11.html
.. _next part: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part12.html

.. toctree::
   :maxdepth: 1

   full_code/part11
>>>>>>> de25cdd8a29ca2bb2c2df08be00b703b967aaed5
