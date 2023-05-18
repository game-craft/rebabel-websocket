package usecase

import (
	"docker-echo-template/api/domain"
)

type ChatInteractor struct {
	ChatRepository ChatRepository
}

func (interactor *ChatInteractor) Add(u domain.Chat) (chat domain.Chat, err error) {
	chat, err = interactor.ChatRepository.Store(u)

	return
}