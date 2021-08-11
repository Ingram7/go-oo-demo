package app

import (
	"context"
	"go-oo-demo/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Engine struct {
	config *Config
	mode *Mode
	db     *Database
	logger *Logger
}

func (engine *Engine) Start(ctx context.Context) {
	engine.loadConfig()
	engine.initMode()
	engine.initLogger()
	engine.initDB()

	engine.run(ctx)
}

func (engine *Engine) loadConfig() {
	engine.config = newConfig()
	engine.config.load()
	//engine.config.print()
}

func (engine *Engine) initMode() {
	engine.mode = newMode(engine.config.mode())
}

func (engine *Engine) initLogger() {
	engine.logger = newLogger(engine.config.logConfig(), engine.mode)
	engine.logger.init()
}

func (engine *Engine) initDB() {
	engine.db = newDatabase(engine.config.dbConfig(), engine.mode)
	engine.db.init()
}

// Run 运行服务
func (engine *Engine) run(ctx context.Context) error {
	state := 1
	sc := make(chan os.Signal, 1)
	// SIGHUP
	// SIGINT 程序终止(interrupt)信号, 在用户键入INTR字符(通常是Ctrl+C)时发出，用于通知前台进程组终止进程
	// SIGTERM 礼貌地要求进程终止.它将正常终止,清理所有资源(文件,套接字,子进程等),删除临时文件等
	// SIGQUIT
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)


	srv := newServer(engine.config.serverConfig(), engine.db.db(), engine.config.mode())
	srv.run(ctx)

EXIT:
	for {
		sig := <-sc
		logger.WithContext(ctx).Infof("接收到信号[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	logger.WithContext(ctx).Infof("服务退出")

	srv.clear(ctx)
	engine.db.clear()
	engine.logger.clear()


	time.Sleep(time.Second)
	os.Exit(state)
	return nil
}

