package serverwrapper

import (
	"archetype/app/shared/configuration"
	"archetype/app/shared/infrastructure/observability"
	logger "archetype/app/shared/logger"
	"log"
	"log/slog"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

type EchoWrapper struct {
	*echo.Echo
	conf configuration.Conf
}

func init() {
	ioc.Registry(echo.New)
	ioc.Registry(
		NewEchoWrapper,
		echo.New,
		configuration.NewConf,
		logger.NewLogger)
}

func NewEchoWrapper(
	e *echo.Echo,
	c configuration.Conf,
	l logger.Logger) EchoWrapper {
	e.Use(otelecho.Middleware(c.PROJECT_NAME))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			spanCtx, span := observability.Tracer.Start(c.Request().Context(), "RequestLogger")
			defer span.End()
			if v.Error == nil {
				l.SpanLogger(span).LogAttrs(spanCtx, slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				l.SpanLogger(span).LogAttrs(spanCtx, slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))
	return EchoWrapper{
		Echo: e,
	}
}

func Start() {
	ioc.Get[EchoWrapper](NewEchoWrapper).start()
}

func (s EchoWrapper) start() {
	s.printRoutes()
	s.Echo.Logger.Fatal(s.Echo.Start(":" + s.conf.PORT))
}

func (s EchoWrapper) printRoutes() {
	routes := s.Echo.Routes()
	for _, route := range routes {
		log.Printf("Method: %s, Path: %s, Name: %s\n", route.Method, route.Path, route.Name)
	}
}
