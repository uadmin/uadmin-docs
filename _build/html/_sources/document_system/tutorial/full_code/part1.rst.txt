Document System Tutorial Part 1 - Build A Project (Current Progress)
====================================================================
`Back to Previous Page`_

.. _Back to Previous Page: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/part1.html

Structure:

* `main.go`_

.. _main.go: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part1.html#id1

main.go
-------
`Back to Top`_

.. _Back To Top: https://uadmin-docs.readthedocs.io/en/latest/document_system/tutorial/full_code/part1.html#document-system-tutorial-part-1-build-a-project-current-progress

.. code-block:: go

    package main

    import (
        "github.com/uadmin/uadmin"
    )

    func main() {
        // Assign Site Name value as "Document System"
        // NOTE: This code works only on first build.
        uadmin.SiteName = "Document System"

        // Activates a uAdmin server
        uadmin.StartServer()
    }
