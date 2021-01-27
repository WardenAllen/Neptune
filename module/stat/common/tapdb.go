package common

type OnlineAry struct {
	Server			string			`json:"server"`
	Online			int				`json:"online"`
	Timestamp		int64			`json:"timestamp"`
}

type OnlineReq struct {
	AppID			string			`json:"appid"`
	Onlines			[]OnlineAry		`json:"onlines"`
}

type OnlineRes struct {
	Code			int				`json:"code"`
	Msg				string			`json:"msg"`
}

type PayProperty struct {
	OrderID			string			`json:"order_id"`
	Amount			int				`json:"amount"`
	ExtraGold		int				`json:"virtual_currency_amount"`
	CurrencyType	string			`json:"currency_type"`
	Product			string			`json:"product"`
	Payment			string			`json:"payment"`
}

type PayReq struct {
	Module			string			`json:"module"`
	IP				string			`json:"ip"`
	Name			string			`json:"name"`
	AppID			string			`json:"index"`
	Account			string			`json:"identify"`
	Properties		PayProperty		`json:"properties"`
}

type PayRes struct {
	Code			int				`json:"code"`
	Msg				string			`json:"msg"`
}
