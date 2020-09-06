package position

import "github.com/sr-2020/gateway/app/domain"

type ServiceInterface interface {
	Location(int) (domain.Location, error)
}
