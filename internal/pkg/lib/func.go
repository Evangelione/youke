package lib

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func Md5Crypt(str string, salt ...interface{}) string {
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func GetFloat64LongAndLat(long, lat string) (float64Long, float64Lat float64, err error) {
	if long != "" {
		float64Long, err = strconv.ParseFloat(long, 64)
	}

	if lat != "" {
		float64Lat, err = strconv.ParseFloat(lat, 64)
	}

	return float64Long, float64Lat, err
}
