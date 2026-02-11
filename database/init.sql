--
-- PostgreSQL database dump
--

-- Dumped from database version 16.11
-- Dumped by pg_dump version 17.4

-- Started on 2026-02-11 17:21:05

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
-- TOC entry 216 (class 1259 OID 16386)
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
-- TOC entry 215 (class 1259 OID 16385)
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
-- TOC entry 3558 (class 0 OID 0)
-- Dependencies: 215
-- Name: cities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: judo
--

ALTER SEQUENCE public.cities_id_seq OWNED BY public.cities.id;


--
-- TOC entry 232 (class 1259 OID 16571)
-- Name: country; Type: TABLE; Schema: public; Owner: judo
--

CREATE TABLE public.country (
    id integer NOT NULL,
    name character varying(32) NOT NULL,
    description text,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.country OWNER TO judo;

--
-- TOC entry 226 (class 1259 OID 16488)
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
-- TOC entry 225 (class 1259 OID 16487)
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
-- TOC entry 3559 (class 0 OID 0)
-- Dependencies: 225
-- Name: judoka_cities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: judo
--

ALTER SEQUENCE public.judoka_cities_id_seq OWNED BY public.judoka_cities.id;


--
-- TOC entry 233 (class 1259 OID 16580)
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
-- TOC entry 234 (class 1259 OID 16596)
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
-- TOC entry 228 (class 1259 OID 16510)
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
-- TOC entry 227 (class 1259 OID 16509)
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
-- TOC entry 3560 (class 0 OID 0)
-- Dependencies: 227
-- Name: judoka_sport_clubs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: judo
--

ALTER SEQUENCE public.judoka_sport_clubs_id_seq OWNED BY public.judoka_sport_clubs.id;


--
-- TOC entry 222 (class 1259 OID 16437)
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
-- TOC entry 221 (class 1259 OID 16436)
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
-- TOC entry 3561 (class 0 OID 0)
-- Dependencies: 221
-- Name: judokas_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: judo
--

ALTER SEQUENCE public.judokas_id_seq OWNED BY public.judokas.id;


--
-- TOC entry 230 (class 1259 OID 16532)
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
-- TOC entry 229 (class 1259 OID 16531)
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
-- TOC entry 3562 (class 0 OID 0)
-- Dependencies: 229
-- Name: sport_club_cities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: judo
--

ALTER SEQUENCE public.sport_club_cities_id_seq OWNED BY public.sport_club_cities.id;


--
-- TOC entry 218 (class 1259 OID 16399)
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
-- TOC entry 217 (class 1259 OID 16398)
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
-- TOC entry 3563 (class 0 OID 0)
-- Dependencies: 217
-- Name: sport_clubs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: judo
--

ALTER SEQUENCE public.sport_clubs_id_seq OWNED BY public.sport_clubs.id;


--
-- TOC entry 224 (class 1259 OID 16463)
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
-- TOC entry 223 (class 1259 OID 16462)
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
-- TOC entry 3564 (class 0 OID 0)
-- Dependencies: 223
-- Name: tournament_results_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: judo
--

ALTER SEQUENCE public.tournament_results_id_seq OWNED BY public.tournament_results.id;


--
-- TOC entry 220 (class 1259 OID 16417)
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
    country_id integer
);


ALTER TABLE public.tournaments OWNER TO judo;

--
-- TOC entry 219 (class 1259 OID 16416)
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
-- TOC entry 3565 (class 0 OID 0)
-- Dependencies: 219
-- Name: tournaments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: judo
--

ALTER SEQUENCE public.tournaments_id_seq OWNED BY public.tournaments.id;


--
-- TOC entry 231 (class 1259 OID 16557)
-- Name: ussr_republic; Type: TABLE; Schema: public; Owner: judo
--

CREATE TABLE public.ussr_republic (
    id integer NOT NULL,
    name character varying(32) NOT NULL,
    description text,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.ussr_republic OWNER TO judo;

--
-- TOC entry 3315 (class 2604 OID 24857)
-- Name: cities id; Type: DEFAULT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.cities ALTER COLUMN id SET DEFAULT nextval('public.cities_id_seq'::regclass);


--
-- TOC entry 3329 (class 2604 OID 16491)
-- Name: judoka_cities id; Type: DEFAULT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_cities ALTER COLUMN id SET DEFAULT nextval('public.judoka_cities_id_seq'::regclass);


--
-- TOC entry 3331 (class 2604 OID 16513)
-- Name: judoka_sport_clubs id; Type: DEFAULT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_sport_clubs ALTER COLUMN id SET DEFAULT nextval('public.judoka_sport_clubs_id_seq'::regclass);


--
-- TOC entry 3324 (class 2604 OID 24820)
-- Name: judokas id; Type: DEFAULT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judokas ALTER COLUMN id SET DEFAULT nextval('public.judokas_id_seq'::regclass);


--
-- TOC entry 3333 (class 2604 OID 16535)
-- Name: sport_club_cities id; Type: DEFAULT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.sport_club_cities ALTER COLUMN id SET DEFAULT nextval('public.sport_club_cities_id_seq'::regclass);


--
-- TOC entry 3318 (class 2604 OID 16402)
-- Name: sport_clubs id; Type: DEFAULT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.sport_clubs ALTER COLUMN id SET DEFAULT nextval('public.sport_clubs_id_seq'::regclass);


--
-- TOC entry 3327 (class 2604 OID 16466)
-- Name: tournament_results id; Type: DEFAULT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.tournament_results ALTER COLUMN id SET DEFAULT nextval('public.tournament_results_id_seq'::regclass);


--
-- TOC entry 3321 (class 2604 OID 16420)
-- Name: tournaments id; Type: DEFAULT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.tournaments ALTER COLUMN id SET DEFAULT nextval('public.tournaments_id_seq'::regclass);


--
-- TOC entry 3343 (class 2606 OID 24859)
-- Name: cities cities_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.cities
    ADD CONSTRAINT cities_pkey PRIMARY KEY (id);


--
-- TOC entry 3389 (class 2606 OID 24755)
-- Name: country country_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.country
    ADD CONSTRAINT country_pkey PRIMARY KEY (id);


--
-- TOC entry 3371 (class 2606 OID 16496)
-- Name: judoka_cities judoka_cities_judoka_id_city_id_key; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_cities
    ADD CONSTRAINT judoka_cities_judoka_id_city_id_key UNIQUE (judoka_id, city_id);


--
-- TOC entry 3373 (class 2606 OID 16494)
-- Name: judoka_cities judoka_cities_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_cities
    ADD CONSTRAINT judoka_cities_pkey PRIMARY KEY (id);


--
-- TOC entry 3391 (class 2606 OID 16585)
-- Name: judoka_countries judoka_countries_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_countries
    ADD CONSTRAINT judoka_countries_pkey PRIMARY KEY (id);


--
-- TOC entry 3393 (class 2606 OID 16601)
-- Name: judoka_republics judoka_republics_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_republics
    ADD CONSTRAINT judoka_republics_pkey PRIMARY KEY (id);


--
-- TOC entry 3377 (class 2606 OID 16518)
-- Name: judoka_sport_clubs judoka_sport_clubs_judoka_id_sport_club_id_key; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_sport_clubs
    ADD CONSTRAINT judoka_sport_clubs_judoka_id_sport_club_id_key UNIQUE (judoka_id, sport_club_id);


--
-- TOC entry 3379 (class 2606 OID 16516)
-- Name: judoka_sport_clubs judoka_sport_clubs_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_sport_clubs
    ADD CONSTRAINT judoka_sport_clubs_pkey PRIMARY KEY (id);


--
-- TOC entry 3359 (class 2606 OID 24822)
-- Name: judokas judokas_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judokas
    ADD CONSTRAINT judokas_pkey PRIMARY KEY (id);


--
-- TOC entry 3383 (class 2606 OID 16538)
-- Name: sport_club_cities sport_club_cities_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.sport_club_cities
    ADD CONSTRAINT sport_club_cities_pkey PRIMARY KEY (id);


--
-- TOC entry 3385 (class 2606 OID 16540)
-- Name: sport_club_cities sport_club_cities_sport_club_id_city_id_key; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.sport_club_cities
    ADD CONSTRAINT sport_club_cities_sport_club_id_city_id_key UNIQUE (sport_club_id, city_id);


--
-- TOC entry 3348 (class 2606 OID 16408)
-- Name: sport_clubs sport_clubs_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.sport_clubs
    ADD CONSTRAINT sport_clubs_pkey PRIMARY KEY (id);


--
-- TOC entry 3365 (class 2606 OID 16470)
-- Name: tournament_results tournament_results_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.tournament_results
    ADD CONSTRAINT tournament_results_pkey PRIMARY KEY (id);


--
-- TOC entry 3367 (class 2606 OID 16472)
-- Name: tournament_results tournament_results_tournament_id_judoka_id_weight_category_key; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.tournament_results
    ADD CONSTRAINT tournament_results_tournament_id_judoka_id_weight_category_key UNIQUE (tournament_id, judoka_id, weight_category);


--
-- TOC entry 3354 (class 2606 OID 16426)
-- Name: tournaments tournaments_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.tournaments
    ADD CONSTRAINT tournaments_pkey PRIMARY KEY (id);


--
-- TOC entry 3387 (class 2606 OID 24773)
-- Name: ussr_republic ussr_republic_pkey; Type: CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.ussr_republic
    ADD CONSTRAINT ussr_republic_pkey PRIMARY KEY (id);


--
-- TOC entry 3344 (class 1259 OID 16396)
-- Name: idx_cities_name; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_cities_name ON public.cities USING btree (name);


--
-- TOC entry 3368 (class 1259 OID 16508)
-- Name: idx_judoka_cities_city_id; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_judoka_cities_city_id ON public.judoka_cities USING btree (city_id);


--
-- TOC entry 3369 (class 1259 OID 16507)
-- Name: idx_judoka_cities_judoka_id; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_judoka_cities_judoka_id ON public.judoka_cities USING btree (judoka_id);


--
-- TOC entry 3374 (class 1259 OID 16529)
-- Name: idx_judoka_sport_clubs_judoka_id; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_judoka_sport_clubs_judoka_id ON public.judoka_sport_clubs USING btree (judoka_id);


--
-- TOC entry 3375 (class 1259 OID 16530)
-- Name: idx_judoka_sport_clubs_sport_club_id; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_judoka_sport_clubs_sport_club_id ON public.judoka_sport_clubs USING btree (sport_club_id);


--
-- TOC entry 3355 (class 1259 OID 16459)
-- Name: idx_judokas_country; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_judokas_country ON public.judokas USING btree (country);


--
-- TOC entry 3356 (class 1259 OID 16457)
-- Name: idx_judokas_name; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_judokas_name ON public.judokas USING btree (name);


--
-- TOC entry 3357 (class 1259 OID 16458)
-- Name: idx_judokas_name_rus; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_judokas_name_rus ON public.judokas USING btree (name_rus);


--
-- TOC entry 3380 (class 1259 OID 16552)
-- Name: idx_sport_club_cities_city_id; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_sport_club_cities_city_id ON public.sport_club_cities USING btree (city_id);


--
-- TOC entry 3381 (class 1259 OID 16551)
-- Name: idx_sport_club_cities_sport_club_id; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_sport_club_cities_sport_club_id ON public.sport_club_cities USING btree (sport_club_id);


--
-- TOC entry 3345 (class 1259 OID 16415)
-- Name: idx_sport_clubs_city_id; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_sport_clubs_city_id ON public.sport_clubs USING btree (city_id);


--
-- TOC entry 3346 (class 1259 OID 16414)
-- Name: idx_sport_clubs_name; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_sport_clubs_name ON public.sport_clubs USING btree (name);


--
-- TOC entry 3360 (class 1259 OID 16484)
-- Name: idx_tournament_results_judoka_id; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_tournament_results_judoka_id ON public.tournament_results USING btree (judoka_id);


--
-- TOC entry 3361 (class 1259 OID 16486)
-- Name: idx_tournament_results_rank; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_tournament_results_rank ON public.tournament_results USING btree (rank);


--
-- TOC entry 3362 (class 1259 OID 16483)
-- Name: idx_tournament_results_tournament_id; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_tournament_results_tournament_id ON public.tournament_results USING btree (tournament_id);


--
-- TOC entry 3363 (class 1259 OID 16485)
-- Name: idx_tournament_results_weight_category; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_tournament_results_weight_category ON public.tournament_results USING btree (weight_category);


--
-- TOC entry 3349 (class 1259 OID 16434)
-- Name: idx_tournaments_city_id; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_tournaments_city_id ON public.tournaments USING btree (city_id);


--
-- TOC entry 3350 (class 1259 OID 16435)
-- Name: idx_tournaments_gender; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_tournaments_gender ON public.tournaments USING btree (gender);


--
-- TOC entry 3351 (class 1259 OID 16432)
-- Name: idx_tournaments_name; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_tournaments_name ON public.tournaments USING btree (name);


--
-- TOC entry 3352 (class 1259 OID 16433)
-- Name: idx_tournaments_year; Type: INDEX; Schema: public; Owner: judo
--

CREATE INDEX idx_tournaments_year ON public.tournaments USING btree (year DESC);


--
-- TOC entry 3394 (class 2606 OID 24774)
-- Name: cities cities_republic_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.cities
    ADD CONSTRAINT cities_republic_id_fkey FOREIGN KEY (republic_id) REFERENCES public.ussr_republic(id) ON DELETE SET NULL;


--
-- TOC entry 3400 (class 2606 OID 24870)
-- Name: judoka_cities judoka_cities_city_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_cities
    ADD CONSTRAINT judoka_cities_city_id_fkey FOREIGN KEY (city_id) REFERENCES public.cities(id) ON DELETE CASCADE;


--
-- TOC entry 3401 (class 2606 OID 24828)
-- Name: judoka_cities judoka_cities_judoka_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_cities
    ADD CONSTRAINT judoka_cities_judoka_id_fkey FOREIGN KEY (judoka_id) REFERENCES public.judokas(id) ON DELETE CASCADE;


--
-- TOC entry 3406 (class 2606 OID 24756)
-- Name: judoka_countries judoka_countries_country_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_countries
    ADD CONSTRAINT judoka_countries_country_id_fkey FOREIGN KEY (country_id) REFERENCES public.country(id) ON DELETE CASCADE;


--
-- TOC entry 3407 (class 2606 OID 24838)
-- Name: judoka_countries judoka_countries_judoka_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_countries
    ADD CONSTRAINT judoka_countries_judoka_id_fkey FOREIGN KEY (judoka_id) REFERENCES public.judokas(id) ON DELETE CASCADE;


--
-- TOC entry 3408 (class 2606 OID 24843)
-- Name: judoka_republics judoka_republics_judoka_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_republics
    ADD CONSTRAINT judoka_republics_judoka_id_fkey FOREIGN KEY (judoka_id) REFERENCES public.judokas(id) ON DELETE CASCADE;


--
-- TOC entry 3409 (class 2606 OID 24779)
-- Name: judoka_republics judoka_republics_republic_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_republics
    ADD CONSTRAINT judoka_republics_republic_id_fkey FOREIGN KEY (republic_id) REFERENCES public.ussr_republic(id) ON DELETE CASCADE;


--
-- TOC entry 3402 (class 2606 OID 24833)
-- Name: judoka_sport_clubs judoka_sport_clubs_judoka_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_sport_clubs
    ADD CONSTRAINT judoka_sport_clubs_judoka_id_fkey FOREIGN KEY (judoka_id) REFERENCES public.judokas(id) ON DELETE CASCADE;


--
-- TOC entry 3403 (class 2606 OID 16524)
-- Name: judoka_sport_clubs judoka_sport_clubs_sport_club_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.judoka_sport_clubs
    ADD CONSTRAINT judoka_sport_clubs_sport_club_id_fkey FOREIGN KEY (sport_club_id) REFERENCES public.sport_clubs(id) ON DELETE CASCADE;


--
-- TOC entry 3404 (class 2606 OID 24875)
-- Name: sport_club_cities sport_club_cities_city_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.sport_club_cities
    ADD CONSTRAINT sport_club_cities_city_id_fkey FOREIGN KEY (city_id) REFERENCES public.cities(id) ON DELETE CASCADE;


--
-- TOC entry 3405 (class 2606 OID 16541)
-- Name: sport_club_cities sport_club_cities_sport_club_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.sport_club_cities
    ADD CONSTRAINT sport_club_cities_sport_club_id_fkey FOREIGN KEY (sport_club_id) REFERENCES public.sport_clubs(id) ON DELETE CASCADE;


--
-- TOC entry 3395 (class 2606 OID 24860)
-- Name: sport_clubs sport_clubs_city_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.sport_clubs
    ADD CONSTRAINT sport_clubs_city_id_fkey FOREIGN KEY (city_id) REFERENCES public.cities(id) ON DELETE SET NULL;


--
-- TOC entry 3398 (class 2606 OID 24823)
-- Name: tournament_results tournament_results_judoka_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.tournament_results
    ADD CONSTRAINT tournament_results_judoka_id_fkey FOREIGN KEY (judoka_id) REFERENCES public.judokas(id) ON DELETE CASCADE;


--
-- TOC entry 3399 (class 2606 OID 16473)
-- Name: tournament_results tournament_results_tournament_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.tournament_results
    ADD CONSTRAINT tournament_results_tournament_id_fkey FOREIGN KEY (tournament_id) REFERENCES public.tournaments(id) ON DELETE CASCADE;


--
-- TOC entry 3396 (class 2606 OID 24865)
-- Name: tournaments tournaments_city_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.tournaments
    ADD CONSTRAINT tournaments_city_id_fkey FOREIGN KEY (city_id) REFERENCES public.cities(id) ON DELETE SET NULL;


--
-- TOC entry 3397 (class 2606 OID 24761)
-- Name: tournaments tournaments_country_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: judo
--

ALTER TABLE ONLY public.tournaments
    ADD CONSTRAINT tournaments_country_id_fkey FOREIGN KEY (country_id) REFERENCES public.country(id) ON DELETE SET NULL;


-- Completed on 2026-02-11 17:21:06

--
-- PostgreSQL database dump complete
--

