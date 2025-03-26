package utils

import "os"

type EnvVars struct {
	Port       string
	StaticPath string
}

func SetEnvVars() *EnvVars {
	envVars := EnvVars{
		Port:       os.Getenv("APP_PORT"),
		StaticPath: os.Getenv("STATIC_PATH"),
	}

	return &envVars
}
