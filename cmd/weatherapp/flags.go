package main

import (
	"flag"
	"log"
	"weatherapp/internal/config"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var envVars = config.EnvVars{}

func init() {
	var configPath string
	flag.StringVar(&configPath, "config", ".env", "path to .env file")

	err := godotenv.Load(configPath)
	if err != nil {
		log.Println(".env file not found")
	}

	err = envconfig.Process("vp_splinter", &envVars)
	if err != nil {
		log.Fatal(err.Error())
	}
}
