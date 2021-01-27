package config

type ServerInfo struct {
	Name			string
	Channel			int
	ServerID		int
	ServerType		int
	IsTest			int
	TimeZone		int
	Version			string
	Addr			string
	Status			int
	CreateTime		int64
	StartTime		int64
	CloseHint		string
}

type ServerListConf struct {
	List			[]ServerInfo
}