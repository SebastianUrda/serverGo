package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var baseURL = "https://maps.googleapis.com/maps/api/geocode/json?latlng="
var apiKey = "AIzaSyA25Q1CdhUgKbOCaL1flmtg05EMD6vZ1YY"
var USER_AGENT = "Mozilla/5.0"

func handleAlertRequest(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:sebi@tcp(127.0.0.1:3306)/licenta?parseTime=true")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()
	switch r.Method {
	case "GET":

		results, err := db.Query("SELECT * FROM alert")
		if err != nil {
			log.Fatal(err.Error()) // proper error handling instead of panic in your app
		}

		var alerts []Alert
		for results.Next() {
			var alert Alert
			err = results.Scan(&alert.Id, &alert.Address, &alert.Description, &alert.Type, &alert.Timestamp, &alert.Latitude, &alert.Longitude)
			if err != nil {
				log.Fatal(err.Error()) // proper error handling instead of panic in your app
			}
			alerts = append(alerts, alert)
		}

		js, err := json.Marshal(alerts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	default:
		fmt.Println(w, "Sorry, only GET and POST methods are supported.")
	}

}
//func getAddress(lat float64, lng float64) string {
//	latStr := fmt.Sprintf("%f", lat)
//	lngStr := fmt.Sprintf("%f", lng)
//	finalURL := baseURL + latStr + "," + lngStr + "&key=" + apiKey
//	resp, err := http.Get(finalURL)
//	if err != nil {
//		log.Fatal(err.Error()) // proper error handling instead of panic in your app
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	//var response map[string][string]
//	//json.Unmarshal([]byte(body), &response)
//	//fmt.Println("results :", response["results"])
//
//	//for key, r := range result["results"] {
//	//	fmt.Println(key,r)
//	//}
//	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
//	if err != nil {
//		log.Fatalf("fatal error: %s", err)
//	}
//	r := &maps.GeocodingResult{}
//	route, _, err := c.Directions(context.Background(), r)
//	if err != nil {
//		log.Fatalf("fatal error: %s", err)
//	}
//	fmt.Println(route)
//	return string(body)
//}
