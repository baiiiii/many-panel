package viper

import (
	"github.com/1Panel-dev/1Panel/backend/configs"
	"github.com/1Panel-dev/1Panel/backend/global"
	"github.com/1Panel-dev/1Panel/cmd/server/conf"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"path"
)

func Init() {
	baseDir := "/opt"
	port := "9999"
	version := "v1.0.0"
	username, password, entrance := "", "", ""
	v := viper.NewWithOptions()
	v.SetConfigType("yaml")
	serverConfig := configs.ServerConfig{}
	if err := yaml.Unmarshal(conf.AppYaml, &serverConfig); err != nil {
		panic(err)
	}
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(&global.CONF); err != nil {
			panic(err)
		}
	})

	if serverConfig.System.BaseDir != "" {
		baseDir = serverConfig.System.BaseDir
	}
	if serverConfig.System.Port != "" {
		port = serverConfig.System.Port
	}
	if serverConfig.System.Version != "" {
		version = serverConfig.System.Version
	}
	if serverConfig.System.Username != "" {
		username = serverConfig.System.Username
	}
	if serverConfig.System.Password != "" {
		password = serverConfig.System.Password
	}
	if serverConfig.System.Entrance != "" {
		entrance = serverConfig.System.Entrance
	}

	global.CONF = serverConfig
	global.CONF.System.BaseDir = baseDir
	global.CONF.System.IsDemo = v.GetBool("system.is_demo")
	global.CONF.System.DataDir = path.Join(global.CONF.System.BaseDir, "many-panel")
	global.CONF.System.Cache = path.Join(global.CONF.System.DataDir, "cache")
	global.CONF.System.Backup = path.Join(global.CONF.System.DataDir, "backup")
	global.CONF.System.DbPath = path.Join(global.CONF.System.DataDir, "db")
	global.CONF.System.LogPath = path.Join(global.CONF.System.DataDir, "log")
	global.CONF.System.TmpDir = path.Join(global.CONF.System.DataDir, "tmp")
	global.CONF.System.Port = port
	global.CONF.System.Version = version
	global.CONF.System.Username = username
	global.CONF.System.Password = password
	global.CONF.System.Entrance = entrance
	global.CONF.System.ChangeUserInfo = ""
	global.Viper = v
}
