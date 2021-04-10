package aktime

import (
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

var (
	tReduceVal time.Duration
)

func InitAkTime(redconn redis.Conn) {
	nowt := time.Now()
	redtime, err := redconn.Do("TIME")
	if err != nil {
		panic("redis get time err: " + err.Error())
		return
	}

	var (
		t1 int64
		t2 int64
	)
	for idx, item := range redtime.([]interface{}) {
		t, err := strconv.Atoi(string(item.([]byte)))
		if err != nil {
			continue
		}

		if idx == 0 {
			t1 += int64(t)
		} else {
			t2 += int64(t) * 1e3
		}
	}

	tReduceVal = nowt.Sub(time.Unix(int64(t1), t2))
}

func Now() time.Time {
	return time.Now().Add(tReduceVal)
}

//----------------------- YEAR ----------------------------
// check cross year for now time.
func IsCrossYear4Now(t1 int64) (yes bool) {
	ot1 := time.Unix(t1, 0)
	y1, _, _ := ot1.Date()
	nowy, _, _ := Now().Date()
	if y1 != nowy {
		yes = true
	}
	return
}

// check cross year for compare time
func IsCrossYear(t1 int64, t2 int64) (yes bool) {
	ot1 := time.Unix(t1, 0)
	ot2 := time.Unix(t2, 0)

	y1, _, _ := ot1.Date()
	y2, _, _ := ot2.Date()
	if y1 != y2 {
		yes = true
	}
	return
}

//----------------------- MONTH ----------------------------
// check cross month for now time.
func IsCrossMonth4Now(t1 int64) (yes bool) {
	ot1 := time.Unix(t1, 0)
	y1, m1, _ := ot1.Date()
	nowy, nowm, _ := Now().Date()
	if y1 != nowy || (y1 == nowy && m1 != nowm) {
		yes = true
	}
	return
}

// check cross month for compare time
func IsCrossMonth(t1 int64, t2 int64) (yes bool) {
	ot1 := time.Unix(t1, 0)
	ot2 := time.Unix(t2, 0)

	y1, m1, _ := ot1.Date()
	y2, m2, _ := ot2.Date()
	if (y1 != y2) || (y1 == y2 && m1 != m2) {
		yes = true
	}
	return
}

//----------------------- DAY ----------------------------
// check cross day for zero(0/24) clock.
func IsCrossDay4Zero(t1 int64, t2 int64) (yes bool) {
	ot1 := time.Unix(t1, 0)
	ot2 := time.Unix(t2, 0)

	y1, m1, d1 := ot1.Date()
	y2, m2, d2 := ot2.Date()
	if (y1 != y2) || (y1 == y2 && m1 != m2) || (y1 == y2 && m1 == m2 && d1 != d2) {
		yes = true
	}
	return
}

//check cross "now" day for zero(0/24) clock
func IsCrossDay4ZeroNow(t1 int64) (yes bool) {
	ot1 := time.Unix(t1, 0)
	y1, m1, d1 := ot1.Date()
	nowy, nowm, nowd := Now().Date()
	if (y1 != nowy) || (y1 == nowy && m1 != nowm) || (y1 == nowy && m1 == nowm && d1 != nowd) {
		yes = true
	}
	return
}

//----------------------- HOUR ----------------------------
// check cross 24 hours for diff cross day.
func IsCross24Hour(t1 int64, t2 int64) (yes bool) {
	subv := (t1 - t2)
	if subv < 0 {
		subv = 0 - subv
	}
	yes = (CstOneDaySecs-subv <= 0)
	return
}

// check cross hours.
func IsCrossHours(t1 int64, t2 int64, hours int64) (yes bool) {
	subv := (t1 - t2)
	if subv < 0 {
		subv = 0 - subv
	}
	yes = (hours*CstOneHourSecs-subv <= 0)
	return
}

//----------------------- Min ----------------------------
// check cross mins.
func IsCrossMins(t1 int64, t2 int64, mins int64) (yes bool) {
	subv := (t1 - t2)
	if subv < 0 {
		subv = 0 - subv
	}
	yes = (mins*CstOneMinSecs-subv <= 0)
	return
}

//----------------------- Second ----------------------------
// check cross seconds.
func IsCrossSecs(t1 int64, t2 int64, secs int64) (yes bool) {
	subv := (t1 - t2)
	if subv < 0 {
		subv = 0 - subv
	}
	yes = (secs-subv <= 0)
	return
}
