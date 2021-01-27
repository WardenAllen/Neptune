package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"pluto/log"
	"runtime"
	"sync"
)

var MySQLConns []*sql.DB
var MysqlMutex sync.Mutex

const (
	connStr = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=Local"
)

func InitMysql(uname string, pwd string, host string, port string, dbname string) bool {

	name := fmt.Sprintf(connStr, uname, pwd, host, port, dbname)

	for i := 0; i <= runtime.NumCPU(); i++ {

		MySQL, err := sql.Open("mysql", name)
		if err != nil {
			log.Error("MySQL connect error ", err)
			return false
		}

		err = MySQL.Ping()
		if err != nil {
			log.Warn("MySQL ping error ", err)
			return false
		}

		MySQLConns = append(MySQLConns, MySQL)
	}

	log.Info("MySQL init success.")

	return true
}

func LockMysql() {
	MysqlMutex.Lock()
}

func UnlockMysql() {
	MysqlMutex.Unlock()
}
