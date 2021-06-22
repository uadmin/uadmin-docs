<<<<<<< HEAD
uAdmin Tutorial Part 6 - Applying uAdmin Tags
=============================================
Create a file named **item.go** inside your models folder with the following codes below:

.. code-block:: go

    package models

    import (
        "github.com/uadmin/uadmin"
    )

    // Item Model !
    type Item struct {
        uadmin.Model
        Name        string `uadmin:"required"`
        Description string
        Cost        int
        Rating      int
    }

Now register the model on main.go where `models` is the package name and `Item` is the model name:

.. code-block:: go

    func main() {
        uadmin.Register(
            models.Todo{},
            models.Category{},
            models.Friend{},
            models.Item{},  //  <-- place it here
        )

        // Existing RegisterInlines code

        // ----------- ADD THIS CODE -----------
        uadmin.RegisterInlines(models.Item{}, map[string]string{
            "Todo": "ItemID",
        })
        // ----------- ADD THIS CODE -----------

        uadmin.StartServer()
    }

Set the foreign key of an Item model to the Todo model and apply the tag `help` to inform the user what are the requirements needed in order to accomplish his activity.

.. code-block:: go

    package models

    import (
        "time"

        "github.com/uadmin/uadmin"
    )

    // Todo Model !
    type Todo struct {
        uadmin.Model
        Name        string
        Description string `uadmin:"html"`
        Category    Category
        CategoryID  uint
        Friend      Friend `uadmin:"help:Who will be a part of your activity?"`
        FriendID    uint
        
        // Item Model
        Item        Item `uadmin:"help:What are the requirements needed in order to accomplish your activity?"`
        // Item ID
        ItemID      uint

        TargetDate  time.Time
        Progress    int `uadmin:"progress_bar"`
    }

Now let's try something much cooler that we can apply in the Item model by adding different types of tags. Before we proceed, add more data in your Item model. Once you are done, let's add the `search` tag in the name field of item.go and see what happens.

.. code-block:: go

    package models

    import (
        "github.com/uadmin/uadmin"
    )

    // Item Model !
    type Item struct {
        uadmin.Model
        Name        string `uadmin:"required;search"` // <-- place it here
        Description string
        Cost        int
        Rating      int
    }

Result

.. image:: assets/searchtagapplied.png

|

Search the word **mini** and see what happens.

.. image:: assets/searchtagappliedoutput.png

|

Nice! Now go back to item.go and apply the tag `categorical_filter` and `filter` in the Name field and see what happens.

.. code-block:: go

    `categorical_filter;filter`

.. code-block:: go

	Name string `uadmin:"required;search;categorical_filter;filter"`

Rebuild your application. In Item model, click the filter button on the upper right.

Result

.. image:: assets/filtertagapplied.png

|

Now let's filter the word **iPad** and see what happens.

.. image:: assets/filtertagappliedoutput.png

|

We can also apply `display_name` tag with a given value such as **Product Name**.

.. code-block:: go

    `display_name:Product Name`

.. code-block:: go

    Name string `uadmin:"required;search;categorical_filter;filter;display_name:Product Name"`

|

Result

.. image:: assets/displaynametagapplied.png
   :align: center

|

uAdmin has a `default_value` tag which will generate a value automatically in the field. Let's say **Computer**.

.. code-block:: go

    `default_value:Computer`

.. code-block:: go

    Name string `uadmin:"required;search;categorical_filter;filter;display_name:Product Name;default_value:Computer"`

|

Result

.. image:: assets/defaultvaluetagapplied.png
   :align: center

|

You can also add `multilingual` tag in the Description field.

.. code-block:: go

    Description string `uadmin:"multilingual"`

|

Result

.. image:: assets/multilingualtagapplied.png

|

If you want to add more languages in your model, go to the Languages in the uAdmin dashboard.

.. image:: assets/languageshighlighted.png

|

Let's say I want to add Chinese and Tagalog in the Items model. In order to do that, set the Active as enabled.

.. image:: assets/activehighlighted.png
   :align: center

|

Now go back to the Items model and see what happens.

.. image:: assets/multilingualtagappliedmultiple.png

To customize your own languages, visit `Language`_ for the instructions.

.. _Language: https://uadmin-docs.readthedocs.io/en/latest/system_reference.html#language

|

In the Cost field, set the `money` tag and see what happens.

.. code-block:: go

    Cost int `uadmin:"money"`

|

Result

.. image:: assets/moneytagapplied.png

|

You can also set `pattern` and `pattern_msg` tag in the Cost field. This means the user must input numbers only. If he inputs otherwise, the pattern message will show up on the screen.

.. code-block:: go

    Cost int `uadmin:"money;pattern:^[0-9]*$;pattern_msg:Your input must be a number."`

|

Result

.. image:: assets/patterntagapplied.png

|

To solve this case, we can use a help tag feature in order to give users a solution to the complex tasks encountered in the model.

.. code-block:: go

    `help:Input numeric characters only in this field.`

.. code-block:: go

    Cost int `uadmin:"money;pattern:^[0-9]*$;pattern_msg:Your input must be a number.;help:Input numeric characters only in this field."`

|

Result

.. image:: assets/helptagapplied.png
   :align: center

|

We can also use min and max tags in the Rating field. Min tag means the minimum value that a user can input and the max one means the maximum value. Let's set the min value as 1 and the max value as 5.

.. code-block:: go

    Rating int `uadmin:"min:1;max:5"`

|

See what happens if the user inputs the value outside the range.

.. image:: assets/minmaxtagapplied.png

|

Well done! Now you know how to apply most of the tags available in our uAdmin framework that are functional in our Todo List project.

See `Tag Reference`_ for more examples.

Click `here`_ to view our progress so far.

In the `next part`_, we will talk on the concept of M2M and how is it useful in our project.


.. _Tag Reference: https://uadmin-docs.readthedocs.io/en/latest/tags.html
.. _here: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part6.html
.. _next part: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part7.html

.. toctree::
   :maxdepth: 1

   full_code/part6
=======
uAdmin Tutorial Part 6 - Applying uAdmin Tags
=============================================
Create a file named **item.go** inside your models folder with the following codes below:

.. code-block:: go

    package models

    import (
        "github.com/uadmin/uadmin"
    )

    // Item Model !
    type Item struct {
        uadmin.Model
        Name        string `uadmin:"required"`
        Description string
        Cost        int
        Rating      int
    }

Now register the model on main.go where `models` is the package name and `Item` is the model name:

.. code-block:: go

    func main() {
        uadmin.Register(
            models.Todo{},
            models.Category{},
            models.Friend{},
            models.Item{},  //  <-- place it here
        )

        // Existing RegisterInlines code

        // ----------- ADD THIS CODE -----------
        uadmin.RegisterInlines(models.Item{}, map[string]string{
            "Todo": "ItemID",
        })
        // ----------- ADD THIS CODE -----------

        uadmin.StartServer()
    }

Set the foreign key of an Item model to the Todo model and apply the tag `help` to inform the user what are the requirements needed in order to accomplish his activity.

.. code-block:: go

    package models

    import (
        "time"

        "github.com/uadmin/uadmin"
    )

    // Todo Model !
    type Todo struct {
        uadmin.Model
        Name        string
        Description string `uadmin:"html"`
        Category    Category
        CategoryID  uint
        Friend      Friend `uadmin:"help:Who will be a part of your activity?"`
        FriendID    uint
        
        // Item Model
        Item        Item `uadmin:"help:What are the requirements needed in order to accomplish your activity?"`
        // Item ID
        ItemID      uint

        TargetDate  time.Time
        Progress    int `uadmin:"progress_bar"`
    }

Now let's try something much cooler that we can apply in the Item model by adding different types of tags. Before we proceed, add more data in your Item model. Once you are done, let's add the `search` tag in the name field of item.go and see what happens.

.. code-block:: go

    package models

    import (
        "github.com/uadmin/uadmin"
    )

    // Item Model !
    type Item struct {
        uadmin.Model
        Name        string `uadmin:"required;search"` // <-- place it here
        Description string
        Cost        int
        Rating      int
    }

Result

.. image:: assets/searchtagapplied.png

|

Search the word **mini** and see what happens.

.. image:: assets/searchtagappliedoutput.png

|

Nice! Now go back to item.go and apply the tag `categorical_filter` and `filter` in the Name field and see what happens.

.. code-block:: go

    `categorical_filter;filter`

.. code-block:: go

	Name string `uadmin:"required;search;categorical_filter;filter"`

Rebuild your application. In Item model, click the filter button on the upper right.

Result

.. image:: assets/filtertagapplied.png

|

Now let's filter the word **iPad** and see what happens.

.. image:: assets/filtertagappliedoutput.png

|

We can also apply `display_name` tag with a given value such as **Product Name**.

.. code-block:: go

    `display_name:Product Name`

.. code-block:: go

    Name string `uadmin:"required;search;categorical_filter;filter;display_name:Product Name"`

|

Result

.. image:: assets/displaynametagapplied.png
   :align: center

|

uAdmin has a `default_value` tag which will generate a value automatically in the field. Let's say **Computer**.

.. code-block:: go

    `default_value:Computer`

.. code-block:: go

    Name string `uadmin:"required;search;categorical_filter;filter;display_name:Product Name;default_value:Computer"`

|

Result

.. image:: assets/defaultvaluetagapplied.png
   :align: center

|

You can also add `multilingual` tag in the Description field.

.. code-block:: go

    Description string `uadmin:"multilingual"`

|

Result

.. image:: assets/multilingualtagapplied.png

|

If you want to add more languages in your model, go to the Languages in the uAdmin dashboard.

.. image:: assets/languageshighlighted.png

|

Let's say I want to add Chinese and Tagalog in the Items model. In order to do that, set the Active as enabled.

.. image:: assets/activehighlighted.png
   :align: center

|

Now go back to the Items model and see what happens.

.. image:: assets/multilingualtagappliedmultiple.png

To customize your own languages, visit `Language`_ for the instructions.

.. _Language: https://uadmin-docs.readthedocs.io/en/latest/system_reference.html#language

|

In the Cost field, set the `money` tag and see what happens.

.. code-block:: go

    Cost int `uadmin:"money"`

|

Result

.. image:: assets/moneytagapplied.png

|

You can also set `pattern` and `pattern_msg` tag in the Cost field. This means the user must input numbers only. If he inputs otherwise, the pattern message will show up on the screen.

.. code-block:: go

    Cost int `uadmin:"money;pattern:^[0-9]*$;pattern_msg:Your input must be a number."`

|

Result

.. image:: assets/patterntagapplied.png

|

To solve this case, we can use a help tag feature in order to give users a solution to the complex tasks encountered in the model.

.. code-block:: go

    `help:Input numeric characters only in this field.`

.. code-block:: go

    Cost int `uadmin:"money;pattern:^[0-9]*$;pattern_msg:Your input must be a number.;help:Input numeric characters only in this field."`

|

Result

.. image:: assets/helptagapplied.png
   :align: center

|

We can also use min and max tags in the Rating field. Min tag means the minimum value that a user can input and the max one means the maximum value. Let's set the min value as 1 and the max value as 5.

.. code-block:: go

    Rating int `uadmin:"min:1;max:5"`

|

See what happens if the user inputs the value outside the range.

.. image:: assets/minmaxtagapplied.png

|

Well done! Now you know how to apply most of the tags available in our uAdmin framework that are functional in our Todo List project.

See `Tag Reference`_ for more examples.

Click `here`_ to view our progress so far.

In the `next part`_, we will talk on the concept of M2M and how is it useful in our project.


.. _Tag Reference: https://uadmin-docs.readthedocs.io/en/latest/tags.html
.. _here: https://uadmin-docs.readthedocs.io/en/latest/tutorial/full_code/part6.html
.. _next part: https://uadmin-docs.readthedocs.io/en/latest/tutorial/part7.html

.. toctree::
   :maxdepth: 1

   full_code/part6
>>>>>>> de25cdd8a29ca2bb2c2df08be00b703b967aaed5
