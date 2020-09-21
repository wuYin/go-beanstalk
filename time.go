package beanstalk

import (
	"strconv"
	"time"
)

type dur time.Duration

// NOTE: convert duration to integer seconds
func (d dur) String() string {
	return strconv.FormatInt(int64(time.Duration(d)/time.Second), 10)
}
