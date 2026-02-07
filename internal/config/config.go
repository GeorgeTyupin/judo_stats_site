package config

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const (
	configPath = "configs/config.yaml"
	envPath    = "configs/database.env"
)

type Config struct {
	HTTPServer HTTPServer `yaml:"http_server"`
	TLS        TLS        `yaml:"tls"`
	Database   Database   `yaml:"database"`
	Env        string     `yaml:"env" env:"ENV" env-default:"local"`
}

type HTTPServer struct {
	Timeouts HTTPTimeouts `yaml:"timeouts"`
	Address  string       `yaml:"address" env-default:"0.0.0.0:5000"`
}

type HTTPTimeouts struct {
	Request  time.Duration `yaml:"request" env-default:"4s"`
	Idle     time.Duration `yaml:"idle" env-default:"60s"`
	Shutdown time.Duration `yaml:"shutdown" env-default:"15s"`
}

type TLS struct {
	HTTPSPort  string `yaml:"https_port" env-default:":443"`
	Domain     string `yaml:"domain"`
	CacheDir   string `yaml:"cache_dir" env-default:"./certs-cache"`
	HTTPPort   string `yaml:"http_port" env-default:":80"`
	UseStaging bool   `yaml:"use_staging" env-default:"false"`
	Enabled    bool   `yaml:"enabled" env-default:"false"`
}

type Database struct {
	Host            string        `env:"DB_HOST"`
	Name            string        `env:"DB_NAME"`
	User            string        `env:"DB_USER"`
	Password        string        `env:"DB_PASSWORD"`
	SSLMode         string        `yaml:"sslmode"`
	DBIdle          time.Duration `yaml:"db_idle"`
	MaxConnLifetime time.Duration `yaml:"max_conn_life_time"`
	MaxConnIdleTime time.Duration `yaml:"max_conn_idle_time"`
	Timeout         time.Duration `yaml:"timeout"`
	MaxConns        int32         `yaml:"max_conns"`
	MinConns        int32         `yaml:"min_conns"`
	Port            int8          `env:"DB_PORT"`
}

func (d *Database) GetConnString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		d.User, d.Password, d.Host, d.Port, d.Name, d.SSLMode,
	)
}

func MustLoad(log *slog.Logger) *Config {
	const op = "config.MustLoad"
	log = log.With(slog.String("op", op))

	file, err := os.Open(configPath)
	if err != nil {
		log.Error("Не удалось открыть файл с конфигом", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer file.Close()

	conf, err := LoadConf(file)
	if err != nil {
		log.Error("Не удалось загрузить конфигурацию", slog.String("error", err.Error()))
		os.Exit(1)
	}

	return conf
}

func LoadConf(file *os.File) (*Config, error) {
	var conf Config

	if err := godotenv.Load(envPath); err != nil {
		return nil, fmt.Errorf("не удалось загрузить env. Возникла ошибка %w", err)
	}

	if err := cleanenv.ParseYAML(file, &conf); err != nil {
		return nil, fmt.Errorf("не удалось прочитать конфиг. Возникла ошибка %w", err)
	}

	if err := cleanenv.ReadEnv(&conf); err != nil {
		return nil, fmt.Errorf("не удалось прочитать env. Возникла ошибка %w", err)
	}

	return &conf, nil
}
