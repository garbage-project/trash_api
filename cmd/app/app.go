package app

import (
	"os"
	"os/signal"
	"syscall"

	util "github.com/04Akaps/go-util/log"
	"github.com/garbage-project/trash_api.git/config"
	"github.com/garbage-project/trash_api.git/network"
	"github.com/garbage-project/trash_api.git/repository"
	"github.com/garbage-project/trash_api.git/service"
)

type App struct {
	config *config.Config
	log    *util.Log
	stop   chan struct{}

	network    *network.Network
	repository *repository.Repository
	service    *service.Service
}

func NewApp(cfg *config.Config) *App {
	app := &App{
		config: cfg,
		stop:   make(chan struct{}),
	}

	app.log = util.SetLog(cfg.ServerInfo.Log)

	app.repository = repository.NewRepository(cfg, app.log)
	app.service = service.NewService(cfg, app.log, app.repository)
	app.network = network.NewNetwork(cfg, app.log, app.service)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		<-c
		app.Stop()
	}()

	return app
}

func (a *App) Run() {
	a.network.Run()
}

func (a *App) Wait() {
	a.log.InfoLog("Server started")
	<-a.stop
	os.Exit(1)
}

func (a *App) Stop() {
	a.log.InfoLog("Server stopped")
	a.stop <- struct{}{}
}
