package config

import (
	"os"
)

// Config represents configuration options for the app.
type Config struct {
	Secret           []byte
	Host             string
	Port             string
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPass     string
	PostgresDatabase string
}

// FromEnv creates and returns a configuration object from the environment.
func FromEnv() *Config {
	return &Config{
		Secret:           []byte(os.Getenv("BEANSTALK_SECRET")),
		Host:             os.Getenv("BEANSTALK_HOST"),
		Port:             os.Getenv("BEANSTALK_PORT"),
		PostgresHost:     os.Getenv("BEANSTALK_POSTGRESHOST"),
		PostgresPort:     os.Getenv("BEANSTALK_POSTGRESPORT"),
		PostgresUser:     os.Getenv("BEANSTALK_POSTGRESUSER"),
		PostgresPass:     os.Getenv("BEANSTALK_POSTGRESPASS"),
		PostgresDatabase: os.Getenv("BEANSTALK_POSTGRESDB"),
	}
}
