package config

import "github.com/dora-exku/shorturl/pkg/config"

func init() {
	config.Add("database", config.MapStr{
		"mysql": map[string]interface{}{
			"host":     config.Env("DB_HOST", "localhost"),
			"port":     config.Env("DB_PORT", "3306"),
			"database": config.Env("DB_DATABASE", ""),
			"username": config.Env("DB_USERNAME", "root"),
			"password": config.Env("DB_PASSWORD", ""),
			"charset":  config.Env("DB_CHARSET", "utf8mb4"),
		},
	})
}
