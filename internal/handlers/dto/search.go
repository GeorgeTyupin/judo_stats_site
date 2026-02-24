package dto

// SearchCategory представляет категорию поиска
type SearchCategory string

const (
	CategoryAll        SearchCategory = "all" // Общий поиск по всем категориям
	CategoryJudoka     SearchCategory = "judoka"
	CategoryTournament SearchCategory = "tournament"
	CategorySportClub  SearchCategory = "sportclub"
	CategoryCity       SearchCategory = "city"
)

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
	Type     string // National Championships, National Open WC Championships, National Tournament
	Gender   string // Men, Women
	Year     int16  // 1970-1992
	City     string // Город проведения
	Country  string // Страна проведения (countries.name)
	Republic string // Республика СССР (ussr_republic.name)
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
	Republic      string // Республика СССР (ussr_republic.name)
	MinPopulation int    // Минимальное население
	MinSportClubs int    // Минимальное количество спортклубов
	MinAthletes   int    // Минимальное количество спортсменов
}

type JudokaResponse struct {
	ID               int64    `json:"id"`
	Name             string   `json:"name"`
	Country          string   `json:"country,omitempty"`
	BirthPlace       string   `json:"birth_place,omitempty"`
	WeightCategories []string `json:"weight_categories,omitempty"`
	BirthDate        string   `json:"birth_date,omitempty"`
}

type TournamentResponse struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type,omitempty"`
	Place  string `json:"place,omitempty"`
	Date   string `json:"date,omitempty"`
	Gender string `json:"gender,omitempty"`
}

type SportClubResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	FullName  string `json:"full_name,omitempty"`
	Region    string `json:"region,omitempty"`
	Founded   string `json:"founded,omitempty"`
	HeadCoach string `json:"head_coach,omitempty"`
}

type CityResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}
