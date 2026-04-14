-- +goose Up
SELECT 'up SQL query';

ALTER TABLE public.judokas
    ADD COLUMN gender character varying(10);

-- +goose Down
SELECT 'down SQL query';

ALTER TABLE public.judokas
    DROP COLUMN IF EXISTS gender;

