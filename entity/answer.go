package entity

import "time"

type Answer struct {
	Id         int64      `json:"id"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Timestamp  time.Time `json:"date"`
	Value      int64       `json:"answer"`
	QuestionId int64       `json:"questionId"`
	UserId     int64       `json:"userId"`
}