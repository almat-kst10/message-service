package models

type Chat struct {
	Id          int    `json:"id"`
	ProfilesId  int    `json:"profiles_id"`
	LastMessage string `json:"last_message"`
	IsRead      bool   `json:"is_read"`
	CountNewMsg int    `json:"count_new_msg"`
}

type Message struct {
	Id         int    `json:"id"`
	SenderId   int    `json:"sender_id"`
	ReceiverId int    `json:"receiver_id"`
	Text       string `json:"text"`
	Timestamp  string `json:"timestamp"`
}
