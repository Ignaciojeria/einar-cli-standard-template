package firebasewrapper

import (
	"archetype/app/shared/configuration"
	"archetype/app/shared/constants"
	"archetype/app/shared/slog"

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
