package harbourcore

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

//response for an unauthenticated user
type errorResponse struct {
	Type        string `json:"type"`
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

type infoResponse struct {
	Type        string    `json:"type"`
	UserID      uuid.UUID `json:"userid"`
	Code        int       `json:"code"`
	Message     string    `json:"message"`
	Description string    `json:"description"`
}

//stdResponseForForbidden
func newErrResponseForbidden() errorResponse {
	return errorResponse{
		"error",
		1,
		"noValidSession",
		"You currently dont have a valid session, please log in to create a session",
	}
}

//stdResponseForForbidden
func newErrResponseParameterNotSatisfied() errorResponse {
	return errorResponse{
		"error",
		2,
		"ParameterNotSatisfied",
		"The required parameters were not satisfied",
	}
}

//stdResponsettlExpired
func newErrResponsettlExpired() errorResponse {
	return errorResponse{
		"error",
		3,
		"ttlExpired",
		"Youre Session time is expired, please log in again",
	}
}

//stdResponseAlreadyLoggedin
func newInfoResponseAlreadyLoggedin(pUserID uuid.UUID) infoResponse {
	return infoResponse{
		"info",
		pUserID,
		1,
		"AlreadyLoggedIn",
		"You are alredy logged in",
	}
}

//API Error Response
func apiError(w http.ResponseWriter, r *http.Request, pl errorResponse) {

	js, err := json.Marshal(pl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

//API Info Response
func apiInfo(w http.ResponseWriter, r *http.Request, pl infoResponse) {

	js, err := json.Marshal(pl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}
