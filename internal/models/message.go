package models

type Chat struct {
	Id             int    `json:"id"`
	MyProfileId    int    `json:"my_profile_id"`
	User2ProfileId int    `json:"user2_profile_id"`
	LastMessage    string `json:"last_message"`
	IsRead         bool   `json:"is_read"`
	CountNewMsg    int    `json:"count_new_msg"`
	IsVisible      bool   `json:"is_visible"`
}

type Message struct {
	Id         int    `json:"id"`
	SenderId   int    `json:"sender_id"`
	ReceiverId int    `json:"receiver_id"`
	Text       string `json:"text"`
	Timestamp  string `json:"timestamp"`
}
