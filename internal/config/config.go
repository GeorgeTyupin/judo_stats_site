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
