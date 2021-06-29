package storage

import "time"

type Mock struct {
	data map[string][]string
}

func NewMock(data map[string][]string) *Mock {
	return &Mock{
		data: data,
	}
}

func (m Mock) Check(key string, token string) bool {
	if v, ok := m.data[key]; ok {
		length := len(v)
		for i, v := range v {
			if v == token {
				return i == length - 1
			}
		}
	}

	return true
}

func (m Mock) ReadCache(key string) (string, error) {

	return key, nil
}

func (m Mock) WriteCache(key string, val interface{}, exp time.Duration) error {

	return nil
}