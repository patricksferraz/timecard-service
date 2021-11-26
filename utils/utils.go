package utils

import (
	"bytes"
	"os"
	"regexp"
	"time"
	"unicode"
)

func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func CleanNonDigits(str *string) {
	buf := bytes.NewBufferString("")
	for _, r := range *str {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	*str = buf.String()
}

func IsClock(str *string) bool {
	match, _ := regexp.MatchString("^([0-1]?[0-9]|[2][0-3]):?([0-5][0-9])(:?[0-5][0-9])?$", *str)
	return match
}
