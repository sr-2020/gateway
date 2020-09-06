package position

import (
	"github.com/sr-2020/gateway/app/domain"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) Location(id int) (domain.Location, error) {
	args := m.Called(id)

	return args.Get(0).(domain.Location), args.Error(1)
}
