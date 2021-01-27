package ta

import (
	"errors"
	"github.com/ThinkingDataAnalytics/go-sdk/thinkingdata"
	"pluto/log"
	"pluto/module/stat/config"
	Proto "pluto/protobuf"
	"pluto/util"
	"strconv"
)

var (
	ta	thinkingdata.TDAnalytics
)

func init() {

	var (
		err			error
		consumer	thinkingdata.Consumer
	)

	// 创建按天切分的 LogConsumer，并设置单个日志文件上限为 10 G
	consumer, err = thinkingdata.NewLogConsumerWithFileSize("./data", thinkingdata.ROTATE_DAILY, 10 * 1024)

	if err != nil {
		log.Error("TA consumer init error %s", err)
	}

	ta = thinkingdata.New(consumer)

}

func TrackEvent(account string, distinct string, ip string,
	uid uint64, id Proto.TA_ID, t int64,
	intList []int64, floatList []float64, strList []string) error {

	var (
		err 		error
		se			*config.StatEvent
		ptys		map[string]interface{}
		accountId	string
	)

	se = config.GetStatEvent(int(id))

	if se == nil {
		return errors.New("get stat event config failed")
	}

	if len(intList) != len(se.IntList) {
		return errors.New("stat event int list len mismatch")
	}

	if len(floatList) != len(se.FloatList) {
		return errors.New("stat event float list len mismatch")
	}

	if len(strList) != len(se.StrList) {
		return errors.New("stat event str list len mismatch")
	}

	accountId = strconv.FormatUint(uid, 10)

	ptys = make(map[string]interface{}, 10)
	ptys["account"] = account
	ptys["#time"] = util.GetTimeFromTimeStampMS(t)
	ptys["#ip"] = ip

	for i, item := range se.IntList {

		val := intList[i]

		switch item.Type {

		case config.StatEvtTimeStamp:
			ptys[item.Name] = util.GetTimeFromTimeStampMS(val)
			break

		case config.StatEvtIP:
			ptys[item.Name] = util.InetNtoA(uint32(val))
			break

		case config.StatEvtAmount:
			ptys[item.Name] = float32(val) / 100.0
			break

		default:
			ptys[item.Name] = val
		
		}

	}

	for i, item := range se.FloatList {
		ptys[item.Name] = floatList[i]
	}

	for i, item := range se.StrList {

		val := strList[i]

		switch item.Type {

		case config.StatEvtArray:
			ptys[item.Name] = GetTAArray(val)
			break

		default:
			ptys[item.Name] = val

		}

	}

	err = ta.Track(accountId, distinct, se.Name, ptys)

	if err != nil {
		return err
	}

	ta.Flush()

	return nil

}

func TrackProperty(account string, distinct string, ip string,
	uid uint64, id Proto.TA_ID,
	intVal int64, floatVal float64, strVal string, t int64) error {

	var (
		err 		error
		sp			*config.StatProperty
		ptys		map[string]interface{}
		accountId	string
	)

	sp = config.GetStatProperty(int(id))

	if sp == nil {
		return errors.New("get stat property config failed")
	}

	accountId = strconv.FormatUint(uid, 10)

	ptys = make(map[string]interface{}, 2)
	ptys["account"] = account
	ptys["#time"] = util.GetTimeFromTimeStampMS(t)
	ptys["#ip"] = ip

	switch sp.Type {

	case config.StatPtyInt:
		ptys[sp.Name] = intVal
		break

	case config.StatPtyFloat:
		ptys[sp.Name] = floatVal
		break

	case config.StatPtyString:
		ptys[sp.Name] = strVal
		break

	case config.StatPtyTimeStamp:
		ptys[sp.Name] = util.GetTimeFromTimeStampMS(intVal)
		break

	case config.StatPtyIP:
		ptys[sp.Name] = util.InetNtoA(uint32(intVal))
		break

	case config.StatPtyArray:
		ptys[sp.Name] = GetTAArray(strVal)
		break

	case config.StatPtyAmount:
		ptys[sp.Name] = float32(intVal) / 100.0
		break

	default:
		break

	}

	if sp.Unique == 1 {

		err = ta.UserSetOnce(accountId, distinct, ptys)

	} else {

		err = ta.UserSet(accountId, distinct, ptys)

	}

	if err != nil {
		return err
	}

	return nil
}
