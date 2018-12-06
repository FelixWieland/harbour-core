package harbourcore

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Start starts the API Server
func Start() {
	router := httprouter.New()

	/*ROUTES*/
	router.GET("/stats", showPublicStats)
	router.POST("/login", login)
	router.POST("/pvt", checkAuth(login, isLoggedin))

	http.ListenAndServe(":5000", router)
}

func showPublicStats(w http.ResponseWriter, r *http.Request, rt httprouter.Params) {
	w.Write([]byte("Nothing to see here"))
}
