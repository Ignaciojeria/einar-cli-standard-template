package firebasewrapper

import (
	"archetype/app/adapter/out/slog"
	"archetype/app/configuration"
	"archetype/app/constants"
	"context"

	firebase "firebase.google.com/go"
	ioc "github.com/Ignaciojeria/einar-ioc"
)

func init() {
	ioc.Registry(NewFirebaseAPP, configuration.NewConf)
}

func NewFirebaseAPP(conf configuration.Conf) (*firebase.App, error) {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: conf.GOOGLE_PROJECT_ID,
	})
	if err != nil {
		slog.Logger().Error("error initializing firebase app", constants.Error, err.Error())
		return nil, err
	}
	return app, nil
}
