package models_manager

import "github.com/sr-2020/gateway/tests/domain"

type Service interface {
	Check() bool
	CharacterModel() (domain.CharacterModelResponse, error)
	SentEvent(event domain.Event) (domain.CharacterModelResponse, error)
}
