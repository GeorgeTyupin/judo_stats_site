-- +goose Up

ALTER TABLE public.judokas     ADD COLUMN external_id bigint UNIQUE;
ALTER TABLE public.cities      ADD COLUMN external_id bigint UNIQUE;
ALTER TABLE public.countries   ADD COLUMN external_id bigint UNIQUE;
ALTER TABLE public.sport_clubs ADD COLUMN external_id bigint UNIQUE;

-- +goose Down

ALTER TABLE public.judokas     DROP COLUMN external_id;
ALTER TABLE public.cities      DROP COLUMN external_id;
ALTER TABLE public.countries   DROP COLUMN external_id;
ALTER TABLE public.sport_clubs DROP COLUMN external_id;
