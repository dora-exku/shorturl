package config

import "github.com/dora-exku/shorturl/pkg/config"

func init() {
	config.Add("app", config.MapStr{
		"env":   config.Env("APP_ENV", "local"),
		"host":  config.Env("APP_HOST", "0.0.0.0"),
		"port":  config.Env("APP_PORT", "8080"),
		"debug": config.Env("debug", true),
		"url":   config.Env("APP_URL", "http://127.0.0.1:8080"),
	})
}
