package models

// TournamentData представляет данные турнира
type TournamentData struct {
	Name         string
	Type         string
	Date         string
	Location     string
	City         string
	Participants int
	Categories   int
	Organizer    string
	Venue        string
	AgeGroup     string
	System       string
	Regions      int
	Matches      int
}
