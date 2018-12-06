package harbourcore

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func checkAuth(next httprouter.Handle, toCheck func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (auth, error)) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Printf("Incoming connection from %v", r.RemoteAddr)

		w.Header().Set("Content-Type", "application/json")
		_, err := toCheck(w, r, ps)
		if err != nil {
			//notloggedin
			forbidden(w, r, ps)
		} else {
			next(w, r, ps)
		}
	}
}
