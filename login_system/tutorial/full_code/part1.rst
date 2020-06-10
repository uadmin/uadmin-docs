Login System Tutorial Part 1 - Build A Project (Current Progress)
=================================================================
`Back to Previous Page`_

.. _Back to Previous Page: https://uadmin-docs.readthedocs.io/en/latest/login_system/tutorial/part1.html

Structure:

* `main.go`_

.. _main.go: https://uadmin-docs.readthedocs.io/en/latest/login_system/tutorial/full_code/part1.html#id1

main.go
-------
`Back to Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/login_system/tutorial/full_code/part1.html#login-system-tutorial-part-1-build-a-project-current-progress

.. code-block:: go

    package main

    import (
        "net/http"

        "github.com/uadmin/uadmin"
    )

    func main() {
        // Assign RootURL value as "/admin/" and Site Name as "Login System"
        // NOTE: This code works only on first build.
        uadmin.RootURL = "/admin/"
        uadmin.SiteName = "Login System"

        // Run the server
        uadmin.StartServer()
    }
