package position

import "github.com/sr-2020/gateway/tests/domain"

type Service interface {
	Check() bool
	Locations() ([]domain.Location, error)
	AddPosition(beacons []domain.Beacon) (domain.Position, error)
	ManaLevel() (domain.ManaLevel, error)
}
