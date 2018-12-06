package harbourcore

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/julienschmidt/httprouter"
)

type auth struct {
	uuid         uuid.UUID
	session      []byte
	latestAction time.Time
}

var activeSessions []auth

func login(w http.ResponseWriter, r *http.Request, rt httprouter.Params) {
	username := rt.ByName("username")
	password := rt.ByName("password")
	_ = username
	_ = password
}

func forbidden(w http.ResponseWriter, r *http.Request, rt httprouter.Params) {
	//http.Error(w, "You have no access to this location. Please log in to use this service", 403)

	js, err := json.Marshal(respForbidden{
		1,
		"NoValidSession",
		"You currently dont have a valid session, please log in to create a session",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func isLoggedin(w http.ResponseWriter, r *http.Request, rt httprouter.Params) (auth, error) {

	cookies := r.Cookies()

	for _, elm := range cookies {
		if elm.Name == "session" {
			for i, auth := range activeSessions {
				if elm.Value == base64.StdEncoding.EncodeToString(auth.session) {
					//loggedin
					if ttlexpired(auth.latestAction) {
						break
					}
					activeSessions[i].latestAction = time.Now()
					return auth, nil
				}
			}
		}
	}
	//notloggedin

	authElm := auth{}
	return authElm, errors.New("noValidSession")

}

func ttlexpired(ptime time.Time) bool {
	return time.Since(ptime).Minutes() > 9
}
