package entity

import "time"

type Observation struct {
	Id              int       `json:"id"`
	Latitude        float64   `json:"latitude"`
	Longitude       float64   `json:"longitude"`
	MeasurementUnit string    `json:"measurement_unit"`
	Timestamp       time.Time `json:"timestamp"`
	Value           float64   `json:"value"`
	SensorId        int       `json:"sensor_id"`
	Measuring       string    `json:"measuring"`
}
