package harbourcore

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	/*ROUTES*/
	router.GET("/stats", showPublicStats)
	router.POST("/login", Login)

	http.ListenAndServe(":5000", router)
}

func showPublicStats(w http.ResponseWriter, r *http.Request, rt httprouter.Params) {

}
