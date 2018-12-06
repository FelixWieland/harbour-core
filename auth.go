package harbourcore

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Login function
func Login(w http.ResponseWriter, r *http.Request, rt httprouter.Params) {
	username := rt.ByName("username")
	password := rt.ByName("password")
	_ = username
	_ = password
}
