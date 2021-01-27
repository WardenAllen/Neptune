package config

type CDKeyGeneralInfo struct {
	ID				int
	Channel			int
	AwardID			int
	CDKey			string
}

type CDKeyGeneralConf struct {
	Key			[]CDKeyGeneralInfo
}
