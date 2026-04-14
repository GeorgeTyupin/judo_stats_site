-- +goose Up
SELECT 'up SQL query';

ALTER TABLE public.judokas
    DROP COLUMN IF EXISTS country;

-- +goose Down
SELECT 'down SQL query';

ALTER TABLE public.judokas
    ADD COLUMN country character varying(100);
