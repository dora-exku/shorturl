package config

import "github.com/dora-exku/shorturl/pkg/config"

func init() {
	config.Add("jwt", config.MapStr{
		"secret": config.Env("JWT_SECRET"),
	})
}
