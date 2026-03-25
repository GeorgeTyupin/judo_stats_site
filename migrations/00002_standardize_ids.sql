-- +goose Up
-- +goose StatementBegin
-- Standardize all primary keys and foreign keys to bigint

-- Primary keys
ALTER TABLE public.cities ALTER COLUMN id TYPE bigint;
ALTER SEQUENCE public.cities_id_seq AS bigint;

ALTER TABLE public.countries ALTER COLUMN id TYPE bigint;

ALTER TABLE public.judoka_cities ALTER COLUMN id TYPE bigint;
ALTER SEQUENCE public.judoka_cities_id_seq AS bigint;

ALTER TABLE public.judoka_countries ALTER COLUMN id TYPE bigint;

ALTER TABLE public.judoka_republics ALTER COLUMN id TYPE bigint;

ALTER TABLE public.judoka_sport_clubs ALTER COLUMN id TYPE bigint;
ALTER SEQUENCE public.judoka_sport_clubs_id_seq AS bigint;

ALTER TABLE public.sport_club_cities ALTER COLUMN id TYPE bigint;
ALTER SEQUENCE public.sport_club_cities_id_seq AS bigint;

ALTER TABLE public.sport_clubs ALTER COLUMN id TYPE bigint;
ALTER SEQUENCE public.sport_clubs_id_seq AS bigint;

ALTER TABLE public.tournament_results ALTER COLUMN id TYPE bigint;
ALTER SEQUENCE public.tournament_results_id_seq AS bigint;

ALTER TABLE public.tournaments ALTER COLUMN id TYPE bigint;
ALTER SEQUENCE public.tournaments_id_seq AS bigint;

ALTER TABLE public.ussr_republics ALTER COLUMN id TYPE bigint;

-- Foreign keys (integer → bigint)
ALTER TABLE public.judoka_cities ALTER COLUMN judoka_id TYPE bigint;
ALTER TABLE public.judoka_cities ALTER COLUMN city_id TYPE bigint;

ALTER TABLE public.judoka_sport_clubs ALTER COLUMN judoka_id TYPE bigint;
ALTER TABLE public.judoka_sport_clubs ALTER COLUMN sport_club_id TYPE bigint;

ALTER TABLE public.sport_club_cities ALTER COLUMN sport_club_id TYPE bigint;
ALTER TABLE public.sport_club_cities ALTER COLUMN city_id TYPE bigint;

ALTER TABLE public.sport_clubs ALTER COLUMN city_id TYPE bigint;

ALTER TABLE public.tournament_results ALTER COLUMN tournament_id TYPE bigint;
ALTER TABLE public.tournament_results ALTER COLUMN judoka_id TYPE bigint;

ALTER TABLE public.tournaments ALTER COLUMN city_id TYPE bigint;
ALTER TABLE public.tournaments ALTER COLUMN country_id TYPE bigint;
ALTER TABLE public.tournaments ALTER COLUMN republic_id TYPE bigint;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.tournaments ALTER COLUMN republic_id TYPE integer;
ALTER TABLE public.tournaments ALTER COLUMN country_id TYPE integer;
ALTER TABLE public.tournaments ALTER COLUMN city_id TYPE integer;
ALTER TABLE public.tournament_results ALTER COLUMN judoka_id TYPE integer;
ALTER TABLE public.tournament_results ALTER COLUMN tournament_id TYPE integer;
ALTER TABLE public.sport_clubs ALTER COLUMN city_id TYPE integer;
ALTER TABLE public.sport_club_cities ALTER COLUMN city_id TYPE integer;
ALTER TABLE public.sport_club_cities ALTER COLUMN sport_club_id TYPE integer;
ALTER TABLE public.judoka_sport_clubs ALTER COLUMN sport_club_id TYPE integer;
ALTER TABLE public.judoka_sport_clubs ALTER COLUMN judoka_id TYPE integer;
ALTER TABLE public.judoka_cities ALTER COLUMN city_id TYPE integer;
ALTER TABLE public.judoka_cities ALTER COLUMN judoka_id TYPE integer;
ALTER TABLE public.ussr_republics ALTER COLUMN id TYPE integer;
ALTER TABLE public.tournaments ALTER COLUMN id TYPE integer;
ALTER SEQUENCE public.tournaments_id_seq AS integer;
ALTER TABLE public.tournament_results ALTER COLUMN id TYPE integer;
ALTER SEQUENCE public.tournament_results_id_seq AS integer;
ALTER TABLE public.sport_clubs ALTER COLUMN id TYPE integer;
ALTER SEQUENCE public.sport_clubs_id_seq AS integer;
ALTER TABLE public.sport_club_cities ALTER COLUMN id TYPE integer;
ALTER SEQUENCE public.sport_club_cities_id_seq AS integer;
ALTER TABLE public.judoka_sport_clubs ALTER COLUMN id TYPE integer;
ALTER SEQUENCE public.judoka_sport_clubs_id_seq AS integer;
ALTER TABLE public.judoka_republics ALTER COLUMN id TYPE integer;
ALTER TABLE public.judoka_countries ALTER COLUMN id TYPE integer;
ALTER TABLE public.judoka_cities ALTER COLUMN id TYPE integer;
ALTER SEQUENCE public.judoka_cities_id_seq AS integer;
ALTER TABLE public.countries ALTER COLUMN id TYPE integer;
ALTER TABLE public.cities ALTER COLUMN id TYPE integer;
ALTER SEQUENCE public.cities_id_seq AS integer;
-- +goose StatementEnd
