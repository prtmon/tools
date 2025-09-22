package tools

import (
	"time"
)

// TimeFrame 定义了不同的K线时间周期
type TimeFrame string

const (
	TimeFrameMinute   TimeFrame = "1m"
	TimeFrame5Minute  TimeFrame = "5m"
	TimeFrame15Minute TimeFrame = "15m"
	TimeFrame30Minute TimeFrame = "30m"
	TimeFrameHour     TimeFrame = "1h"
	TimeFrame4Hour    TimeFrame = "4h"
	TimeFrameDay      TimeFrame = "1d"
)

// ToDuration 将时间转换为K线周期
func (tf TimeFrame) ToDuration() time.Duration {
	switch tf {
	case TimeFrameMinute:
		return time.Minute
	case TimeFrame5Minute:
		return 5 * time.Minute
	case TimeFrame15Minute:
		return 15 * time.Minute
	case TimeFrame30Minute:
		return 30 * time.Minute
	case TimeFrameHour:
		return time.Hour
	case TimeFrame4Hour:
		return 4 * time.Hour
	case TimeFrameDay:
		return 24 * time.Hour
	default:
		panic("unsupported time frame")
	}
}

// Floor 将时间戳转换为K线时间周期的起始时间戳
func (tf TimeFrame) Floor(t time.Time) time.Time {
	return t.Truncate(tf.ToDuration())
}

// FloorSec 将sec时间戳转换为K线时间周期的起始时间戳
func (tf TimeFrame) FloorSec(sec int64) time.Time {
	unixT := time.Unix(sec, 0)
	return unixT.Truncate(tf.ToDuration())
}

// FloorMsec 将msec时间戳转换为K线时间周期的起始时间戳
func (tf TimeFrame) FloorMsec(msec int64) time.Time {
	unixT := time.UnixMilli(msec)
	return unixT.Truncate(tf.ToDuration())
}
