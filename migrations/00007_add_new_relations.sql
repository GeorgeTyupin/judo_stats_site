-- +goose Up

-- Добавление связей в tournament_results
ALTER TABLE public.tournament_results
ADD COLUMN city_id bigint REFERENCES public.cities(id) ON DELETE SET NULL,
ADD COLUMN country_id bigint REFERENCES public.countries(id) ON DELETE SET NULL,
ADD COLUMN sport_club_id bigint REFERENCES public.sport_clubs(id) ON DELETE SET NULL;

CREATE INDEX idx_tournament_results_city_id ON public.tournament_results(city_id);
CREATE INDEX idx_tournament_results_country_id ON public.tournament_results(country_id);
CREATE INDEX idx_tournament_results_sport_club_id ON public.tournament_results(sport_club_id);


-- Создание таблицы для связи стран и городов
CREATE TABLE public.countries_cities (
    country_id bigint REFERENCES public.countries(id) ON DELETE CASCADE,
    city_id bigint REFERENCES public.cities(id) ON DELETE CASCADE,
    PRIMARY KEY (country_id, city_id)
);

CREATE INDEX idx_countries_cities_country_id ON public.countries_cities(country_id);
CREATE INDEX idx_countries_cities_city_id ON public.countries_cities(city_id);

-- Создание таблицы для связи республик СССР и городов
CREATE TABLE public.ussr_republics_cities (
    ussr_republic_id bigint REFERENCES public.ussr_republics(id) ON DELETE CASCADE,
    city_id bigint REFERENCES public.cities(id) ON DELETE CASCADE,
    PRIMARY KEY (ussr_republic_id, city_id)
);

CREATE INDEX idx_ussr_republics_cities_ussr_republic_id ON public.ussr_republics_cities(ussr_republic_id);
CREATE INDEX idx_ussr_republics_cities_city_id ON public.ussr_republics_cities(city_id);


-- +goose Down

DROP TABLE IF EXISTS public.ussr_republics_cities;
DROP TABLE IF EXISTS public.countries_cities;

ALTER TABLE public.tournament_results
DROP COLUMN city_id,
DROP COLUMN country_id,
DROP COLUMN sport_club_id;