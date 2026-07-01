package config

import "os"

type Config struct {
	Port string
}

func NewConfig(port string) *Config {
	return &Config{
		Port: port,
	}
}

func (c *Config) LoadPort() *Config {

	port := os.Getenv("PORT")

	if port == "" {
		port = ":8080"
	}

	c.Port = port

	return c
}