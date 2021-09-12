package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

// HomeHandler handles the home page.
func HomeHandler(w http.ResponseWriter, r *http.Request, session *uadmin.Session) {
	// Initialize the fields that we need in the custom struct.
	type Context struct {
		User        string
		OTPRequired bool
	}

	// Call the custom struct and assign the full name in the User field under the context object.
	c := Context{}
	c.User = session.User.FirstName + " " + session.User.LastName

	// Check if the user has OTPRequired enabled in the database.
	if session.User.OTPRequired {
		// Assign a boolean value to OTPRequired field. We will use this to manipulate the grammar in the UI.
		c.OTPRequired = true
	}

	// Render the home filepath and pass the context data object to the HTML file.
	uadmin.RenderHTML(w, r, "templates/home.html", c)
	return
}
