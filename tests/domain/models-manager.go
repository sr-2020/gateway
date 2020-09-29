package domain

type CharacterModelResponse struct {
	BaseModel BaseModel
}

type BaseModel struct {
	ModelId string
}

type Event struct {
	EventType string `json:"eventType"`
	Data map[string]interface{} `json:"data"`
}
