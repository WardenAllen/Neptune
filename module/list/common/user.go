package common

import (
	"errors"
	"pluto/database"
	"pluto/log"
	"time"
)

type User struct {
	UserID		uint64
	Account		string
	RoleName	string
	Channel		int
	ServerID	int
	Level		int
	Icon		int
	RegTime		int64
	ActiveTime	int64
}

func (user *User) Insert(idx int) error {

	t := time.Unix(user.RegTime, 0)

	rs, err := database.MySQLConns[idx].Exec(
		"INSERT into users (account, user_id, role_name, icon, channel, server_id, reg_time, active_time) value (?,?,?,?,?,?,?,?)",
		user.Account, user.UserID, user.RoleName, user.Icon, user.Channel, user.ServerID, t, t)

	if err != nil {
		log.Debug("Insert user to db error %d", user.UserID)
		return err
	}

	if rs != nil {
		id, _ := rs.LastInsertId()
		log.Debug("Insert user id %d %s.", id, time.Now().Format("2006-01-02 15:04:05"))
	} else {
		return errors.New("User " + string(user.UserID) + "exist in db")
	}

	return nil
}

func (user *User) UpdateLevel(idx int, level int, icon int) {

	t := time.Now()

	rs, err := database.MySQLConns[idx].Exec(
		"UPDATE users SET level=?,icon=?,active_time=? WHERE account=? and server_id=?",
		level, icon, t, user.Account, user.ServerID)

	if err != nil {
		log.Error("Update user level error %d", user.UserID)
		return
	}

	if rs == nil {
		return
	}

	num, err := rs.RowsAffected()

	if num == 0 {
		//log.Info("No rows update!")
	}

	user.Level = level
	user.Icon = icon
	user.ActiveTime = t.Unix()
}

func (user *User) UpdateRoleName(idx int, rolename string) {

	t := time.Now()

	rs, err := database.MySQLConns[idx].Exec(
		"UPDATE users SET role_name=?,active_time=? WHERE account=? and server_id=?",
		rolename, t, user.Account, user.ServerID)

	if err != nil {
		log.Error("Update user role name error %d", user.UserID)
		return
	}

	if rs == nil {
		return
	}

	num, err := rs.RowsAffected()

	if num == 0 {
		//log.Info("No rows update!")
	}

	user.RoleName = rolename
	user.ActiveTime = t.Unix()
}

func UserWLock()  {
	UsersMutex.Lock()
}

func UserWUnlock()  {
	UsersMutex.Unlock()
}

func UserRLock()  {
	UsersMutex.RLock()
}

func UserRUnlock()  {
	UsersMutex.RUnlock()
}
