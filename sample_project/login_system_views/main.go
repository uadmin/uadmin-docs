package main

import (
	"net/http"

	"github.com/uadmin/uadmin"
	"github.com/uadmin/uadmin-docs/sample_project/login_system_views/views"
)

func main() {
	// Assign RootURL value as "/admin/" and Site Name as "Login System"
	// NOTE: This code works only if database does not exist yet.
	uadmin.RootURL = "/admin/"
	uadmin.SiteName = "Login System"

	// Login System Main Handler
	http.HandleFunc("/login_system/", uadmin.Handler(views.MainHandler))

	// Run the server
	uadmin.StartServer()
}
