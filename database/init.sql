-- Judo Archive Database Schema
-- PostgreSQL 16

CREATE TABLE IF NOT EXISTS cities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    region VARCHAR(255) NOT NULL,
    population INTEGER,
    founded VARCHAR(50),
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_cities_name ON cities(name);
CREATE INDEX idx_cities_region ON cities(region);

CREATE TABLE IF NOT EXISTS sport_clubs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    full_name VARCHAR(500),
    founded VARCHAR(50),
    city_id INTEGER REFERENCES cities(id) ON DELETE SET NULL,
    region VARCHAR(255),
    head_coach VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_sport_clubs_name ON sport_clubs(name);
CREATE INDEX idx_sport_clubs_city_id ON sport_clubs(city_id);

CREATE TABLE IF NOT EXISTS tournaments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(500) NOT NULL,
    type VARCHAR(100),
    place VARCHAR(255),
    city_id INTEGER REFERENCES cities(id) ON DELETE SET NULL,
    date VARCHAR(100),
    year SMALLINT,
    month SMALLINT,
    gender VARCHAR(10),
    city_last VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_tournaments_name ON tournaments(name);
CREATE INDEX idx_tournaments_year ON tournaments(year DESC);
CREATE INDEX idx_tournaments_city_id ON tournaments(city_id);
CREATE INDEX idx_tournaments_gender ON tournaments(gender);

CREATE TABLE IF NOT EXISTS judokas (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    name_rus VARCHAR(255),
    country VARCHAR(100),
    country_flag VARCHAR(10),
    weight_category VARCHAR(50),
    birth_date VARCHAR(50),
    birth_place VARCHAR(255),
    city_id INTEGER REFERENCES cities(id) ON DELETE SET NULL,
    sport_club_id INTEGER REFERENCES sport_clubs(id) ON DELETE SET NULL,
    sport_club VARCHAR(255),
    coach VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_judokas_name ON judokas(name);
CREATE INDEX idx_judokas_name_rus ON judokas(name_rus);
CREATE INDEX idx_judokas_country ON judokas(country);
CREATE INDEX idx_judokas_city_id ON judokas(city_id);
CREATE INDEX idx_judokas_sport_club_id ON judokas(sport_club_id);

CREATE TABLE IF NOT EXISTS tournament_results (
    id SERIAL PRIMARY KEY,
    tournament_id INTEGER NOT NULL REFERENCES tournaments(id) ON DELETE CASCADE,
    judoka_id INTEGER NOT NULL REFERENCES judokas(id) ON DELETE CASCADE,
    weight_category VARCHAR(50) NOT NULL,
    rank INTEGER NOT NULL CHECK (rank BETWEEN 1 AND 3),
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(tournament_id, judoka_id, weight_category)
);

CREATE INDEX idx_tournament_results_tournament_id ON tournament_results(tournament_id);
CREATE INDEX idx_tournament_results_judoka_id ON tournament_results(judoka_id);
CREATE INDEX idx_tournament_results_weight_category ON tournament_results(weight_category);
CREATE INDEX idx_tournament_results_rank ON tournament_results(rank);

CREATE TABLE IF NOT EXISTS judoka_cities (
    id SERIAL PRIMARY KEY,
    judoka_id INTEGER NOT NULL REFERENCES judokas(id) ON DELETE CASCADE,
    city_id INTEGER NOT NULL REFERENCES cities(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(judoka_id, city_id)
);

CREATE INDEX idx_judoka_cities_judoka_id ON judoka_cities(judoka_id);
CREATE INDEX idx_judoka_cities_city_id ON judoka_cities(city_id);

CREATE TABLE IF NOT EXISTS judoka_sport_clubs (
    id SERIAL PRIMARY KEY,
    judoka_id INTEGER NOT NULL REFERENCES judokas(id) ON DELETE CASCADE,
    sport_club_id INTEGER NOT NULL REFERENCES sport_clubs(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(judoka_id, sport_club_id)
);

CREATE INDEX idx_judoka_sport_clubs_judoka_id ON judoka_sport_clubs(judoka_id);
CREATE INDEX idx_judoka_sport_clubs_sport_club_id ON judoka_sport_clubs(sport_club_id);

CREATE TABLE IF NOT EXISTS sport_club_cities (
    id SERIAL PRIMARY KEY,
    sport_club_id INTEGER NOT NULL REFERENCES sport_clubs(id) ON DELETE CASCADE,
    city_id INTEGER NOT NULL REFERENCES cities(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(sport_club_id, city_id)
);

CREATE INDEX idx_sport_club_cities_sport_club_id ON sport_club_cities(sport_club_id);
CREATE INDEX idx_sport_club_cities_city_id ON sport_club_cities(city_id);
