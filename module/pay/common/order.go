package common

import (
	"errors"
	"pluto/database"
	"pluto/log"
	"time"
)

type Order struct {
	OrderID		string			// 订单ID
	UserID		uint64			// 用户ID
	PayID		int				// 支付ID
	State		int				// 状态 1-下单 2-支付完成 3-发放完成
	Amount		int				// 充值金额(分)
	ServerID	int				// 服务器ID
	ChannelID	int				// 渠道ID
	OrderTime	int64			// 下单时间
}

func (order *Order) Insert(idx int) error {

	rs, err := database.MySQLConns[idx].Exec(
		"INSERT INTO order (order_id, user_id, pay_id, state, amount, server_id, channel_id, order_time) VALUE (?,?,?,?,?,?,?,?)",
		order.OrderID, order.UserID, order.PayID, order.State, order.Amount, order.ServerID, order.ChannelID, order.OrderTime)

	if err != nil {
		log.Debug("Insert order to db error %s.", order.OrderID)
		return err
	}

	if rs != nil {
		id, _ := rs.LastInsertId()
		log.Debug("Insert order id %d %s.", id, time.Now().Format("2006-01-02 15:04:05"))
	} else {
		return errors.New("Order " + order.OrderID + "exist in db")
	}

	return nil
}

func UpdateState(idx int, orderId string, state int) {

	rs, err := database.MySQLConns[idx].Exec("UPDATE order SET state=? WHERE order_id=?", state, orderId)

	if err != nil {
		log.Error("Update order state error %s", orderId)
		return
	}

	if rs == nil {
		return
	}

	num, err := rs.RowsAffected()

	if num == 0 {
		//log.Info("No rows update!")
	}

}