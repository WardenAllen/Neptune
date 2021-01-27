package config

type MysqlConf struct {
	Host     	string
	Port     	string
	Username 	string
	Password 	string
	DBName   	string
}

type RedisConf struct {
	Host		string
	Port		string
	Password	string
}

type ServerConf struct {
	HttpHost	string
	TcpHost		string
}