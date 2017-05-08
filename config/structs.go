package config

import "time"

type configuration struct {
	Port        string
	RedisHost   string
	RedisPort   string
	Sources     []source
	ReadTimeout time.Duration
}

type source struct {
	Host       string
	Database   string
	Collection string
}
