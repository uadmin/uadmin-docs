package main

import (
	"net/http"

	"github.com/uadmin/uadmin-docs/sample_project/login_system/views"
	"github.com/uadmin/uadmin"
)

func main() {
	// Assign RootURL value as "/admin/" and Site Name as "Login System"
	// NOTE: This code works only on first build.
	uadmin.RootURL = "/admin/"
	uadmin.SiteName = "Login System"

	// Login Handler
	http.HandleFunc("/login/", uadmin.Handler(views.LoginHandler))

	// Run the server
	uadmin.StartServer()
}
