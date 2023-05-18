package domain

import (
    "time"
)

type Chats []Chat

type Chat struct {
	ID int `json:"id"`
	WorldsId int `json:"worlds_id"`
	ChatsContent string `json:"chats_content"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}