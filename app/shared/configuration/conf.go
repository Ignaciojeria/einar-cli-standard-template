package configuration

import (
	"archetype/app/shared/constants"
	"log/slog"
	"os"
	"strings"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/joho/godotenv"
)

type Conf struct {
	PORT                        string `required:"true"`
	VERSION                     string `required:"true"`
	COUNTRY                     string `required:"true"`
	GEMINI_API_KEY              string `required:"false"`
	PROJECT_NAME                string `required:"false"`
	GOOGLE_PROJECT_ID           string `required:"false"`
	OTEL_EXPORTER_OTLP_ENDPOINT string `required:"false"`
	DD_SERVICE                  string `required:"false"`
	DD_ENV                      string `required:"false"`
	DD_VERSION                  string `required:"false"`
	DD_AGENT_HOST               string `required:"false"`
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
		VERSION:                     os.Getenv(constants.Version),
		COUNTRY:                     strings.ToUpper(os.Getenv("COUNTRY")),
		PROJECT_NAME:                os.Getenv("PROJECT_NAME"),
		GEMINI_API_KEY:              os.Getenv("GEMINI_API_KEY"),
		GOOGLE_PROJECT_ID:           os.Getenv("GOOGLE_PROJECT_ID"),
		OTEL_EXPORTER_OTLP_ENDPOINT: os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"),
	}
	setupDatadog(&conf)
	if conf.DD_SERVICE != "" && conf.DD_ENV != "" &&
		conf.DD_VERSION != "" && conf.DD_AGENT_HOST != "" &&
		conf.OTEL_EXPORTER_OTLP_ENDPOINT != "" {
		conf.OTEL_EXPORTER_OTLP_ENDPOINT = conf.DD_AGENT_HOST + ":4317"
	}

	if conf.PORT == "" {
		conf.PORT = "8080"
	}

	return validateConfig(conf)
}

func setupDatadog(c *Conf) {
	os.Setenv("DD_SERVICE", c.PROJECT_NAME)
	c.DD_SERVICE = c.PROJECT_NAME
	if os.Getenv("DD_ENV") == "" {
		os.Setenv("DD_ENV", "unknown")
	}
	c.DD_ENV = os.Getenv("DD_ENV")
	c.DD_AGENT_HOST = os.Getenv("DD_AGENT_HOST")
	os.Setenv("DD_VERSION", c.VERSION)
	c.DD_VERSION = c.VERSION
}
