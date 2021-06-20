package domain

const (
	RolePlayer = "ROLE_PLAYER"
	RoleMaster = "ROLE_MASTER"
)

type Payload struct {
	Auth    string
	ModelId int
	Exp     int64
}

func (p Payload) Valid() error {
	return nil
}
