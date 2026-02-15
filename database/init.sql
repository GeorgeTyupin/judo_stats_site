--
-- PostgreSQL database dump
--

-- Dumped from database version 16.11
-- Dumped by pg_dump version 17.5

-- Started on 2026-02-15 21:31:48 MSK

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 215 (class 1259 OID 16385)
-- Name: cities; Type: TABLE; Schema: public; Owner: judo
--

CREATE TABLE public.cities (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    republic_id bigint
);


ALTER TABLE public.cities OWNER TO judo;

--
-- TOC entry 216 (class 1259 OID 16392)
-- Name: cities_id_seq; Type: SEQUENCE; Schema: public; Owner: judo
--

CREATE SEQUENCE public.cities_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cities_id_seq OWNER TO judo;

--
-- TOC entry 3559 (class 0 OID 0)
-- Dependencies: 216
-- Name: cities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: judo
--

ALTER SEQUENCE public.cities_id_seq OWNED BY public.cities.id;


--
-- TOC entry 217 (class 1259 OID 16393)
-- Name: countries; Type: TABLE; Schema: public; Owner: judo
--

CREATE TABLE public.countries (
    id integer NOT NULL,
    name character varying(32) NOT NULL,
    description text,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.countries OWNER TO judo;

--
-- TOC entry 218 (class 1259 OID 16400)
-- Name: judoka_cities; Type: TABLE; Schema: public; Owner: judo
--

CREATE TABLE public.judoka_cities (
    id integer NOT NULL,
    judoka_id integer NOT NULL,
    city_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.judoka_cities OWNER TO judo;

--
-- TOC entry 219 (class 1259 OID 16404)
-- Name: judoka_cities_id_seq; Type: SEQUENCE; Schema: public; Owner: judo
--

CREATE SEQUENCE public.judoka_cities_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.judoka_cities_id_seq OWNER TO judo;

--
-- TOC entry 3560 (class 0 OID 0)
-- Dependencies: 219
-- Name: judoka_cities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: judo
--

ALTER SEQUENCE public.judoka_cities_id_seq OWNED BY public.judoka_cities.id;


--
-- TOC entry 220 (class 1259 OID 16405)
-- Name: judoka_countries; Type: TABLE; Schema: public; Owner: judo
--

CREATE TABLE public.judoka_countries (
    id integer NOT NULL,
    judoka_id bigint NOT NULL,
    country_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.judoka_countries OWNER TO judo;

--
-- TOC entry 221 (class 1259 OID 16409)
-- Name: judoka_republics; Type: TABLE; Schema: public; Owner: judo
--

CREATE TABLE public.judoka_republics (
    id integer NOT NULL,
    judoka_id bigint NOT NULL,
    republic_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.judoka_republics OWNER TO judo;

--
-- TOC entry 222 (class 1259 OID 16413)
-- Name: judoka_sport_clubs; Type: TABLE; Schema: public; Owner: judo
--

CREATE TABLE public.judoka_sport_clubs (
    id integer NOT NULL,
    judoka_id integer NOT NULL,
    sport_club_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.judoka_sport_clubs OWNER TO judo;

--
-- TOC entry 223 (class 1259 OID 16417)
-- Name: judoka_sport_clubs_id_seq; Type: SEQUENCE; Schema: public; Owner: judo
--

CREATE SEQUENCE public.judoka_sport_clubs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.judoka_sport_clubs_id_seq OWNER TO judo;

--
-- TOC entry 3561 (class 0 OID 0)
-- Dependencies: 223
-- Name: judoka_sport_clubs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: judo
--

ALTER SEQUENCE public.judoka_sport_clubs_id_seq OWNED BY public.judoka_sport_clubs.id;


--
-- TOC entry 224 (class 1259 OID 16418)
-- Name: judokas; Type: TABLE; Schema: public; Owner: judo
--

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


ALTER TABLE public.judokas OWNER TO judo;

--
-- TOC entry 225 (class 1259 OID 16425)
-- Name: judokas_id_seq; Type: SEQUENCE; Schema: public; Owner: judo
--

CREATE SEQUENCE public.judokas_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.judokas_id_seq OWNER TO judo;

--
-- TOC entry 3562 (class 0 OID 0)
-- Dependencies: 225
-- Name: judokas_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: judo
--

ALTER SEQUENCE public.judokas_id_seq OWNED BY public.judokas.id;


--
-- TOC entry 226 (class 1259 OID 16426)
-- Name: sport_club_cities; Type: TABLE; Schema: public; Owner: judo
--

CREATE TABLE public.sport_club_cities (
    id integer NOT NULL,
    sport_club_id integer NOT NULL,
    city_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.sport_club_cities OWNER TO judo;

--
-- TOC entry 227 (class 1259 OID 16430)
-- Name: sport_club_cities_id_seq; Type: SEQUENCE; Schema: public; Owner: judo
--

CREATE SEQUENCE public.sport_club_cities_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sport_club_cities_id_seq OWNER TO judo;

--
-- TOC entry 3563 (class 0 OID 0)
-- Dependencies: 227
-- Name: sport_club_cities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: judo
--

ALTER SEQUENCE public.sport_club_cities_id_seq OWNED BY public.sport_club_cities.id;


--
-- TOC entry 228 (class 1259 OID 16431)
-- Name: sport_clubs; Type: TABLE; Schema: public; Owner: judo
--

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


ALTER TABLE public.sport_clubs OWNER TO judo;

--
-- TOC entry 229 (class 1259 OID 16438)
-- Name: sport_clubs_id_seq; Type: SEQUENCE; Schema: public; Owner: judo
--

CREATE SEQUENCE public.sport_clubs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sport_clubs_id_seq OWNER TO judo;

--
-- TOC entry 3564 (class 0 OID 0)
-- Dependencies: 229
-- Name: sport_clubs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: judo
--

ALTER SEQUENCE public.sport_clubs_id_seq OWNED BY public.sport_clubs.id;


--
-- TOC entry 230 (class 1259 OID 16439)
-- Name: tournament_results; Type: TABLE; Schema: public; Owner: judo
--

CREATE TABLE public.tournament_results (
    id integer NOT NULL,
    tournament_id integer NOT NULL,
    judoka_id integer NOT NULL,
    weight_category character varying(50) NOT NULL,
    rank integer NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    CONSTRAINT tournament_results_rank_check CHECK (((rank >= 1) AND (rank <= 3)))
);


ALTER TABLE public.tournament_results OWNER TO judo;

--
-- TOC entry 231 (class 1259 OID 16444)
-- Name: tournament_results_id_seq; Type: SEQUENCE; Schema: public; Owner: judo
--

CREATE SEQUENCE public.tournament_results_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tournament_results_id_seq OWNER TO judo;

--
-- TOC entry 3565 (class 0 OID 0)
-- Dependencies: 231
-- Name: tournament_results_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: judo
--

ALTER SEQUENCE public.tournament_results_id_seq OWNED BY public.tournament_results.id;


--
-- TOC entry 232 (class 1259 OID 16445)
-- Name: tournaments; Type: TABLE; Schema: public; Owner: judo
--

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


ALTER TABLE public.tournaments OWNER TO judo;

--
-- TOC entry 233 (class 1259 OID 16452)
-- Name: tournaments_id_seq; Type: SEQUENCE; Schema: public; Owner: judo
--

CREATE SEQUENCE public.tournaments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tournaments_id_seq OWNER TO judo;

--
-- TOC entry 3566 (class 0 OID 0)
-- Dependencies: 233
-- Name: tournaments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: judo
--

ALTER SEQUENCE public.tournaments_id_seq OWNED BY public.tournaments.id;


--
-- TOC entry 234 (class 1259 OID 16453)
-- Name: ussr_republics; Type: TABLE; Schema: public; Owner: judo
--

CREATE TABLE public.ussr_republics (
    id integer NOT NULL,
    name character varying(32) NOT NULL,
    description text,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.ussr_republics OWNER TO judo;

--
-- TOC entry 3315 (class 2604 OID 16460)
-- Name: cities id; Type: DEFAULT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.cities ALTER COLUMN id SET DEFAULT nextval('public.cities_id_seq'::regclass);


--
-- TOC entry 3320 (class 2604 OID 16461)
-- Name: judoka_cities id; Type: DEFAULT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_cities ALTER COLUMN id SET DEFAULT nextval('public.judoka_cities_id_seq'::regclass);


--
-- TOC entry 3324 (class 2604 OID 16462)
-- Name: judoka_sport_clubs id; Type: DEFAULT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_sport_clubs ALTER COLUMN id SET DEFAULT nextval('public.judoka_sport_clubs_id_seq'::regclass);


--
-- TOC entry 3326 (class 2604 OID 16463)
-- Name: judokas id; Type: DEFAULT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judokas ALTER COLUMN id SET DEFAULT nextval('public.judokas_id_seq'::regclass);


--
-- TOC entry 3329 (class 2604 OID 16464)
-- Name: sport_club_cities id; Type: DEFAULT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.sport_club_cities ALTER COLUMN id SET DEFAULT nextval('public.sport_club_cities_id_seq'::regclass);


--
-- TOC entry 3331 (class 2604 OID 16465)
-- Name: sport_clubs id; Type: DEFAULT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.sport_clubs ALTER COLUMN id SET DEFAULT nextval('public.sport_clubs_id_seq'::regclass);


--
-- TOC entry 3334 (class 2604 OID 16466)
-- Name: tournament_results id; Type: DEFAULT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.tournament_results ALTER COLUMN id SET DEFAULT nextval('public.tournament_results_id_seq'::regclass);


--
-- TOC entry 3336 (class 2604 OID 16467)
-- Name: tournaments id; Type: DEFAULT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.tournaments ALTER COLUMN id SET DEFAULT nextval('public.tournaments_id_seq'::regclass);


--
-- TOC entry 3343 (class 2606 OID 16469)
-- Name: cities cities_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.cities
    ADD CONSTRAINT cities_pkey PRIMARY KEY (id);


--
-- TOC entry 3346 (class 2606 OID 16471)
-- Name: countries country_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.countries
    ADD CONSTRAINT country_pkey PRIMARY KEY (id);


--
-- TOC entry 3350 (class 2606 OID 16473)
-- Name: judoka_cities judoka_cities_judoka_id_city_id_key; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_cities
    ADD CONSTRAINT judoka_cities_judoka_id_city_id_key UNIQUE (judoka_id, city_id);


--
-- TOC entry 3352 (class 2606 OID 16475)
-- Name: judoka_cities judoka_cities_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_cities
    ADD CONSTRAINT judoka_cities_pkey PRIMARY KEY (id);


--
-- TOC entry 3354 (class 2606 OID 16477)
-- Name: judoka_countries judoka_countries_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_countries
    ADD CONSTRAINT judoka_countries_pkey PRIMARY KEY (id);


--
-- TOC entry 3356 (class 2606 OID 16479)
-- Name: judoka_republics judoka_republics_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_republics
    ADD CONSTRAINT judoka_republics_pkey PRIMARY KEY (id);


--
-- TOC entry 3360 (class 2606 OID 16481)
-- Name: judoka_sport_clubs judoka_sport_clubs_judoka_id_sport_club_id_key; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_sport_clubs
    ADD CONSTRAINT judoka_sport_clubs_judoka_id_sport_club_id_key UNIQUE (judoka_id, sport_club_id);


--
-- TOC entry 3362 (class 2606 OID 16483)
-- Name: judoka_sport_clubs judoka_sport_clubs_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_sport_clubs
    ADD CONSTRAINT judoka_sport_clubs_pkey PRIMARY KEY (id);


--
-- TOC entry 3367 (class 2606 OID 16485)
-- Name: judokas judokas_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judokas
    ADD CONSTRAINT judokas_pkey PRIMARY KEY (id);


--
-- TOC entry 3371 (class 2606 OID 16487)
-- Name: sport_club_cities sport_club_cities_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.sport_club_cities
    ADD CONSTRAINT sport_club_cities_pkey PRIMARY KEY (id);


--
-- TOC entry 3373 (class 2606 OID 16489)
-- Name: sport_club_cities sport_club_cities_sport_club_id_city_id_key; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.sport_club_cities
    ADD CONSTRAINT sport_club_cities_sport_club_id_city_id_key UNIQUE (sport_club_id, city_id);


--
-- TOC entry 3377 (class 2606 OID 16491)
-- Name: sport_clubs sport_clubs_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.sport_clubs
    ADD CONSTRAINT sport_clubs_pkey PRIMARY KEY (id);


--
-- TOC entry 3383 (class 2606 OID 16493)
-- Name: tournament_results tournament_results_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.tournament_results
    ADD CONSTRAINT tournament_results_pkey PRIMARY KEY (id);


--
-- TOC entry 3385 (class 2606 OID 16495)
-- Name: tournament_results tournament_results_tournament_id_judoka_id_weight_category_key; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.tournament_results
    ADD CONSTRAINT tournament_results_tournament_id_judoka_id_weight_category_key UNIQUE (tournament_id, judoka_id, weight_category);


--
-- TOC entry 3391 (class 2606 OID 16497)
-- Name: tournaments tournaments_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.tournaments
    ADD CONSTRAINT tournaments_pkey PRIMARY KEY (id);


--
-- TOC entry 3393 (class 2606 OID 16499)
-- Name: ussr_republics ussr_republic_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.ussr_republics
    ADD CONSTRAINT ussr_republic_pkey PRIMARY KEY (id);


--
-- TOC entry 3344 (class 1259 OID 16500)
-- Name: idx_cities_name; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_cities_name ON public.cities USING btree (name);


--
-- TOC entry 3347 (class 1259 OID 16501)
-- Name: idx_judoka_cities_city_id; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_judoka_cities_city_id ON public.judoka_cities USING btree (city_id);


--
-- TOC entry 3348 (class 1259 OID 16502)
-- Name: idx_judoka_cities_judoka_id; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_judoka_cities_judoka_id ON public.judoka_cities USING btree (judoka_id);


--
-- TOC entry 3357 (class 1259 OID 16503)
-- Name: idx_judoka_sport_clubs_judoka_id; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_judoka_sport_clubs_judoka_id ON public.judoka_sport_clubs USING btree (judoka_id);


--
-- TOC entry 3358 (class 1259 OID 16504)
-- Name: idx_judoka_sport_clubs_sport_club_id; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_judoka_sport_clubs_sport_club_id ON public.judoka_sport_clubs USING btree (sport_club_id);


--
-- TOC entry 3363 (class 1259 OID 16505)
-- Name: idx_judokas_country; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_judokas_country ON public.judokas USING btree (country);


--
-- TOC entry 3364 (class 1259 OID 16506)
-- Name: idx_judokas_name; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_judokas_name ON public.judokas USING btree (name);


--
-- TOC entry 3365 (class 1259 OID 16507)
-- Name: idx_judokas_name_rus; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_judokas_name_rus ON public.judokas USING btree (name_rus);


--
-- TOC entry 3368 (class 1259 OID 16508)
-- Name: idx_sport_club_cities_city_id; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_sport_club_cities_city_id ON public.sport_club_cities USING btree (city_id);


--
-- TOC entry 3369 (class 1259 OID 16509)
-- Name: idx_sport_club_cities_sport_club_id; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_sport_club_cities_sport_club_id ON public.sport_club_cities USING btree (sport_club_id);


--
-- TOC entry 3374 (class 1259 OID 16510)
-- Name: idx_sport_clubs_city_id; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_sport_clubs_city_id ON public.sport_clubs USING btree (city_id);


--
-- TOC entry 3375 (class 1259 OID 16511)
-- Name: idx_sport_clubs_name; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_sport_clubs_name ON public.sport_clubs USING btree (name);


--
-- TOC entry 3378 (class 1259 OID 16512)
-- Name: idx_tournament_results_judoka_id; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_tournament_results_judoka_id ON public.tournament_results USING btree (judoka_id);


--
-- TOC entry 3379 (class 1259 OID 16513)
-- Name: idx_tournament_results_rank; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_tournament_results_rank ON public.tournament_results USING btree (rank);


--
-- TOC entry 3380 (class 1259 OID 16514)
-- Name: idx_tournament_results_tournament_id; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_tournament_results_tournament_id ON public.tournament_results USING btree (tournament_id);


--
-- TOC entry 3381 (class 1259 OID 16515)
-- Name: idx_tournament_results_weight_category; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_tournament_results_weight_category ON public.tournament_results USING btree (weight_category);


--
-- TOC entry 3386 (class 1259 OID 16516)
-- Name: idx_tournaments_city_id; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_tournaments_city_id ON public.tournaments USING btree (city_id);


--
-- TOC entry 3387 (class 1259 OID 16517)
-- Name: idx_tournaments_gender; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_tournaments_gender ON public.tournaments USING btree (gender);


--
-- TOC entry 3388 (class 1259 OID 16518)
-- Name: idx_tournaments_name; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_tournaments_name ON public.tournaments USING btree (name);


--
-- TOC entry 3389 (class 1259 OID 16519)
-- Name: idx_tournaments_year; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_tournaments_year ON public.tournaments USING btree (year DESC);


--
-- TOC entry 3394 (class 2606 OID 16520)
-- Name: cities cities_republic_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.cities
    ADD CONSTRAINT cities_republic_id_fkey FOREIGN KEY (republic_id) REFERENCES public.ussr_republics(id) ON DELETE SET NULL;


--
-- TOC entry 3395 (class 2606 OID 16525)
-- Name: judoka_cities judoka_cities_city_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_cities
    ADD CONSTRAINT judoka_cities_city_id_fkey FOREIGN KEY (city_id) REFERENCES public.cities(id) ON DELETE CASCADE;


--
-- TOC entry 3396 (class 2606 OID 16530)
-- Name: judoka_cities judoka_cities_judoka_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_cities
    ADD CONSTRAINT judoka_cities_judoka_id_fkey FOREIGN KEY (judoka_id) REFERENCES public.judokas(id) ON DELETE CASCADE;


--
-- TOC entry 3397 (class 2606 OID 16535)
-- Name: judoka_countries judoka_countries_country_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_countries
    ADD CONSTRAINT judoka_countries_country_id_fkey FOREIGN KEY (country_id) REFERENCES public.countries(id) ON DELETE CASCADE;


--
-- TOC entry 3398 (class 2606 OID 16540)
-- Name: judoka_countries judoka_countries_judoka_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_countries
    ADD CONSTRAINT judoka_countries_judoka_id_fkey FOREIGN KEY (judoka_id) REFERENCES public.judokas(id) ON DELETE CASCADE;


--
-- TOC entry 3399 (class 2606 OID 16545)
-- Name: judoka_republics judoka_republics_judoka_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_republics
    ADD CONSTRAINT judoka_republics_judoka_id_fkey FOREIGN KEY (judoka_id) REFERENCES public.judokas(id) ON DELETE CASCADE;


--
-- TOC entry 3400 (class 2606 OID 16550)
-- Name: judoka_republics judoka_republics_republic_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_republics
    ADD CONSTRAINT judoka_republics_republic_id_fkey FOREIGN KEY (republic_id) REFERENCES public.ussr_republics(id) ON DELETE CASCADE;


--
-- TOC entry 3401 (class 2606 OID 16555)
-- Name: judoka_sport_clubs judoka_sport_clubs_judoka_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_sport_clubs
    ADD CONSTRAINT judoka_sport_clubs_judoka_id_fkey FOREIGN KEY (judoka_id) REFERENCES public.judokas(id) ON DELETE CASCADE;


--
-- TOC entry 3402 (class 2606 OID 16560)
-- Name: judoka_sport_clubs judoka_sport_clubs_sport_club_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_sport_clubs
    ADD CONSTRAINT judoka_sport_clubs_sport_club_id_fkey FOREIGN KEY (sport_club_id) REFERENCES public.sport_clubs(id) ON DELETE CASCADE;


--
-- TOC entry 3403 (class 2606 OID 16565)
-- Name: sport_club_cities sport_club_cities_city_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.sport_club_cities
    ADD CONSTRAINT sport_club_cities_city_id_fkey FOREIGN KEY (city_id) REFERENCES public.cities(id) ON DELETE CASCADE;


--
-- TOC entry 3404 (class 2606 OID 16570)
-- Name: sport_club_cities sport_club_cities_sport_club_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.sport_club_cities
    ADD CONSTRAINT sport_club_cities_sport_club_id_fkey FOREIGN KEY (sport_club_id) REFERENCES public.sport_clubs(id) ON DELETE CASCADE;


--
-- TOC entry 3405 (class 2606 OID 16575)
-- Name: sport_clubs sport_clubs_city_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.sport_clubs
    ADD CONSTRAINT sport_clubs_city_id_fkey FOREIGN KEY (city_id) REFERENCES public.cities(id) ON DELETE SET NULL;


--
-- TOC entry 3406 (class 2606 OID 16580)
-- Name: tournament_results tournament_results_judoka_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.tournament_results
    ADD CONSTRAINT tournament_results_judoka_id_fkey FOREIGN KEY (judoka_id) REFERENCES public.judokas(id) ON DELETE CASCADE;


--
-- TOC entry 3407 (class 2606 OID 16585)
-- Name: tournament_results tournament_results_tournament_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.tournament_results
    ADD CONSTRAINT tournament_results_tournament_id_fkey FOREIGN KEY (tournament_id) REFERENCES public.tournaments(id) ON DELETE CASCADE;


--
-- TOC entry 3408 (class 2606 OID 16590)
-- Name: tournaments tournaments_city_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.tournaments
    ADD CONSTRAINT tournaments_city_id_fkey FOREIGN KEY (city_id) REFERENCES public.cities(id) ON DELETE SET NULL;


--
-- TOC entry 3409 (class 2606 OID 16595)
-- Name: tournaments tournaments_country_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.tournaments
    ADD CONSTRAINT tournaments_country_id_fkey FOREIGN KEY (country_id) REFERENCES public.countries(id) ON DELETE SET NULL;


--
-- TOC entry 3410 (class 2606 OID 16604)
-- Name: tournaments tournaments_republic_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.tournaments
    ADD CONSTRAINT tournaments_republic_id_fkey FOREIGN KEY (republic_id) REFERENCES public.ussr_republics(id) ON DELETE SET NULL;


-- Completed on 2026-02-15 21:31:48 MSK

--
-- PostgreSQL database dump complete
--

