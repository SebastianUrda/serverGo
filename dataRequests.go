package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func handleDataRequest(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:sebi@tcp(127.0.0.1:3306)/licenta?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	switch r.Method {

	case "GET":

		results, err := db.Query("SELECT * FROM data")
		if err != nil {
			log.Fatal(err)
		}

		var dates []Data
		for results.Next() {
			var data Data
			err = results.Scan(&data.Id, &data.BackTemp, &data.Co, &data.Co2,
				&data.Timestamp, &data.Dust, &data.FrontTemp, &data.Humidity,
				&data.Ir, &data.Latitude, &data.Longitude, &data.Lpg,
				&data.Pressure, &data.Smoke, &data.Uv, &data.Vis,
				&data.FrontHumidity, &data.FrontTempDht)
			if err != nil {
				panic(err.Error())
			}

			dates = append(dates, data)
		}

		js, err := json.Marshal(dates)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

	case "POST":

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		w.Write([]byte("Received a POST request\n"))
		var observation Data
		if err := json.Unmarshal(reqBody, &observation); err != nil {
			log.Fatal(err)
		}
		stmtIns, err := db.Prepare("INSERT INTO data (back_temp, co, co2, date, dust, front_temp, humidity, ir, latitude, longitude, lpg, pressure, smoke, uv, vis, frontHumidity, frontTempDht) " +
			"VALUES( ?,?,?,?,?,?,?,?, ?,?,?,?,?,?,?,?,? )")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(observation)
		_, err = stmtIns.Exec(observation.BackTemp, observation.Co, observation.Co2,
			observation.Timestamp, observation.Dust, observation.FrontTemp, observation.Humidity,
			observation.Ir, observation.Latitude, observation.Longitude, observation.Lpg,
			observation.Pressure, observation.Smoke, observation.Uv, observation.Vis,
			observation.FrontHumidity, observation.FrontTempDht)
		if err != nil {
			log.Fatal(err)
		}

		defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

		//fmt.Println(getSensorId(observation.DeviceId, "MQ2"))
		stmtInsObs, err := db.Prepare("INSERT INTO sensor_observation(latitude, longitude, measurement_unit, timestamp, value, sensor_id, measuring) VALUES (?,?,?,?,?,?,?)")

		if err != nil {
			log.Fatal(err) // proper error handling instead of panic in your app
		}

		t := getTimestamp(observation.Timestamp)
		fmt.Println(t)
		latitude := observation.Latitude
		longitude := observation.Longitude
		_, err = stmtInsObs.Exec(latitude, longitude, "PPM", t, observation.Lpg, getSensorId(observation.DeviceId, "MQ2"), "LPG")
		_, err = stmtInsObs.Exec(latitude, longitude, "PPM", t, observation.Co, getSensorId(observation.DeviceId, "MQ2"), "CO")
		_, err = stmtInsObs.Exec(latitude, longitude, "PPM", t, observation.Smoke, getSensorId(observation.DeviceId, "MQ2"), "SMOKE")
		_, err = stmtInsObs.Exec(latitude, longitude, "PPM", t, observation.Co2, getSensorId(observation.DeviceId, "MQ135"), "CO2")
		_, err = stmtInsObs.Exec(latitude, longitude, "C", t, observation.BackTemp, getSensorId(observation.DeviceId, "DHT11"), "BackTemp")
		_, err = stmtInsObs.Exec(latitude, longitude, "%", t, observation.Humidity, getSensorId(observation.DeviceId, "DHT11"), "BackHumidity")
		_, err = stmtInsObs.Exec(latitude, longitude, "Pa", t, observation.Pressure, getSensorId(observation.DeviceId, "MPL3115A2"), "Pressure")
		_, err = stmtInsObs.Exec(latitude, longitude, "C", t, observation.FrontTemp, getSensorId(observation.DeviceId, "MPL3115A2"), "FrontTemp")
		_, err = stmtInsObs.Exec(latitude, longitude, " ", t, observation.Ir, getSensorId(observation.DeviceId, "SI1145"), "Infrared")
		_, err = stmtInsObs.Exec(latitude, longitude, "UV Index", t, observation.Uv, getSensorId(observation.DeviceId, "SI1145"), "Ultraviolet")
		_, err = stmtInsObs.Exec(latitude, longitude, "Lux", t, observation.Vis, getSensorId(observation.DeviceId, "SI1145"), "Visible")
		_, err = stmtInsObs.Exec(latitude, longitude, "ug/m3", t, observation.Dust, getSensorId(observation.DeviceId, "GP2Y1014"), "DustDensity")
		_, err = stmtInsObs.Exec(latitude, longitude, "C", t, observation.FrontTempDht, getSensorId(observation.DeviceId, "DHT112"), "FrontTempDht")
		_, err = stmtInsObs.Exec(latitude, longitude, "%", t, observation.FrontHumidity, getSensorId(observation.DeviceId, "DHT112"), "FrontHumidity")
		if err != nil {
			log.Fatal(err)
		}

		defer stmtInsObs.Close() // Close the statement when we leave main() / the program terminates
		fmt.Println("Data Inserted!")
	default:
		fmt.Println(w, "Sorry, only GET and POST methods are supported.")
	}
}

func getSensorId(deviceId int, deviceName string) int {
	db, err := sql.Open("mysql", "root:sebi@tcp(127.0.0.1:3306)/licenta?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}

	results, err := db.Query("SELECT id FROM sensor WHERE device_id=? AND name=?", deviceId, deviceName)
	var sensorId int
	for results.Next() {
		err = results.Scan(&sensorId)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	if err != nil {
		log.Fatal(err)
		return 0
	}
	db.Close()
	return sensorId
}

func getTimestamp(timestamp string) time.Time {
	ts := strings.Split(timestamp, " ")

	dt := strings.Split(ts[0], "-")
	str := dt[2] + "-" + dt[1] + "-" + dt[0] + "T" + ts[1] + ".000Z"
	t, err := time.Parse("2006-01-02T15:04:05.000Z", str)
	if err != nil {
		log.Fatal(err)
	}
	return t
}
