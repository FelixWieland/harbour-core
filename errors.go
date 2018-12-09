package harbourcore

import (
	"errors"
	"net/http"
)

var errNoValidSession = errors.New("noValidSession")
var errttlExpired = errors.New("ttlExpired")
var errCookieNotFound = errors.New("cookieNotFound")
var errCantConnectToDB = errors.New("cantConnectToDB")

func apiErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	switch err {
	case errNoValidSession:
		apiError(w, r, newErrResponseForbidden())
		break
	case errttlExpired:
		apiError(w, r, newErrResponsettlExpired())
		break
	default:
		return
	}
}
