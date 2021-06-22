uAdmin Tutorial Part 5 - Register Inlines and Drop Down List
============================================================
Inlines is where we keep all registered models' inlines. It allows you to merge a parent model to a submodel where the foreign key(s) are specified.

**Why do we use Register Inlines?** We use them to show that the field of a model is related to another model as long as there is a foreign key specified.

Structure:

.. code-block:: go

    uadmin.RegisterInlines(/package_name/./struct_name of a parent model/{}, map[string]string{
        "/sub_model name/": "/parent_model name/ID",
    })

Now let's apply it in the main.go. Copy the codes below and paste it after the uadmin.Register function.

.. code-block:: go

    uadmin.RegisterInlines(models.Category{}, map[string]string{
        "Todo": "CategoryID",
    })
    uadmin.RegisterInlines(models.Friend{}, map[string]string{
        "Todo": "FriendID",
    })

Let's run the application and see what happens.

.. image:: assets/registerinlinetodo.png

|

The parent model Todo is now included in the Category submodel as shown above. You can go to Friends model and it will display the same result.

Drop Down List in a Field
^^^^^^^^^^^^^^^^^^^^^^^^^
Go to the friend.go in the models folder. Let's create a drop down list selection in the Nationality field. In order to do that, initialize a variable with the type int. Create a function that will set an item and return the integer value inside it. One function is equivalent to one item. Put it above the Friend model.

.. code-block:: go

    // Nationality Field !
    type Nationality int

    // Chinese !
    func (Nationality) Chinese() Nationality {
        return 1
    }

    // Filipino !
    func (Nationality) Filipino() Nationality {
        return 2
    }

    // Others !
    func (Nationality) Others() Nationality {
        return 3
    }

Now inside the Friend model, initialize a Nationality field so that it will be created.

.. code-block:: go

    // Friend Model !
    type Friend struct {
        uadmin.Model
        Name        string `uadmin:"required"`
        Email       string `uadmin:"email"`
        Password    string `uadmin:"password;list_exclude"`
        Nationality Nationality // <-- place it here
    }

Result

.. image:: assets/nationalityhighlighted.png

|

We can also add an Invite field that will direct you to his website. In order to do that, set the field name as "Invite" with the tag "link".

.. code-block:: go

    // Friend Model !
    type Friend struct {
        uadmin.Model
        Name        string `uadmin:"required"`
        Email       string `uadmin:"email"`
        Password    string `uadmin:"password;list_exclude"`
        Nationality Nationality
        Invite      string `uadmin:"link"` // <-- place it here
    }

Add the overriding save function after the Friend struct.

.. code-block:: go

    // Save !
    func (f *Friend) Save() {
        f.Invite = "https://www.google.com/"
        uadmin.Save(f)
    }

Run your application, go to the Friends model and update the elements inside. Afterwards, click the Invite button in the list view and see what happens.

.. image:: assets/invitebuttonhighlighted.png

|

Result

.. image:: assets/googlewebsitescreen.png
   :align: center

|

Congrats, now you know how to do the following:

* Register Inlines
* Adding drop down list manually to a field
* How to use "link" tag in our project
* Applying override save function

Click `here`_ to view our progress so far.

In the `next part`_ we will discuss about applying different uadmin tags through our project.

.. _here: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part5.html

.. _next part: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part6.html

.. toctree::
   :maxdepth: 1

   full_code/part5
