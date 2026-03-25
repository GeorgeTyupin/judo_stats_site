-- +goose Up
-- +goose StatementBegin

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;
SET default_tablespace = '';
SET default_table_access_method = heap;

CREATE TABLE public.cities (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    republic_id bigint
);

CREATE SEQUENCE public.cities_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.cities_id_seq OWNED BY public.cities.id;

CREATE TABLE public.countries (
    id integer NOT NULL,
    name character varying(32) NOT NULL,
    description text,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);

CREATE TABLE public.judoka_cities (
    id integer NOT NULL,
    judoka_id integer NOT NULL,
    city_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);

CREATE SEQUENCE public.judoka_cities_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.judoka_cities_id_seq OWNED BY public.judoka_cities.id;

CREATE TABLE public.judoka_countries (
    id integer NOT NULL,
    judoka_id bigint NOT NULL,
    country_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);

CREATE TABLE public.judoka_republics (
    id integer NOT NULL,
    judoka_id bigint NOT NULL,
    republic_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);

CREATE TABLE public.judoka_sport_clubs (
    id integer NOT NULL,
    judoka_id integer NOT NULL,
    sport_club_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);

CREATE SEQUENCE public.judoka_sport_clubs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.judoka_sport_clubs_id_seq OWNED BY public.judoka_sport_clubs.id;

CREATE TABLE public.judokas (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    name_rus character varying(255),
    country character varying(100),
    weight_category text[],
    birth_date character varying(50),
    birth_place character varying(255),
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);

CREATE SEQUENCE public.judokas_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.judokas_id_seq OWNED BY public.judokas.id;

CREATE TABLE public.sport_club_cities (
    id integer NOT NULL,
    sport_club_id integer NOT NULL,
    city_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);

CREATE SEQUENCE public.sport_club_cities_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.sport_club_cities_id_seq OWNED BY public.sport_club_cities.id;

CREATE TABLE public.sport_clubs (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    full_name character varying(500),
    founded character varying(50),
    city_id integer,
    region character varying(255),
    head_coach character varying(255),
    description text,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);

CREATE SEQUENCE public.sport_clubs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.sport_clubs_id_seq OWNED BY public.sport_clubs.id;

CREATE TABLE public.tournament_results (
    id integer NOT NULL,
    tournament_id integer NOT NULL,
    judoka_id integer NOT NULL,
    weight_category character varying(50) NOT NULL,
    rank integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    CONSTRAINT tournament_results_rank_check CHECK (((rank >= 1) AND (rank <= 3)))
);

CREATE SEQUENCE public.tournament_results_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.tournament_results_id_seq OWNED BY public.tournament_results.id;

CREATE TABLE public.tournaments (
    id integer NOT NULL,
    name character varying(500) NOT NULL,
    type character varying(100),
    place character varying(255),
    city_id integer,
    date character varying(100),
    year smallint,
    month smallint,
    gender character varying(10),
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    country_id integer,
    republic_id integer
);

CREATE SEQUENCE public.tournaments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.tournaments_id_seq OWNED BY public.tournaments.id;

CREATE TABLE public.ussr_republics (
    id integer NOT NULL,
    name character varying(32) NOT NULL,
    description text,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);

ALTER TABLE ONLY public.cities ALTER COLUMN id SET DEFAULT nextval('public.cities_id_seq'::regclass);
ALTER TABLE ONLY public.judoka_cities ALTER COLUMN id SET DEFAULT nextval('public.judoka_cities_id_seq'::regclass);
ALTER TABLE ONLY public.judoka_sport_clubs ALTER COLUMN id SET DEFAULT nextval('public.judoka_sport_clubs_id_seq'::regclass);
ALTER TABLE ONLY public.judokas ALTER COLUMN id SET DEFAULT nextval('public.judokas_id_seq'::regclass);
ALTER TABLE ONLY public.sport_club_cities ALTER COLUMN id SET DEFAULT nextval('public.sport_club_cities_id_seq'::regclass);
ALTER TABLE ONLY public.sport_clubs ALTER COLUMN id SET DEFAULT nextval('public.sport_clubs_id_seq'::regclass);
ALTER TABLE ONLY public.tournament_results ALTER COLUMN id SET DEFAULT nextval('public.tournament_results_id_seq'::regclass);
ALTER TABLE ONLY public.tournaments ALTER COLUMN id SET DEFAULT nextval('public.tournaments_id_seq'::regclass);

ALTER TABLE ONLY public.cities ADD CONSTRAINT cities_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.countries ADD CONSTRAINT country_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.judoka_cities ADD CONSTRAINT judoka_cities_judoka_id_city_id_key UNIQUE (judoka_id, city_id);
ALTER TABLE ONLY public.judoka_cities ADD CONSTRAINT judoka_cities_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.judoka_countries ADD CONSTRAINT judoka_countries_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.judoka_republics ADD CONSTRAINT judoka_republics_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.judoka_sport_clubs ADD CONSTRAINT judoka_sport_clubs_judoka_id_sport_club_id_key UNIQUE (judoka_id, sport_club_id);
ALTER TABLE ONLY public.judoka_sport_clubs ADD CONSTRAINT judoka_sport_clubs_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.judokas ADD CONSTRAINT judokas_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.sport_club_cities ADD CONSTRAINT sport_club_cities_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.sport_club_cities ADD CONSTRAINT sport_club_cities_sport_club_id_city_id_key UNIQUE (sport_club_id, city_id);
ALTER TABLE ONLY public.sport_clubs ADD CONSTRAINT sport_clubs_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.tournament_results ADD CONSTRAINT tournament_results_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.tournament_results ADD CONSTRAINT tournament_results_tournament_id_judoka_id_weight_category_key UNIQUE (tournament_id, judoka_id, weight_category);
ALTER TABLE ONLY public.tournaments ADD CONSTRAINT tournaments_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.ussr_republics ADD CONSTRAINT ussr_republic_pkey PRIMARY KEY (id);

CREATE INDEX idx_cities_name ON public.cities USING btree (name);
CREATE INDEX idx_judoka_cities_city_id ON public.judoka_cities USING btree (city_id);
CREATE INDEX idx_judoka_cities_judoka_id ON public.judoka_cities USING btree (judoka_id);
CREATE INDEX idx_judoka_sport_clubs_judoka_id ON public.judoka_sport_clubs USING btree (judoka_id);
CREATE INDEX idx_judoka_sport_clubs_sport_club_id ON public.judoka_sport_clubs USING btree (sport_club_id);
CREATE INDEX idx_judokas_country ON public.judokas USING btree (country);
CREATE INDEX idx_judokas_name ON public.judokas USING btree (name);
CREATE INDEX idx_judokas_name_rus ON public.judokas USING btree (name_rus);
CREATE INDEX idx_sport_club_cities_city_id ON public.sport_club_cities USING btree (city_id);
CREATE INDEX idx_sport_club_cities_sport_club_id ON public.sport_club_cities USING btree (sport_club_id);
CREATE INDEX idx_sport_clubs_city_id ON public.sport_clubs USING btree (city_id);
CREATE INDEX idx_sport_clubs_name ON public.sport_clubs USING btree (name);
CREATE INDEX idx_tournament_results_judoka_id ON public.tournament_results USING btree (judoka_id);
CREATE INDEX idx_tournament_results_rank ON public.tournament_results USING btree (rank);
CREATE INDEX idx_tournament_results_tournament_id ON public.tournament_results USING btree (tournament_id);
CREATE INDEX idx_tournament_results_weight_category ON public.tournament_results USING btree (weight_category);
CREATE INDEX idx_tournaments_city_id ON public.tournaments USING btree (city_id);
CREATE INDEX idx_tournaments_gender ON public.tournaments USING btree (gender);
CREATE INDEX idx_tournaments_name ON public.tournaments USING btree (name);
CREATE INDEX idx_tournaments_year ON public.tournaments USING btree (year DESC);

ALTER TABLE ONLY public.cities ADD CONSTRAINT cities_republic_id_fkey FOREIGN KEY (republic_id) REFERENCES public.ussr_republics(id) ON DELETE SET NULL;
ALTER TABLE ONLY public.judoka_cities ADD CONSTRAINT judoka_cities_city_id_fkey FOREIGN KEY (city_id) REFERENCES public.cities(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.judoka_cities ADD CONSTRAINT judoka_cities_judoka_id_fkey FOREIGN KEY (judoka_id) REFERENCES public.judokas(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.judoka_countries ADD CONSTRAINT judoka_countries_country_id_fkey FOREIGN KEY (country_id) REFERENCES public.countries(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.judoka_countries ADD CONSTRAINT judoka_countries_judoka_id_fkey FOREIGN KEY (judoka_id) REFERENCES public.judokas(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.judoka_republics ADD CONSTRAINT judoka_republics_judoka_id_fkey FOREIGN KEY (judoka_id) REFERENCES public.judokas(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.judoka_republics ADD CONSTRAINT judoka_republics_republic_id_fkey FOREIGN KEY (republic_id) REFERENCES public.ussr_republics(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.judoka_sport_clubs ADD CONSTRAINT judoka_sport_clubs_judoka_id_fkey FOREIGN KEY (judoka_id) REFERENCES public.judokas(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.judoka_sport_clubs ADD CONSTRAINT judoka_sport_clubs_sport_club_id_fkey FOREIGN KEY (sport_club_id) REFERENCES public.sport_clubs(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.sport_club_cities ADD CONSTRAINT sport_club_cities_city_id_fkey FOREIGN KEY (city_id) REFERENCES public.cities(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.sport_club_cities ADD CONSTRAINT sport_club_cities_sport_club_id_fkey FOREIGN KEY (sport_club_id) REFERENCES public.sport_clubs(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.sport_clubs ADD CONSTRAINT sport_clubs_city_id_fkey FOREIGN KEY (city_id) REFERENCES public.cities(id) ON DELETE SET NULL;
ALTER TABLE ONLY public.tournament_results ADD CONSTRAINT tournament_results_judoka_id_fkey FOREIGN KEY (judoka_id) REFERENCES public.judokas(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.tournament_results ADD CONSTRAINT tournament_results_tournament_id_fkey FOREIGN KEY (tournament_id) REFERENCES public.tournaments(id) ON DELETE CASCADE;
ALTER TABLE ONLY public.tournaments ADD CONSTRAINT tournaments_city_id_fkey FOREIGN KEY (city_id) REFERENCES public.cities(id) ON DELETE SET NULL;
ALTER TABLE ONLY public.tournaments ADD CONSTRAINT tournaments_country_id_fkey FOREIGN KEY (country_id) REFERENCES public.countries(id) ON DELETE SET NULL;
ALTER TABLE ONLY public.tournaments ADD CONSTRAINT tournaments_republic_id_fkey FOREIGN KEY (republic_id) REFERENCES public.ussr_republics(id) ON DELETE SET NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.tournament_results  CASCADE;
DROP TABLE IF EXISTS public.judoka_sport_clubs  CASCADE;
DROP TABLE IF EXISTS public.judoka_republics    CASCADE;
DROP TABLE IF EXISTS public.judoka_countries    CASCADE;
DROP TABLE IF EXISTS public.judoka_cities       CASCADE;
DROP TABLE IF EXISTS public.sport_club_cities   CASCADE;
DROP TABLE IF EXISTS public.tournaments         CASCADE;
DROP TABLE IF EXISTS public.sport_clubs         CASCADE;
DROP TABLE IF EXISTS public.judokas             CASCADE;
DROP TABLE IF EXISTS public.cities              CASCADE;
DROP TABLE IF EXISTS public.countries           CASCADE;
DROP TABLE IF EXISTS public.ussr_republics      CASCADE;
DROP SEQUENCE IF EXISTS public.cities_id_seq;
DROP SEQUENCE IF EXISTS public.judoka_cities_id_seq;
DROP SEQUENCE IF EXISTS public.judoka_sport_clubs_id_seq;
DROP SEQUENCE IF EXISTS public.judokas_id_seq;
DROP SEQUENCE IF EXISTS public.sport_club_cities_id_seq;
DROP SEQUENCE IF EXISTS public.sport_clubs_id_seq;
DROP SEQUENCE IF EXISTS public.tournament_results_id_seq;
DROP SEQUENCE IF EXISTS public.tournaments_id_seq;
-- +goose StatementEnd
