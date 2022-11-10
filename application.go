package application

import (
	"github.com/whling/go-config"
)

type Application struct {
	Name   string
	Config *config.TomlValue
}
