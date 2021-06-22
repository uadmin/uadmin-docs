package main

import (
	"net/http"

	"github.com/uadmin/uadmin-docs/sample_project/todo/api"
	"github.com/uadmin/uadmin-docs/sample_project/todo/models"
	"github.com/uadmin/uadmin-docs/sample_project/todo/views"

	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.Register(
		models.Todo{},
		models.Category{},
		models.Friend{},
		models.Item{},
	)

	uadmin.RegisterInlines(models.Category{}, map[string]string{
		"Todo": "CategoryID",
	})
	uadmin.RegisterInlines(models.Friend{}, map[string]string{
		"Todo": "FriendID",
	})
	uadmin.RegisterInlines(models.Item{}, map[string]string{
		"Todo": "ItemID",
	})

	// Initialize Setting model
	setting := uadmin.Setting{}

	// Get the code
	uadmin.Get(&setting, "code = ?", "uAdmin.RootURL")

	// Assign the value as "/admin/"
	setting.ParseFormValue([]string{"/admin/"})

	// Save changes
	setting.Save()

	// Assign Site Name in the Settings
	setting = uadmin.Setting{}
	uadmin.Get(&setting, "code = ?", "uAdmin.SiteName")
	setting.ParseFormValue([]string{"Todo List"})
	setting.Save()

	// API Handler
	http.HandleFunc("/api/", uadmin.Handler(api.Handler))

	// HTTP UI Handler
	http.HandleFunc("/http_handler/", uadmin.Handler(views.HTTPHandler))

	uadmin.StartServer()
}
