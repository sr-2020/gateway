package domain

type BalanceResponse struct {
	Data    Balance
	Status  bool
	Message *string
}

type Balance struct {
	ModelId           int
	SIN               string
	CurrentBalance    float64
	PersonName        string
	CurrentScoring    float64
	LifeStyle         string
	ForecastLifeStyle string
}
