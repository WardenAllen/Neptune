package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"pluto/log"
)

type GlobalConf struct {
	Redis 		*RedisConf
	Mysql 		*MysqlConf
	Server 		*ServerConf
	NoticeList	*NoticeListConf
	ServerList	*ServerListConf
}

var (
	conf 		*GlobalConf
	serverMap	map[int]ServerInfo
	noticeMap	map[int]NoticeInfo
)

func Init()  {
	serverMap = make(map[int]ServerInfo, 10)
	noticeMap = make(map[int]NoticeInfo, 10)
}

func parseConfig() *GlobalConf {

	var (
		file string
		err error
	)

	// connect info.
	newConf := new(GlobalConf)
	file = "./module/list/config/connect.toml"
	_, err = toml.DecodeFile(file, newConf)
	if err != nil {
		log.Error("parse connect config error! ", err.Error())
		return nil
	}
	// server-list info.
	file = "./module/list/config/server_list.toml"
	_, err = toml.DecodeFile(file, newConf)
	if err != nil {
		log.Error("parse server_list config error! ", err.Error())
		return nil
	}

	// notice info.
	file = "./module/list/config/notice.toml"
	_, err = toml.DecodeFile(file, newConf)
	if err != nil {
		log.Error("parse notice config error! ", err.Error())
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

	list := config.ServerList.List
	for i := 0; i < len(list); i++ {
		serverMap[int(list[i].ServerID)] = list[i]
	}

	nl := config.NoticeList.List
	for i := 0; i < len(nl); i++ {
		noticeMap[nl[i].Channel] = nl[i]
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

func GetServer(sid int) (addr string) {

	s, ok := serverMap[sid]

	if !ok {
		return
	}

	addr = s.Addr

	return
}

func GetNotice(channel int) (notice string) {

	n, ok := noticeMap[channel]

	if !ok {
		return
	}

	notice = n.Content

	return
}