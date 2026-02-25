package logger

import (
	"context"
	"log/slog"
	"os"
)


func New(env string) *slog.Logger{
	var handler slog.Handler
	if env == "production"{
		handler = slog.NewJSONHandler(os.Stdout, nil)
	}else{
		handler = slog.NewTextHandler(os.Stdout, nil)
	}
	return slog.New(handler)
}

func WithContext(ctx context.Context, log *slog.Logger) context.Context{
	return context.WithValue(ctx, "logger", log)
}

func FromContext(ctx context.Context) *slog.Logger{
	return ctx.Value("logger").(*slog.Logger)
}