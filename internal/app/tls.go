package app

import (
	"crypto/tls"
	"judo_stats_site/internal/config"
	"log/slog"

	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
)

func NewAutocertManager(logger *slog.Logger, cfg *config.TLS) *autocert.Manager {
	const op = "app.NewAutocertManager"
	logger = logger.With(slog.String("op", op))

	directoryURL := acme.LetsEncryptURL

	if cfg.UseStaging {
		directoryURL = "https://acme-staging-v02.api.letsencrypt.org/directory"
		logger.Warn("Используется Let's Encrypt STAGING - сертификаты НЕ будут доверенными")
	}

	manager := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,                 // Автоматически принять ToS Let's Encrypt
		HostPolicy: autocert.HostWhitelist(cfg.Domain), // Только указанный домен
		Cache:      autocert.DirCache(cfg.CacheDir),    // Сохранять сертификаты на диск
		Client: &acme.Client{
			DirectoryURL: directoryURL,
		},
	}

	return manager
}

func NewTLSConfig(manager *autocert.Manager) *tls.Config {
	return &tls.Config{
		// Получать сертификаты через autocert
		GetCertificate: manager.GetCertificate,

		// Минимальная версия TLS 1.2 (современный стандарт)
		MinVersion: tls.VersionTLS12,

		// Рекомендуемые cipher suites (ECDHE для forward secrecy)
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},

		// Предпочитать cipher suites сервера
		PreferServerCipherSuites: true,
	}
}
