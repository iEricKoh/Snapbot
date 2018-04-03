package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Specification struct {
	TELEGRAM_TOKEN string
}

var Config Specification

func init() {
	envconfig.Process("", &Config)
}
