package ta

import (
	"pluto/util"
	"strings"
	"time"
)

func GetTATimeStr(ts int64) string {

	return time.Unix(ts, 0).Format("01-02 15:04")

}

func GetTAArray(str string) (ary []string) {

	var (
		items		[]string
	)

	ary = make([]string, 0)

	str = util.DeleteSpace(str)

	if (len(str)) != 0 {

		items = strings.Split(str, ",")

		for _, t := range items {

			ary = append(ary, t)

		}

	}

	return ary
}
