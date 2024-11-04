package logging

import (
	"log/slog"
	"os"
)

const (
	envProd  = "production"
	envDev   = "development"
	envLocal = "local"
)

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "Error",
		Value: slog.StringValue(err.Error()),
	}
}

func SetupLogger(envType string) (log *slog.Logger) {
	switch envType {
	case envProd:
		{
			log = slog.New(
				slog.NewJSONHandler(
					os.Stdout, &slog.HandlerOptions{
						Level: slog.LevelWarn,
					},
				),
			)
		}
	case envDev:
		{
			log = slog.New(
				slog.NewJSONHandler(
					os.Stdout, &slog.HandlerOptions{
						Level: slog.LevelInfo,
					},
				),
			)
		}
	case envLocal:
		fallthrough
	default:
		{
			log = slog.New(
				slog.NewTextHandler(
					os.Stdout, &slog.HandlerOptions{
						Level: slog.LevelDebug,
					},
				),
			)
		}
	}
	return
}
