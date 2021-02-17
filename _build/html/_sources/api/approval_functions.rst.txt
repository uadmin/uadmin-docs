Approval Functions
==================
`Back To uAdmin Functions List`_

.. _Back To uAdmin Functions List: https://uadmin-docs.readthedocs.io/en/latest/api.html#api-reference

In this section, we will cover the following functions in-depth listed below:

* `uadmin.Approval`_
    * `func (*Approval) Save`_
    * `func (*Approval) String`_
* `uadmin.ApprovalAction`_
    * `func (ApprovalAction) Approved`_
    * `func (ApprovalAction) Rejected`_
* `uadmin.ApprovalHandleFunc`_

uadmin.Approval
---------------
`Back To Top`_

.. code-block:: go

    type Approval struct {
        Model
        ModelName           string `uadmin:"read_only"`
        ModelPK             uint   `uadmin:"read_only"`
        ColumnName          string `uadmin:"read_only"`
        OldValue            string `uadmin:"read_only"`
        NewValue            string
        NewValueDescription string    `uadmin:"read_only"`
        ChangedBy           string    `uadmin:"read_only"`
        ChangeDate          time.Time `uadmin:"read_only"`
        ApprovalAction      ApprovalAction
        ApprovalBy          string     `uadmin:"read_only"`
        ApprovalDate        *time.Time `uadmin:"read_only"`
        ViewRecord          string     `uadmin:"link"`
        UpdatedBy           string     `uadmin:"read_only;hidden;list_exclude"`
    }

Approval is a model that stores approval data.

Here are the following fields and their definitions:

* **ModelName** - The name of the Model in small letters
* **ModelPK** - Used to uniquely identify each row in the table
* **ColumnName** - The name of the column in the model
* **OldValue** - A value that was assigned before
* **NewValue** - A value that you want to replace from the old value
* **NewValueDescription** - A value that was stored from the new value after saving
* **ChangedBy** - Returns the username who changed the value of the field record
* **ChangeDate** - The date when the value of the field record was changed
* **ApprovalAction** - A selection of approval actions. There are two selections: Approved and Declined.
* **ApprovalBy** - Returns the username who approved the value of the field record
* **ApprovalDate** - The date when the value of the field record was approved
* **ViewRecord** - A link to view the information of the actual record
* **UpdatedBy** - Returns the username who updated the record

**func (\*Approval) Save**
^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back To Top`_

.. code-block:: go

    func (a *Approval) Save()

Save overrides save.

**func (\*Approval) String**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back To Top`_

.. code-block:: go

    func (a *Approval) String() string

String returns the Model Name, Model PK, and Column Name.

There are 2 ways you can do for initialization process using this function: one-by-one and by group.

One-by-one initialization:

.. code-block:: go

    func main(){
        // Some codes
        approval := uadmin.Approval{}
        approval.ModelName = "Model Name"
        approval.ModelPK = 1
        approval.ColumnName = "Column Name"
    }

By group initialization:

.. code-block:: go

    func main(){
        // Some codes
        approval := uadmin.Approval{
            ModelName:  "Model Name",
            ModelPK:    1,
            ColumnName: "Column Name",
        }
    }

In this example, we will use "by group” initialization process.

Suppose the user has assigned a name by mistake and the approver has rejected the record by mistake.

.. image:: ../assets/approvalrejectedbymistake.png

|

Go to the main.go and apply the following codes below:

.. code-block:: go

    func main(){
        // Some codes

        // Get the record in the user model where the approver is an admin
        approver := uadmin.User{}
        uadmin.Get(&approver, "id = ?", 1)

        // Get the record in the user model where johndoe is the user
        user := uadmin.User{}
        uadmin.Get(&user, "id = ?", 2)

        // Get an old approval record
        a := uadmin.Approval{}
        uadmin.Get(&a, "id = ?", 1)

        // Assign the date and time today
        now := time.Now()

        // Subtract the time by one hour
        then := now.Add(-time.Hour)

        // Assign a value that you want to update in the old record
        b := uadmin.Approval{
            OldValue:            a.NewValueDescription,
            NewValue:            "Jane Doe",
            NewValueDescription: "Jane Doe",
            ChangedBy:           user.Username,
            ChangeDate:          then,
            ApprovalAction:      uadmin.ApprovalAction(0).Approved(),
            ApprovalBy:          approver.Username,
            ApprovalDate:        &now,
        }

        // Update the record here
        uadmin.Update(&a, "old_value", b.OldValue, "id = ?", 1)
        uadmin.Update(&a, "new_value", b.NewValue, "id = ?", 1)
        uadmin.Update(&a, "new_value_description", b.NewValueDescription, "id = ?", 1)
        uadmin.Update(&a, "changed_by", b.ChangedBy, "id = ?", 1)
        uadmin.Update(&a, "change_date", b.ChangeDate, "id = ?", 1)
        uadmin.Update(&a, "approval_action", b.ApprovalAction, "id = ?", 1)
        uadmin.Update(&a, "approval_by", b.ApprovalBy, "id = ?", 1)
        uadmin.Update(&a, "approval_date", b.ApprovalDate, "id = ?", 1)

        // Get the record in the friend model based on the model primary key
        friend := models.Friend{}
        uadmin.Get(&friend, "id = ?", a.ModelPK)

        // Update the name to the Friend model based on the model primary key
        uadmin.Update(&friend, "name", a.NewValue, "id = ?", a.ModelPK)
    }

Run your application and see what happens.

.. image:: ../assets/approvalrecordupdatedapi.png

|

As expected, the old approval record has been updated. Now click on View Record button to see if the name of the Friend and the approval status were updated.

.. image:: ../assets/friendnameapprovalupdated.png
   :align: center

Quiz:

* `Approval Functions`_

uadmin.ApprovalAction
---------------------
`Back To Top`_

.. code-block:: go

    type ApprovalAction int

ApprovalAction is a selection of approval actions.

**func (ApprovalAction) Approved**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back To Top`_

.. code-block:: go

    func (ApprovalAction) Approved() ApprovalAction

Approved is an accepted change.

**func (ApprovalAction) Rejected**
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`Back To Top`_

.. code-block:: go

    func (ApprovalAction) Rejected() ApprovalAction

Rejected is a rejected change.

See `uadmin.Approval`_  and `uadmin.ApprovalHandleFunc`_ for the examples.

Quiz:

* `Approval Functions`_

uadmin.ApprovalHandleFunc
-------------------------
`Back To Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/api/approval_functions.html#approval-functions

.. code-block:: go

    var ApprovalHandleFunc func(*Approval) bool

ApprovalHandleFunc is a function that could be called during the save process of each approval.

Before you proceed to this example, see `uadmin.Approval`_, `Approval Tag`_, or `Approval System`_.

.. _Approval Tag: https://uadmin-docs.readthedocs.io/en/latest/tags.html#approval
.. _Approval System: https://uadmin-docs.readthedocs.io/en/latest/system_reference.html#approval

Suppose you have this email in the Friend model that has an ID of 1.

.. image:: ../assets/johndoeoldemail.png
   :align: center

|

Go to the main.go and apply the following codes below:

.. code-block:: go

    func main() {
        // Some codes

        // Generic approval validation based on the model that was called
        uadmin.ApprovalHandleFunc = func(a *uadmin.Approval) bool {

            // Check whether a model name is friend and approval action is
            // Approved
            if a.ModelName == "friend" && a.ApprovalAction == a.ApprovalAction.Approved() {

                // Initialize a Friend model
                f := models.Friend{}

                // Get the record based on the model primary key
                uadmin.Get(&f, "id = ?", a.ModelPK)

                // Assign an email
                f.Email = "uadmin-support@email.com"

                // Save changes
                uadmin.Save(&f)

                return true
            }
            return false
        }
    }

From uAdmin dashboard, go to Approvals model, click the existing record in the list, choose Approved in Approval Action then click Save button below.

.. image:: ../assets/friendemailapproved.png

|

Click View Button on the right side of the record to see the result.

.. image:: ../assets/janedoeviewrecord.png

|

As expected, the email on that record was changed.

.. image:: ../assets/uadminsupportemail.png
   :align: center

Quiz:

* `Approval Functions`_

.. _Approval Functions: https://uadmin-docs.readthedocs.io/en/latest/_static/quiz/approval-functions.html
