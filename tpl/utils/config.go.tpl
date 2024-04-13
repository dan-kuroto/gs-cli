package utils

import "github.com/dan-kuroto/gin-stronger/gs"

type Configuration struct {
	gs.Configuration `yaml:",inline"`
	Message          string `yaml:"message"`
}

var Config Configuration
