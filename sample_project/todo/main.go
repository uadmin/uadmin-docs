package main

import (
	"net/http"

	"github.com/uadmin/uadmin-docs/sample_project/todo/api"
	"github.com/uadmin/uadmin-docs/sample_project/todo/models"
	"github.com/uadmin/uadmin-docs/sample_project/todo/views"

	"github.com/uadmin/uadmin"
)

func main() {
	// Register Models
	uadmin.Register(
		models.Todo{},
		models.Category{},
		models.Friend{},
		models.Item{},
	)

	// Register Inlines
	uadmin.RegisterInlines(models.Category{}, map[string]string{
		"Todo": "CategoryID",
	})
	uadmin.RegisterInlines(models.Friend{}, map[string]string{
		"Todo": "FriendID",
	})
	uadmin.RegisterInlines(models.Item{}, map[string]string{
		"Todo": "ItemID",
	})

	// Call InitializeRootURL function to change the RootURL value in the Settings model.
	InitializeRootURL()

	// Call InitializeSiteName function to assign the SiteName value in the Settings model.
	InitializeSiteName()

	// API Handler
	http.HandleFunc("/api/", uadmin.Handler(api.Handler))

	// Page Handler
	http.HandleFunc("/page/", uadmin.Handler(views.PageHandler))

	uadmin.StartServer()
}

func InitializeRootURL() {
	// Initialize Setting model
	setting := uadmin.Setting{}

	// Get the code
	uadmin.Get(&setting, "code = ?", "uAdmin.RootURL")

	// Assign the value as "/admin/"
	setting.ParseFormValue([]string{"/admin/"})

	// Save changes
	setting.Save()
}

func InitializeSiteName() {
	// Assign Site Name in the Settings
	setting := uadmin.Setting{}
	uadmin.Get(&setting, "code = ?", "uAdmin.SiteName")
	setting.ParseFormValue([]string{"Todo List"})
	setting.Save()
}
