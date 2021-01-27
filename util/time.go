package util

import (
	"time"
)

// Get current timestamp.

func GetNowSec() int64 {
	return time.Now().Unix()
}

func GetNowMS() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetNowNS() int64 {
	return time.Now().UnixNano()
}

func GetNowString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Timestamp -> Time

func GetTimeFromTimeStampSec(ts int64) time.Time {
	return time.Unix(ts, 0)
}

func GetTimeFromTimeStampMS(ts int64) time.Time {
	return time.Unix(ts / 1000, (ts % 1000) * 1e6)
}

// Timestamp -> TimeString

func GetTimeStrFromTimeStampSec(ts int64) string {
	return GetTimeFromTimeStampSec(ts).Format("2006-01-02 15:04:05")
}

func GetTimeStrFromTimeStampMS(ts int64) string {
	return GetTimeFromTimeStampMS(ts).Format("2006-01-02 15:04:05")
}

// TimeString -> Time

func GetTimeFromString(timeStr string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", timeStr)
}

// TimeString -> Timestamp

func GetTimeStampSecFromString(timeStr string) (int64, error) {

	t, err := time.Parse("2006-01-02 15:04:05", timeStr)

	if err == nil {
		return 0, err
	}

	return t.Unix(), nil
}


func GetTimeStampMSFromString(timeStr string) (int64, error) {

	t, err := time.Parse("2006-01-02 15:04:05", timeStr)

	if err == nil {
		return 0, err
	}

	return t.UnixNano() / 1e6, nil
}