package server

import (
	"encoding/gob"
	"fmt"
	"github.com/1Panel-dev/1Panel/backend/global"
	"github.com/1Panel-dev/1Panel/backend/i18n"
	"github.com/1Panel-dev/1Panel/backend/init/cache"
	"github.com/1Panel-dev/1Panel/backend/init/db"
	"github.com/1Panel-dev/1Panel/backend/init/log"
	"github.com/1Panel-dev/1Panel/backend/init/migration"
	"github.com/1Panel-dev/1Panel/backend/init/router"
	"github.com/1Panel-dev/1Panel/backend/init/session"
	"github.com/1Panel-dev/1Panel/backend/init/session/psession"
	"github.com/1Panel-dev/1Panel/backend/init/validator"
	"github.com/1Panel-dev/1Panel/backend/init/viper"
	"net"
	"net/http"
)

func Start() {
	viper.Init()
	i18n.Init()
	log.Init()
	db.Init()
	migration.Init()
	// app.Init()
	validator.Init()
	gob.Register(psession.SessionUser{})
	cache.Init()
	session.Init()
	// cron.Run()
	// business.Init()
	// hook.Init()
	// InitOthers()

	rootRouter := router.Routers()

	tcpItem := "tcp4"
	if global.CONF.System.Ipv6 == "enable" {
		tcpItem = "tcp"
		global.CONF.System.BindAddress = fmt.Sprintf("[%s]", global.CONF.System.BindAddress)
	}
	server := &http.Server{
		Addr:    global.CONF.System.BindAddress + ":" + global.CONF.System.Port,
		Handler: rootRouter,
	}
	ln, err := net.Listen(tcpItem, server.Addr)
	if err != nil {
		panic(err)
	}
	type tcpKeepAliveListener struct {
		*net.TCPListener
	}
	global.LOG.Infof("listen at http://%s:%s [%s]", global.CONF.System.BindAddress, global.CONF.System.Port, tcpItem)
	if err := server.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)}); err != nil {
		panic(err)
	}
}
