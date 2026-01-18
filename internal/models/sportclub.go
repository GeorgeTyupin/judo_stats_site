package models

// SportClubData представляет данные спортивного общества
type SportClubData struct {
	Name           string
	FullName       string
	Founded        string
	City           string
	Region         string
	HeadCoach      string
	Athletes       int
	ActiveAthletes int
	Coaches        int
	GoldMedals     int
	SilverMedals   int
	BronzeMedals   int
	TotalMedals    int
	Description    string
}
