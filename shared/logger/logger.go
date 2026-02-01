package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func init() {
	logger = zerolog.New(os.Stdout)
}
func Info(message string) {
	logger.Info().Msg(message)
}
func Infof(format string, v ...interface{}) {
	logger.Info().Msgf(format, v...)
}
func Error(err error) {
	logger.Error().Err(err).Send()
}
func Debug(message string) {
	logger.Info().Msg(message)
}
func Trace(message string) {
	logger.Info().Msg(message)
}
