package configuration

import (
	"os"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

type StorjConfiguration struct {
	STORJ_ACCESS_GRANT string `required:"true"`
}

func init() {
	ioc.Registry(NewStorjConfiguration)
}
func NewStorjConfiguration() (StorjConfiguration, error) {
	conf := StorjConfiguration{
		STORJ_ACCESS_GRANT: os.Getenv("STORJ_ACCESS_GRANT"),
	}
	return validateConfig(conf)
}
