package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
)

type ILogger interface {
	Info(msg string)
	Error(msg string)
	Debug(msg string)
	Trace(msg string)
	Fatal(msg string)
}

type ZeroLoggerAdapter struct {
	Logger zerolog.Logger
}

func NewZeroLoggerAdapter() *ZeroLoggerAdapter {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		_, filename := filepath.Split(file)
		return fmt.Sprintf("%s:%d", filename, line)
	}
	zlogger := zerolog.New(os.Stdout).With().CallerWithSkipFrameCount(3).Timestamp().Logger()
	return &ZeroLoggerAdapter{
		Logger: zlogger,
	}
}

func (z ZeroLoggerAdapter) Info(msg string) {
	z.Logger.Info().Msg(msg)
}

func (z ZeroLoggerAdapter) Error(msg string) {
	z.Logger.Error().Msg(msg)
}

func (z ZeroLoggerAdapter) Debug(msg string) {
	z.Logger.Debug().Msg(msg)
}

func (z ZeroLoggerAdapter) Trace(msg string) {
	z.Logger.Trace().Msg(msg)
}

func (z ZeroLoggerAdapter) Fatal(msg string) {
	z.Logger.Fatal().Msg(msg)
}
