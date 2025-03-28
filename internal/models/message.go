package models

type Message struct {
	Id         int    `json:"id"`
	SenderId   int    `json:"sender_id"`
	ReceiverId int    `json:"receiver_id"`
	Text       string `json:"text"`
	Timestamp  string `json:"timestamp"`
}