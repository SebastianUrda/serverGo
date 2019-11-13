package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handleUserLogIn(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:sebi@tcp(127.0.0.1:3306)/licenta?parseTime=true")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	defer db.Close()
	switch r.Method {

	case "POST":

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		var user User
		if err := json.Unmarshal(reqBody, &user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		results, err := db.Query("SELECT * FROM user WHERE username=? AND password=?", user.Username, GetMD5Hash(user.Password))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		for results.Next() {
			err = results.Scan(&user.Id, &user.Password, &user.Role, &user.Username, &user.Email,&user.Sex,&user.Birthday)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		res, err := db.Query("SELECT id FROM device WHERE owner_id=?", user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		for res.Next() {
			err = res.Scan(&user.DeviceId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		js, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

	default:
		fmt.Println(w, "Sorry, only GET and POST methods are supported.")
	}

}
func handleUserRegister(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:sebi@tcp(127.0.0.1:3306)/licenta?parseTime=true")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	defer db.Close()
	switch r.Method {

	case "POST":

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Println(reqBody)
		var user User
		if err := json.Unmarshal(reqBody, &user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Println(user)

		stmtInsAns, _ := db.Prepare("INSERT INTO user ( password, role, username, email, sex, birthday)" +
			" VALUES( ?,?,?,?,?,?)")

		_, err = stmtInsAns.Exec(GetMD5Hash(user.Password), "USER", user.Username, user.Email, user.Sex, user.Birthday)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		results, err := db.Query("SELECT * FROM user WHERE username=? AND password=?", user.Username, GetMD5Hash(user.Password))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		for results.Next() {
			err = results.Scan(&user.Id, &user.Password, &user.Role, &user.Username, &user.Email,&user.Sex,&user.Birthday)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
		js, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

	default:
		http.Error(w, "Not Supported", http.StatusInternalServerError)
	}

}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
