package config

import "os"

var (
	Port       = os.Getenv("APP_PORT")
	StaticPath = os.Getenv("STATIC_PATH")
)
