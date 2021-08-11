package app

import log "go-oo-demo/pkg/logger"

type Logger struct {
	config *LogConfig
	mode *Mode
	hook  *log.HookMysql
}

func newLogger(config *LogConfig, mode *Mode) *Logger {
	logger := new(Logger)
	logger.config = config
	logger.mode = mode
	return logger
}

func (logger *Logger) init() {
	logger.hook = log.NewHookMysql(
		logger.config.Database.Username,
		logger.config.Database.Password,
		logger.config.Database.Host,
		logger.config.Database.Name,
		logger.config.Database.MaxIdle,
		logger.config.Database.MaxOpen)

	log.AddHook(logger.hook)
	log.SetReportCaller(logger.config.ReportCaller)
}

func (logger *Logger) clear() {
	logger.hook.Clear()
}