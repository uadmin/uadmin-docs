package api

import (
	"net/http"
	"strconv"

	"github.com/uadmin/uadmin"
	"github.com/uadmin/uadmin-docs/sample_project/todo/models"
)

// AddFriendAPIHandler !
func AddFriendAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch data from Friend DB
	friend := models.Friend{}

	// Set the parameters of Name, Email, Password, and Nationality such that where nationality is equivalent
	// to the following:
	// 1 - Chinese
	// 2 - Filipino
	// 3 - Others
	friendName := r.FormValue("name")
	friendEmail := r.FormValue("email")
	friendPassword := r.FormValue("password")
	friendNationalityRaw := r.FormValue("nationality")

	// Convert the nationality to an integer.
	friendNationality, err := strconv.Atoi(friendNationalityRaw)

	// Validate if the friendName variable is empty.
	if friendName == "" {
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Name is required.",
		})
		return
	}

	// Store input into the Name, Email, and Password fields
	friend.Name = friendName
	friend.Email = friendEmail
	friend.Password = friendPassword
	friend.Nationality = models.Nationality(friendNationality)

	// Store input in the Friend model
	err = uadmin.Save(&friend)
	if err != nil {
		// Return an error message if the database did not save properly.
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Error saving the database : " + err.Error(),
		})
		return
	}

	// Return JSON to the client.
	uadmin.ReturnJSON(w, r, map[string]interface{}{
		"status": "ok",
	})
}
