-- +goose Up

ALTER TABLE public.cities
ADD COLUMN name_rus      character varying(255),
ADD COLUMN name_rus_last character varying(255);

ALTER TABLE public.countries
ADD COLUMN iso_code  character varying(10),
ADD COLUMN name_rus  character varying(255);

ALTER TABLE public.sport_clubs
ADD COLUMN name_rus character varying(255);

-- +goose Down

ALTER TABLE public.cities
DROP COLUMN name_rus,
DROP COLUMN name_rus_last;

ALTER TABLE public.countries
DROP COLUMN iso_code,
DROP COLUMN name_rus;

ALTER TABLE public.sport_clubs
DROP COLUMN name_rus;
