package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-oo-demo/pkg/orm"
	"time"
)

type HookMysql struct {
	orm *orm.Orm
}

func NewHookMysql(username string, password string, host string, name string, maxIdle int, maxOpen int) *HookMysql {
	hook := new(HookMysql)
	hook.orm = orm.New(username, password, host, name, maxIdle, maxOpen, false)
	return hook
}

func (hook *HookMysql) Fire(entry *logrus.Entry) error {

	logger := &Logger{
		Level:     entry.Level.String(),
		Message:   entry.Message,
		CreatedAt: entry.Time,
	}

	if entry.HasCaller() {
		logger.File = fmt.Sprintf("%s line:%d func:%s",
			entry.Caller.File,
			entry.Caller.Line,
			entry.Caller.Function)
	}

	data := entry.Data
	if v, ok := data[TagKey]; ok {
		logger.Tag, _ = v.(string)
		delete(data, TagKey)
	}
	if v, ok := data[VersionKey]; ok {
		logger.Version, _ = v.(string)
		delete(data, VersionKey)
	}

	return hook.orm.DB.Save(logger).Error
}

func (hook *HookMysql) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *HookMysql) Clear() {
	hook.orm.Clear()
}

type Logger struct {
	Id         int
	Level      string
	TraceId    string
	UserId     string
	Tag        string
	Version    string
	Message    string
	Data       string
	File       string
	ErrorStack string
	CreatedAt  time.Time
}
