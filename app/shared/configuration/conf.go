package configuration

import (
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"strings"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/joho/godotenv"
)

type Conf struct {
	PORT                        string `required:"true"`
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
		COUNTRY:                     strings.ToUpper(os.Getenv("COUNTRY")),
		PROJECT_NAME:                os.Getenv("PROJECT_NAME"),
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

	return conf, validateConfig(conf)
}

func Values() Conf {
	return ioc.Get[Conf](NewConf)
}

func validateConfig(conf Conf) error {
	var validationErrors []error

	val := reflect.ValueOf(conf)
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i).String()

		requiredTag := field.Tag.Get("required")
		if requiredTag == "true" && value == "" {
			validationErrors = append(validationErrors, fmt.Errorf("%s is required but not set", field.Name))
		}
	}
	if len(validationErrors) > 0 {
		// Convert errors to strings
		var errorStrings []string
		for _, err := range validationErrors {
			errorStrings = append(errorStrings, err.Error())
		}
		return fmt.Errorf("configuration errors:\n%s", strings.Join(errorStrings, "\n"))
	}

	return nil
}
