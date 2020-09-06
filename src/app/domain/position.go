package domain

type PositionLocation struct {
	Id      int                    `json:"id"`
	Label   string                 `json:"label"`
	Options map[string]interface{} `json:"options"`
}

type PositionUser struct {
	Id         int               `json:"id"`
	LocationId *int              `json:"location_id"`
	CreatedAt  string            `json:"created_at"`
	UpdatedAt  string            `json:"updated_at"`
	Location   *PositionLocation `json:"location"`
}
