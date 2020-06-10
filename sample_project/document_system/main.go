package main

import (
	"github.com/uadmin/uadmin"
	"github.com/uadmin/uadmin-docs/sample_project/document_system/models"
)

func main() {
	// Register models to uAdmin
	uadmin.Register(
		models.Folder{},
		models.FolderGroup{},
		models.FolderUser{},
		models.Channel{},
		models.Document{},
		models.DocumentGroup{},
		models.DocumentUser{},
		models.DocumentVersion{},
	)

	// Register FolderGroup and FolderUser to Folder model
	uadmin.RegisterInlines(
		models.Folder{},
		map[string]string{
			"foldergroup": "FolderID",
			"folderuser":  "FolderID",
		},
	)

	// Register DocumentVersion, DocumentGroup, and DocumentUser to Document
	// model
	uadmin.RegisterInlines(
		models.Document{},
		map[string]string{
			"documentgroup":   "DocumentID",
			"documentuser":    "DocumentID",
			"documentversion": "DocumentID",
		},
	)

	// Initialize docS variable that calls the document model in the schema
	docS := uadmin.Schema["document"]

	// FormModifier makes CreatedBy read only if the user is not an admin
	// and the CreatedBy is not an empty string.
	docS.FormModifier = func(s *uadmin.ModelSchema, m interface{}, u *uadmin.User) {
		// Casts an interface to the Document model
		d, _ := m.(*models.Document)

		// Check whether the user is not an admin and the CreatedBy Field of
		// Document model is not an empty string
		if !u.Admin && d.CreatedBy != "" {
			// Set the CreatedBy Field to read only
			s.FieldByName("CreatedBy").ReadOnly = "true"
		}
	}

	// ListModifier is based on the user ID where the admin status is active
	// or not. If the user is not an admin, he has limited access to the
	// models and its records.
	docS.ListModifier = func(s *uadmin.ModelSchema, u *uadmin.User) (string, []interface{}) {
		// Checks whether the user is not an admin
		if !u.Admin {
			// Returns the user ID
			return "user_id = ?", []interface{}{u.ID}
		}
		// Returns nothing
		return "", []interface{}{}
	}

	// Pass back to the schema of document model
	uadmin.Schema["document"] = docS

	// Assign Site Name value as "Document System"
	// NOTE: This code works only on first build.
	uadmin.SiteName = "Document System"

	// Activates a uAdmin server
	uadmin.StartServer()
}
