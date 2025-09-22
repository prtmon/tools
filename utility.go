package tools

import (
	"fmt"
	"github.com/spf13/cast"
	"strconv"
	"strings"
)

func IsExistString(in string, list []string) bool {
	for _, value := range list {
		if in == value {
			return true
		}
	}
	return false
}

func AbsInt64(in int64) (out int64) {
	if in < 0 {
		out = -in
	} else {
		out = in
	}
	return
}

func ToFloat64(v interface{}) float64 {
	if v == nil {
		return 0.0
	}
	switch v.(type) {
	case float64:
		return v.(float64)
	case string:
		vStr := v.(string)
		vF, _ := strconv.ParseFloat(vStr, 64)
		return vF
	case int64:
		return float64(v.(int64))
	case int32:
		return float64(v.(int32))
	case int:
		return float64(v.(int))
	default:
		panic("to float64 error.")
	}
}

func Int64ToString(in int64) (out string) {
	out = fmt.Sprintf("%d", in)
	return out
}

func IntToString(in int) (out string) {
	out = strconv.Itoa(in)
	return out
}

func ToInt(v interface{}) int {
	if v == nil {
		return 0
	}

	switch v.(type) {

	case string:
		vStr := v.(string)
		vInt, _ := strconv.Atoi(vStr)
		return vInt
	case int:
		return v.(int)
	case int64:
		return int(v.(int64))
	case float64:
		vF := v.(float64)
		return int(vF)
	default:
		panic("to int error.")
	}
}

func ToInt64(v interface{}) int64 {
	if v == nil {
		return 0
	}

	switch v.(type) {
	case float64:
		return int64(v.(float64))
	default:
		vv := fmt.Sprint(v)

		if vv == "" {
			return 0
		}

		vvv, err := strconv.ParseInt(vv, 0, 64)
		if err != nil {
			return 0
		}

		return vvv
	}
}

// 四舍五入
func Float64Rand(v float64, dig int) float64 {
	cDig := strconv.Itoa(dig)
	val := fmt.Sprintf("%0."+cDig+"f", v)
	return cast.ToFloat64(val)
}

// 四舍五入
func Float64RandStr(v float64, dig int) string {
	cDig := strconv.Itoa(dig)
	val := fmt.Sprintf("%0."+cDig+"f", v)
	return val
}

// 四舍五入
func Float64StrRand(v string, dig int) string {
	fv := ToFloat64(v)
	cDig := strconv.Itoa(dig)
	val := fmt.Sprintf("%0."+cDig+"f", fv)
	return val
}

// 小数点位数
func GetRoundIDig(in string) int {
	floatString := strings.ReplaceAll(in, ",", "")
	return len(strings.Split(floatString, ".")[1])
}

// 小数点位数
func GetRoundSDig(in string) string {
	floatString := strings.ReplaceAll(in, ",", "")
	return strconv.Itoa(len(strings.Split(floatString, ".")[1]))
}

func IntAdd(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func AddFloat64(a, b interface{}) float64 {
	return ToFloat64(a) + ToFloat64(b)
}

func SubtractFloat64(a, b interface{}) float64 {
	return ToFloat64(a) - ToFloat64(b)
}

// 辅助函数
func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func Min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
