-- +goose Up
-- +goose StatementBegin

-- Переименовываем существующие поля (каждый RENAME — отдельный оператор)
ALTER TABLE public.judokas RENAME COLUMN name TO last_name;
ALTER TABLE public.judokas RENAME COLUMN name_rus TO last_name_rus;

-- Добавляем новые поля
ALTER TABLE public.judokas
    ADD COLUMN first_name     character varying(100),
    ADD COLUMN first_name_rus character varying(100);

-- Ограничения на непустые значения
ALTER TABLE public.judokas
    ALTER COLUMN last_name SET NOT NULL,
    ALTER COLUMN first_name SET NOT NULL;

-- Индексы для поиска по фамилии
CREATE INDEX idx_judokas_last_name     ON public.judokas USING btree (last_name);
CREATE INDEX idx_judokas_last_name_rus ON public.judokas USING btree (last_name_rus);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX IF EXISTS public.idx_judokas_last_name;
DROP INDEX IF EXISTS public.idx_judokas_last_name_rus;

ALTER TABLE public.judokas
    ALTER COLUMN last_name DROP NOT NULL,
    ALTER COLUMN first_name DROP NOT NULL;

ALTER TABLE public.judokas
    DROP COLUMN IF EXISTS first_name,
    DROP COLUMN IF EXISTS first_name_rus;

ALTER TABLE public.judokas RENAME COLUMN last_name TO name;
ALTER TABLE public.judokas RENAME COLUMN last_name_rus TO name_rus;

-- +goose StatementEnd
