package harbourcore

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type sqlQuery string

func (sqlQuery sqlQuery) prep(vals ...string) string {
	if len(vals) > strings.Count(string(sqlQuery), "?") {
		panic("Too many arguments in prep call")
	}

	buffer := string(sqlQuery)
	for _, elm := range vals {
		buffer = strings.Replace(string(buffer), "?", "'"+elm+"'", 1)
	}
	return buffer
}

func testCreateDemoSession(w http.ResponseWriter, r *http.Request) {

	activeCache.appendAuth(auth{
		uuid.New(),
		[]byte("demoSession"),
		time.Now(),
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: "demoSession",
	})

	s, err1 := r.Cookie("session")
	s2, err2 := r.Cookie("session")
	if err1 == nil {
		log.Printf("%v", s.Value)
	}
	if err2 == nil {
		log.Printf("%v", s2.Value)
	}
}

func testDisplayDemoSession(w http.ResponseWriter, r *http.Request) {

	s, err1 := r.Cookie("session")
	if err1 == nil {
		log.Printf("%v", s.Value)
		w.Write([]byte(s.Value))
	}

	testForward(w, r)
}

func testForward(w http.ResponseWriter, r *http.Request) {
	s, err1 := r.Cookie("session")
	if err1 == nil {
		log.Printf("%v", s.Value)
		w.Write([]byte(s.Value))
	}
}

func testDBInsert() {
	log.Printf("%v", db.Ping())

	//insert, err := db.Query("INSERT INTO harbour_tests VALUES ( '" + time.Now().String() + "', 'TEST', 'TEST')")

	stmt, err := db.Prepare("INSERT INTO harbour_tests SET userid=?, settings=?, latest_update=?")
	_, err = stmt.Exec(time.Now(), "test", "2012-12-09")

	// if there is an error inserting, handle it
	if err != nil {
		log.Printf(err.Error())
	}
}

func testDBSelect() {
	log.Printf("%v", db.Ping())
	// query

	rows, err := db.Query(sqlQuery("SELECT * FROM harbour_tests WHERE settings=?").prep("TEST"))
	if err != nil {
		log.Printf("Error in Select")
	}

	for rows.Next() {
		var col1 string
		var col2 string
		var col3 string
		err = rows.Scan(&col1, &col2, &col3)

		fmt.Println(col1)
		fmt.Println(col2)
		fmt.Println(col3)
	}
}
