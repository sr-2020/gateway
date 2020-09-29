package billing

import "github.com/sr-2020/gateway/tests/domain"

type Service interface {
	Check() bool
	Balance() (domain.BalanceResponse, error)
}
