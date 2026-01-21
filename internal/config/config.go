package config

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	configPath = "configs/config.yaml"
)

type Config struct {
	Env        string     `yaml:"env" env:"ENV" env-default:"local"`
	HTTPServer HTTPServer `yaml:"http_server"`
	TLS        TLS        `yaml:"tls"`
}

type HTTPServer struct {
	Address  string       `yaml:"address" env-default:"0.0.0.0:5000"`
	Timeouts HTTPTimeouts `yaml:"timeouts"`
}

type HTTPTimeouts struct {
	Request  time.Duration `yaml:"request" env-default:"4s"`
	Idle     time.Duration `yaml:"idle" env-default:"60s"`
	Shutdown time.Duration `yaml:"shutdown" env-default:"15s"`
}

type TLS struct {
	Enabled    bool   `yaml:"enabled" env:"TLS_ENABLED" env-default:"false"`
	Domain     string `yaml:"domain" env:"TLS_DOMAIN"`
	CacheDir   string `yaml:"cache_dir" env:"TLS_CACHE_DIR" env-default:"./certs-cache"`
	UseStaging bool   `yaml:"use_staging" env:"TLS_USE_STAGING" env-default:"false"`
	HTTPPort   string `yaml:"http_port" env:"TLS_HTTP_PORT" env-default:":80"`
	HTTPSPort  string `yaml:"https_port" env:"TLS_HTTPS_PORT" env-default:":443"`
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

	if err := cleanenv.ParseYAML(file, &conf); err != nil {
		return nil, fmt.Errorf("не удалось прочитать конфиг. Возникла ошибка %w", err)
	}

	return &conf, nil
}
