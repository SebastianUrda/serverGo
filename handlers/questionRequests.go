package main

import (
	"awesomeProject/repository"
	"encoding/json"
	"log"
	"net/http"
)

var (
	questionRepository repository.QuestionRepository = repository.NewQuestionRepository()
)

func handleQuestionRequest(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		var questionType = r.URL.Query().Get("type")

		err, questions := questionRepository.FindAllInMysql(questionType)
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


