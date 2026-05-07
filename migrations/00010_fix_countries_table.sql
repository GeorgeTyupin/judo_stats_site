-- +goose Up
-- +goose StatementBegin

CREATE SEQUENCE IF NOT EXISTS public.countries_id_seq
    AS bigint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.countries_id_seq OWNED BY public.countries.id;
ALTER TABLE public.countries ALTER COLUMN id SET DEFAULT nextval('public.countries_id_seq'::regclass);
SELECT setval('public.countries_id_seq', COALESCE((SELECT MAX(id) FROM public.countries), 0) + 1, false);

ALTER TABLE public.countries ALTER COLUMN name TYPE character varying(255);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE public.countries ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE IF EXISTS public.countries_id_seq;

ALTER TABLE public.countries ALTER COLUMN name TYPE character varying(32);

-- +goose StatementEnd
