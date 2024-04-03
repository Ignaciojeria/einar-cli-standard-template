package configuration

import (
	"log/slog"
	"os"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/joho/godotenv"
)

func init() {
	ioc.Registry(NewConf)
}

type Conf struct {
	PORT           string
	API_PREFIX     string
	GEMINI_API_KEY string
}

func NewConf() (Conf, error) {
	if err := godotenv.Load(); err != nil {
		slog.Warn(".env not found, loading environment from system.")
	}
	conf := Conf{
		PORT:           os.Getenv("PORT"),
		API_PREFIX:     os.Getenv("API_PREFIX"),
		GEMINI_API_KEY: os.Getenv("GEMINI_API_KEY"),
	}
	if conf.API_PREFIX == "" {
		conf.API_PREFIX = "/api"
	}
	if conf.PORT == "" {
		conf.PORT = "8080"
	}
	return conf, nil
}
