package entity

type Question struct {
	Id   int64    `json:"id"`
	Text string `json:"text"`
	Type string `json:"type"`
	PossibleAnswers []PossibleAnswer `json:"possibleAnswers"`
}

type PossibleAnswer struct{
	Ans string `json:"ans"`
	Value float64 `json:"value"`
}
