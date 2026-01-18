package models

// JudokaData представляет данные дзюдоиста
type JudokaData struct {
	Name           string
	Country        string
	CountryFlag    string
	WeightCategory string
	BirthDate      string
	Age            int
	BirthPlace     string
	SportClub      string
	Coach          string
	GoldMedals     int
	SilverMedals   int
	BronzeMedals   int
}
