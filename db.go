package harbourcore

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql" //sql driver
)

type loginCredentials struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func loadCredentials(path string) loginCredentials {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully Opened " + path)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	data := loginCredentials{}
	json.Unmarshal(byteValue, &data)
	return data
}

func (loginCredentials *loginCredentials) toString() string {
	//username:password@tcp(127.0.0.1:3306)/test
	return loginCredentials.User + ":" + loginCredentials.Password + "@tcp(" + loginCredentials.Host + ":" + loginCredentials.Port + ")/" + loginCredentials.Database + "?charset=utf8mb4"
}

func connectToDB(connString string) (*sql.DB, error) {
	l, err := sql.Open("mysql", connString) //"astaxie:astaxie@/test?charset=utf8"
	if err == nil && nil == l.Ping() {
		return l, nil
	}
	return &sql.DB{}, errCantConnectToDB
}
