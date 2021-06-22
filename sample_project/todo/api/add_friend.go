<<<<<<< HEAD
package api

import (
	"net/http"

	// Specify the username that you used inside github.com folder
	"github.com/uadmin/uadmin"
	"github.com/uadmin/uadmin-docs/sample_project/todo/models"
)

// AddFriendAPIHandler !
func AddFriendAPIHandler(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{}

	// Fetch data from Friend DB
	friend := models.Friend{}

	// Set the parameters of Name, Email, and Password
	friendName := r.FormValue("name")
	friendEmail := r.FormValue("email")
	friendPassword := r.FormValue("password")

	// Validate if the friendName variable is empty.
	if friendName == "" {
		res["status"] = "ERROR"
		res["err_msg"] = "Name is required."
		uadmin.ReturnJSON(w, r, res)
		return
	}

	// Store input into the Name, Email, and Password fields
	friend.Name = friendName
	friend.Email = friendEmail
	friend.Password = friendPassword

	// Store input in the Friend model
	uadmin.Save(&friend)

	res["status"] = "ok"
	uadmin.ReturnJSON(w, r, res)
}
=======
package api

import (
	"net/http"

	// Specify the username that you used inside github.com folder
	"github.com/uadmin/uadmin"
	"github.com/uadmin/uadmin-docs/sample_project/todo/models"
)

// AddFriendAPIHandler !
func AddFriendAPIHandler(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{}

	// Fetch data from Friend DB
	friend := models.Friend{}

	// Set the parameters of Name, Email, and Password
	friendName := r.FormValue("name")
	friendEmail := r.FormValue("email")
	friendPassword := r.FormValue("password")

	// Validate if the friendName variable is empty.
	if friendName == "" {
		res["status"] = "ERROR"
		res["err_msg"] = "Name is required."
		uadmin.ReturnJSON(w, r, res)
		return
	}

	// Store input into the Name, Email, and Password fields
	friend.Name = friendName
	friend.Email = friendEmail
	friend.Password = friendPassword

	// Store input in the Friend model
	uadmin.Save(&friend)

	res["status"] = "ok"
	uadmin.ReturnJSON(w, r, res)
}
>>>>>>> de25cdd8a29ca2bb2c2df08be00b703b967aaed5
