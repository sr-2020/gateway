package domain

type Token struct {
	Id     int    `json:"id"`
	ApiKey string `json:"api_key"`
}

type AuthUser struct {
	Id                int         `json:"id"`
	Amount            int         `json:"amount"`
	Followers         []int       `json:"followers"`
	Following         []int       `json:"following"`
	Admin             bool        `json:"admin"`
	Status            string      `json:"status"`
	Role              string      `json:"role"`
	Items             interface{} `json:"items"`
	Name              string      `json:"name"`
	CreatedAt         string      `json:"created_at"`
	UpdatedAt         string      `json:"updated_at"`
	LocationUpdatedAt string      `json:"location_updated_at"`
	LocationId        *int        `json:"location_id"`
	Location          *Location   `json:"location"`
}
