package harbourcore

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type auth struct {
	uuid         uuid.UUID
	session      []byte
	latestAction time.Time
}

type cache struct {
	sessions []auth
}

var activeCache cache

func login(w http.ResponseWriter, r *http.Request) {

	//ctx := httpway.GetContext(r)

	username := r.FormValue("username") //ctx.ParamByName("username")
	password := r.FormValue("password") //ctx.ParamByName("password")

	log.Printf("%v", password)

	if !(len(username) > 0 && len(password) > 0) {
		//error -> parameter not satisfied
		apiError(w, r, newErrResponseParameterNotSatisfied())
		return
	}

	session, err := r.Cookie("session")
	if err == nil {
		//has session cookie
		authElm, err := activeCache.getAuthBySession(session.Value)
		if err == nil {
			//has valid session
			_, err := authElm.updatettl()
			if err != nil {
				apiErrorHandler(w, r, err)
				return
			}
			apiInfo(w, r, newInfoResponseAlreadyLoggedin(authElm.uuid))
		}

	}

	//login logic

	_ = username
	_ = password
}

func registrate(w http.ResponseWriter, r *http.Request) {

}

func isLoggedin(w http.ResponseWriter, r *http.Request) (*auth, error) {
	cookie, err := r.Cookie("session")
	if err == nil {
		authElm, err := activeCache.getAuthBySession(cookie.Value)
		if err == nil {
			//activeSession
			_, err := authElm.updatettl()
			if err != nil {
				//ttl expired
				authElm := &auth{}
				return authElm, err
			}
			//ttl updated
			return authElm, nil
		}
	}
	//notloggedin
	authElm := &auth{}
	log.Printf("IP %v don't has a valid Session", r.RemoteAddr)
	return authElm, errNoValidSession
}

func (cache *cache) getAuthBySession(pSessionKey string) (*auth, error) {
	for i := range cache.sessions {
		if pSessionKey == string(cache.sessions[i].session) {
			//loggedin
			//cache.sessions[i].latestAction = time.Now()
			return &cache.sessions[i], nil
		}
	}
	authElm := auth{}
	return &authElm, errNoValidSession
}

func (cache *cache) getAuthByUserID(pUserID uuid.UUID) (auth, error) {
	for i := range cache.sessions {
		if pUserID == cache.sessions[i].uuid {
			//loggedin
			//cache.sessions[i].latestAction = time.Now()
			return cache.sessions[i], nil
		}
	}
	authElm := auth{}
	return authElm, errNoValidSession
}

func (cache *cache) appendAuth(pAuth auth) {
	cache.sessions = append(cache.sessions, pAuth)
}

//updatettl returns an ttlExpired error if the ttl is expired
func (auth *auth) updatettl() (bool, error) {
	if !ttlexpired(auth.latestAction) {
		return true, nil
	}
	return false, errttlExpired
}

func (auth *auth) delete() {
	auth.session = []byte{}
	auth.uuid = uuid.UUID{}
	auth.latestAction = time.Time{}
}

func ttlexpired(ptime time.Time) bool {
	return time.Since(ptime).Minutes() > 9
}
