package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"pluto/log"
	"pluto/util"
	"strings"
)

type GlobalConf struct {
	Mysql				MysqlConf
	Server				ServerConf
	TapDB				TapDBConf
	StatEventList		StatEventConf
	StatPropertyList	StatPropertyConf
}

var (
	conf 			*GlobalConf
	statEventMap	map[int]StatEvent
	statPropertyMap	map[int]StatProperty
)

func Init() {
	statEventMap = make(map[int]StatEvent, 20)
	statPropertyMap = make(map[int]StatProperty, 20)
}

func parseConfig() *GlobalConf {

	var (
		file string
		err error
	)

	// connect info.
	newConf := new(GlobalConf)

	file = "./module/stat/config/connect.toml"
	_, err = toml.DecodeFile(file, newConf)
	if err != nil {
		log.Error("parse connect config error! ", err.Error())
		return nil
	}

	file = "./module/stat/config/stat_event.toml"
	_, err = toml.DecodeFile(file, newConf)
	if err != nil {
		log.Error("parse stat event config error! ", err.Error())
		return nil
	}

	file = "./module/stat/config/stat_property.toml"
	_, err = toml.DecodeFile(file, newConf)
	if err != nil {
		log.Error("parse stat property config error! ", err.Error())
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

	events := config.StatEventList.List
	for _, e := range events {

		var (
			intNameList		[]string
			floatNameList	[]string
			strNameList		[]string
			intList			[]StatEventItem
			floatList		[]StatEventItem
			strList			[]StatEventItem
		)

		e.IntList = util.DeleteSpace(e.IntList)
		e.FloatList = util.DeleteSpace(e.FloatList)
		e.StrList = util.DeleteSpace(e.StrList)

		if (len(e.IntList)) != 0 {

			intNameList = strings.Split(e.IntList, ",")

			for _, name := range intNameList {

				itemType := CheckEvtItemType(&name)
				if itemType == StatEvtNone {
					itemType = StatEvtInt
				}

				intList = append(intList, StatEventItem{
					Name: name,
					Type: itemType,
				})

			}

		}

		if (len(e.FloatList)) != 0 {

			floatNameList = strings.Split(e.FloatList, ",")

			for _, name := range floatNameList {

				itemType := CheckEvtItemType(&name)
				if itemType == StatEvtNone {
					itemType = StatEvtFloat
				}

				floatList = append(floatList, StatEventItem{
					Name: name,
					Type: itemType,
				})

			}

		}

		if (len(e.StrList)) != 0 {

			strNameList = strings.Split(e.StrList, ",")

			for _, name := range strNameList {

				itemType := CheckEvtItemType(&name)
				if itemType == StatEvtNone {
					itemType = StatEvtString
				}

				strList = append(strList, StatEventItem{
					Name: name,
					Type: itemType,
				})

			}

		}

		statEventMap[e.ID] = StatEvent{
			ID:        e.ID,
			Name:      e.Name,
			IntList:   intList,
			FloatList: floatList,
			StrList:   strList,
		}

	}

	properties := config.StatPropertyList.List
	for _, p := range properties {

		statPropertyMap[p.ID] = p

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

func GetStatEvent(eid int) *StatEvent {

	se, ok := statEventMap[eid]

	if !ok {
		return nil
	}

	return &se
}

func GetStatProperty(pid int) *StatProperty {

	sp, ok := statPropertyMap[pid]

	if !ok {
		return nil
	}

	return &sp
}
