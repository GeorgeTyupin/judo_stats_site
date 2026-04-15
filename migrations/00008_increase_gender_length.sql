-- +goose Up
ALTER TABLE public.tournaments
    ALTER COLUMN gender TYPE character varying(50);

-- +goose Down
ALTER TABLE public.tournaments
    ALTER COLUMN gender TYPE character varying(10);
