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
	GEMINI_API_KEY string
}

func NewConf() (Conf, error) {
	if err := godotenv.Load(); err != nil {
		slog.Warn(".env not found, loading environment from system.")
	}
	conf := Conf{
		PORT:           os.Getenv("PORT"),
		GEMINI_API_KEY: os.Getenv("GEMINI_API_KEY"),
	}
	if conf.PORT == "" {
		conf.PORT = "8080"
	}
	return conf, nil
}
