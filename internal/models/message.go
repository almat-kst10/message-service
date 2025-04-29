package models

import "time"

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

type ClientRoom struct {
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
	FullName  string    `json:"full_name"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type MessageRead struct {
	Id        int       `json:"id"`
	MessageId int       `json:"message_id"`
	ProfileId int       `json:"profile_id"`
	CreatedAt time.Time `json:"created_at"`
}
