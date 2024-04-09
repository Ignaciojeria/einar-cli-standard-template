package configuration

import (
	"log/slog"
	"os"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/joho/godotenv"
)

type Conf struct {
	PORT              string
	GEMINI_API_KEY    string
	GOOGLE_PROJECT_ID string
	DD_SERVICE        string
	DD_ENV            string
	DD_VERSION        string
}

func init() {
	ioc.Registry(NewConf)
}
func NewConf() (Conf, error) {
	if err := godotenv.Load(); err != nil {
		slog.Warn(".env not found, loading environment from system.")
	}
	conf := Conf{
		PORT:              os.Getenv("PORT"),
		GEMINI_API_KEY:    os.Getenv("GEMINI_API_KEY"),
		GOOGLE_PROJECT_ID: os.Getenv("GOOGLE_PROJECT_ID"),
	}
	if conf.PORT == "" {
		conf.PORT = "8080"
	}
	return conf, nil
}

func Values() Conf {
	return ioc.Get[Conf](NewConf)
}
