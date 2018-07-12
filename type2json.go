package type2json

import (
	"time"
	"errors"
	"strconv"
)

var (
	ErrType = errors.New("err type")
)


// int64 => time
func i64totime(i int64) (time.Time, error) {
	return time.Unix(i, 0), nil
}

// string => time
func atotime(s string) (time.Time, error) {
	return time.ParseInLocation(
		"2006-01-02 15:04:05",
		s,
		time.Local,
	)
}

// string=>date
func atodate(s string) (time.Time, error) {
	return time.ParseInLocation(
		"2006-01-02",
		s,
		time.Local,
	)
}

// string > int64
// "-1" => -1
func atoi64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}


// time => string
func timetoa(s time.Time) string {
	return s.Format("2006-01-02 15:04:05")
}

// time => string
func timetodate(s time.Time) string {
	return s.Format("2006-01-02")
}

// time => string
func timetoint64(s time.Time) int64 {
	return s.Unix()
}



