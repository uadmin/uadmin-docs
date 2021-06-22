<<<<<<< HEAD
package views

import (
	"net/http"
	"strings"
)

// HTTPHandler !
func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called /http_handler
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/http_handler")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

	if strings.HasPrefix(r.URL.Path, "/todo") {
		TodoHandler(w, r)
		return
	}
}
=======
package views

import (
	"net/http"
	"strings"
)

// HTTPHandler !
func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called /http_handler
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/http_handler")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

	if strings.HasPrefix(r.URL.Path, "/todo") {
		TodoHandler(w, r)
		return
	}
}
>>>>>>> de25cdd8a29ca2bb2c2df08be00b703b967aaed5
