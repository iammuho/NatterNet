// Package config aims to provide a config package
package config

import (
	"github.com/caarlos0/env"
)

func init() {
	Config = &config{}

	if err := env.Parse(&Config.Application); err != nil {
		panic(err)
	}
	if err := env.Parse(&Config.Logger); err != nil {
		panic(err)
	}
	if err := env.Parse(&Config.HTTPServer); err != nil {
		panic(err)
	}
	if err := env.Parse(&Config.MongoDB); err != nil {
		panic(err)
	}
	if err := env.Parse(&Config.JWT); err != nil {
		panic(err)
	}
	if err := env.Parse(&Config.Nats); err != nil {
		panic(err)
	}
}
