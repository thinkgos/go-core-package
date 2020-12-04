package exttime

import (
	"testing"
	"time"
)

func TestServiceTime(t *testing.T) {
	time.Sleep(time.Second * 5)
	ts := ServiceStartupTime()
	t.Log(Millisecond(ts))
	t.Log(Microsecond(ts))
	t.Log(Time(Millisecond(time.Now())))
	t.Log(ServiceStartupTime())
	t.Log(ServiceElapseTime())
	t.Log(ServiceUptime())
}
