package harbourcore

import (
	"log"
	"net/http"

	"github.com/corneldamian/httpway"
)

/*
func checkAuth(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, rt httprouter.Params) {
		log.Printf("Incoming connection from %v", r.RemoteAddr)

		s, err1 := r.Cookie("session")
		if err1 == nil {
			log.Printf("%v", s.Value)
		}

		w.Header().Set("Content-Type", "application/json")
		_, err := isLoggedin(w, r, rt)
		if err != nil {
			//notloggedin
			apiErrorHandler(w, r, rt, err)
		} else {
			isLoggedin(w, r, rt)
		}
	}
}
*/

func accessLogger(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming connection from %v", r.RemoteAddr)
	httpway.GetContext(r).Next(w, r)
}

func authCheck(w http.ResponseWriter, r *http.Request) {

	ctx := httpway.GetContext(r)

	w.Header().Set("Content-Type", "application/json")
	_, err := isLoggedin(w, r)
	if err != nil {
		//notloggedin
		apiErrorHandler(w, r, err)
	} else {
		ctx.Next(w, r)
	}

}

/*
func julienHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// do stuff
	}
}
*/
