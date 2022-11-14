package api

import "time"

// 将RFC3339格式时间转换为time.Time格式
func ConvertDatetime(t string) (tm time.Time, err error) {
	tm, err = time.Parse(time.RFC3339, t)
	return
}

// 将字符串类型的切片转换为字符串
func SpliceSlice(s []string, separator string) (res string) {
	for i, v := range s {
		if i+1 < len(s) {
			res += v + separator + " "
		} else {
			res += v
		}
	}
	res = "[" + res + "]"
	return
}
