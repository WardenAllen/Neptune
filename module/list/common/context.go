package common

import (
	"pluto/database"
	"pluto/log"
	"sync"
	"time"
)

type Context struct {
	
}

var (
	UsersMap		map[string] map[int]*User	// <account, <serverid, User>>
	UsersMutex		sync.RWMutex
	UserSyncIdx 	int
)

func Init()  {

	var (
		t		time.Time
	)

	UsersMap = make(map[string] map[int]*User, 10000)

	// load users from db

	rows, err := database.MySQLConns[0].Query(
		"SELECT user_id, account, channel, server_id, level, reg_time, active_time FROM users")

	if err != nil {
		log.Error("Load user from data failed!")
		return
	}

	if rows == nil {
		return
	}

	for rows.Next() {

		u := &User{}

		var reg, active string

		err := rows.Scan(&u.UserID, &u.Account, &u.Channel, &u.ServerID, &u.Level, &reg, &active)
		if err != nil {
			log.Error("Load user from data failed %s!", err.Error())
			return
		}

		t, err = time.ParseInLocation("2006-01-02 15:04:05", reg, time.Local)
		if err == nil {
			u.RegTime = t.Unix()
		}

		t, err = time.ParseInLocation("2006-01-02 15:04:05", active, time.Local)
		if err == nil {
			u.ActiveTime = t.Unix()
		}

		InsertUser(u)
	}

	_ = rows.Close()

	log.Info("Load user from database success.")

}

func InsertUser(u *User) bool {

	var ok bool

	// check whether exist.

	_, ok = UsersMap[u.Account]

	if ok {

		_, ok = UsersMap[u.Account][u.ServerID]

		if ok {

			log.Debug("Insert user failed, user %s already exist.", u.Account)

			return false

		}

	} else {

		UsersMap[u.Account] = make(map[int]*User, 1)

	}

	UsersMap[u.Account][u.ServerID] = u

	return true

}

func IsUserExist(account string, sid int, init bool) bool {

	var ok bool

	_, ok = UsersMap[account]

	if ok {

		_, ok = UsersMap[account][sid]

		if ok {
			return true
		}

	} else {

		if init {
			UsersMap[account] = make(map[int]*User, 1)
		}

	}

	return false

}