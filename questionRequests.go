package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func handleQuestionRequest(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:sebi@tcp(127.0.0.1:3306)/licenta?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	switch r.Method {
	case "GET":
		var questionType = r.URL.Query().Get("type")
		results, err := db.Query("SELECT * FROM question WHERE type=?", questionType)
		defer results.Close()
		var questions []Question
		if results != nil {
			for results.Next() {
				var question Question
				err = results.Scan(&question.Id, &question.Text, &question.Type)
				if err != nil {
					log.Fatal(err)
				}
				questions= append(questions,question)
			}

		}
		js, err := json.Marshal(questions)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}
