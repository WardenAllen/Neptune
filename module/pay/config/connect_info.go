package config

type MysqlConf struct {
	Host     	string
	Port     	string
	Username 	string
	Password 	string
	DBName   	string
}

type ServerConf struct {
	HttpHost	string
	TcpHost		string
}

type NASDKConf struct {
	AppID		string
	SecretKey	string
}
