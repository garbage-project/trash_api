package network

import (
	"fmt"
	"time"

	util "github.com/04Akaps/go-util/log"
	"github.com/garbage-project/trash_api.git/config"
	"github.com/garbage-project/trash_api.git/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Network struct {
	cfg     *config.Config
	engine  *gin.Engine
	log     *util.Log
	service *service.Service
	port    string
}

func NewNetwork(cfg *config.Config, log *util.Log, service *service.Service) *Network {
	n := &Network{
		engine:  gin.New(),
		port:    fmt.Sprintf(":%s", cfg.ServerInfo.Port),
		log:     log,
		cfg:     cfg,
		service: service,
	}

	n.engine.Use(gin.Logger())
	n.engine.Use(gin.Recovery())
	n.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"ORIGIN", "Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "expires"},
		ExposeHeaders:    []string{"ORIGIN", "Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "expires"},
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           12 * time.Hour,
	}))

	n.registerRouter()

	return n
}

func (r *Network) Run() error {
	r.log.InfoLog("Http Server Start", "endpoint", r.port)
	return r.engine.Run(r.port)
}
