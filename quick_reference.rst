<<<<<<< HEAD
Quick Reference
===============

Overriding Save Function
------------------------

.. code-block:: go

    func (m *Model) Save() {
        // business logic
        uadmin.Save(m)
    }

Validation
----------

.. code-block:: go

    func (v Validate) Validate() (ret map[string]string) {
        ret = map[string]string{}
        if v.Name != "test" {
            ret["Name"] = "Error name not found"
        }
        return
    }
=======
Quick Reference
===============

Overriding Save Function
------------------------

.. code-block:: go

    func (m *Model) Save() {
        // business logic
        uadmin.Save(m)
    }

Validation
----------

.. code-block:: go

    func (v Validate) Validate() (ret map[string]string) {
        ret = map[string]string{}
        if v.Name != "test" {
            ret["Name"] = "Error name not found"
        }
        return
    }
>>>>>>> de25cdd8a29ca2bb2c2df08be00b703b967aaed5
