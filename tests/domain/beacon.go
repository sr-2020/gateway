package domain

type Beacon struct {
	SSID  string `json:"ssid"`
	BSSID string `json:"bssid"`
	Level int    `json:"level"`
}

type Beacons struct {
	Beacons []Beacon `json:"beacons"`
}
