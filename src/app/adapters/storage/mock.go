package storage

type Mock struct {
	data map[string]int64
}

func NewMock() *Mock {
	data := make(map[string]int64)
	return &Mock{
		data: data,
	}
}

func (m Mock) Check(key string, exp int64) bool {
	if v, ok := m.data[key]; ok {
		if v <= exp {
			m.data[key] = exp
			return true
		}

		return false
	}

	m.data[key] = exp
	return true
}
