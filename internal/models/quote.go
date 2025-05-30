package models

type Quote struct {
	Author string `json:"author"`
	Text string   `json:"quote"`
	ID   int64    `json:"id"`
}