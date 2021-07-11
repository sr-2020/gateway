package domain

type Location struct {
	Id        int   `json:"id"`
	ManaLevel int   `json:"manaLevel"`
	Label     string `json:"label"`
}

type LocationWithoutLabel struct {
	Id        int   `json:"id"`
	ManaLevel int   `json:"manaLevel"`
}

func (l LocationWithoutLabel) Apply(location Location) LocationWithoutLabel {
	l.Id = location.Id
	l.ManaLevel = location.ManaLevel

	return l
}
