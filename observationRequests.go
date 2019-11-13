package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)


func handleObservationsRequest(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:sebi@tcp(127.0.0.1:3306)/licenta?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	switch r.Method {
	case "GET":

		results, err := db.Query("SELECT * FROM sensor_observation")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		var observations []Observation
		for results.Next() {
			var observation Observation
			err = results.Scan(&observation.Id, &observation.Latitude, &observation.Longitude, &observation.MeasurementUnit, &observation.Timestamp, &observation.Value, &observation.SensorId, &observation.Measuring)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			observations = append(observations, observation)
		}

		js, err := json.Marshal(observations)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	default:
		fmt.Println(w, "Sorry, only GET and POST methods are supported.")
	}

}



