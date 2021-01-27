package config

import Proto "pluto/protobuf"

type CDKeyAwardInfo struct {
	ID				int
	Channel			int
	Type			int
	Award			string
}

type AwardItem struct {
	ItemID			int
	ItemNum			int
}

type CDKeyAward struct {
	ID				int
	Channel			int
	Type			int
	Award			[]*Proto.AwardItem
}

type CDKeyAwardConf struct {
	Award			[]CDKeyAwardInfo
}
