package freespacecmd

import (
	"os"

	"github.com/Duansg/freespace/manager/base/conf"
	"github.com/Duansg/freespace/manager/base/constant"
	"github.com/Duansg/freespace/manager/cli"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman"
	"github.com/segmentfault/pacman/contrib/log/zap"
	"github.com/segmentfault/pacman/contrib/server/http"
	"github.com/segmentfault/pacman/log"
)

var (
	// Name is the name of the project
	Name = "freespace"
	// Version is the version of the project
	Version = "0.0.0"
	// GoVersion is the go version of the project
	GoVersion = "1.24.4"
	// log level
	logLevel = os.Getenv("LOG_LEVEL")
	// log path
	logPath = os.Getenv("LOG_PATH")
)

func Main() {
	log.SetLogger(zap.NewLogger(log.ParseLevel(logLevel), zap.WithName("freespace"), zap.WithPath(logPath)))
	Execute()
}

func runApp() {
	configPath := cli.GetConfigFilePath()
	c, err := conf.ReadConfig(configPath)
	if err != nil {
		panic(err)
	}
	app, cleanup, err := initApplication(
		c.Debug, c.Server, c.Data.Database, c.Data.Cache, c.I18n, c.Swaggerui, c.ServiceConfig, log.GetLogger())
	if err != nil {
		panic(err)
	}
	constant.Version = Version
	constant.GoVersion = GoVersion
	log.Info("Freesapce Version:", constant.Version, " GoVersion:", constant.GoVersion)
	defer cleanup()
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newApplication(serverConf *conf.Server, server *gin.Engine) *pacman.Application {
	log.Info("serverConf: %+v\n", serverConf)
	if serverConf.HTTP == nil {
		panic("serverConf.HTTP is nil")
	}
	log.Info("serverConf.HTTP.Addr: %s\n", serverConf.HTTP.Addr)
	return pacman.NewApp(
		pacman.WithName(Name),
		pacman.WithVersion(Version),
		pacman.WithServer(http.NewServer(server, serverConf.HTTP.Addr)),
	)
}
