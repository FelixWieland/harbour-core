package harbourcore

import (
	"database/sql"
	"net/http"

	"github.com/corneldamian/httpway"
)

var server *httpway.Server
var db *sql.DB

//Start starts the API Server
func Start() {

	credentials := loadCredentials("../auth.json")
	if ldb, err := connectToDB(credentials.toString()); err == nil {
		db = ldb
		defer db.Close()
	} else {
		println("Cant connect to Database")
	}

	testDBInsert()
	testDBSelect()

	router := httpway.New()

	public := router.Middleware(accessLogger)
	private := public.Middleware(authCheck)
	test := public

	/*PUBLIC ROUTES*/
	public.GET("/stats", showPublicStats)
	public.POST("/login", login)

	//PRIVATE ROUTES*/
	private.POST("/pvt", testDisplayDemoSession)

	/*TEST ROUTES*/
	test.POST("/createDemoSession", testCreateDemoSession) //cookies are bound to path
	test.POST("/test/displayDemoSession", testDisplayDemoSession)

	http.ListenAndServe(":5000", router)

	server = httpway.NewServer(nil)
	server.Addr = ":5000"
	server.Handler = router

	server.Start()
}

func showPublicStats(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Nothing to see here"))
}
