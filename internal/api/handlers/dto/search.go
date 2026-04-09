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

// SortParams содержит параметры сортировки
type SortParams struct {
	SortBy  string // Поле для сортировки
	SortDir string // Направление: asc или desc
}

// JudokaFilters представляет фильтры для поиска дзюдоистов
type JudokaFilters struct {
	WeightCategories []string // -48, -52, -57, -63, -70, -78, +78 (множественный выбор)
	Country          string   // Страна
	Gender           string   // Мужской/Женский
	BirthYear        string   // Год рождения
	SportClub        string   // Название спортивного общества
	City             string   // Город
	SortParams
}

// TournamentFilters представляет фильтры для поиска турниров
type TournamentFilters struct {
	Type     string // National Championships, National Open WC Championships, National Tournament
	Gender   string // Men, Women
	Year     int    // Год проведения (точное значение)
	Month    int    // Месяц проведения (1-12)
	City     string // Город проведения
	Republic string // Республика СССР (ussr_republic.name)
	Country  string // Страна проведения (countries.name)
	SortParams
}

// SportClubFilters представляет фильтры для поиска спортивных обществ
type SportClubFilters struct {
	City     string // Город
	Republic string // Республика СССР
	SortParams
}

// CityFilters представляет фильтры для поиска городов
type CityFilters struct {
	Republic string // Республика СССР (ussr_republic.name)
	Oblast   string // Область
	SortParams
}

type JudokaResponse struct {
	ID               int64    `json:"id"`
	Name             string   `json:"name"`
	Gender           string   `json:"gender,omitempty"`
	Country          string   `json:"country,omitempty"`
	BirthPlace       string   `json:"birth_place,omitempty"`
	WeightCategories []string `json:"weight_categories,omitempty"`
	BirthDate        string   `json:"birth_date,omitempty"`
}

type TournamentResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type,omitempty"`
	Place    string `json:"place,omitempty"`
	Date     string `json:"date,omitempty"`
	Gender   string `json:"gender,omitempty"`
}

type SportClubResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name,omitempty"`
	Republic string `json:"republic,omitempty"`
	City     string `json:"city,omitempty"`
}

type CityResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}
