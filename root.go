package harbourcore

import (
	"crypto/rsa"
	"database/sql"
	"net/http"

	harbourauth "github.com/FelixWieland/harbour-auth"
	"github.com/corneldamian/httpway"
)

var signKey *rsa.PrivateKey
var server *httpway.Server
var db *sql.DB
var secret string

const (
	privKeyPath = "keys/app.rsa" //openssl genrsa -out app.rsa 1024
)

//Start starts the API Server
func Start() {

	signKey, _ = harbourauth.LoadAsPrivateRSAKey("")
	secret = "demoSecret"
	credentials := loadCredentials("../auth.json")

	if ldb, err := connectToDB(credentials.toString()); err == nil {
		db = ldb
		defer db.Close()
	} else {
		println("Cant connect to Database")
	}

	router := httpway.New()

	public := router.Middleware(accessLogger)
	private := public.Middleware(authCheck)
	allowed := private.Middleware(permissionCheck)

	/*PUBLIC ROUTES*/
	public.GET("/stats", showPublicStats)

	//PRIVATE ROUTES*/
	private.POST("/pvt", testJWTLogin)

	/*ALLOWED ROUTES*/
	allowed.POST("/alw", testAllowedRoute)

	http.ListenAndServe(":5000", router)

	server = httpway.NewServer(nil)
	server.Addr = ":5000"
	server.Handler = router

	server.Start()
}

func showPublicStats(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Nothing to see here"))
}
