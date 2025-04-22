package models

import "time"

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

// --- new
type RoomGeneralInfo struct {
	RoomId         int    `json:"room_id"`
	RoomTitle      string `json:"room_title"`
	ClientId       int    `json:"client_id"`
	ProfileId      int    `json:"profile_id"`
	ProfileName    string `json:"profile_name"`
	ProfileSurname string `json:"profile_surname"`
	RoleId         int    `json:"role_id"`
	RoleName       string `json:"role_name"`
	IsMuted        bool   `json:"is_muted"`
	IsTyping       bool   `json:"is_typing"`
}

type Room struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type RoomRole struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type RoomClient struct {
	Id        int  `json:"id"`
	RoomId    int  `json:"room_id"`
	ProfileId int  `json:"profile_id"`
	RoleId    int  `json:"role_id"`
	IsMuted   bool `json:"is_muted"`
	IsTyping  bool `json:"is_typing"`
}

type MessageClientRoom struct {
	Id        int       `json:"id"`
	RoomId    int       `json:"room_id"`
	ProfileId int       `json:"profile_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type MessageRead struct {
	Id        int       `json:"id"`
	MessageId int       `json:"message_id"`
	ProfileId int       `json:"profile_id"`
	CreatedAt time.Time `json:"created_at"`
}
