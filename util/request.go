package util

import (
	"errors"
	"github.com/gin-gonic/gin"
	"pluto/log"
	"strconv"
)

func parseKey(proto string, key string, c *gin.Context) (string, error) {

	val := c.Request.FormValue(key)

	if len(val) == 0 {
		log.Proto(proto, false, "key %s not exist %s %s", key, c.Request.RemoteAddr, c.Request.URL)
		return "", errors.New("Key" + key + "Not Exist")
	}

	return val, nil

}

func parseKeyCanBeNil(_, key string, c *gin.Context) (string, error) {

	var (
		err	error
		val	string
	)

	err = c.Request.ParseForm()

	if err != nil {
		return "", err
	}

	val = c.Request.FormValue(key)

	return val, nil

}

func GetParamString(api string, key string, c *gin.Context) (string, error) {

	return parseKey(api, key, c)

}

func GetParamStringCanBeNil(api string, key string, c *gin.Context) (string, error) {

	return parseKeyCanBeNil(api, key, c)

}

func GetParamInt(api string, key string, c *gin.Context) (int, error) {

	val, err := parseKey(api, key, c)

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(val)

}

func GetParamUInt64(api string, key string, c *gin.Context) (uint64, error) {

	val, err := parseKey(api, key, c)

	if err != nil {
		return 0, err
	}

	return strconv.ParseUint(val, 0, 64)

}

func GetParamUInt32(api string, key string, c *gin.Context) (u32 uint32, err error) {

	var (
		val	string
		u64 uint64
	)

	val, err = parseKey(api, key, c)

	if err != nil {
		return 0, err
	}

	u64, err = strconv.ParseUint(val, 0, 32)

	u32 = uint32(u64)

	return

}