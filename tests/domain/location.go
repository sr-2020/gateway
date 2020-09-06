package domain

type Location struct {
	Id      int                    `json:"id"`
	Label   string                 `json:"label"`
	Options map[string]interface{} `json:"options"`
}
