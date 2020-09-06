package domain

type Payload struct {
	Sub string
	Auth string
	ModelId int
}

func (p Payload) Valid() error {
	return nil
}
