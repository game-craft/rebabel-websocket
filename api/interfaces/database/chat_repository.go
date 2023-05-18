package database

import (
	"docker-echo-template/api/domain"
)

type ChatRepository struct {
	SqlHandler
}

func (repo *ChatRepository) Store(u domain.Chat) (chat domain.Chat, err error) {
	if err = repo.Create(&u).Error; err != nil {
		return
	}
	chat = u

	return
}