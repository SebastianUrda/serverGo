package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func handleAnswersToObservationRequest(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:sebi@tcp(127.0.0.1:3306)/licenta?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	switch r.Method {

	case "GET":

		var userId = r.URL.Query().Get("userId")
		var gDist = r.URL.Query().Get("geoDistance")
		var geoDistance, _ = strconv.ParseFloat(gDist, 64)
		var timeDist = r.URL.Query().Get("timeDistance")
		var timeDistance, _ = strconv.ParseFloat(timeDist, 64)

		//http://licenta.ddns.net:8080/mightWork/NewAnswersToObservations?geoDistance=30&timeDistance=100&userId=1
		fmt.Println("Get Received " + userId)
		results, err := db.Query("SELECT * FROM answer WHERE user_id=? ", userId)

		if err != nil {
			log.Fatal(err)
		}

		var answers []Answer
		for results.Next() {
			var answer Answer
			err = results.Scan(&answer.Id, &answer.Latitude, &answer.Longitude, &answer.Timestamp,
				&answer.Value, &answer.QuestionId, &answer.UserId)
			if err != nil {
				log.Fatal(err)
			}
			answers = append(answers, answer)
			var observations []Observation
			if strings.Compare(getQuestion(db, answer.QuestionId).Type, "temperature") == 0 {
				observations = append(observations, callFindAllInTimeMeasuring(db, answer.Latitude, answer.Longitude, answer.Timestamp, timeDistance, geoDistance, "FrontTemp")...)
				observations = append(observations, callFindAllInTimeMeasuring(db, answer.Latitude, answer.Longitude, answer.Timestamp, timeDistance, geoDistance, "BackTemp")...)
				observations = append(observations, callFindAllInTimeMeasuring(db, answer.Latitude, answer.Longitude, answer.Timestamp, timeDistance, geoDistance, "FrontTempDht")...)
			}
			if strings.Compare(getQuestion(db, answer.QuestionId).Type, "dust") == 0 {
				observations = append(observations, callFindAllInTimeMeasuring(db, answer.Latitude, answer.Longitude, answer.Timestamp, timeDistance, geoDistance, "DustDensity")...)
			}
			if strings.Compare(getQuestion(db, answer.QuestionId).Type, "light") == 0 {
				observations = append(observations, callFindAllInTimeMeasuring(db, answer.Latitude, answer.Longitude, answer.Timestamp, timeDistance, geoDistance, "Infrared")...)
				observations = append(observations, callFindAllInTimeMeasuring(db, answer.Latitude, answer.Longitude, answer.Timestamp, timeDistance, geoDistance, "Ultraviolet")...)
				observations = append(observations, callFindAllInTimeMeasuring(db, answer.Latitude, answer.Longitude, answer.Timestamp, timeDistance, geoDistance, "Visible")...)
			}
			if strings.Compare(getQuestion(db, answer.QuestionId).Type, "humidity") == 0 {
				observations = append(observations, callFindAllInTimeMeasuring(db, answer.Latitude, answer.Longitude, answer.Timestamp, timeDistance, geoDistance, "FrontHumidity")...)
				observations = append(observations, callFindAllInTimeMeasuring(db, answer.Latitude, answer.Longitude, answer.Timestamp, timeDistance, geoDistance, "BackHumidity")...)
			}

			fmt.Println("Type and text " + (getQuestion(db, answer.QuestionId).Type + " " + getQuestion(db, answer.QuestionId).Text))
			fmt.Print("Observations of Answer ")
			fmt.Print(answer.Id)
			fmt.Println(observations)

		}


		js, err := json.Marshal(answers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func callFindAllInTimeMeasuring(db *sql.DB, latitude float64, longitude float64, timestamp time.Time, timeDist float64, spaceDist float64, measuring string) []Observation {
	results, err := db.Query("CALL allObservationsInTimeAndSpaceMeasuring(?,?,?,?,?,?)", latitude, longitude, timestamp,
		timeDist, spaceDist, measuring)
	defer results.Close()
	var observations []Observation
	if results != nil {
		for results.Next() {
			var observation Observation
			err = results.Scan(&observation.Id, &observation.Latitude, &observation.Longitude, &observation.MeasurementUnit, &observation.Timestamp, &observation.Value, &observation.SensorId, &observation.Measuring)
			if err != nil {
				log.Fatal(err)
			}
			observations = append(observations, observation)
		}
	}

	return observations
}
func getQuestion(db *sql.DB, id int) Question {
	results, err := db.Query("SELECT * FROM question WHERE id=?", id)
	defer results.Close()
	var question Question
	if results != nil {
		for results.Next() {
			err = results.Scan(&question.Id, &question.Text, &question.Type)
			if err != nil {
				log.Fatal(err)
			}
		}

	}
	return question
}
