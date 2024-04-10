package configuration

import (
	"log/slog"
	"os"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/joho/godotenv"
)

type Conf struct {
	PORT                        string
	GEMINI_API_KEY              string
	GOOGLE_PROJECT_ID           string
	OTEL_EXPORTER_OTLP_ENDPOINT string
	DD_SERVICE                  string
	DD_ENV                      string
	DD_VERSION                  string
	DD_AGENT_HOST               string
}

func init() {
	ioc.Registry(NewConf)
}
func NewConf() (Conf, error) {
	if err := godotenv.Load(); err != nil {
		slog.Warn(".env not found, loading environment from system.")
	}
	conf := Conf{
		PORT:                        os.Getenv("PORT"),
		GEMINI_API_KEY:              os.Getenv("GEMINI_API_KEY"),
		GOOGLE_PROJECT_ID:           os.Getenv("GOOGLE_PROJECT_ID"),
		OTEL_EXPORTER_OTLP_ENDPOINT: os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"),
		DD_AGENT_HOST:               os.Getenv("DD_AGENT_HOST"),
		DD_SERVICE:                  os.Getenv("DD_SERVICE"),
		DD_ENV:                      os.Getenv("DD_ENV"),
		DD_VERSION:                  os.Getenv("DD_VERSION"),
	}

	if conf.DD_SERVICE != "" && conf.DD_ENV != "" &&
		conf.DD_VERSION != "" && conf.DD_AGENT_HOST != "" &&
		conf.OTEL_EXPORTER_OTLP_ENDPOINT != "" {
		conf.OTEL_EXPORTER_OTLP_ENDPOINT = conf.DD_AGENT_HOST + ":4317"
	}

	if conf.PORT == "" {
		conf.PORT = "8080"
	}
	return conf, nil
}

func Values() Conf {
	return ioc.Get[Conf](NewConf)
}
