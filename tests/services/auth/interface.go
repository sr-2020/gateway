package auth

import (
	"github.com/sr-2020/gateway/tests/domain"
)

type Service interface {
	Check() bool
	Auth(map[string]string) (domain.Token, int, error)
	ModelId(domain.Token) (int, error)
}
