package app
import (
	"lcng/util/logger"
)
type App struct {
	logger *logger.Logger
}
func New(logger *logger.Logger) *App {
	return &App{logger: logger}
}
func (app *App) Logger() *logger.Logger {
	return app.logger
}