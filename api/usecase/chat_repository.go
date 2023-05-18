package usecase

import (
	"docker-echo-template/api/domain"
)

type ChatRepository interface {
	Store(domain.Chat) (domain.Chat, error)
}
