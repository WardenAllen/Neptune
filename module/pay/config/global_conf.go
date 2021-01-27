package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"pluto/log"
)

type GlobalConf struct {
	Mysql	MysqlConf
	Server	*ServerConf
	NASDK	NASDKConf
}

var (
	conf 		*GlobalConf
)

func Init()  {
}

func parseConfig() *GlobalConf {

	var (
		file string
		err error
	)

	// connect info.
	newConf := new(GlobalConf)
	file = "./module/pay/config/connect.toml"
	_, err = toml.DecodeFile(file, newConf)
	if err != nil {
		log.Error("parse connect config error! ", err.Error())
		return nil
	}

	return newConf
}

func LoadConfig() bool {

	config := parseConfig()

	if config == nil {
		log.Error("Load config failed.")
		return false
	}

	conf = config
	log.Info("Load config success.")

	return true
}

func Config() *GlobalConf {

	if conf == nil {

		if !LoadConfig() {
			os.Exit(1)
		}

	}

	return conf
}
