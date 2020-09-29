package domain

type Position struct {
	Id         int `json:"id"`
	UserId     int `json:"user_id"`
	LocationId int `json:"location_id"`
}
