package main

import (
	"database/sql"
	"fmt"
	forecast "github.com/mlbright/darksky/v2"
	"net/http"
	"time"
)

func insertDarkSkyData(w http.ResponseWriter,lat float64, lng float64, timestamp time.Time)  {
	db, err := sql.Open("mysql", "root:sebi@tcp(127.0.0.1:3306)/licenta?parseTime=true")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer db.Close()
	latStr := fmt.Sprintf("%f", lat)
	lngStr := fmt.Sprintf("%f", lng)
	key:="key"
	f, err := forecast.Get(key, latStr, lngStr, "now", forecast.SI, forecast.Romanian)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	stmtInsData, err := db.Prepare("INSERT INTO api_observations(latitude, longitude, measurement_unit, timestamp, value, api_name, measuring) VALUES (?,?,?,?,?,?,?)")
	_, err = stmtInsData.Exec(lat, lng, "C", timestamp, f.Currently.Temperature, "DarkSkyApi", "Temperature")
	_, err = stmtInsData.Exec(lat, lng, "%", timestamp, f.Currently.Humidity, "DarkSkyApi", "Humidity")
	_, err = stmtInsData.Exec(lat, lng, "Hectopascals", timestamp, f.Currently.Pressure, "DarkSkyApi", "Pressure")
	_, err = stmtInsData.Exec(lat, lng, "", timestamp, f.Currently.UVIndex, "DarkSkyApi", "UV Index")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Printf("temperature: %.2f Celsius\n", f.Currently.Temperature)
}
