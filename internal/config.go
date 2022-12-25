package config

import (
	"log"

	env "github.com/Netflix/go-env"
	_ "github.com/joho/godotenv/autoload"
)

type Environment struct {
	DbFile      string `env:"DB_FILE"`
	DbPort      int    `env:"DB_PORT"`
	DbHost      string `env:"DB_HOST"`
	DbUser      string `env:"DB_USER"`
	DbPass      string `env:"DB_PASS"`
	DbName      string `env:"DB_NAME"`
	PORT        int    `env:"PORT"`
	HOST        string `env:"HOST,default=0.0.0.0"`
	SwaggerHost string `env:"SWAGGER_HOST,default=127.0.0.1"`
}

var Env Environment

func init() {
	_, err := env.UnmarshalFromEnviron(&Env)
	if err != nil {
		log.Fatal(err)
	}
}
