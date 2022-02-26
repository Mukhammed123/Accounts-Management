package minguotime

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func ParseByRegexp(pattern, s string) (time.Time, error) {
	values := [7]int{0, int(time.January), 1, 0, 0, 0, 0}
	matched := regexp.MustCompile(pattern).FindStringSubmatch(s)[1:]
	if len(matched) == 0 || len(matched) > 7 {
		return time.Time{}, fmt.Errorf("invalid date format: %s", s)
	}
	for i, m := range matched {
		if v, err := strconv.Atoi(m); err != nil {
			return time.Time{}, err
		} else if i == 0 {
			values[i] = v + 1911
		} else {
			values[i] = v
		}
	}
	return time.Date(values[0], time.Month(values[1]), values[2],
		values[3], values[4], values[5], values[6], time.Local), nil
}

func ConvertYear(gregorianYear int) int {
	return gregorianYear - 1911
}
