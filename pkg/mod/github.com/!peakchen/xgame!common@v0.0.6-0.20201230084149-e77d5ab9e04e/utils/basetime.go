package utils

import "math"

// ----------------------- time ------------------------------
const (
	ONE_THOUSAND int64 = 1000
	SIXTY        int64 = 60
	TWENTYFOUR   int64 = 24
	THIRTY       int64 = 30
	TWELVE       int64 = 12
)

//sec -> ms
func Sec2Ms_Int64(sec int64) int64 {
	return sec * ONE_THOUSAND
}

//min -> ms
func Min2Ms_Int64(min int64) int64 {
	return Sec2Ms_Int64(min * SIXTY)
}

//hour -> ms
func Hour2Ms_Int64(hour int64) int64 {
	return Min2Ms_Int64(hour * SIXTY)
}

//day -> ms
func Day2Ms_Int64(day int64) int64 {
	return Hour2Ms_Int64(day * TWENTYFOUR)
}

//month -> ms
func Month2Ms_Int64(month int64) int64 {
	return Day2Ms_Int64(month * THIRTY)
}

//year -> ms
func Year2Ms_Int64(year int64) int64 {
	return Month2Ms_Int64(year * TWELVE)
}

// ------------ float64 -----------------
//sec -> ms
func Sec2Ms_Float64(sec float64) int64 {
	return int64(math.Floor(sec * float64(ONE_THOUSAND)))
}

//min -> ms
func Min2Ms_Float64(min float64) int64 {
	return Sec2Ms_Float64(min * float64(SIXTY))
}

//hour -> ms
func Hour2Ms_Float64(hour float64) int64 {
	return Min2Ms_Float64(hour * float64(SIXTY))
}

//day -> ms
func Day2Ms_Float64(day float64) int64 {
	return Hour2Ms_Float64(day * float64(TWENTYFOUR))
}

//month -> ms
func Month2Ms_Float64(month float64) int64 {
	return Day2Ms_Float64(month * float64(THIRTY))
}

//year -> ms
func Year2Ms_Float64(year float64) int64 {
	return Month2Ms_Float64(year * float64(TWELVE))
}
