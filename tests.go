package harbourcore

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func testJWTLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func testAllowedRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
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
