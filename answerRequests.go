package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handleAnswerRequest(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:sebi@tcp(127.0.0.1:3306)/licenta?parseTime=true")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	switch r.Method {
	case "POST":
		fmt.Println("au ajuns raspunsurile");
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		var answers []Answer
		fmt.Println(reqBody)
		if err := json.Unmarshal(reqBody, &answers); err != nil {
			fmt.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}

		stmtInsAns, _ := db.Prepare("INSERT INTO answer ( latitude, longitude, timestamp, value, question_id, user_id)" +
			" VALUES( ?,?,?,?,?,?)")

		for _, ans := range answers {
			fmt.Println(ans)
			_, err = stmtInsAns.Exec(ans.Latitude, ans.Longitude, ans.Timestamp, ans.Value, ans.QuestionId, ans.UserId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

		insertDarkSkyData(w,answers[0].Latitude,answers[0].Longitude,answers[0].Timestamp)
		insertAccurateWeatherData(w,answers[0].Latitude,answers[0].Longitude,answers[0].Timestamp)
		fmt.Println(answers)
	}
}
