<<<<<<< HEAD
package api

import (
	"net/http"
	"strings"
)

// Handler !
func Handler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called /api
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

	if strings.HasPrefix(r.URL.Path, "/todo_list") {
		TodoListAPIHandler(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/custom_list") {
		CustomListAPIHandler(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/add_friend") {
		AddFriendAPIHandler(w, r)
		return
	}
}
=======
package api

import (
	"net/http"
	"strings"
)

// Handler !
func Handler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called /api
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

	if strings.HasPrefix(r.URL.Path, "/todo_list") {
		TodoListAPIHandler(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/custom_list") {
		CustomListAPIHandler(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/add_friend") {
		AddFriendAPIHandler(w, r)
		return
	}
}
>>>>>>> de25cdd8a29ca2bb2c2df08be00b703b967aaed5
