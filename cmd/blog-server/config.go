// Application configurations
package main

import (
	"os"
)

type Config struct {
	Port string
}

func LoadConfig() Config {
	return Config{
		Port: os.Getenv("PORT"),
	}
}
