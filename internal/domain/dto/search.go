package dto

// SearchCategory представляет категорию поиска
type SearchCategory string

const (
	CategoryJudoka     SearchCategory = "judoka"
	CategoryTournament SearchCategory = "tournament"
	CategorySportClub  SearchCategory = "sportclub"
	CategoryCity       SearchCategory = "city"
)

// SearchData представляет данные для страницы поиска
type SearchData struct {
	SelectedCategory SearchCategory
	Query            string
}

// JudokaFilters представляет фильтры для поиска спортсменов
type JudokaFilters struct {
	WeightCategory string // -48, -52, -57, -63, -70, -78, +78
	Country        string // Страна
	Gender         string // Мужской/Женский
	AgeGroup       string // Взрослые, Юниоры, Кадеты
	SportClub      string // Название спортклуба
	City           string // Город
}

// TournamentFilters представляет фильтры для поиска турниров
type TournamentFilters struct {
	Type     string // Первенство, Чемпионат, Кубок
	AgeGroup string // Взрослые, Юниоры, Кадеты
	Year     string // 2024, 2023, 2022...
	City     string // Город проведения
	Region   string // Регион
	DateFrom string // Дата от
	DateTo   string // Дата до
}

// SportClubFilters представляет фильтры для поиска спортклубов
type SportClubFilters struct {
	City     string // Город
	Region   string // Регион
	Founded  string // Год основания от
	HasCoach bool   // Есть главный тренер
}

// CityFilters представляет фильтры для поиска городов
type CityFilters struct {
	Region        string // Регион
	MinPopulation int    // Минимальное население
	MinSportClubs int    // Минимальное количество спортклубов
	MinAthletes   int    // Минимальное количество спортсменов
}
