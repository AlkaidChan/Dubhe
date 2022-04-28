package app

import (
	"github.com/alkaidchen/Dubhe/internal/config"
	"github.com/alkaidchen/Dubhe/internal/log"
	"github.com/alkaidchen/Dubhe/internal/router"
	dirUtils "github.com/alkaidchen/Dubhe/internal/utils/dir"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Dubhe application
type Dubhe struct {
	config *config.Config
	server *gin.Engine
	Home   string
	Logger *zap.Logger
}

func New() *Dubhe {
	return &Dubhe{}
}

// Run run blog application
func (d *Dubhe) Run() error {
	err := d.server.Run()
	if err != nil {
		return err
	}
	return nil
}

func (d *Dubhe) Init() error {
	for _, f := range []func() error{
		d.initHome,
		d.initLogger,
		//d.initConfig,
		//d.initDB,
		d.initRouter,
	} {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

func (d *Dubhe) initHome() error {
	home, err := dirUtils.GetOrCreateDubheHome()
	if err != nil {
		return err
	}
	d.Home = home
	return nil
}

func (d *Dubhe) initLogger() error {
	// TODO: set debug mode or not
	logger, err := log.New(true)
	if err != nil {
		return err
	}
	d.Logger = logger
	return nil
}

func (d *Dubhe) initRouter() error {
	server, err := router.NewRouter(d.Logger)
	if err != nil {
		log.Error("init router failed, ", err)
		return err
	}
	d.server = server
	return nil
}
