package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"pluto/log"
	Proto "pluto/protobuf"
	"pluto/util"
	"strconv"
	"strings"
)

type GlobalConf struct {
	Redis 			*RedisConf
	Mysql 			*MysqlConf
	Server 			*ServerConf
	CDKeyAward		*CDKeyAwardConf
	CDKeyGeneral	*CDKeyGeneralConf
}

var (
	conf 			*GlobalConf
	cdkeyAwardMap	map[int]CDKeyAward
	cdkeyGeneralMap	map[string]CDKeyGeneralInfo
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
	file = "./module/gm/config/connect.toml"
	_, err = toml.DecodeFile(file, newConf)
	if err != nil {
		log.Error("parse connect config error! ", err.Error())
		return nil
	}

	file = "./module/gm/config/cdkey_award.toml"
	_, err = toml.DecodeFile(file, newConf)
	if err != nil {
		log.Error("parse cdkey award config error! ", err.Error())
		return nil
	}

	file = "./module/gm/config/cdkey_general.toml"
	_, err = toml.DecodeFile(file, newConf)
	if err != nil {
		log.Error("parse cdkey general config error! ", err.Error())
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

	cdkeyAwardMap = make(map[int]CDKeyAward, 20)
	cdkeyGeneralMap = make(map[string]CDKeyGeneralInfo, 20)

	awards := config.CDKeyAward.Award
	for _, e := range awards {

		var (
			err			error
			itemId		int
			itemNum		int
			items		[]string
			itemData	[]string
			awards		[]*Proto.AwardItem
		)

		e.Award = util.DeleteSpace(e.Award)

		if (len(e.Award)) != 0 {

			items = strings.Split(e.Award, ",")

			for _, t := range items {

				itemData = strings.Split(t, ";")

				if len(itemData) != 2 {
					log.Error("Invalid cdkey award %d", e.ID)
					continue
				}

				itemId, err = strconv.Atoi(itemData[0])
				if err != nil {
					log.Error("Invalid cdkey award %d", e.ID)
					continue
				}

				itemNum, err = strconv.Atoi(itemData[1])
				if err != nil {
					log.Error("Invalid cdkey award %d", e.ID)
					continue
				}

				awards = append(awards, &Proto.AwardItem{
					ItemID:  int32(itemId),
					ItemNum: int32(itemNum),
				})

			}

		}

		cdkeyAwardMap[e.ID] = CDKeyAward{
			ID:      e.ID,
			Channel: e.Channel,
			Type:    e.Type,
			Award:   awards,
		}

	}

	keys := config.CDKeyGeneral.Key
	for _, k := range keys {

		k.CDKey = util.DeleteSpace(k.CDKey)

		cdkeyGeneralMap[k.CDKey] = k

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

func GetCdkeyAward(id int) *CDKeyAward {

	award, ok := cdkeyAwardMap[id]

	if !ok {
		return nil
	}

	return &award
}

func GetCdkeyGenInfo(key string) *CDKeyGeneralInfo {

	info, ok := cdkeyGeneralMap[key]

	if !ok {
		return nil
	}

	return &info
}

